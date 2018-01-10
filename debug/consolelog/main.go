package main

import (
  "log"
)

type Foo struct {
  Bar string
  Baz string
}

func main() {
    foo := &Foo{"bar", "baz"}
    log.Printf("%v", foo)
    log.Printf("%+v", foo)
    log.Printf("%#v", foo)
    log.Printf("%T", foo)
}
