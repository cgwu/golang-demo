package ref_test

import (
	"reflect"
	"fmt"
	"testing"
	"io"
	"os"
	"unsafe"
	"math"
)

func TestAbc(test *testing.T) {
	// Type
	fmt.Println("---type---")
	t := reflect.TypeOf(3)
	fmt.Println(t.String())
	fmt.Println(t)

	var w io.Writer = os.Stdout
	fmt.Println(reflect.TypeOf(w))

	fmt.Printf("%T\n", int64(3))

	// Value
	fmt.Println("---value---")
	v := reflect.ValueOf(3)
	fmt.Println(v)
	fmt.Printf("%v\n", v)
	fmt.Println(v.String())

	vt := v.Type()
	fmt.Println(vt)

	x := v.Interface() // an interface
	i := x.(int)       // an int
	fmt.Printf("x.(int) = %d\n", i)
}


func TestUnsafe(test *testing.T) {
	fmt.Println(unsafe.Sizeof(int(0)))
	fmt.Println(unsafe.Sizeof(float64(0)))
	fmt.Println(unsafe.Sizeof(float32(0)))
	fmt.Println(unsafe.Sizeof(true))

	fmt.Printf("%#016x\n", math.Float64bits(1.0))
}