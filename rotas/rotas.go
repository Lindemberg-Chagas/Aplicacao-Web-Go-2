package rotas

import (
	"models/aplicacao/controle"
	"net/http"
)

func CarregaRotas() {
	http.HandleFunc("/", controle.Index)
}
