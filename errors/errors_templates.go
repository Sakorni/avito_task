package errors

import "fmt"

var NoSuchUser = fmt.Errorf("there is no user with such id")
var NotEnoughMoney = fmt.Errorf("not enough money on the balance to make an operation")
