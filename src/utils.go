package src

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type intTuple [2]int

func Abs(num int) int {
	if num < 0 {
		num = -1 * num
	}
	return num
}

func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

// ReadFile reads a text file to extract the nodes list of it
func ReadFile(name string) (maxvalue int, nodeList []intTuple, lineCount int) {
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Some error just happened.")
		}
	}(file)
	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	scanner.Scan()
	line := scanner.Text()
	_, _ = fmt.Sscanf(line, "%d", &lineCount)

	var anod intTuple
	for scanner.Scan() {
		lineStr := scanner.Text()
		_, err2 := fmt.Sscanf(lineStr, "%d %d", &anod[0], &anod[1])
		if err2 != nil {
			return 0, nil, 0
		}
		nodeList = append(nodeList, anod)
		if maxvalue < anod[0] {
			maxvalue = anod[0]
		}
		if maxvalue < anod[1] {
			maxvalue = anod[1]
		}
	}
	return
}
