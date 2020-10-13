# readme.yaml
Possible keys and values that can be used at `clencli/readme.yaml`.


## logo
Allows you to customize the project
```
logo:
  theme: <string> # Fetch a random photo from Unsplash
  url: <URL or local path>
```

## shields:
Badges from shields.io. Search for your project at shields.io, copy & paste the results in this format:
```
shields:
  badges:
  - description:
    url:
```

## include:
List of files you want to include into README.md. No parsing is performed, only accepts plain text files.
```
include:
  - < URL or local file >
```