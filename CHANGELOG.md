# CHANGELOG

## v1.0.0

Hey everyone, I'm glad to release `clencli` with all of you. This moment has a very special taste, as I always wanted to give back to the community. A legion of anonymous people who always were supportive and kind enough to guide and teach me and many others. So, please accept this gift called `clencli` as a small token of appreciation, thank you everyone!

### Release notes

#### New features and changes

* Create a command to initialize projects

By running `clencli init project --name <value>` you can easily initialize a project. Depending of `--type` value given, it will create the project with the necessary files for your cloud project. A folder `clencli` will be created containing templates and yaml files to ease the rendering of such templates. To know more, please run: `clencli init --help`.

* Create a command to render templates

When you run `clencli render template --name <value>`, `clencli` renders a template located at `clencli/*.tmpl` based on their respective `clencli/*.yaml` database. `clencli` uses [Gomplate](https://github.com/hairyhenderson/gomplate) as template renderer, you will find more about Gomplate docs [here](https://docs.gomplate.ca/).

* Create a command to download pictures from Unsplash.com

I wanted to personalize projects during initialization or when desired. Therefore, I've decided to use [unpslash](https://unsplash.com/) as source to fetch photos. You can download images by using the command `clencli unsplash`, or by defining the `theme` at the `clencli/readme.yaml` file. However, you first will need to create your [Unsplash Developer account](https://unsplash.com/documentation#creating-a-developer-account). After that, you can create a new application as `demo` and copy `Access key` and `Secret key` into your `clencli` config, usually at `$HOME/.clencli.yaml` ( you will need to create this file manually ):

```
unsplash:
  access_key: "XXX"
  secret_key: "XXX"
```

Now, commands can use `unsplash` to download photos in commands such as `clencli unsplash` or `clencli render`.

