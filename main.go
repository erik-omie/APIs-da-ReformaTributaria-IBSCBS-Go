package main

import (
	"api-tributos/internal/services/calculadora"
	"fmt"
)

func main() {
	// Usamos o pacote 'fmt' para imprimir mensagens no terminal
	fmt.Println("🚀 Iniciando o Sistema de Consulta de Tributos...")

	//calculadora.XmlValidate()
	calculadora.XmlGenerate()

	fmt.Println("✅ Execução finalizada com sucesso!")
}
