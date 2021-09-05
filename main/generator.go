package main

import (
	"fmt"
	"html/template"
	"os"
)


var tpl = `package main

import "errors"

type {{.Type}}Stack []{{.Type}}

func New{{.Type}}Stack() {{.Type}}Stack {
	return {{.Type}}Stack{}
}

func(s *{{.Type}}Stack) Add(elem {{.Type}}) {
	*s = append(*s, elem)
	return
}

func(s *{{.Type}}Stack) Pop() ({{.Type}}, error) {
	var elem {{.Type}}
	if len(*s) > 0 {
		elem = (*s)[len(*s)-1]
		*s = (*s)[:len(*s)-1]
		return elem, nil 
	}
	return elem, errors.New("empty stack")
}
`

func main() {
	temp := template.Must(template.New("stack").Parse(tpl))
	for i := 1; i < len(os.Args); i++ {
		fileName := os.Args[i] + "_stack.go"
		file, err := os.Create(fileName)
		if err != nil {
			fmt.Println("create error: ", err)
		}
		tMap := map[string]string{
			"Type":os.Args[i],
		}
		err = temp.Execute(file, tMap)
		if err != nil {
			fmt.Println("execute error : ", err)
		}
		err = file.Close()
		if err != nil {
			fmt.Println("close error : ", err)
		}
	}
}
