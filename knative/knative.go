package knative

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	clientservingv1 "knative.dev/client/pkg/serving/v1"
	"knative.dev/serving/pkg/apis/serving"
	servingv1 "knative.dev/serving/pkg/apis/serving/v1"

	"knative.dev/client/pkg/kn/commands"
)

const (
	// How often to retry
	MaxUpdateRetries                          = 3
	PullAlways              corev1.PullPolicy = "Always"
	knativeDefaultNamespace                   = "default"
)

type Function struct {
	ServiceName     string
	Image           string
	ImagePullPolicy string
	Namespace       string
	EnvVars         map[string]string
}

//DeployFunction deploy knative function
func DeployFunction(f *Function) error {
	if len(f.Namespace) < 1 {
		f.Namespace = knativeDefaultNamespace
	}

	// create kn client
	client, err := f.newKnativeClient()
	if err != nil {
		return err
	}

	service := f.constructService()

	serviceExists, err := f.serviceExists(client, f.ServiceName)
	if err != nil {
		return err
	}

	if !serviceExists {
		err = client.CreateService(service)
		if err != nil {
			return err
		}
		return nil
	}

	err = f.prepareAndUpdateService(client, service)
	if err != nil {
		return err
	}
	return nil
}

func (f *Function) newKnativeClient() (clientservingv1.KnServingClient, error) {
	p := &commands.KnParams{}
	p.Initialize()

	client, err := p.NewServingClient(f.Namespace)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (f *Function) constructService() *servingv1.Service {
	imgPullPolicy := PullAlways
	if len(f.ImagePullPolicy) > 1 {
		imgPullPolicy = corev1.PullPolicy(f.ImagePullPolicy)
	}

	//create env var
	var envVar []corev1.EnvVar
	for key, value := range f.EnvVars {
		n := corev1.EnvVar{Name: key, Value: value}
		envVar = append(envVar, n)
	}

	//create service
	service := &servingv1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      f.ServiceName,
			Namespace: f.Namespace,
		},
	}

	service.Spec.Template.Spec.Containers = []corev1.Container{
		{
			Image:           f.Image,
			Env:             envVar,
			ImagePullPolicy: imgPullPolicy,
		},
	}
	return service
}

//serviceExists return true if function exists , source: https://github.com/knative/client/blob/master/pkg/kn/commands/service/create.go#L238
func (f *Function) serviceExists(client clientservingv1.KnServingClient, name string) (bool, error) {
	_, err := client.GetService(name)
	if apierrors.IsNotFound(err) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

// prepareAndUpdateService,source: https://github.com/knative/client/blob/master/pkg/kn/commands/service/create.go#L238
func (f *Function) prepareAndUpdateService(client clientservingv1.KnServingClient, service *servingv1.Service) error {
	var retries = 0
	for {
		existingService, err := client.GetService(service.Name)
		if err != nil {
			return err
		}

		// Copy over some annotations that we want to keep around. Erase others
		copyList := []string{
			serving.CreatorAnnotation,
			serving.UpdaterAnnotation,
		}

		// If the target Annotation doesn't exist, create it even if
		// we don't end up copying anything over so that we erase all
		// existing annotations
		if service.Annotations == nil {
			service.Annotations = map[string]string{}
		}

		// Do the actual copy now, but only if it's in the source annotation
		for _, k := range copyList {
			if v, ok := existingService.Annotations[k]; ok {
				service.Annotations[k] = v
			}
		}

		service.ResourceVersion = existingService.ResourceVersion
		err = client.UpdateService(service)
		if err != nil {
			// Retry to update when a resource version conflict exists
			if apierrors.IsConflict(err) && retries < MaxUpdateRetries {
				retries++
				continue
			}
			return err
		}
		return nil
	}
}
