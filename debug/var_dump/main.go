package main

import (
  "fmt"
)

type Foo struct {
  Bar string
  Baz string
}

func main () {
  foo := &Foo{"bar","baz"}
  var_dump(1,true,nil,"string",foo)
}

func var_dump(v ...interface{}) {
  for _, vv := range(v) {
    fmt.Printf("%#v\n", vv)
  }
}
