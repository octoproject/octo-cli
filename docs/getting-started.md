---
layout: page
title: Getting Started
permalink: /quick-start/
nav_order: 3
---

# Getting Started
{: .no_toc }

This page summarizes all the essentials you will need to know to start with Octo.

{: .fs-6 .fw-300 }



## Table of contents
{: .no_toc .text-delta }

1. TOC
{:toc}



## Prerequisites

 This tutorial assumes that you already have an OpenFaaS or Knative deployed on either Minikube or the cloud(GKE, Digitalocean..etc).
For more detailed installation instructions,check  [Knative ](https://knative.dev/docs/install/) 
 and [OpenFaaS](https://docs.openfaas.com/deployment/).


## Installation
if you did not install Octo CLI already,please refer [here](/octo/installation/) for instructions.

## Service configuration file explained
Octo will generate a service configuration YAML file that will be used through the deployment process, 
it has the following configurations:-

- `service_name`: name of the service which would be used in HTTP Route.
-  `query`: SQL Statement.
- `parameters`: represent the input for the API and they will be would be passed to the query.
-  `db`:  is the source database for the  API to run the given query against.
- `service_type`: describe if the API will read data or write it to the database.
- `platform`: define the serverless framework that API will be deployed to.


## Create a new service
The below instructions will help you deploy your first service using Octo CLI.

For examples, they can be found in the repo [examples/](https://github.com/octoproject/octo-cli/tree/master/examples) directory. 

### Step1: Generate Octo configuration
Generate Octo configuration by running
```
$ octo-cli init 
```
Then you will be asked to add the query, database credentials, API parameters, service name, and it
going to generate a YAML file with the service name `service-name-config.yml`

### Step3: Create a new service
To create a new service run the following command:
```
$  octo-cli  create -f service-name-config.yml
```
as result, it creates a new service folder in the current path with the service name as a name.

### Step4: Build function Docker container
To build function Docker container type the following command: 
```
 octo-cli  build -f service-name-config.yml  --prefix dev.local --tag v1
```
If no tag is provided, the `:latest` is the default tag .


### Step5: Deploy the service
The `deploy` command have different options based on the platform,to see the full options run `octo-cli deploy --help`

- To deploy to OpenFaaS run the following command: 
```
$ octo-cli  deploy -f service-name-config.yml -i dev.local/get-users:latest  -u admin -p 41d21dfa77da9 -g http://127.0.0.1:8080
```
- To deploy to Knative run the following command: 
```
$ octo-cli deploy -f service-name-config.yml -i dev.local/get-users:latest
```
When deploying to Knative, Octo detects and uses a kubeconfig file to communicate with the Kubernetes cluster. The kubeconfig file should be stored at $HOME/.kube.

### Step6: Test the service
- OpenFaaS
```
$ curl --location --request GET 'http://openfaas-gateway-url/function/service-name'
```

- Knative
```
$ curl --location --request GET 'http://knative-url/service-name'
```