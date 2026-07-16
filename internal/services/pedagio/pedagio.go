package pedagio

/*
Este documento apresenta a especificação técnica e o guia de integração para a API de Pedágio (Versão Beta), desenvolvida integralmente em Go (Golang) para garantir alta performance no mapeamento de rotas e tarifas em tempo real.

Ele centraliza todas as instruções, parâmetros e endpoints necessários para consultar tarifas vigentes e calcular custos de praças de pedágio de forma automatizada e precisa.

⚠️ Nota de Versão: Por se tratar de uma versão BETA, os endpoints e estruturas de retorno estão sujeitos a ajustes e melhorias contínuas.
*/

import (
	"api-tributos/internal/client"
	"bytes"
	"encoding/json"
	"fmt"
)

// Calculo do pedágio
func PostPedagioCalc() {

	urlPost := "https://consumo.tributos.gov.br/servico/calcular-tributos-consumo/api/calculadora/pedagio"

	dadosEnvio := EnvioCalculo{
		DataHoraEmissao:       "2027-01-01T09:50:05-03:00",
		CodigoMunicipioOrigem: 4314902,
		UFMunicipioOrigem:     "RS",
		CST:                   "000",
		CClassTrib:            "000002",
		BaseCalculo:           200,
		Trechos: []Trecho{
			{
				Numero:    1,
				Municipio: 4314902,
				UF:        "RS",
				Extensao:  10,
			},
		},
	}

	dadosJson, err := json.Marshal(dadosEnvio)
	if err != nil {
		fmt.Println("Erro ao converter dados para JSON:", err)
		return
	}

	resposta, err := client.RequisicaoPost(urlPost, dadosJson)
	if err != nil {
		fmt.Println("Erro ao enviar requisição:", err)
		return
	}

	var respostaFormatada bytes.Buffer
	err = json.Indent(&respostaFormatada, resposta, "", "  ")
	if err != nil {
		fmt.Println("Erro ao formatar resposta:", err)
		return
	}

	fmt.Println("Resposta da API:", respostaFormatada.String())
}
