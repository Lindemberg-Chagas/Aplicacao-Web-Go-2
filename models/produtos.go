package models

import (
	"models/aplicacao/db"
)

type Produto struct {
	Id              int
	Nome, Descricao string
	Preco           float64
	Quantidade      int
}

func BuscarTodosOsProdutos() []Produto {

	db := db.ConectaBancoDeDados()
	selectDeTodosProdutos, err := db.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}
	p := Produto{}
	produtos := []Produto{}
	for selectDeTodosProdutos.Next() {
		var id int
		var nome, descricao string
		var preco float64
		var quantidade int
		err = selectDeTodosProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())

		}
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	defer db.Close()
	return produtos
}
func CriaNovoProduto(nome string, descricao string, preco float64, quantidade int) {
	db := db.ConectaBancoDeDados()

	insereNoBanco, err := db.Prepare("insert into produtos (nome, descricao, preco, quantidade) values ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}
	insereNoBanco.Exec(nome, descricao, preco, quantidade)
	defer db.Close()

}
