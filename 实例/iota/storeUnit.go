package main

type ByteSize float64

const (
	_           = iota // 通过_来跳过iota 的初始值
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	println("KB:", KB)
	println("MB:", MB)
	println("GB:", GB)
	println("TB:", TB)
	println("PB:", PB)
	println("EB:", EB)
	println("ZB:", ZB)
	println("YB:", YB)
}
