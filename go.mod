module github.com/octoproject/octo-cli

go 1.14

require (
	github.com/AlecAivazis/survey/v2 v2.0.8
	github.com/go-git/go-git/v5 v5.1.0
	github.com/spf13/cobra v1.0.0
	golang.org/x/net v0.0.0-20200707034311-ab3426394381 // indirect
	golang.org/x/sys v0.0.0-20200625212154-ddb9806d33ae // indirect
	gopkg.in/yaml.v2 v2.3.0
	k8s.io/api v0.18.1
	k8s.io/apimachinery v0.18.1
	knative.dev/client v0.16.0
	knative.dev/serving v0.16.0
)

replace (
	github.com/spf13/cobra => github.com/chmouel/cobra v0.0.0-20191021105835-a78788917390
	k8s.io/api => k8s.io/api v0.17.6
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.17.6
	k8s.io/apimachinery => k8s.io/apimachinery v0.17.6
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.17.6
	k8s.io/client-go => k8s.io/client-go v0.17.6
	k8s.io/code-generator => k8s.io/code-generator v0.17.6
)
