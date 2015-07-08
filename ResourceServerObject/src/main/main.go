package main

import (
	"net/http"
	//"io/ioutil"
	"strings"
	"os"	// OS specific functions; here to open the file
	"bufio"	// Allows to impliment actual buffering logic
)

func main(){
	/*
	http.Handle("/", handler)
    http.Handle tells the http package to handle all requests to the web root ("/")
    with handler object.
	*/
	http.Handle("/",new(MyHandler))
		
	/*
	ListenAndServe starts an HTTP server with a given address and handler. 
	The handler is usually nil, which means to use DefaultServeMux. 
	DefaultServeMux is Default Server Multiplexer. When a request comes, the
	recives it in its primary thread, and then it will spin up a go routing
	that runs concurrently to the primary thread. Then it delegates the handling
	of the request to the new go routing. This process is called DefaultServerMux.
	(In real world this may not be enough where you have to write your own handler)
	Since this function blocks the main thread of the service
	make sure to do all the other stuff, such as configurations etc., in advance.
	*/
	http.ListenAndServe(":8000", nil)
}

// Create a custom type
type MyHandler struct{
	// Use go's compostion to impliment http.Hanler interface
	// Acutally we don't have to explicit in this because 
	// go identifies it by implimenting the interface (see below)
	http.Handler
}

// Implimentation of the http.Handler interface
// Method ServeHTTP(w http.ResponseWriter, req *http.Request)
func (this *MyHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := "public/" + req.URL.Path
	
	// This completly load the file into memory
	// Not good at all for the server
	//data, err := ioutil.ReadFile(string(path))
	
	// Instead do buffering
	file, err := os.Open(path)
	
	if err == nil {
		bufferReader := bufio.NewReader(file)
		
		// Find the content type of the file and put it in the response header.
		// Otherwise the browser doensn't have enough information about the page content
		var contentType string
		
		if strings.HasSuffix(path, ".css") {
			contentType = "text/css"
		} else if strings.HasSuffix(path, ".html") {
			contentType = "text/html"
		} else if strings.HasSuffix(path, ".js") {
			contentType = "application/javascript"
		} else if strings.HasSuffix(path, ".jpg") {
			contentType = "image/jpg"
		} else if strings.HasSuffix(path, ".png") {
			contentType = "image/png"
		} else if strings.HasSuffix(path, ".mp4") {
			contentType = "video/mp4"
		} else {
			contentType = "text/plain"
		} 
		
		w.Header().Add("Content Type", contentType)
		
		//w.Write([]byte(data))
		// We are writing through the bufferReader
		bufferReader.WriteTo(w)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404 - " + http.StatusText(404)))
	}
}