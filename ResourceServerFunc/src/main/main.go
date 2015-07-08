package main

import (
	"net/http"
)

func main(){
	/*
	http.HandleFunc("/", handler)
    http.HandleFunc tells the http package to handle all requests to the web root ("/")
    with handler. Let's put an anonumous function, with the correct form, as the handler.
	*/
	http.HandleFunc("/",func(w http.ResponseWriter, req *http.Request) {
			// Since write require a byte slice we need to explicitly cast string to []byte
			w.Write([]byte("Hello world"))
		})
		
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