package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
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
	exibeIntroducao()

	exibeNomes()

	// if comando == 1 {
	// 	fmt.Println("1")
	// } else if comando == 2 {
	// 	fmt.Println("2")
	// } else {
	// 	fmt.Println("Nao conheco")
	// }

	for {
		comando, _ := leComando()

		switch comando {
		case 1:
			iniciarMonitoranmento()
		case 2:
			imprimeLogs()
		case 0:
			fmt.Println("ok")
			os.Exit(0)
		default:
			fmt.Println("Nao conheco")
			os.Exit(-1)
		}
	}
}

func exibeIntroducao() {
	nome := "Douglas"
	fmt.Println("Hello", nome)
	fmt.Println("type =", reflect.TypeOf(nome))
}

func leComando() (int, string) {
	var comando int
	fmt.Scan(&comando)

	comandoPrint := "typed"

	return comando, comandoPrint
}

func iniciarMonitoranmento() {
	fmt.Println("Monitorando site")

	// var sites [4]string
	// sites[0] = "https://www.alura.com.br"

	sites := leSitesArquivo()

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testSite(site)
		}

		time.Sleep(delay * time.Second)
	}
}

func testSite(site string) {
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "esta com proplemas. Status Code:", resp.StatusCode)
		registraLog(site, false)
	}
}

func exibeNomes() {
	nomes := []string{"Douglas", "Daniel", "Bernado"}
	nomes = append(nomes, "Aparecida")
	fmt.Println(nomes)
}

func leSitesArquivo() []string {
	var sites []string

	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println(err)
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
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {
	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))
}
