/*
bzip2 -k -9 filename
bunzip2 -k -9 filename.bz
sha256sum file1.bz2 file2.bz2
*/

package main
import (
	"io/ioutil"
	"bytes"
	"fmt"
)

func mustNil(e error) {
	if e != nil {
		panic(e)
	}
}

func main(){
	data1, err := ioutil.ReadFile("compressed.txt.bz2")
	mustNil(err)
	data2, err := ioutil.ReadFile("uncompressed.txt.bz2")
	mustNil(err)
	if bytes.Equal(data1, data2) {
		fmt.Println("file content is the same!")
	} else {
		fmt.Println("file content is NOT the same!")
	}
}
