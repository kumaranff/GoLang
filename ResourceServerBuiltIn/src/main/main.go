package main

import (
	"net/http"
)

func main(){
	/*
	Use built in file server of go
	*/
	http.ListenAndServe(":8000", http.FileServer(http.Dir("public")))
}