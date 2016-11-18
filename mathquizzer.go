package main

import (
	//"math/rand"
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {

	var i int
	fmt.Scan(&i)
	fmt.Println("read number", i, "from stdin")

	reader := bufio.NewReader(os.Stdin)

	a := rand.Intn(12)
	b := rand.Intn(12)

	operations := []string{"+", "-", "*", "/"}
	operationNumber := rand.Intn(3)

	fmt.Printf("What is " + strconv.Itoa(a) + operations[operationNumber] + strconv.Itoa(b) + " ? Type exit to exit: ")

	answer, _ := reader.ReadString('\n')

	if strings.EqualFold(answer, "exit") {

	}

	trueAnswer := 0
	switch operations[operationNumber] {
	case "+":
		trueAnswer = a + b
	case "-":
		trueAnswer = a - b
	case "*":
		trueAnswer = a * b
	case "/":
		trueAnswer = a / b
	}

	answerInt, _ := (strconv.Atoi(answer))
	fmt.Printf(answer)
	fmt.Println(trueAnswer)
	fmt.Printf("%v \n", answerInt)

	fmt.Printf("%T, %v \n", trueAnswer, trueAnswer)
	fmt.Printf("%T, %v \n", answerInt, answerInt)
	fmt.Println(trueAnswer == answerInt)

	if strings.EqualFold(answer, strconv.Itoa(trueAnswer)) {
		fmt.Printf("Correct! \n")
	} else {
		fmt.Printf("Incorrect! \n")
	}

}
