package main

import(
	"fmt"
	"net/http"
)

func main(){
	handle()
	//http.Handle("/", http.FileServer(http.Dir("./src/")))
	err:= http.ListenAndServe(":7070", nil)
	if err!=nil {
		fmt.Println("Error at running the server:", err)
	}
}