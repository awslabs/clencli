<!--

  ** DO NOT EDIT THIS FILE
  ** 
  ** This file was automatically generated by the [CLENCLI](https://github.com/awslabs/clencli)
  ** 1) Make all changes on files under clencli/yaml/*.yaml
  ** 2) Run `clencli template` to rebuild this file
  **
  ** By following this practice we ensure standard and high-quality accross multiple projects.
  ** DO NOT EDIT THIS FILE

-->

![Logo](https://images.unsplash.com/photo-1611965581519-1f5a5ee98178?crop=entropy&cs=tinysrgb&fit=max&fm=jpg&ixid=MXwxOTEyNTB8MHwxfHJhbmRvbXx8fHx8fHx8&ixlib=rb-1.2.1&q=80&w=1080)

[![GitHub issues](https://img.shields.io/github/issues/awslabs/clencli)](https://github.com/awslabs/clencli/issues)[![GitHub forks](https://img.shields.io/github/forks/awslabs/clencli)](https://github.com/awslabs/clencli/network)[![GitHub stars](https://img.shields.io/github/stars/awslabs/clencli)](https://github.com/awslabs/clencli/stargazers)[![GitHub license](https://img.shields.io/github/license/awslabs/clencli)](https://github.com/awslabs/clencli/blob/master/LICENSE)[![Twitter](https://img.shields.io/twitter/url?style=social&url=https%3A%2F%2Fgithub.com%2Fawslabs%2Fclencli)](https://twitter.com/intent/tweet?text=Wow:&url=https%3A%2F%2Fgithub.com%2Fawslabs%2Fclencli)

# Cloud Engineer CLI  ( clencli ) 

A CLI built to assist Cloud Engineers.

## Table of Contents
---




 - [Usage](#usage) 

 - [Installing](#installing) 


 - [Acknowledgments](#acknowledgments) 
 - [Contributors](#contributors) 
 - [References](#references) 
 - [License](#license) 
 - [Copyright](#copyright) 




## Usage
---
<details open>
  <summary>Expand</summary>

In a polyglot world where a team can choose it's programming language, often this flexibility can spill into chaos as every repo looks different.
CLENCLI solves this issue by giving developers a quick and easy way to create a standardised repo structure and easily rendering documentation via a YAML file.

### Create a new project
```
  $ clencli init project --project-name foo
  $ tree -a moon/
  foo/
  ├── clencli
  │   ├── readme.tmpl
  │   └── readme.yaml
  └── .gitignore
```

### Create a new CloudFormation project
```
$ clencli init project --project-name foo --project-type cloudformation
$ tree -a sun/
  foo/
  ├── clencli
  │   ├── hld.tmpl
  │   ├── hld.yaml
  │   ├── readme.tmpl
  │   └── readme.yaml
  ├── environments
  │   ├── dev
  │   └── prod
  ├── .gitignore
  ├── skeleton.json
  └── skeleton.yaml
```

### Create a new Terraform project
```
$ clencli init project --project-name foo --project-type terraform
$ tree -a foo/
foo/
├── clencli
│   ├── hld.tmpl
│   ├── hld.yaml
│   ├── readme.tmpl
│   └── readme.yaml
├── environments
│   ├── dev.tf
│   └── prod.tf
├── .gitignore
├── LICENSE
├── main.tf
├── Makefile
├── outputs.tf
└── variables.tf
```

## Render a template
```
$ clencli init project --project-name foo
foo was successfully initialized as a basic project
$ cd foo/
$ clencli render template
Template readme.tmpl rendered as README.md
```

The `README.md` you are reading right now was generated and it's maintained by `CLENCLI` itself.
You can check [readme.yaml](clencli/readme.yaml) for more details. Every time the `README.md` is updated, a new photo is chosen for the project automatically.

## Download random photos from [Unsplash](https://unsplash.com)
```
# first you need to inform your unsplash developer API credentials

$ clencli configure
clencli configuration directory created at /home/valter/.clencli
Would you like to setup credentials? [false]: true
> Credentials
>> Profile: default
>>>> Credential
>>>>> Name: default
>>>>> Description:
>>>>> Enabled [true]:
>>>>> Provider: unsplash
>>>>> Access Key []: XXX
>>>>> Secret Key []: XXX
>>>>> Session Token []:
Would you like to setup another credential? [false]:
Would you like to setup configurations? [false]:

$ clencli unsplash
tree -a downloads/
downloads/
└── unsplash
    └── mountains
        ├── 3gz2hsA1T3s-full.jpeg
        ├── 3gz2hsA1T3s-raw.jpeg
        ├── 3gz2hsA1T3s-regular.jpeg
        ├── 3gz2hsA1T3s-small.jpeg
        └── 3gz2hsA1T3s-thumb.jpeg

$ clencli unplash --query dog
clencli unsplash --query dog --size full
tree -a downloads/
downloads/
└── unsplash
    └── dog
        └── bbjSWtDtHbM.jpeg
```

## Download a .gitignore for your project
```
$ clencli gitignore --input terraform,vscode
.gitignore created successfully
$ less .gitignore

# Created by https://www.toptal.com/developers/gitignore/api/terraform,vscode
# Edit at https://www.toptal.com/developers/gitignore?templates=terraform,vscode

### Terraform ###
# Local .terraform directories
**/.terraform/*

# .tfstate files
*.tfstate
*.tfstate.*

# Crash log files
crash.log

# Ignore any .tfvars files that are generated automatically for each Terraform run. Most
# .tfvars files are managed as part of configuration and so should be included in
# version control.
#
# example.tfvars

# Ignore override files as they are usually used to override resources locally and so
# are not checked in
override.tf
override.tf.json
*_override.tf
*_override.tf.json

# Include override files you do wish to add to version control using negated pattern
# !example_override.tf

# Include tfplan files to ignore the plan output of command: terraform plan -out=tfplan
# example: *tfplan*

### vscode ###
.vscode/*
!.vscode/settings.json
!.vscode/tasks.json
!.vscode/launch.json
!.vscode/extensions.json
*.code-workspace

# End of https://www.toptal.com/developers/gitignore/api/terraform,vscode
```
</details>





## Installing
---
<details open>
  <summary>Expand</summary>

Download the latest version [released](https://github.com/awslabs/clencli/releases) according to your platform and execute it directly. I recommend placing the binary into your `$PATH`, so it's easily accessible.
</details>









## Commands
```
The Cloud Engineer CLI

Usage:
  clencli [command]

Available Commands:
  configure   Configures CLENCLI global settings
  gitignore   Download .gitignore based on the given input
  help        Help about any command
  init        Initialize a project
  render      Render a template
  unsplash    Downloads random photos from Unsplash.com
  version     Displays the version of CLENCLI and all installed plugins

Flags:
  -h, --help                   help for clencli
      --log                    Enable or disable logs (can be found at ./clencli/log.json). Log outputs will be redirected default output if disabled. (default true)
      --log-file-path string   Log file path. Requires log=true, ignored otherwise. (default "clencli/log.json")
  -p, --profile string         Use a specific profile from your credentials and configurations file. (default "default")
  -v, --verbosity string       Valid log level:panic,fatal,error,warn,info,debug,trace). (default "error")

Use "clencli [command] --help" for more information about a command.
```




## Contributors
---
<details open>
  <summary>Expand</summary>

|     Name     |         Email        |       Role      |
|:------------:|:--------------------:|:---------------:|
|  Silva, Valter  |  -  |  AWS Professional Services - Cloud Architect  |

</details>



## Acknowledgments
---
<details>
  <summary>Expand</summary>

Gratitude for assistance:
  * Sia, William - AWS Professional Service - Senior Cloud Architect
  * Dhingra, Prashit - AWS Professional Service - Cloud Architect


</details>



## References
---
<details open>
  <summary>Expand</summary>

  * [cobra](https://github.com/spf13/cobra) - Cobra is both a library for creating powerful modern CLI applications as well as a program to generate applications and command files.
  * [viper](https://github.com/spf13/viper) - Viper is a complete configuration solution for Go applications including 12-Factor apps.
  * [twelve-factor-app](https://12factor.net) - The Twelve-Factor App
  * [gomplate](https://github.com/hairyhenderson/gomplate) - gomplate is a template renderer which supports a growing list of datasources, such as JSON (including EJSON - encrypted JSON), YAML, AWS EC2 metadata, BoltDB, Hashicorp Consul and Hashicorp Vault secrets.
  * [unsplash](https://unsplash.com) - The most powerful photo engine in the world.
  * [placeholder](https://placeholder.com) - The Free Image Placeholder Service Favoured By Designers
  * [pirate-ipsum](https://pirateipsum.me) - The best Lorem Ipsum Generator in all the sea
  * [recordit](https://recordit.co) - Record Fast Screencasts
  * [ttystudio](https://github.com/chjj/ttystudio) - A terminal-to-gif recorder minus the headaches.
  * [gihub-super-linter](https://github.com/github/super-linter) - GitHub Super Linter
  * [github-actions](https://docs.github.com/en/free-pro-team@latest/actions/learn-github-actions/introduction-to-github-actions) - GitHub Actions
  * [gitignore.io](https://www.toptal.com/developers/gitignore) - Create useful .gitignore files for your project


</details>



## License
---
This project is licensed under the Apache License 2.0.

For more information please read [LICENSE](LICENSE).



## Copyright
---
```
Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
```

