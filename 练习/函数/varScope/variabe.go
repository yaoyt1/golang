package varScope

var A = 1
var a = 11

func Add() int {
	return A + a
}

func sub() int {
	return A - a
}
