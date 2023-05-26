package main

import (
	//"fmt"
	//"html/template"
	"net/http"
)

/*
func page(w http.ResponseWriter, r *http.Request){
	pagename:= "./src/login.html"
	/*t, err := template.ParseFiles(pagename)
	if err!=nil{
		fmt.Println(err)
	}
	t.ExecuteTemplate(w, pagename, nil)
	http.ServeFile(w, r, pagename)
}

func handler(w http.ResponseWriter, r *http.Request){
	switch r.URL.Path{
	case "/":
		page(w,r)
	}
}


func pages(w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("src/index.html")
	if err!=nil{
		fmt.Println(err)
	}
	t.Execute(w, nil)
}

func handlers(w http.ResponseWriter, r *http.Request){
	switch r.URL.Path{
	case "/":
		pages(w,r)
	}
}*/

func handle(){
	
	http.Handle("/", http.FileServer(http.Dir("./src/")))
}