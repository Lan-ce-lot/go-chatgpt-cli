package services

import (
	"bufio"
	"fmt"
	"github.com/kyokomi/emoji/v2"
	"os"
)

const DefaultUserName = ":satisfied: You: "

func GetRequest() string {
	fmt.Println("\n")
	scan := bufio.NewScanner(os.Stdin)
	if scan.Scan() {
		return scan.Text()
	}
	return "exit"
}
func PrintUserName(name string) string {
	fmt.Print(emoji.Sprint(name))
	return GetRequest()
}
