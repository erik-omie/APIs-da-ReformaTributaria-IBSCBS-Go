package nfse

import (
	"api-tributos/internal/client"
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
)

/*
Este documento apresenta a especificação técnica e o guia de integração para a API de NFS-e (Versão Beta), desenvolvida integralmente em Go (Golang) para garantir resiliência, baixo consumo de memória e agilidade na comunicação com diversas prefeituras.

Ele centraliza todas as instruções, parâmetros e endpoints necessários para a emissão, consulta e cancelamento de Notas Fiscais de Serviços Eletrônicas de forma automatizada.

⚠️ Nota de Versão: Por se tratar de uma versão BETA, os endpoints e estruturas de retorno estão sujeitos a ajustes e melhorias contínuas.
*/

// Retorna todas as combinações de Classificações Tributárias e Situações Tributárias associadas a um código NBS em uma data específica
func GetClassTributarias() {

	Url := "https://consumo.tributos.gov.br/servico/calcular-tributos-consumo/api/calculadora/nfse/situacoes-classificacoes-tributarias"

	parametros := url.Values{}

	parametros.Add("data", "2026-06-01") // Data no padrão ISO 8601 (yyyy-MM-dd) Exemplo : "2027-01-01"
	parametros.Add("nbs", "101011100")   // Código NBS sem formatação Exemplo : "24021000"

	UrlCompleta := Url + "?" + parametros.Encode()

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

// Consulta o local da operação baseado no código do Indicador de Operação (cIndOp) e data de ocorrência do fato gerador
func GetLocalOperacao() {

	Url := "https://consumo.tributos.gov.br/servico/calcular-tributos-consumo/api/calculadora/nfse/local-operacao"

	parametros := url.Values{}

	parametros.Add("cIndOp", "100301")                        // Código do indicador de Operação (6 dígitos) Exemplo : "100301"
	parametros.Add("dataOcorrenciaFatoGerador", "2026-06-01") // Data no padrão ISO 8601 (yyyy-MM-dd) Exemplo : "2027-01-01"

	UrlCompleta := Url + "?" + parametros.Encode()

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
