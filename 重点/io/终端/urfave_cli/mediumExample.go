package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	var name string
	var boy bool
	var age int

	app := cli.NewApp()
	app.Name = "测试CLi框架"
	app.Usage = "命令行参数"

	app.Flags = []cli.Flag{
		&cli.StringFlag{ // 注意这里一定要是指针类型
			Name:        "studentName",
			Aliases:     []string{"s"}, // 命令行缩写
			Value:       "pete",
			Usage:       "学生姓名",
			Destination: &name,
		},
		&cli.BoolFlag{
			Name:        "boy",
			Value:       true,
			Usage:       "学生性别",
			Destination: &boy,
		},
		&cli.IntFlag{
			Name:        "age",
			Value:       18,
			Usage:       "学生年龄",
			Destination: &age,
		},
	}

	app.Action = func(c *cli.Context) error {
		if c.NArg() > 0 {
			fmt.Printf("%v\n", c.Args())
		}
		fmt.Printf("姓名：%s 是否是男姓：%v 年龄：%d", name, boy, age)
		return nil
	}

	app.Run(os.Args)
}
