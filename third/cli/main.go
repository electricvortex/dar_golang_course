package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

var (
	configFilePath = ""


	flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "config",
			Aliases:     []string{"c"},
			Destination: &configFilePath,
		},
	}

	commands = cli.Commands{}
	command = &cli.Command{
		Name: "printpath",
		Action: PrintConfigFilePath,
	}

)

func main(){
	commands = append(commands, command)
	app := cli.NewApp()
	app.Flags = flags
	app.Commands = commands
	app.Run(os.Args)
}

func PrintConfigFilePath(c *cli.Context) error {
	fmt.Println(c.Args().First())
	fmt.Printf(configFilePath)
	return nil
}