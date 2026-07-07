package client

import (
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
