package services

import (
	"bufio"
	"fmt"
	"os"

	"github.com/kyokomi/emoji/v2"
)

const DefaultUserName = ":satisfied:You: "

func GetRequest() string {
	fmt.Print("\n")
	scan := bufio.NewScanner(os.Stdin)
	if scan.Scan() {
		return scan.Text()
	}
	return "exit"
}
func PrintUserName(name string) string {
	fmt.Println("========================================")
	fmt.Printf("\u001B[33m%s\u001B[0m", emoji.Sprint(name))
	return GetRequest()
}
