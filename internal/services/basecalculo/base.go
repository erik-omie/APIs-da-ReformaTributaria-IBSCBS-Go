package basecalculo

import (
	"api-tributos/internal/client"
	"bytes"
	"encoding/json"
	"fmt"
)

/*
Este documento apresenta a especificação técnica e o guia de integração para a API de Base de Cálculo (Versão Beta), desenvolvida integralmente em Go (Golang) para garantir respostas em milissegundos e alta eficiência no processamento de regras fiscais complexas.

Ele centraliza todas as instruções, parâmetros e endpoints necessários para simular e estruturar as bases tributárias, permitindo que desenvolvedores e parceiros realizem cálculos fiscais precisos de forma automatizada.

⚠️ Nota de Versão: Por se tratar de uma versão BETA, os endpoints e estruturas de retorno estão sujeitos a ajustes e melhorias contínuas.
*/

// Afere a Base de Cálculo do Imposto Seletivo de uma operação de consumo. ATENÇÃO: Os campos ICMS e ISS não podem ser informados a partir de 2033.
func BaseCalculoIsMercadoriasPost() {

	url := "https://consumo.tributos.gov.br/servico/calcular-tributos-consumo/api/calculadora/base-calculo/is-mercadorias"

	dadosEnvio := IsMercadorias{
		AnoFatoGerador:        2027,
		ValorBem:              1000.00,
		AjusteAcrescimo:       50.00,
		Juros:                 10.00,
		Multas:                5.00,
		Encargos:              2.00,
		FreteCobrado:          20.00,
		OutrosTributos:        15.00,
		DemaisImportancias:    8.00,
		Icms:                  18.00,
		Iss:                   5.00,
		Cosip:                 3.00,
		Ipi:                   12.00,
		DescontoIncondicional: 30.00,
		Bonificacao:           25.00,
		DevolucaoVenda:        40.00,
	}

	dadosJson, err := json.Marshal(dadosEnvio)
	if err != nil {
		fmt.Println("Erro ao converter para JSON:", err)
		return
	}

	resosta, err := client.RequisicaoPost(url, dadosJson)
	if err != nil {
		fmt.Println("Erro ao enviar requisição:", err)
		return
	}

	var respostaFormatada bytes.Buffer
	err = json.Indent(&respostaFormatada, resosta, "", "  ")
	if err != nil {
		fmt.Println("Erro ao formatar resposta:", err)
		return
	}
	fmt.Println("Resposta formatada:")
	fmt.Println(respostaFormatada.String())
}

// Afere a Base de Cálculo da CBS/IBS de uma operação de consumo. ATENÇÃO: Os campos PIS, COFINS, PIS Importação e COFINS Importação não podem ser informados a partir de 2027. Os campos ICMS e ISS não podem ser informados a partir de 2033.
func BaseCalculoCbsIbsPost() {

	url := "https://consumo.tributos.gov.br/servico/calcular-tributos-consumo/api/calculadora/base-calculo/cbs-ibs-mercadorias"

	dadosEnvio := IsMercadorias{
		AnoFatoGerador:        2027,
		ValorBem:              1000.00,
		AjusteAcrescimo:       50.00,
		Juros:                 10.00,
		Multas:                5.00,
		Encargos:              2.00,
		FreteCobrado:          20.00,
		OutrosTributos:        15.00,
		DemaisImportancias:    8.00,
		Icms:                  18.00,
		Iss:                   5.00,
		Cosip:                 3.00,
		Ipi:                   12.00,
		DescontoIncondicional: 30.00,
		Bonificacao:           25.00,
		DevolucaoVenda:        40.00,
	}

	dadosJson, err := json.Marshal(dadosEnvio)
	if err != nil {
		fmt.Println("Erro ao converter para JSON:", err)
		return
	}

	resosta, err := client.RequisicaoPost(url, dadosJson)
	if err != nil {
		fmt.Println("Erro ao enviar requisição:", err)
		return
	}

	var respostaFormatada bytes.Buffer
	err = json.Indent(&respostaFormatada, resosta, "", "  ")
	if err != nil {
		fmt.Println("Erro ao formatar resposta:", err)
		return
	}
	fmt.Println("Resposta formatada:")
	fmt.Println(respostaFormatada.String())
}
