package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	for {

		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramente()
		case 2:
			fmt.Println("Exibindo Logs ....")
		case 3:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Nao conheco este comando")
			os.Exit(-1)
		}
	}

}

func exibeMenu() {
	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("3- Sair do programa")
}

func leComando() int {
	var comando int
	fmt.Scanf("%d", &comando)
	fmt.Println("O comando escolhido foi", comando)

	return comando
}

func iniciarMonitoramente() {
	fmt.Println("Monitorando")
	site := "https://alura.com.br/"
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site: ", site, "for carregado com sucesso!")
	} else {
		fmt.Println("Site: ", site, "esta com problemas. Status code: ", resp.StatusCode)
	}
}
