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
)

func main(){
	app := cli.NewApp()
	app.Flags = flags
	app.Commands = cli.Commands{
		&cli.Command{
			Name: "printpath",
			Action: PrintConfigFilePath,
		},
	}
	app.Run(os.Args)
}

func PrintConfigFilePath(c *cli.Context) error {
	fmt.Printf(configFilePath)
	return nil
}