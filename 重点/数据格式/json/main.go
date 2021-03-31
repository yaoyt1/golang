/*
1.构造asset.json资源文件
2.根据资源文件写出对应的数据结构
3.开始序列化和反序列化编码
4.把序列化写入文件
*/
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

//这里结构的tag决定json 文件中的key
type jsondatas struct {
	Id      int         `json:"id"`
	Version string      `josn:"version"`
	Data    []*jsondata `json:"data"`
}

type jsondata struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Class string `json:"class"`
}

func main() {
	//readJson()
	writeJson()
}

func readJson() {
	data, err := ioutil.ReadFile("重点/数据格式/asset/asset.json")
	if err != nil {
		fmt.Printf("反序列化失败，%v\n", err)
		return
	}

	var j jsondatas
	err = json.Unmarshal(data, &j)
	if err != nil {
		fmt.Printf("反序列化失败，%v\n", err)
		return
	}
	fmt.Printf("反序列化：%#v", j)
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

	jsonstr, err := json.Marshal(itmes)
	if err != nil {
		fmt.Printf("序列化失败，错误如下：%#v", err)
		return
	}

	err = ioutil.WriteFile("重点/数据格式/asset/writeJson.json", jsonstr, 0755)
	if err != nil {
		fmt.Printf("写入失败，错误如下：%#v", err)
		return
	}
	fmt.Println("写入成功")
}
