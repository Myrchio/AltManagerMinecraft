package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var fileAccount = "accounts.txt"
var fileSettings = "settings.txt"

type Account struct {
	email, password string
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func GetNumberOfAccounts(fileName string) int {
	file, err := os.Open(fileName)
	Check(err)
	scanner := bufio.NewScanner(file)

	var numberOfAccounts int

	for i := 0; scanner.Scan(); i++ {
		numberOfAccounts = i
	}
	return numberOfAccounts
}

func GetCredentials(fileName string) []Account {
	file, err := os.Open(fileName)
	Check(err)
	scanner := bufio.NewScanner(file)

	var list = make([]Account, GetNumberOfAccounts(fileName)+1)

	for i := 0; scanner.Scan(); i++ {
		list[i].email = scanner.Text()[:strings.IndexByte(scanner.Text(), ':')]
		list[i].password = scanner.Text()[strings.IndexByte(scanner.Text(), ':')+1:]
	}
	return list
}

func GetSettings(fileName string) (string, bool, string) {
	file, err := os.Open(fileName)
	Check(err)
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

func KillScreens() {
	cmd := exec.Command("killall", "screen")
	err := cmd.Start()
	Check(err)
}

func ExecuteLoginLinux(list []Account, server string, screen string) error {
	command := make([]string, 8)
	command[0] = "screen"
	command[1] = "-dmS"
	command[3] = "mono"
	command[4] = "MinecraftClient.exe"
	command[7] = server
	KillScreens()
	for i := 0; i < len(list); i++ {
		command[2] = screen + strconv.Itoa(i+1)
		command[5] = list[i].email
		command[6] = list[i].password
		cmd := exec.Command(command[0], command[1:]...)
		cmd.Dir = "ConsoleClient"
		fmt.Println("Logging", list[i].email)
		err := cmd.Start()
		if err != nil {
			return err
		}
		time.Sleep(3 * time.Second)

	}
	return nil
}

func main() {
	serverIP, _, screenPrefix := GetSettings(fileSettings)
	var altList = make([]Account, GetNumberOfAccounts(fileAccount)+1)
	altList = GetCredentials(fileAccount)

	switch runtime.GOOS {
	case "linux":
		err := ExecuteLoginLinux(altList, serverIP, screenPrefix)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		} else {
			//todo server
		}
	case "windows":
		//todo ExecuteLoginWindows()

	default:
		fmt.Println("OS not supported")
	}
}
