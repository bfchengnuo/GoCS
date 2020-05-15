package main

import (
	"GoCS/study/basic/entity"
	"html/template"
	"log"
	"os"
)

// fmt 使用模版来格式化输出
func main() {
	// 可以使用逻辑控制 if-else
	// 可以使用循环：{{range .users}}  {{end}}
	// 也可以使用 HTML 作为模板
	temple := `结构体展示：
name: {{.Name}}
age:  {{.Age}}
desc: {{.Desc | printf "%.5s"}}
other: {{.Other | add}}
END`

	// 使用 Must 包裹便于处理 err
	must := template.Must(
		// 创建模板
		template.New("tName").
			// 将自定义方法注册到模板
			Funcs(template.FuncMap{"add": otherFunc}).
			Parse(temple))

	p := entity.Person{Name: "mps", Age: 12, Desc: "desc........", Other: "OH"}
	if err := must.Execute(os.Stdout, p); err != nil {
		log.Fatal(err)
	}
}

func otherFunc(s string) string {
	return "处理后..." + s
}

// 使用 HTML 模板时，处理特殊字符
type data struct {
	A string
	B template.HTML
}
