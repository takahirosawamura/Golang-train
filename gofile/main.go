package main

import (
  "fmt"
  "os"
//  "syscall"
)

//func main() {
//   defaultUmask := syscall.Umask(0)
//   os.Mkdir("/tmp/go-dir2", 0777)
//   syscall.Umask(defaultUmask)
//} 

//os.MkdirAllを使うと、親ディレクトリがなければ親ディレクトリを作ってくるので、今回はそれを使ってコードを書く。
//参考ページ　https://qiita.com/suin/items/af8f306dc6b38a293ef://qiita.com/suin/items/af8f306dc6b38a293ef://qiita.com/suin/items/af8f306dc6b38a293ef555
func main() {
  for i := 1; i <= 10; i++ {
    os.MkdirAll(fmt.Sprintf("/tmp/go-dirs/%02d", i), os.ModePerm)
  }
}
