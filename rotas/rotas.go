package rotas

import (
	"models/aplicacao/controle"
	"net/http"
)

func CarregaRotas() {
	http.HandleFunc("/", controle.Index)
	http.HandleFunc("/new", controle.New)
	http.HandleFunc("/insert", controle.Insert)
	http.HandleFunc("/delete", controle.Delete)
	http.HandleFunc("/edit", controle.Edit)
<<<<<<< HEAD
	http.HandleFunc("/update", controle.Update)
=======
>>>>>>> 8c9ba8837fbfff19f07a9c2442fc1d07564aa00b
}
