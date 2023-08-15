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
}
