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

const monitoringTimes = 3
const delay = 5

func main() {
	introduction()
	for {
		showOptions()
		selectedOption := readCommand()
		handleOptions(selectedOption)
	}
}

func introduction() {
	var autor string = "Lucas"
	fmt.Println("Olá, sr", autor)

	var version float32 = 1.1
	fmt.Println("Versão", version)
}

func readCommand() int {
	var selectedCommand int
	// fmt.Scanf("%d", &selectedCommand)
	fmt.Scan(&selectedCommand)
	return selectedCommand
}

func showOptions() {
	fmt.Println("")
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair do Programa")
	fmt.Println("")
}

func handleOptions(option int) {
	// if selectedCommand == 1 {

	// } else if selectedCommand == 2 {

	// } else if selectedCommand == 0 {

	// } else {
	// 	fmt.Println("Comando desconhecido")
	// }

	switch option {
	case 1:
		monitoring()
	case 2:
		fmt.Println("Exibindo logs...")
		readLogs()
	case 0:
		fmt.Println("Saindo do programa")
		os.Exit(0)
	default:
		fmt.Println("Comando desconhecido")
		os.Exit(-1)
	}
}

func monitoring() {
	fmt.Println("Monitorando...")

	sites := readFile()

	for i := 0; i < monitoringTimes; i++ {
		for _, site := range sites {
			validateSiteStatus(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
}

func validateSiteStatus(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site", site, "foi carregado")
		logRegister(site, true)
	} else {
		fmt.Println("Site", site, "esta com problemas")
		logRegister(site, false)
	}
}

func readFile() []string {
	var sites []string

	file, err := os.Open("sites.txt")

	//file, err := ioutil.ReadFile("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	leitor := bufio.NewReader(file)

	for {
		line, err := leitor.ReadString('\n')
		line = strings.TrimSpace(line)

		sites = append(sites, line)

		if err == io.EOF {
			break
		}
	}

	file.Close()

	return sites
}

func readLogs() {
	file, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(file))
}

func logRegister(site string, status bool) {
	//OpenFile permite criar um arquivo caso ele não exista
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	file.WriteString(time.Now().Format("[02/01/2006 15:04:05] ") + site + " - online: " + strconv.FormatBool(status) + "\n")

	file.Close()
}

// func showNamesSlice() {
// 	nomes := []string{"Lucas", "Ruan"}

// 	fmt.Println(nomes)
// }

// func devolveNomeEIdade() (string, int) {
// 	name := "Lucas"
// 	idade := 28
// 	return name, idade
// }
