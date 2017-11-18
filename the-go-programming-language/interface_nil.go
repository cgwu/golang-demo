/*
P204
7.5.1 Caveat: An Interface Containing a Nil Pointer Is Non-Nil
A nil interface value, which contains no value at all, is not the same as an interface value
containing a pointer that happens to be nil. This subtle distinction creates a trap into
which every Go programmer has stumbled.
*/

package main

import (
	"bytes"
	"io"
	"fmt"
)

const debug = false

func main() {
	//var buf *bytes.Buffer
	var buf io.Writer
	if debug {
		buf = new(bytes.Buffer)
	}

	if buf != nil {
		fmt.Println("main buf is NOT nil")
	} else {
		fmt.Println("main buf is nil")
	}
	FnLog(buf)	// NOTE: subtly incorrect
}

func FnLog(out io.Writer) {
	fmt.Printf("%T\n", out)
	if out != nil {
		fmt.Println("!= nil")
		out.Write([]byte("done!\n"))
	}else {
		fmt.Println("== nil")
	}
}

/*
debug false output:
<nil>
== nil
main buf is nil

debug true output:
*bytes.Buffer
!= nil
main buf is NOT nil
*/
