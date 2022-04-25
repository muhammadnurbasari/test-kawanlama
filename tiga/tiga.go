package tiga

import "fmt"

func SoalTiga(j, k int32) {
	fmt.Println("JAWABAN SOAL TEST NO 3")
	fmt.Println("======================")
	fmt.Println("NILAI J = ", j)
	fmt.Println("NILAI K = ", k)

	var hasil int32
	for j > 0 {
		hasil = hasil + k

		j--
	}

	fmt.Println("HASIL PERKALIAN = ", hasil)
	fmt.Println("======================")
	fmt.Println("======================")
}
