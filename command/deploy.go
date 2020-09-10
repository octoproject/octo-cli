package command

import (
	"errors"
	"fmt"

	"github.com/octoproject/octo-cli/config"
	"github.com/octoproject/octo-cli/faas"
	"github.com/octoproject/octo-cli/knative"
	"github.com/spf13/cobra"
)

var (
	ErrBadGatewayURL       = errors.New("error empty gateway url")
	ErrOpenfaasCredentials = errors.New("error empty Openfaas username or password")
	ErrUnknownPlatform     = errors.New("error unknown platform")
)

type deplpyOptions struct {
	fileName        string
	username        string
	password        string
	gatewayURL      string
	namespace       string
	image           string
	imagePullPolicy string
}

func NewDeployCommand() *cobra.Command {
	options := deplpyOptions{}

	cmd := &cobra.Command{
		Use:   "deploy",
		Short: "Deploy a new service",
		Example: `octo-cli deploy -f example.yml -i functions/nodeinfo-http:latest
octo-cli deploy -f example.yml -g http://127.0.0.1:8080 -i functions/nodeinfo-http:latest -u admin -p 123456 `,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runDeploy(options)
		},
		SilenceUsage: true,
	}

	flags := cmd.Flags()
	flags.StringVarP(&options.fileName, "file", "f", "", "Name of the service yml file")
	flags.StringVarP(&options.username, "username", "u", "", "Openfaas gateway username")
	flags.StringVarP(&options.password, "password", "p", "", "Openfaas gateway password")
	flags.StringVarP(&options.gatewayURL, "gateway", "g", "http://127.0.0.1:8080", "Openfaas gateway URL")
	flags.StringVarP(&options.namespace, "namespace", "n", "", "Namespace for deployed function")
	flags.StringVarP(&options.image, "image", "i", "", "Docker image name")
	flags.StringVarP(&options.imagePullPolicy, "pullPolicy", "", "", "Docker image pull policy")

	return cmd
}

func runDeploy(o deplpyOptions) error {
	s, err := config.LoadService(o.fileName)
	if err != nil {
		return err
	}

	env := map[string]string{
		"DB_USER":     s.DB.User,
		"DB_PASSWORD": s.DB.Password,
		"DB_NAME":     s.DB.Name,
		"DB_HOST":     s.DB.Host,
		"DB_DIALECT":  s.DB.Type,
		"DB_PORT":     s.DB.Port,
		"DB_TIMEOUT":  s.DB.RequestTimeout,
	}

	switch s.Platform {
	case "openfaas":
		if len(o.username) < 1 || len(o.password) < 1 {
			return ErrOpenfaasCredentials
		}

		c := faas.New(o.username, o.password, o.gatewayURL)

		f := faas.Function{
			ServiceName: s.ServiceName,
			Image:       o.image,
			Namespace:   o.namespace,
			EnvVars:     env,
		}

		err := c.DeployFunction(&f)
		if err != nil {
			return err
		}

		fmt.Printf("Servive %s deployed successfully.\n", s.ServiceName)

	case "knative":
		f := knative.Function{
			ServiceName: s.ServiceName,
			Image:       o.image,
			Namespace:   o.namespace,
			EnvVars:     env,
		}

		err := knative.DeployFunction(&f)
		if err != nil {
			return err
		}

		fmt.Printf("Servive %s deployed successfully.\n", s.ServiceName)
	default:
		return ErrUnknownPlatform

	}
	return nil
}
