package main

import "fmt"

type Pessoa struct{
	Name  string
	CPF   string
	Email string 
}

func (p *Pessoa)Andar(){
	fmt.Println(p.Name, "Andando...")	 
}
type Funcionario struct{
	Pessoa
	Departamento  string
	Salario   string
	Turno string 
}

func main(){
	pessoa1 := Pessoa{
		Name: "F",
		CPF: "568",
		Email: "adf@asdf.com",
	}

	fmt.Println("Essa pessoa veio aqui:",pessoa1)
	pessoa1.Andar()
}