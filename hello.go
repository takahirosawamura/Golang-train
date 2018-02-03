//変数
//変数名：1文字目に注意
package main

import "fmt"//format

func main() {
	//var msg string
	//msg = "hello world"
	//var msg = "hello world"
	msg := "hello world"//宣言と値の代入を同時に行う。
	fmt.Println("hello world")

	//var a,b int
	//a,b = 10,15
	a,b := 10,15 //宣言と値の紐付け

	var (
		c int
		d string
	)
	c = 20
	d = "hoge"
}




