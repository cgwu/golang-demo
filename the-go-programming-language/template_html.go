package main

import (
	"html/template"
	"os"
	"log"
	"fmt"
)

func foo(a int, b interface{}) {
	fmt.Println(a, b)
}
func bar() (a, b int){
	return 1,2
}

func main() {
	const templ = `<p>A: {{.A}}</p><p>A: {{.B}}</p>   {{23 -}} < {{- 45}}`
	// Must方法，当error不为nil时，中止执行; .Parse方法返回2个值(*Template, error)
	var t = template.Must(template.New("escape").Parse(templ))
	var data struct {
		A string
		B template.HTML
	}
	data.A = "<b>Hello你好</b>"
	data.B = "<b>Hello你好</b>"
	if err := t.Execute(os.Stdout, data); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("\n--测试返回值作为函数参数--")
	foo(bar())
}
