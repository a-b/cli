# Cloud Foundry CLI Plugin System

This package contains the core interfaces and types for the Cloud Foundry CLI plugin system.

## Overview

The CF CLI plugin system allows developers to extend the functionality of the CF CLI by creating plugins that can be installed and run alongside the core CLI commands.

## Plugin Interface

All plugins must implement the `Plugin` interface:

```go
type Plugin interface {
	Run(cliConnection CliConnection, args []string)
	GetMetadata() PluginMetadata
}
```

- `Run`: The entry point for your plugin. This is called when a user executes your plugin command.
- `GetMetadata`: Returns metadata about your plugin, including commands, help text, and version information.

## CLI Connection

Plugins interact with the CF CLI through the `CliConnection` interface, which provides methods to:
- Execute CF CLI commands
- Get information about the current context (org, space, user)
- Access API endpoints
- Query CF resources (apps, services, orgs, spaces)

## Plugin Development

### Getting Started

1. Create a new Go package
2. Implement the `Plugin` interface
3. Build your plugin as an executable
4. Install your plugin with `cf install-plugin PATH_TO_PLUGIN`

### Example Plugin

```go
package main

import (
	"fmt"
	"code.cloudfoundry.org/cli/plugin"
)

type MyPlugin struct{}

func (c *MyPlugin) Run(cliConnection plugin.CliConnection, args []string) {
	// Your plugin logic here
	fmt.Println("My plugin is running!")
}

func (c *MyPlugin) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "MyPlugin",
		Version: plugin.VersionType{
			Major: 1,
			Minor: 0,
			Build: 0,
		},
		Commands: []plugin.Command{
			{
				Name:     "my-command",
				HelpText: "This is my custom command",
			},
		},
	}
}

// Main function must be implemented by the plugin
func main() {
	plugin.Start(new(MyPlugin))
}
```

### Plugin Versioning

Plugins include version information for both the plugin itself and the CLI library version it was built against.
This allows the CLI to check compatibility before running a plugin.

## Installation and Distribution

Plugins can be installed from:
- Local files: `cf install-plugin PATH_TO_PLUGIN`
- Plugin repositories: `cf install-plugin PLUGIN_NAME -r REPO_NAME`

To share your plugin, you can:
1. Publish it to a plugin repository
2. Share the binary directly with users
3. Open source your code for others to build

## Best Practices

1. Keep your plugin focused on a specific task
2. Provide clear help text and usage examples
3. Handle errors gracefully
4. Follow CF CLI UX patterns for consistency
5. Test your plugin against multiple CF API versions
