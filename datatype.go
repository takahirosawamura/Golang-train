//基本的なデータ型
/*
string	"hello"
int	S3
float64	10.2
bool	false/true
nil

var s string //**
var a int //0
var f bool //false
*/

package main

import "fmt"

func main(){
	a:=10
	b:=12.3
	c:="hoge"
	var d bool
	fmt.Printf("a: %v, b:%f, c:%s, d:%t\n",a,b,c,d) 
} 
