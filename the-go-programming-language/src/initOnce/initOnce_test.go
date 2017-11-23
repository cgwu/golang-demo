package initOnce_test

import (
	"testing"
	"fmt"
	"initOnce"
)

func TestGetString(t *testing.T) {
	fmt.Println(initOnce.GetString(1))
	fmt.Println(initOnce.GetString(2))
	fmt.Println(initOnce.GetString(3))
}

func TestGetString2(t *testing.T) {
	go func() {
		fmt.Println(initOnce.GetString(1))
	}()
	go func() {
		fmt.Println(initOnce.GetString(2))
	}()
	go func() {
		fmt.Println(initOnce.GetString(3))
	}()
}
