package satu

import (
	"fmt"
	"strings"
)

func SoalSatu(s string) {
	// make slice string from string parameter
	mySliceString := strings.Split(s, "")

	var sliceSiapa = []string{
		"S", "I", "A", "P", "A",
	}

	// make a closure
	var totalSiapa = func(randomSLice, s []string) (int32, []string) {

		var (
			total int32
		)

		// loop slice siapa
		for i := 0; i < len(s); i++ {

			for j := 0; j < len(randomSLice); j++ {
				if randomSLice[j] == s[i] {
					total = total + 1
					randomSLice = append(randomSLice[:j], randomSLice[j+1:]...)
					break
				} else {
					continue
				}
			}

		}

		if total == 5 {
			return 1, randomSLice
		}

		return 0, randomSLice

	}

	var result int32
	for true {

		myTotal, myRandom := totalSiapa(mySliceString, sliceSiapa)

		mySliceString = myRandom
		if myTotal == 1 {
			result = result + myTotal
			continue
		} else {
			break
		}

	}

	fmt.Println("JAWABAN SOAL TEST NO 1")
	fmt.Println("======================")
	fmt.Println("CONTOH STRING USOMAANAPAIUMASYDNIP")
	fmt.Println("BERAPA BANYAK KATA SIAPA YANG DAPAT DIBENTUK ?")
	fmt.Println("JAWABANNYA ADALAH ", result)
	fmt.Println("======================")
}
