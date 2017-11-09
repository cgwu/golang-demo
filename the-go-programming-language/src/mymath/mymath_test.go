package mymath_test

import (
	"mymath"
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	c,ok :=  mymath.Add(1,2)
	if ok == nil {
		fmt.Println(c)
	}
}
