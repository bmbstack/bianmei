package main

import (
	"github.com/codegangsta/cli"
	"os"
	"github.com/bmbstack/bianmei/scripts"
)

func main() {
	app := cli.NewApp()
	app.Name = "bianmei"
	app.Usage = "A bianmei application powered by Ripple framework"
	app.Author = "wangmingjob"
	app.Email = "wangmingjob@icloud.com"
	app.Version = "0.0.1"
	app.Commands = scripts.Commands()
	app.Run(os.Args)
}
