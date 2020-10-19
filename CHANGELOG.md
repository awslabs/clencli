# CHANGELOG

## v1.1.0

A few bug fixes and a new features added.

### Release notes

#### New features and changes

* Customize your project initialization

You can customize the structure of your projects by defining it into `clencli` config.

```
init:
  types:
    - type: basic # clencli will always execute basic init
      name: default # name of initialization, allow user to name different configs, will always execute init with name default
      enabled: true # allow to keep the configuration but disable it in any moment
      files: 
        - file:
            path: docs
            state: directory # to create a directory
        - file:
            path: a/b/c/d/e/f/g/h # to create nested directories
            state: directory
        - file: 
            src: /tmp/CODE_OF_CONDUCT.md # to copy a file locally
            dest: CODE_OF_CONDUCT.md
            state: file
        - file:
            src: https://raw.githubusercontent.com/valter-silva-au/company-master-template/main/LICENSE # to copy a file remotelly
            dest: LICENSE
            state: file
        - file:
            src: https://raw.githubusercontent.com/valter-silva-au/company-master-template/919a814bc44fa86a72a004fa99f2319a84838790/readme.tmpl # possible to use versioning capabilities such as commits, or tags
            dest: clencli/readme.tmpl
            state: file
        - file:
            src: https://raw.githubusercontent.com/valter-silva-au/company-master-template/ec22acd40123b413e05751b92c07d8fc244ea282/readme.yaml
            dest: clencli/readme.yaml
            state: file
```

You can combine different project types into `init` definition:

```
init:
  types:
    - type: terraform # https://www.hashicorp.com/resources/a-practitioner-s-guide-to-using-hashicorp-terraform-cloud-with-github
      name: default
      enabled: true
      files:
        - file:
            src: https://raw.githubusercontent.com/hashicorp/learn-terraform-modules/master/main.tf
            dest: main.tf
            state: file
        - file:
            src: https://raw.githubusercontent.com/hashicorp/learn-terraform-modules/master/outputs.tf
            dest: outputs.tf
            state: file
        - file:
            src: https://raw.githubusercontent.com/hashicorp/learn-terraform-modules/master/variables.tf
            dest: variables.tf
            state: file
        - file:
            src: https://raw.githubusercontent.com/hashicorp/learn-terraform-modules/master/.gitignore
            dest: .gitignore
            state: file
    - type: cloudformation
      name: default
      enabled: true
      files:
        - file:
            src: https://raw.githubusercontent.com/awslabs/aws-cloudformation-templates/master/aws/solutions/WordPress_Single_Instance.yaml
            dest: stack.yml
            state: file
```

You can also name you initializations and only execute them for the type of project you want:
```
init:
  types:
    - type: terraform # https://www.hashicorp.com/resources/a-practitioner-s-guide-to-using-hashicorp-terraform-cloud-with-github
      name: module
      enabled: true
      files:
        - file:
            path: modules
            state: directory
        - file:
            src: https://raw.githubusercontent.com/hashicorp/learn-terraform-modules/master/main.tf
            dest: modules/main.tf
            state: file
        - file:
            src: https://raw.githubusercontent.com/hashicorp/learn-terraform-modules/master/outputs.tf
            dest: modules/outputs.tf
            state: file
        - file:
            src: https://raw.githubusercontent.com/hashicorp/learn-terraform-modules/master/variables.tf
            dest: modules/variables.tf
            state: file
        - file:
            src: https://raw.githubusercontent.com/hashicorp/learn-terraform-modules/master/.gitignore
            dest: .gitignore
            state: file
    - type: cloudformation
      name: nested
      enabled: true
      files:
        - file:
            src: https://raw.githubusercontent.com/valter-silva-au/company-master-template/main/cloudformation-nested-stack.yml
            dest: nested-stack.yml
            state: file
```