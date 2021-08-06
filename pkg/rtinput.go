package input

import (
	"bufio"
	"fmt"
	"os"
)

func GetInput(question string) string {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println(question)
	inputText, _ := reader.ReadString('\n')
	return inputText
}
