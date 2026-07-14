package main

import (
	//"api-tributos/internal/services/calculadora"
	dados "api-tributos/internal/services/dadosabertos"
	"fmt"
)

func main() {
	// Usamos o pacote 'fmt' para imprimir mensagens no terminal
	fmt.Println("🚀 Iniciando o Sistema de Consulta de Tributos...")

	//calculadora.XmlValidate()
	//calculadora.XmlGenerate()
	//dados.GetUf()
	//dados.GetVersao()
	//dados.GetMunicipios()
	//dados.GetTransferenciasIbs()
	//dados.GetTransferenciasCbs()
	dados.GetImpostoSeletivo()

	fmt.Println("✅ Execução finalizada com sucesso!")
}
