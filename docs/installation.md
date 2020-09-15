---
layout: page
title: Installation
permalink: /installation/
nav_order: 1
---



# Install Octo CLI
Octo CLI is available as source code or a pre-compiled binary,You can download pre-built binaries for Linux, macOS, and Windows on the [releases page](https://github.com/octoproject/octo-cli/releases).

## Homebrew install instructions


On macOS or Linux, you can install Octo CLI via [Homebrew](https://brew.sh/https://brew.sh/):

- Install the tap via

```
$ brew tap octoproject/octoproject
```

- Install octo-cli

```
$ brew install octo-cli
```

# Verifying the Installation

After installing Octo CLI, verify that the installation worked by opening a new terminal session and typing `octo-cli` you should see help output similar to the following:


``` bash
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


