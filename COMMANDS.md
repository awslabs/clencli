# Command reference

This page documents every clencli command, its flags, and examples. For an overview and installation instructions, see [README.md](README.md).

## Global flags

These flags apply to every command.

| Flag | Default | Description |
| --- | --- | --- |
| `-p`, `--profile` | `default` | Use a specific profile from your credentials and configurations files. |
| `-v`, `--verbosity` | `error` | Log level. Valid values: `panic`, `fatal`, `error`, `warn`, `info`, `debug`, `trace`. |
| `--log` | `true` | Enable or disable logging to a file. When disabled, log output goes to the default output. |
| `--log-file-path` | `clencli/log.json` | Log file path. Requires `--log=true`. |
| `-h`, `--help` | | Show help for the command. |

## configure

```text
clencli configure [delete] [flags]
```

Creates or updates clencli credentials and configurations interactively. Credentials and configurations are stored as YAML files in your user configuration directory, under a `clencli` subdirectory.

On first run, `configure` creates the configuration directory and prompts you to set up credentials and configurations. On later runs, it prompts you to update the existing files. Pass the `delete` argument to remove a profile.

Use `--profile` to target a profile other than `default`.

| Flag | Default | Description |
| --- | --- | --- |
| `-p`, `--profile` | `default` | Profile to create, update, or delete. |

Examples:

```bash
# Create or update the "work" profile
clencli configure --profile work

# Delete the "work" profile
clencli configure delete --profile work
```

## gitignore

```text
clencli gitignore [list] [flags]
```

Downloads a `.gitignore` file from the gitignore.io API based on the types you provide, and writes it to `.gitignore` in the current directory. Pass the `list` argument to print the valid types instead.

You must provide either the `list` argument or the `--input` flag.

| Flag | Default | Description |
| --- | --- | --- |
| `-i`, `--input` | | One or more gitignore types. If multiple, comma-separated. |

Examples:

```bash
# List the valid types
clencli gitignore list

# Download a .gitignore for a Terraform project
clencli gitignore --input terraform

# Combine multiple types
clencli gitignore --input terraform,vscode
```

## init

```text
clencli init [project] [--project-name <value>] [--project-type <value>] [flags]
```

Initializes a project with a standardized directory structure and template files. The command creates a directory named after the project, scaffolds it for the chosen type, and reports success.

The `--project-name` flag is required. The `--project-type` flag selects the layout.

| Flag | Default | Description |
| --- | --- | --- |
| `--project-name` | | Project name. Required. Used as the new directory name. |
| `--project-type` | `basic` | Project layout. Valid values: `basic`, `cloud`, `cloudformation`, `terraform`. |

Project types:

- `basic` — a `clencli/` directory with `readme.yaml` and `readme.tmpl`, plus a `.gitignore`.
- `cloud` — the `basic` layout plus a high-level design template (`hld.yaml`, `hld.tmpl`).
- `cloudformation` — the `cloud` layout plus `environments/dev`, `environments/prod`, `skeleton.yaml`, and `skeleton.json`.
- `terraform` — the `cloud` layout plus `main.tf`, `variables.tf`, `outputs.tf`, `Makefile`, `LICENSE`, and `environments/dev.tf` and `environments/prod.tf`.

If a `configurations.yaml` file with a customized initialization exists for the active profile, clencli applies those additional files after scaffolding the chosen type.

Examples:

```bash
# Basic project
clencli init project --project-name foo

# Cloud project
clencli init project --project-name foo --project-type cloud

# AWS CloudFormation project
clencli init project --project-name foo --project-type cloudformation

# Terraform project
clencli init project --project-name foo --project-type terraform
```

## render

```text
clencli render template [--name <value>] [flags]
```

Renders a Go template into a Markdown file using a YAML data file as input. The command reads `clencli/<name>.yaml` and `clencli/<name>.tmpl` from the current directory and writes the result to `<NAME>.md` (uppercased). Both files must exist, or the command reports an error.

Before rendering, clencli trims trailing whitespace from the data file and updates the README logo. If an `unsplash.yaml` file is present, or if the active profile has an Unsplash credential and random-photo configuration, clencli downloads a photo and sets it as the logo.

| Flag | Default | Description |
| --- | --- | --- |
| `-n`, `--name` | `readme` | Name of the template and data file under `clencli/`, without extension. |

Examples:

```bash
# Render clencli/readme.tmpl to README.md
clencli render template

# Render clencli/hld.tmpl to HLD.md
clencli render template --name hld
```

## unsplash

```text
clencli unsplash [flags]
```

Downloads a photo from Unsplash. With no `--id`, it retrieves a random photo filtered by the flags below. With `--id`, it retrieves that specific photo. This command requires an `unsplash` credential in the active profile; configure one with `clencli configure`.

| Flag | Default | Description |
| --- | --- | --- |
| `--id` | | Photo ID. Leave empty to download a random photo instead. |
| `--collections` | | Public collection ID(s) to filter selection. If multiple, comma-separated. |
| `--featured` | `false` | Limit selection to featured photos. |
| `--filter` | `low` | Limit results by content safety. Valid values: `low`, `high`. |
| `--orientation` | `landscape` | Filter by orientation. Valid values: `landscape`, `portrait`, `squarish`. |
| `--query` | `mountains` | Limit selection to photos matching a search term. |
| `--size` | `all` | Photo size. Valid values: `all`, `thumb`, `small`, `regular`, `full`, `raw`. |
| `--username` | | Limit selection to a single user. |

Examples:

```bash
# Download a random landscape photo matching "mountains"
clencli unsplash

# Download a random portrait photo matching a search term
clencli unsplash --query desert --orientation portrait

# Download a specific photo by ID
clencli unsplash --id Dwu85P9SOIk
```

## version

```text
clencli version
```

Prints the clencli version, Go version, operating system, and architecture.

Example:

```bash
clencli version
```

```text
clencli v0.3.6 go1.16 linux amd64
```
