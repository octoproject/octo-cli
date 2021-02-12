# Octo CLI

`octo-cli`  makes the data available from any database as a serverless web service, simplifying the process of building data-driven applications. 

Knative and OpenFaaS are the only supported serverless frameworks in `octo-cli` for now.

<img src="https://user-images.githubusercontent.com/20528562/92412306-1b8d1880-f154-11ea-974e-8b610cbb4ea4.png" max-width="100%" />



Octo will create an endpoint that will expose your data as service, all you need to provide is yml file that describes your service.

![overview](https://user-images.githubusercontent.com/20528562/92733888-b9652b00-f380-11ea-9643-9845953050dd.png)

# Supported Databases
- PostgreSQL
- MSSQL
- MySQL

# Supported Serverless Frameworks
- OpenFaaS
- Knative

# Installation
[Download Latest Binary](https://github.com/octoproject/octo-cli/releases/latest)

Alternatively you can install using go:

```bash
go get github.com/octoproject/octo-cli
```

# Documentation
Documentation can be found on [here](https://octoproject.github.io/octo-cli/quick-start/).

# Examples
Examples can be found in the [examples/](https://github.com/octoproject/octo-cli/tree/master/examples) directory. They are step-by-step examples that will help you to deploy your first service using 
 `octo-cli` 

# Usage

```
$ octo-cli      
Expose data from any database as web service

Usage:
  octo-cli [flags]
  octo-cli [command]

Available Commands:
  build       Build function Docker container
  create      Create a new service
  deploy      Deploy a new service
  help        Help about any command
  init        Generate service configuration YAML file

Flags:
  -h, --help   help for octo-cli

Use "octo-cli [command] --help" for more information about a command.

```

