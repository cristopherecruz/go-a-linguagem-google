package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {

	for {
		nome, idade := devolverNomeIdade()
		exibirIntroducao(nome, idade)
		exibirMenu()
		/*if comando == 1 {
			println("Monitorando...")
		} else if comando == 2 {
			println("Exibindo Logs...")
		} else if comando == 0 {
			println("Saindo do Programa...")
		} else {
			fmt.Println("Não conmheço este comando")
		}*/

		comando, _ := lerComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			imprimirLogs()
		case 0:
			fmt.Println("Saindo do Programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}

}

func devolverNomeIdade() (string, int) {
	nome := "Cristopher"
	idade := 29

	return nome, idade
}

func exibirIntroducao(nome string, idade int) {

	versao := 1.1
	fmt.Println("Hello World!", nome, "sua idade é", idade)
	fmt.Println("Este programa está na versão", versao)
	fmt.Println("O tipo da variável é", reflect.TypeOf(versao))
}

func exibirMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func lerComando() (int, error) {

	var comando int

	scan, err := fmt.Scan(&comando)
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	fmt.Println("O endereço da minha variável comando é", &comando)
	fmt.Println("O comando escolhido foi", scan)
	fmt.Println("")

	return comando, nil
}
func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	/*sites := []string{"https://random-status-code.herokuapp.com/",
		"https://www.alura.com.br",
		"https://www.caelum.com.br",
		"https://cristopher.dev.br",
	}*/

	sites := lerSitesArquivo()

	for i := 0; i < monitoramentos; i++ {

		for _, site := range sites {
			fmt.Println("Testando site:", site)
			testarSite(site)

		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")

}

func testarSite(site string) {
	response, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if response.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registrarLog(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas, status code:", response.StatusCode)
		registrarLog(site, false)
	}
}

func lerSitesArquivo() []string {

	var sites []string

	arquivo, err := os.Open("sites.txt")
	//arquivo, err := os.ReadFile("sites.txt")

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

	err = arquivo.Close()
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	return sites
}

func registrarLog(site string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	_, err = arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "- online: " + strconv.FormatBool(status) + "\n")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	err = arquivo.Close()
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

}

func imprimirLogs() {
	fmt.Println("Exibindo Logs...")

	arquivo, err := os.ReadFile("log.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(string(arquivo))

}
