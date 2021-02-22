/*
1.构造asset.xml资源文件
2.根据资源文件写出对应的数据结构
3.开始序列化和反序列化编码
4.把序列化写入文件
*/
package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type Servers struct {
	Name       xml.Name `xml:"servers,name"`
	Version    string   `xml:"version,attr"`
	ServerData []Server `xml:"server"`
}

type Server struct {
	ServerName string `xml:"serverName"`
	ServerIP   string `xml:"serverIP"`
}

func main() {
	//readJson()
	writeJson()
}

func readJson() {
	data, err := ioutil.ReadFile("重点/数据格式/asset/asset.xml")
	if err != nil {
		fmt.Printf("反序列化失败，%v\n", err)
		return
	}

	var s Servers
	err = xml.Unmarshal(data, &s)
	if err != nil {
		fmt.Printf("反序列化失败，%v\n", err)
		return
	}
	fmt.Printf("反序列化：%#v", s)
}
func writeJson() {
	var itmes Servers = Servers{
		Name: xml.Name{
			Space: "s",
			Local: "ss",
		},
		Version: "1.0",
	}

	for i := 0; i < 5; i++ {
		itmes.ServerData = append(itmes.ServerData, Server{
			ServerName: "mysql",
			ServerIP:   fmt.Sprintf("192.168.1.%d", i),
		})
	}
	fmt.Printf("xml数据：%#v", itmes)
	xmlstr, err := xml.Marshal(&itmes)
	if err != nil {
		fmt.Printf("序列化失败，错误如下：%#v", err)
		return
	}

	err = ioutil.WriteFile("重点/数据格式/asset/writeXml.xml", xmlstr, 0755)
	if err != nil {
		fmt.Printf("写入失败，错误如下：%#v", err)
		return
	}
	fmt.Println("写入成功")
}
