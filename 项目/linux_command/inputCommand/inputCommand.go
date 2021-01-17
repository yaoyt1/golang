package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func terminalInput() {
	var operatingTree string

	cmd := cli.NewApp()
	cmd.Name = "GoLinux"
	cmd.Usage = "使用go来实现linux基本命令"
	cmd.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "tree",
			Aliases:     []string{"t"},
			Usage:       "展开目录下所有的文件夹和文件",
			Destination: &operatingTree,
		},
	}

	cmd.Action = func(c *cli.Context) error {
		if c.NArg() == 0 {
			fmt.Println("暂无参数")
			fmt.Println(c.Args())
			return nil
		}

		fmt.Println(c.Args())
		return nil
	}

	cmd.Run(os.Args)
}

func main() {
	terminalInput()
}
