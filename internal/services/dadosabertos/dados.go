package dados

import (
	"api-tributos/internal/client"
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
)

// Obtém a versão do aplicativo e do banco de dados
func GetVersao() {

	Url := "https://consumo.tributos.gov.br/servico/calcular-tributos-consumo/api/calculadora/dados-abertos/versao"

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
	fmt.Println("✅ Resposta da API de obtenção das UFs:", jsonFormatado.String())

}

func GetUf() {

	Url := "https://consumo.tributos.gov.br/servico/calcular-tributos-consumo/api/calculadora/dados-abertos/ufs"

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
	fmt.Println("✅ Resposta da API de obtenção das UFs:", jsonFormatado.String())

}

// Obtém a lista dos municípios cadastrados com base na sigla de uma unidade federativa
func GetMunicipios() {

	Url := "https://consumo.tributos.gov.br/servico/calcular-tributos-consumo/api/calculadora/dados-abertos/ufs/municipios"

	stringUf := url.Values{}

	stringUf.Add("siglaUf", "SP")

	UrlCompleta := Url + "?" + stringUf.Encode()

	// Fazendo a requisução via client
	dados, err := client.RequisicaoGet(UrlCompleta)

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
	fmt.Println("✅ Resposta da API de obtenção das UFs:", jsonFormatado.String())

}

// Obtém a lista de todos os percentuais de transferência IBS cadastrados
func GetTransferenciasIbs() {

	Url := "https://consumo.tributos.gov.br/servico/calcular-tributos-consumo/api/calculadora/dados-abertos/transferencias-ibs"

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
	fmt.Println("✅ Resposta da API de obtenção das UFs:", jsonFormatado.String())

}

// Obtém a lista de todos os percentuais de transferência CBS cadastrados
func GetTransferenciasCbs() {

	Url := "https://consumo.tributos.gov.br/servico/calcular-tributos-consumo/api/calculadora/dados-abertos/transferencias-cbs"

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
	fmt.Println("✅ Resposta da API de obtenção das UFs:", jsonFormatado.String())

}

// Obtém a lista das situações tributárias cadastradas para Imposto Seletivo
func GetImpostoSeletivo() {

	Url := "https://consumo.tributos.gov.br/servico/calcular-tributos-consumo/api//calculadora/dados-abertos/situacoes-tributarias/imposto-seletivo"

	data := url.Values{}

	data.Add("data", "2027-01-01")

	UrlCompleta := Url + "?" + data.Encode()

	// Fazendo a requisução via client
	dados, err := client.RequisicaoGet(UrlCompleta)

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
	fmt.Println("✅ Resposta da API de obtenção das UFs:", jsonFormatado.String())

}
