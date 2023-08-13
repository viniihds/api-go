package schemas

import ("gorm.io/gorm")

type Vaga struct{
	gorm.Model
	Funcao string
	Empresa string
	Localizacao string
	Remoto bool
	Link string
	Salario int64
}