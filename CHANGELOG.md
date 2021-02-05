# CHANGELOG

## v0.2.1

A few bug fixes and a new features added.

### Release notes

### Bugfixes

* Remove trailing characters from template's database

```
.. shorten for brevity
usage: |-
  In a polyglot world where a team can choose it's programming language, often this flexibility can spill into chaos as every repo looks different.
  CLENCLI solves this issue by giving developers a quick and easy way to create a standardised repo structure and easily rendering documentation via a YAML file. <whitespace>
.. shorten for brevity
```

If you had a whitespace in your `clencli/readme.yaml` and executed `clencli render template` it would overwrite its original content to conform with YAML format.

#### New features and changes

* Download photo from unsplash by ID

```
$ clencli unsplash --id=jLjfAWwHdB8
```

A file named `unsplash.yaml` will be created and placed in the current directory. It contains the response from Unsplash API.
It's useful to give credit to the photo's author and provides information about the photo and the author.
