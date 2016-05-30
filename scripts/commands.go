package scripts

import (
	"github.com/codegangsta/cli"
	"io"
	"os"
	"os/exec"
	"strings"
	"sync"
	"github.com/bmbstack/bianmei/logger"
)


// Commands
func Commands() []cli.Command {
	return []cli.Command{
		//Init application
		{
			Name:  "init",
			Usage: "Init application(go packages / DB migration /Frontend compiler)",
			Action: func(c *cli.Context) {
				commands := GetInitCommands()
				RunScript(commands)
				logger.Logger.Info("Init application done")
			},
		},
		//Run server
		{
			Name:    "server",
			Aliases: []string{"s"},
			Usage:   "Run server",
			Action: func(c *cli.Context) {
				db := ""
				if len(c.Args()) > 0 {
					db = c.Args()[0]
				}
				commands := GetServerCommands(db)
				RunScript(commands)
				RunServer()
				logger.Logger.Info("Run server done")
			},
		},
	}
}

// Run Script
func RunScript(commands []string) {
	entireScript := strings.NewReader(strings.Join(commands, "\n"))
	bash := exec.Command(Bash)
	stdin, _ := bash.StdinPipe()
	stdout, _ := bash.StdoutPipe()
	stderr, _ := bash.StderrPipe()

	wait := sync.WaitGroup{}
	wait.Add(3)
	go func() {
		io.Copy(stdin, entireScript)
		stdin.Close()
		wait.Done()

	}()
	go func() {
		io.Copy(os.Stdout, stdout)
		wait.Done()

	}()
	go func() {
		io.Copy(os.Stderr, stderr)
		wait.Done()

	}()

	bash.Start()
	wait.Wait()
	bash.Wait()
}
