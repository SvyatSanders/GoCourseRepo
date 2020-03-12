package main

import (
	"html/template"
	"log"
	"os"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {
	tmpl, err := template.ParseFiles("template.html")
	if err != nil {
		log.Fatal(err)
	}

	data := TodoPageData{
		PageTitle: "TODO list",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}
	tmpl.Execute(os.Stdout, data)

	// s1 := `"Fran & Freddie's Diner" <tasty@example.com>`
	// s2 := html.EscapeString(s1)
	// s3 := html.UnescapeString(s2)
	// fmt.Println(s2) // &#34;Fran &amp; Freddie&#39;s Diner&#34; &lt;tasty@example.com&gt;
	// fmt.Println(s3) // "Fran & Freddie's Diner" <tasty@example.com>

	// const tmpl = "{{.Greeting}} {{.Name}}"
	// data := struct {
	// 	Greeting string
	// 	Name     string
	// }{"Hello", "Joe"}
	// t := template.Must(template.New("").Parse(tmpl))
	// err := t.Execute(os.Stdout, data)
	// if err != nil {
	// 	log.Fatal(err)
	// }

}
