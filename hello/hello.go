package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	numeroDeMonitoramentos = 3
	delay                  = 5
)

func main() {
	leArquivoSites()
	exibeIntroducao()

	for {
		exibeMenu()

		comando := lerComando()

		switch comando {
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		case 1:
			monitorar()
		case 2:
			fmt.Println("Exibindo logs")
			imprimeLogs()
		default:
			fmt.Println("Comando invalido")
			os.Exit(-1)
		}
	}
}

func exibeIntroducao() {
	nome := "Raissa"
	versao := 1.1

	fmt.Println("Olá, ", nome)
	fmt.Println("Este programa está na versão ", versao)
}

func exibeMenu() {
	fmt.Println("0 - Sair do programa")
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")

}

func lerComando() int {
	var comandoEscolhido int
	fmt.Scan(&comandoEscolhido)
	fmt.Println("Comando escolhido: ", comandoEscolhido)
	fmt.Println()
	return comandoEscolhido
}

func monitorar() {
	fmt.Println("Iniciando monitoramento...")
	sites := leArquivoSites()

	for i := 0; i < numeroDeMonitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println()
	}
	fmt.Println()
}

func testaSite(site string) {
	response, err := http.Get(site)
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	statusCode := response.StatusCode

	if statusCode == 200 {
		fmt.Println("O site:", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("O site:", site, "esta com problema. StatusCode", statusCode)
		registraLog(site, false)
	}
}

func leArquivoSites() []string {
	var sites []string

	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}

	arquivo.Close()

	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05 - ") + site + " - online: " + strconv.FormatBool(status) + "\n")
	arquivo.Close()
}

func imprimeLogs() {
	arquivo, err := ioutil.ReadFile("logs.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(string(arquivo))

}
