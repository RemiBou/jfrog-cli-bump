package main

import (
	"github.com/jfrog/jfrog-cli-core/v2/plugins"
	"github.com/jfrog/jfrog-cli-core/v2/plugins/components"
	"github.com/jfrog/jfrog-cli-plugin-template/commands/configs/add_dep"
	"github.com/jfrog/jfrog-cli-plugin-template/commands/configs/set_vcs"
)

func main() {
	plugins.PluginMain(getApp())
}

func getApp() components.App {
	app := components.App{}
	app.Name = "bump"
	app.Description = "Bump them all."
	app.Version = "v0.1.0"
	app.Commands = getCommands()
	return app
}

func getCommands() []components.Command {
	return []components.Command{
		set_vcs.GetSetVcsCommand(),
		add_dep.GetAddDepCommand(),
	}
}
