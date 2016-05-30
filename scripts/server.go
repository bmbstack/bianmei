package scripts

import (
	_ "github.com/bmbstack/bianmei/controllers"
	_ "github.com/bmbstack/bianmei/models"
	"github.com/bmbstack/bianmei/logger"
	"github.com/bmbstack/ripple"
)

// Server commands
func GetServerCommands(db string) []string {
	commands := make([]string, 2)
	switch db {
	case "mysql":
		commands = append(commands, "/usr/local/bin/mysqld_safe &")
		commands = append(commands, "sleep 10s")
	}
	return commands
}

// Run server
func RunServer() {
	logger.Logger.Info("Run server ....")

	ripple.Run()
}
