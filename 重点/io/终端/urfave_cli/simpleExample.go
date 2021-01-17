package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "name"
	app.Usage = "学生姓名"
	app.Action = func(c *cli.Context) error {
		fmt.Println("执行")
		return nil
	}

	app.Run(os.Args)
}
