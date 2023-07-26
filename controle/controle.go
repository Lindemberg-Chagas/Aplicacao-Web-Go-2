package controle

import (
	"html/template"
	"models/aplicacao/models"
	"net/http"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscarTodosOsProdutos()
	temp.ExecuteTemplate(w, "Index", todosOsProdutos)

}
