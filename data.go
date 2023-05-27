package main

import (
	"encoding/json"
	"fmt"
	//"io/ioutil"
	"net/http"
)

type User struct{
	//struct de usuario
	ID int `json:"ID"`
	Username string `json:"Username"`
	Password string `json:"Password"`
}

var Users []User = []User{
	//slice de struct de usuários
	{
		ID: 0,
		Username: "rebecamadi",
		Password: "senha123",
	},
	{
		ID: 1,
		Username: "davimadi",
		Password: "senha123",
	},
}

//Adiciona um novo cadastro
func SignUpUser(w http.ResponseWriter, r *http.Request){
	//Especifica o tipo de conteúdo no cabeçalho
	w.Header().Set("Content-Type", "application/json") 

	r.ParseForm()

	//Converte o conteudo do json para Go
	var u User

	//Atualiza os dados
	//Solução apenas para o exercicio
	u.ID = len(Users)
	u.Username = r.PostForm.Get("username")
	u.Password = r.PostForm.Get("password")
	Users = append(Users, u)

	//Cria umnovo encoder para a resposta
	encoder := json.NewEncoder(w)
	err := encoder.Encode(u)

	if err!=nil {
		//tratamento de erros
		fmt.Println(err)
	}else{
		//Status
		fmt.Println(http.StatusCreated, http.StatusText(http.StatusCreated))
	}
	
}

func ListUsers(w http.ResponseWriter, r *http.Request){
	//Lista os usuários
	
	//Especifica o tipo de conteudo na header
	w.Header().Set("Content-Type", "application/json")

	//Cria um encoder para listar os usuários
	encoder := json.NewEncoder(w)
	err := encoder.Encode(Users)
	if err!=nil {
		//tratamento de erros
		fmt.Println(err)
	}else{
		//Status
		fmt.Println(http.StatusOK, http.StatusText(http.StatusOK))
	}
}
