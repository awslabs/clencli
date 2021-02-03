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
