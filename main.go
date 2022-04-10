package main

import (
	"ayush-100ms/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {

	var (
		size    uint64
		command string
		counter uint64 = 0
	)

	fmt.Scanf("%d", &size)

	numToLogMap := make(map[uint64]string)
	wordPosMap := make(map[string][]int)

	arr := make([]uint64, size)

	for i := true; i; {
		command = utils.GetCommand()

		if strings.Contains(command, "END") {
			fmt.Println("END")
			i = false
		} else if strings.Contains(command, "ADD") {

			command = strings.Replace(command, "ADD ", "", 1)
			commandArray := strings.Split(command, " ")

			num := utils.GetNumFromString(commandArray[0])
			commandArray = commandArray[1:]
			command = strings.Join(commandArray, " ")

			counter, arr = utils.AddNumCommand(num, command, numToLogMap, arr[:], counter, size, wordPosMap)

		} else if strings.Contains(command, "SEARCH") {
			command = strings.Replace(command, "SEARCH ", "", 1)

			searchArray := strings.Split(command, " ")
			term := searchArray[0]
			count, _ := strconv.Atoi(searchArray[1])
			foundArr := utils.FindWord(numToLogMap, arr[:], term, uint64(count))
			if len(foundArr) == 0 {
				fmt.Println("NONE")
			} else {
				for _, v := range foundArr {
					fmt.Print(v, " ")
				}
				fmt.Println(" ")
			}

		}
	}
}
