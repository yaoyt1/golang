package main

import "os"

func main() {
	var b [16]byte
	os.Stdin.Read(b[:])
	os.Stdout.WriteString(string(b[:]))
}
