package main

import(
  "fmt"
  "log"
)

type Money struct {
  amount uint
  currency string
}

func NewMoney(amount uint) *Money {
  return &Money{amount,"yen"}
}

/*状態を変更しないクエリメソッドFormatを定義する。*/
func (this *Money) Format() string {
  return fmt.Sprintf("%d %s", this.amount, this.currency)
}
/*"%d:基数10  %s:文字列"*/

func (this *Money) Add(that *Money) {
  this.amount += that.amount
}

func main() {
  money := NewMoney(120)
  log.Printf(money.Format())
  money.Add(NewMoney(180))
  log.Print(money.Format())
}
