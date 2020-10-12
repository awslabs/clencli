# Configuration

## unsplash

1. Register as a developer at [unsplash][https://unsplash.com/documentation#creating-a-developer-account]
2. Create a new application as `demo`

Copy `Access key` and `Secret key` into your `clencli` config, usually at `$HOME/.clencli.yaml`:

```
unsplash:
  access_key: "XXX"
  secret_key: "XXX"
```

Now, commands can use `unsplash` to download photos, such as `clencli unsplash` or `clencli render` for example.