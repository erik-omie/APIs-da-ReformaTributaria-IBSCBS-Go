package client

import (
	"bytes"
	"io"
	"net/http"
	"time"
)

func RequisicaoGet(Url string) ([]byte, error) {

	// Criando um cliente HTTP com timeout
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	// Fazendo a requisição GET
	resp, err := client.Get(Url)
	if err != nil {
		return nil, err
	}

	// Garantindo que o corpo da resposta será fechado após a leitura
	defer resp.Body.Close()

	// Lendo o corpo da resposta
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Retornando o corpo da resposta e nenhum erro
	return body, nil
}

func RequisicaoPost(urlCompleta string, dadosJSON []byte) ([]byte, error) {
	// 1. Criando um cliente HTTP com timeout de segurança
	cliente := &http.Client{
		Timeout: 30 * time.Second,
	}

	// 2. Montando a requisição POST.
	// O bytes.NewBuffer transforma nossos bytes soltos num fluxo que viaja pela rede
	req, err := http.NewRequest("POST", urlCompleta, bytes.NewBuffer(dadosJSON))
	if err != nil {
		return nil, err
	}

	// 3. REGRA DE OURO: Avisar o governo que o idioma do nosso pacote é JSON
	req.Header.Set("Content-Type", "application/json")

	// 4. Disparando a requisição para o servidor
	resposta, err := cliente.Do(req)
	if err != nil {
		return nil, err
	}
	// Garantindo que a conexão será fechada ao final
	defer resposta.Body.Close()

	// 5. Lendo a resposta que o governo nos devolveu
	corpo, err := io.ReadAll(resposta.Body)
	if err != nil {
		return nil, err
	}

	// 6. Retornando os dados lidos
	return corpo, nil
}
