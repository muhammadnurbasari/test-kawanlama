package main

import (
	"test-kawanlama/empat"
	"test-kawanlama/satu"
	"test-kawanlama/tiga"
)

func main() {

	//  menyelesaikan soal 1
	var myRandom string = "USOMAANAPAIUMASYDNIP"

	satu.SoalSatu(myRandom)

	// menyelesaikan soal 3
	var j int32 = 5
	var k int32 = 4
	tiga.SoalTiga(j, k)

	//  menyelesaikan soal 4
	var jml empat.JumlahAnting
	jml.JumlahRuby = 5
	jml.JumlahTopaz = 3
	jml.JumlahPermata = 1

	var hrg empat.HargaAnting
	hrg.HargaRuby = 1000000
	hrg.HargaTopaz = 1250000
	hrg.HargaPermata = 3000000

	var c empat.Calculate
	c.JumlahAnting = jml
	c.HargaAnting = hrg

	c.SoalEmpat()

}
