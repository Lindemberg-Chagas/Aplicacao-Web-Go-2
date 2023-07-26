package main

import (
	"models/aplicacao/rotas"
	"net/http"
)

func main() {

	rotas.CarregaRotas()
	http.ListenAndServe(":8000", nil)

}
