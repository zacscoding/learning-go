package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
	"strings"
	"time"
)

func main() {
	var app = cli.NewApp()
	app.Name = "sample"
	app.Usage = "Usage.."
	app.Author = "zacscoding"
	app.Version = "1.0.0"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "network, n",
			Usage: "name of the network",
			Value: "Default network",
		},
		cli.IntFlag{
			Name:  "repeat, r",
			Usage: "repeat..",
			Value: 5,
		},
	}
	app.Commands = []cli.Command{
		{
			Name:    "command1",
			Aliases: []string{"c1"},
			Usage:   "print network with lowercase",
			Action: func(c *cli.Context) {
				network := strings.ToLower(c.String("network"))
				repeat := c.Int("repeat")
				fmt.Println("command1 action is running.. network :", network, "and repeat :", repeat)

				for i := 0; i < repeat; i++ {
					fmt.Println(i, "->", network)
					time.Sleep(time.Second)
				}
			},
			Flags: app.Flags,
		},
		{
			Name:    "command2",
			Aliases: []string{"c2"},
			Usage:   "print network with uppercase",
			Action: func(c *cli.Context) {
				network := strings.ToUpper(c.String("network"))
				repeat := c.Int("repeat")

				fmt.Println("command2 action is running.. network :", network, "and repeat :", repeat)

				for i := 0; i < repeat; i++ {
					fmt.Println(i, "->", network)
					time.Sleep(time.Second)
				}
			},
			Flags: app.Flags,
		},
	}

	app.Before = beforeAction
	app.After = afterAction

	err := app.Run(os.Args)

	if err != nil {
		fmt.Println(err)
	}
}

func beforeAction(c *cli.Context) error {
	fmt.Println("beforeAction is called()..")
	return nil
}

func afterAction(c *cli.Context) error {
	fmt.Println("afterAction is called()..")
	return nil
}
