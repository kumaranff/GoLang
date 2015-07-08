package main

import (
	"net/http"
	"text/template"
)

func main() {
	// Handle the request using http.HandleFunc
	http.HandleFunc("/", func (w http.ResponseWriter, req *http.Request) {
			w.Header().Add("Content Type", "text/html")
			/*
			New function allows to create a new template
			Since go struct types doesn't have consturctors this
			allows to initialize the struct.
			Parse parses a string into a template.
			*/
			tmpl, err := template.New("StaticTemplate").Parse(doc)
			
			if err == nil {
				// Execute applies a parsed template to the specified data object,
				// and writes the output to w.
				tmpl.Execute(w,nil)
			}
		})
	
	http.ListenAndServe(":8000", nil)
}

const doc = `
<!DOCTYPE html>
<html>
	<head>
		<title>Blue Wave Template</title>
	</head>
	<body>
		<h1>Hello Template</h1>
	</body>
</html>`