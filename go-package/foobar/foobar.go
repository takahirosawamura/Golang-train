package foobar

import "log"

/*foobar.goに二つの関数を移植。関数は大文字から始めるようにする。大文字から始まる関数は外部に公開される。*/
func Foo() {
  log.Print("Hello world from foo!")
}

func Bar() {
  log.Print("Hello world from bar!")
}
