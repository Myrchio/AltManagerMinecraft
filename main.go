package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Account struct {
	email, password string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func takeNumberOfAccounts(fileName string) int {
	var numberOfAccounts int
	file, err := os.Open(fileName)
	check(err)

	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		numberOfAccounts = i
	}
	return numberOfAccounts
}

func takeCredentials(fileName string) []Account {
	file, err := os.Open(fileName)
	check(err)
	var list = make([]Account, takeNumberOfAccounts(fileName)+1)

	scanner := bufio.NewScanner(file)

	for i := 0; scanner.Scan(); i++ {
		list[i].email = scanner.Text()[:strings.IndexByte(scanner.Text(), ':')]
		list[i].password = scanner.Text()[strings.IndexByte(scanner.Text(), ':')+1:]
	}
	return list
}

func main() {
	var accountFile = "accounts.txt"
	var altList = make([]Account, takeNumberOfAccounts(accountFile)+1)
	altList = takeCredentials(accountFile)

	fmt.Println(altList)
}
