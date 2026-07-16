package main

import (
	//"api-tributos/internal/services/calculadora"
	//dados "api-tributos/internal/services/dadosabertos"
	pedagio "api-tributos/internal/services/pedagio"
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
	//dados.GetImpostoSeletivo()
	//dados.GetCbsIbs()
	//dados.GetRedutores()
	//dados.GetNcm()
	//dados.GetNbs()
	//dados.GetNbsLista()
	//dados.GetNbsAplicaveis()
	//dados.GetFundamentacoesLegais()
	//dados.GetClassNbs()
	//dados.GetClassImpostoSeletivo()
	//dados.GetImpostoSeletivoCst()
	//dados.GetClassTribCbsIbs()
	//dados.GetClassTribCbsIbsUf()
	//dados.GetAliqUniao()
	//dados.GetAliqUf()
	//dados.GetAliqMun()
	pedagio.PostPedagioCalc()

	fmt.Println("✅ Execução finalizada com sucesso!")
}
