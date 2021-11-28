package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var fileAccount = "accounts.txt"
var fileSettings = "settings.txt"

type Account struct {
	email, password string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getNumberOfAccounts(fileName string) int {
	file, err := os.Open(fileName)
	check(err)
	scanner := bufio.NewScanner(file)

	var numberOfAccounts int

	for i := 0; scanner.Scan(); i++ {
		numberOfAccounts = i
	}
	return numberOfAccounts
}

func getCredentials(fileName string) []Account {
	file, err := os.Open(fileName)
	check(err)
	scanner := bufio.NewScanner(file)

	var list = make([]Account, getNumberOfAccounts(fileName)+1)

	for i := 0; scanner.Scan(); i++ {
		list[i].email = scanner.Text()[:strings.IndexByte(scanner.Text(), ':')]
		list[i].password = scanner.Text()[strings.IndexByte(scanner.Text(), ':')+1:]
	}
	return list
}

func getSettings(fileName string) (string, bool, string) {
	file, err := os.Open(fileName)
	check(err)
	scanner := bufio.NewScanner(file)

	var server, screenPrefix string
	var isPremium bool

	for i := 0; scanner.Scan(); i++ {
		switch strings.ReplaceAll(strings.ToUpper(scanner.Text()[:strings.IndexByte(scanner.Text(), '=')]), " ", "") {
		case "SERVER":
			server = strings.ReplaceAll(strings.ToLower(scanner.Text()[strings.IndexByte(scanner.Text(), '=')+1:]), " ", "")
		case "PREMIUM":
			if strings.ReplaceAll(strings.ToLower(scanner.Text()[strings.IndexByte(scanner.Text(), '=')+1:]), " ", "") == "true" {
				isPremium = true
			} else {
				isPremium = false
			}
		case "SCREEN_PREFIX":
			screenPrefix = strings.ReplaceAll(scanner.Text()[strings.IndexByte(scanner.Text(), '=')+1:], " ", "")
		}
	}
	return server, isPremium, screenPrefix
}

/*
func executeLoginLinux(list []Account, server string, screen string) bool{
	command := make([]string, 10)
	command[0] = "screen"
	command[1] = "-S"
	command[3] = "-d"
	command[4] = "-m"
	command[5] = "mono"
	command[6] = "MinecraftClient.exe"
	command[9] = server

	for i:=0; i < len(list); i++ {
		command[2] = screen + string(rune(i))
		command[7] = list[i].email
		command[8] = list[i].password
		cmd := exec.Command(command[0], command[1:]...)
		cmd.Dir = "ConsoleClient"
		err := cmd.Start()
		check(err)
	}
	return true
}
*/

func main() {
	serverIP, isPremium, screenPrefix := getSettings(fileSettings)
	var altList = make([]Account, getNumberOfAccounts(fileAccount)+1)
	altList = getCredentials(fileAccount)
	fmt.Println(altList)
	fmt.Println("SERVER=", serverIP, "\nPREMIUM=", isPremium, "\nSCREEN=", screenPrefix)
}
