package service

import (
	"github.com/johncgriffin/overflow"
)

func Add(x int64, y int64)(int64, bool) {
	return overflow.Add64(x,y)
}

func Multiply(x int64, y int64) (int64, bool) {
	return overflow.Mul64(x,y)
}
