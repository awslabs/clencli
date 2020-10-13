PROJECT_NAME=generated-project-0

.PHONY: test
test: test/all ## Test all commands

.PHONY: test/all
test/all: clean build test/init/terraform test/build/template

.PHONY: test/init/terraform
test/init/terraform:
	make run command='init terraform --name=$(PROJECT_NAME)'

.PHONY: test/build/template
test/build/template: test/build/template/type/inputs test/build/template/type/outputs
	cd $(PROJECT_NAME) && clencli template

.PHONY: test/build/template/type/inputs
test/build/template/type/inputs:
	cd $(PROJECT_NAME) && clencli template --type inputs

.PHONY: test/build/template/type/outputs
test/build/template/type/outputs:
	cd $(PROJECT_NAME) && clencli template --type outputs