package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func AddNumCommand(num uint32, s string, m map[uint32]string, a []uint32, counter uint32, size uint32, wordPosMap map[string][]int) (uint32, []uint32) {
	_, pres := m[num]

	m[num] = s

	if counter == size {
		if !pres {
			delete(m, a[0])
			return counter, a
		} else {
			return counter, a
		}
	} else {
		if !pres {
			counter++
		} else {
			return counter, a
		}
	}
	a[counter-1] = num

	return counter, a
}

func GetNumFromString(s string) uint32 {
	num, _ := strconv.Atoi(s)
	return uint32(num)
}

func GetCommand() string {

	reader := bufio.NewReader(os.Stdin)
	command, _ := reader.ReadString('\n')
	command = strings.Replace(command, "\n", "", 1)
	return command
}

func FindWord(m map[uint32]string, a []uint32, word string, limit uint32) []uint32 {
	foundArr := make([]uint32, 0)
	var counter uint32 = 0
	for i := (len(a) - 1); i >= 0; i-- {
		s := m[a[i]]
		if strings.Contains(s, word) {
			foundArr = append(foundArr, a[i])
			if counter+1 == limit {
				break
			}
			counter++
		}
	}
	return foundArr
}
