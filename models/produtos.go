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
	selectDeTodosProdutos, err := db.Query("select * from produtos order by id asc")
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
		p.Id = id
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
func DeletaProduto(id string) {
	db := db.ConectaBancoDeDados()
	deletarProduto, err := db.Prepare("delete from produtos where id=$1")
	if err != nil {
		panic(err.Error())
	}
	deletarProduto.Exec(id)
	defer db.Close()
}
func EditaProduto(id string) Produto {
	db := db.ConectaBancoDeDados()
	produtoDoBanco, err := db.Query("select * from produtos where id=$1", id)
	if err != nil {
		panic(err.Error())
	}
	produtoParaAtualizar := Produto{}
	for produtoDoBanco.Next() {
		var id int
		var nome string
		var descricao string
		var preco float64
		var quantidade int
		err = produtoDoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
<<<<<<< HEAD
		produtoParaAtualizar.Id = id
=======
>>>>>>> 8c9ba8837fbfff19f07a9c2442fc1d07564aa00b
		produtoParaAtualizar.Nome = nome
		produtoParaAtualizar.Descricao = descricao
		produtoParaAtualizar.Preco = preco
		produtoParaAtualizar.Quantidade = quantidade

	}
	defer db.Close()
	return produtoParaAtualizar
}
<<<<<<< HEAD
func AtualizaProduto(id int, nome string, descricao string, preco float64, quantidade int) {
	db := db.ConectaBancoDeDados()
	AtualizaProduto, err := db.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}
	AtualizaProduto.Exec(nome, descricao, preco, quantidade, id)
	defer db.Close()
}
=======
>>>>>>> 8c9ba8837fbfff19f07a9c2442fc1d07564aa00b
