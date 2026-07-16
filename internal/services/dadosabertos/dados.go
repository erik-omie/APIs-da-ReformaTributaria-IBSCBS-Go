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

// Obtém a lista das unidades federativas cadastradas
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

	parametros := url.Values{}

	parametros.Add("siglaUf", "SP") // Sigla da unidade federativa

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

	Url := "https://consumo.tributos.gov.br/servico/calcular-tributos-consumo/api/calculadora/dados-abertos/situacoes-tributarias/imposto-seletivo"

	parametros := url.Values{}

	parametros.Add("data", "2027-01-01") //Data no padrão ISO 8601 (yyyy-MM-dd) Exemplo : "2027-01-01"

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

// Obtém a lista das situações tributárias cadastradas vigentes em uma determinada data para CBS/IBS
func GetCbsIbs() {

	Url := "https://consumo.tributos.gov.br/servico/calcular-tributos-consumo/api/calculadora/dados-abertos/situacoes-tributarias/cbs-ibs"

	parametros := url.Values{}

	parametros.Add("data", "2027-01-01") // Data no padrão ISO 8601 (yyyy-MM-dd) Exemplo : "2027-01-01"

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

// Obtém a lista de todos os redutores de compra governamental cadastrados
func GetRedutores() {

	Url := "https://consumo.tributos.gov.br/servico/calcular-tributos-consumo/api/calculadora/dados-abertos/redutores-compra-governamental"

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

// Obtém informações sobre a NCM em relação ao Imposto Seletivo
func GetNcm() {

	Url := "https://consumo.tributos.gov.br/servico/calcular-tributos-consumo/api/calculadora/dados-abertos/ncm"

	parametros := url.Values{}

	parametros.Add("ncm", "24021000")    // Código NCM sem formatação Exemplo : "24021000"
	parametros.Add("data", "2027-01-01") // Data no padrão ISO 8601 (yyyy-MM-dd) Exemplo : "2027-01-01"

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

// Obtém informações sobre a NBS em relação ao Imposto Seletivo
func GetNbs() {

	Url := "https://consumo.tributos.gov.br/servico/calcular-tributos-consumo/api/calculadora/dados-abertos/nbs"

	parametros := url.Values{}

	parametros.Add("nbs", "114052200")   // Código NBS sem formatação Exemplo : "24021000"
	parametros.Add("data", "2027-01-01") // Data no padrão ISO 8601 (yyyy-MM-dd) Exemplo : "2027-01-01"

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

// Obtém a lista de todas as NBS válidas em uma data, incluindo todos os níveis hierárquicos (capítulo, posição, subposições e itens)
func GetNbsLista() {

	Url := "https://consumo.tributos.gov.br/servico/calcular-tributos-consumo/api/calculadora/dados-abertos/nbs/lista"

	parametros := url.Values{}

	parametros.Add("data", "2027-01-01") // Data no padrão ISO 8601 (yyyy-MM-dd) Exemplo : "2027-01-01"

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

// Obtém a lista de NBS aplicáveis a um cClassTrib em uma determinada data. Se não houver NBS associados e a nomenclatura for aplicável a serviços (NBS ou 'NBS ou NCM'),
// retorna todos os NBS vigentes. Se não for aplicável a serviços (apenas NCM), retorna lista vazia.
func GetNbsAplicaveis() {

	Url := "https://consumo.tributos.gov.br/servico/calcular-tributos-consumo/api/calculadora/dados-abertos/nbs-aplicaveis"

	parametros := url.Values{}

	parametros.Add("cClassTrib", "000001") // Código da Classificação Tributária (cClassTrib) Exemplo : "000001"
	parametros.Add("data", "2027-01-01")   // Data no padrão ISO 8601 (yyyy-MM-dd) Exemplo : "2027-01-01"

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

// Obtém informações sobre as fundamentações legais
func GetFundamentacoesLegais() {

	Url := "https://consumo.tributos.gov.br/servico/calcular-tributos-consumo/api/calculadora/dados-abertos/fundamentacoes-legais"

	parametros := url.Values{}

	parametros.Add("data", "2027-01-01") // Data no padrão ISO 8601 (yyyy-MM-dd) Exemplo : "2027-01-01"

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

// Obtém a lista dos códigos cClassTrib aplicáveis a um NBS em uma determinada data, considerando vínculos e regras de exceção. Também inclui classificações de serviço vigentes não vinculadas a nenhum NBS.
func GetClassNbs() {

	Url := "https://consumo.tributos.gov.br/servico/calcular-tributos-consumo/api/calculadora/dados-abertos/classificacoes-tributarias/nbs"

	parametros := url.Values{}

	parametros.Add("nbs", "114052200")   // Código NBS sem formatação Exemplo : "24021000"
	parametros.Add("data", "2027-01-01") // Data no padrão ISO 8601 (yyyy-MM-dd) Exemplo : "2027-01-01"

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

// Obtém a lista das classificações tributárias (cClassTrib) para Imposto Seletivo
func GetClassImpostoSeletivo() {

	Url := "https://consumo.tributos.gov.br/servico/calcular-tributos-consumo/api/calculadora/dados-abertos/classificacoes-tributarias/imposto-seletivo"

	parametros := url.Values{}

	parametros.Add("data", "2027-01-01") //Data no padrão ISO 8601 (yyyy-MM-dd) Exemplo : "2027-01-01"

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

// Obtém a lista das classificações tributárias por CST para Imposto Seletivo
func GetImpostoSeletivoCst() {

	cst := "000" // Código da Situação Tributária (CST) Exemplo : "000001"

	Url := "https://consumo.tributos.gov.br/servico/calcular-tributos-consumo/api/calculadora/dados-abertos/classificacoes-tributarias/imposto-seletivo/" + cst

	parametros := url.Values{}

	parametros.Add("data", "2027-01-01") // Data no padrão ISO 8601 (yyyy-MM-dd) Exemplo : "2027-01-01"

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

// Obtém a lista das classificações tributárias (cClassTrib) para CBS e IBS
func GetClassTribCbsIbs() {

	Url := "https://consumo.tributos.gov.br/servico/calcular-tributos-consumo/api/calculadora/dados-abertos/classificacoes-tributarias/cbs-ibs"

	parametros := url.Values{}

	parametros.Add("data", "2027-01-01") // Data no padrão ISO 8601 (yyyy-MM-dd) Exemplo : "2027-01-01"

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

// Obtém a classificação tributária CBS/IBS com base na sigla do DFe e código da classificação
func GetClassTribCbsIbsUf() {

	siglaDfe := "NFSE"     // Sigla do tipo de Documento Fiscal Eletrônico Exemplo : "NFSE"
	cClassTrib := "000001" // Código da Classificação Tributária (cClassTrib) Exemplo : "000001"

	Url := "https://consumo.tributos.gov.br/servico/calcular-tributos-consumo/api/calculadora/dados-abertos/classificacoes-tributarias/imposto-seletivo/" + siglaDfe + "/" + cClassTrib

	parametros := url.Values{}

	parametros.Add("data", "2027-01-01") // Data no padrão ISO 8601 (yyyy-MM-dd) Exemplo : "2027-01-01"

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

// Obtém a alíquota padrão ou de referência para CBS
func GetAliqUniao() {

	Url := "https://consumo.tributos.gov.br/servico/calcular-tributos-consumo/api/calculadora/dados-abertos/aliquota-uniao"

	parametros := url.Values{}

	parametros.Add("data", "2026-07-16") // Data no padrão ISO 8601 (yyyy-MM-dd) Exemplo : "2027-01-01"

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

// Obtém a alíquota padrão ou de referência para IBS Estadual
func GetAliqUf() {

	Url := "https://consumo.tributos.gov.br/servico/calcular-tributos-consumo/api/calculadora/dados-abertos/aliquota-uf"

	parametros := url.Values{}

	parametros.Add("codigoUf", "43")     // Código da UF Exemplo : 43
	parametros.Add("data", "2027-01-01") //Data no padrão ISO 8601 (yyyy-MM-dd) Exemplo : "2027-01-01"

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

// Obtém a alíquota padrão ou de referência para IBS Municipal
func GetAliqMun() {

	Url := "https://consumo.tributos.gov.br/servico/calcular-tributos-consumo/api/calculadora/dados-abertos/aliquota-municipio"

	parametros := url.Values{}

	parametros.Add("codigoMunicipio", "4314902") // Código do Município (Tabela IBGE) Exemplo : 4314902
	parametros.Add("data", "2027-01-01")         //Data no padrão ISO 8601 (yyyy-MM-dd) Exemplo : "2027-01-01"

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
