package calculadora

import (
	"api-tributos/internal/client"
	"bytes"
	"encoding/json"
	"fmt"
)

func XmlValidate() {

	Url := "https://consumo.tributos.gov.br/servico/calcular-tributos-consumo/api/calculadora/xml/validate"

	// Fazendo a requisução via client
	dados, err := client.RequisicaoGet(Url)

	//tratamento de erro
	if err != nil {
		fmt.Println("Erro ao fazer a requisição:", err)
		return
	}

	// Guardando o resultado da requisição em um buffer para formatação JSON
	var jsonFormatado bytes.Buffer

	erroIndent := json.Indent(&jsonFormatado, dados, "", "  ")

	if erroIndent != nil {
		fmt.Println("Erro ao formatar o JSON:", erroIndent)
		return
	}

	//Convertemos os bytes recebidos para texto e imprimimos na tela
	fmt.Println("✅ Resposta da API de validação do XML:", jsonFormatado.String())

}

func XmlGenerate() {

	url := "https://consumo.tributos.gov.br/servico/calcular-tributos-consumo/api/calculadora/xml/generate"

	// Fazendo a requisição via client
	dados, err := client.RequisicaoGet(url)

	//tratamento de erro
	if err != nil {
		fmt.Println("Erro ao fazer a requisição:", err)
		return
	}

	// Guardando o resultado da requisição em um buffer para formatação JSON
	var jsonFormatado bytes.Buffer

	erroIndent := json.Indent(&jsonFormatado, dados, "", "  ")

	if erroIndent != nil {
		fmt.Println("Erro ao formatar o JSON:", erroIndent)
		return
	}

	fmt.Println("✅ Resposta da API de geração do XML:", jsonFormatado.String())

}
