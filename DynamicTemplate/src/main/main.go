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
				/*
				func (t *Template) Execute(wr io.Writer, data interface{}) (err error)
				Execute applies a parsed template to the specified data object, and writes 
				the output to w. If an error occurs executing the template or writing its 
				output, execution stops, but partial results may already have been written 
				to the output writer. To inject the data into template pipleline, put it 
				as the second argument.
				*/
				tmpl.Execute(w, req.URL.Path)
			}
		})
	
	http.ListenAndServe(":8000", nil)
}

/*
In the context of template {{.}} is called pipeline
where the contents can be injected by the template.
*/
const doc = `
<!DOCTYPE html>
<html>
	<head>
		<title>Blue Wave Template</title>
	</head>
	<body>
		<h1>Hello {{.}}</h1>
	</body>
</html>`