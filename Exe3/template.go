package main

import (
	"log"
	"net/http"
	"text/template"
)

var s *Story

const tpl = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>{{.Title}}</title>
	</head>
	<body>
		<h1>Intro : {{.Title}}</h1>
		<h2>Story:</h2>
		<ul>
			{{range .Story}}
				<li>{{.}}</li>			
			{{end}}
		</ul>
		{{if .Options}}  
			<h2>Options:</h2>
			<ul>
				{{range .Options}}
					<li>{{.Text}} => <a href=/{{.Arc}}>{{.Arc}}</a></li>			
				{{end}}
			</ul>
		{{else}}
			<a href=/>Try a again !</a>
		{{end}}
	</body>
</html>`

func templatedHandler(w http.ResponseWriter, r *http.Request) {
	tmplt := template.New("hello template")
	tmplt, _ = tmplt.Parse(tpl)
	if s == nil {
		s = getStory()
		log.Println("getStory()")
	}
	url := r.URL.Path
	// log.Println(url)
	data := mapStory(url, s)
	tmplt.Execute(w, data) //merge template ‘t’ with content of ‘p’
}
