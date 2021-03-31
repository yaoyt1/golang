//msgpack 官网https://msgpack.org/
package main

import (
	"fmt"
	"github.com/vmihailenco/msgpack"
	"io/ioutil"
)

type jsondatas struct {
	Id      int         `msgpack:"id"`
	Version string      `msgpack:"version"`
	Data    []*jsondata `msgpack:"data"`
}

type jsondata struct {
	Id    int    `msgpack:"id"`
	Name  string `msgpack:"name"`
	Age   int    `msgpack:"age"`
	Class string `msgpack:"class"`
}

func main() {
	writeJson()
	readJson()
}

func readJson() {
	data, err := ioutil.ReadFile("重点/数据格式/asset/asset.dat")
	if err != nil {
		fmt.Printf("读取文件失败，%v\n", err)
		return
	}

	var j jsondatas
	err = msgpack.Unmarshal(data, &j)
	if err != nil {
		fmt.Printf("反序列化失败，%v\n", err)
		return
	}
	fmt.Printf("反序列化成功：%#v", j)
}
func writeJson() {
	var itmes = jsondatas{
		Id:      1,
		Version: "1.0",
	}

	for i := 0; i < 5; i++ {
		itmes.Data = append(itmes.Data, &jsondata{
			Id:    i,
			Name:  fmt.Sprintf("yyt%d", i),
			Age:   18,
			Class: fmt.Sprintf("高一（%d）班", i),
		})
	}

	bytes, err := msgpack.Marshal(&itmes)
	if err != nil {
		fmt.Printf("序列化失败，错误如下：%#v", err)
		return
	}

	err = ioutil.WriteFile("重点/数据格式/asset/writeMsgPack.dat", bytes, 0755)
	if err != nil {
		fmt.Printf("写入失败，错误如下：%#v", err)
		return
	}
	fmt.Println("写入成功")
}
