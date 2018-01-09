package main

import(
  "fmt"
  "log"
)

type Money struct{
  amount uint
  currency string
}

/*状態を変更しないクエリメソッドFormatを定義する。*/
func (this *Money) Format() string {
  return fmt.Sprintf("%d %s", this.amount, this.currency)
}
/*"%d:基数10  %s:文字列"*/

func (this *Money) Add(that *Money){
  this.amount += that.amount
}

func main() {
  money := &Money{120, "yen"}
  log.Printf("%#v",money)
  money.Add(&Money{180,"yen"})
  log.Print(money.Format())
}
