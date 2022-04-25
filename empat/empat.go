package empat

import "fmt"

type (
	JumlahAnting struct {
		JumlahRuby    int32
		JumlahTopaz   int32
		JumlahPermata int32
	}

	HargaAnting struct {
		HargaRuby    int32
		HargaTopaz   int32
		HargaPermata int32
	}

	Calculate struct {
		JumlahAnting
		HargaAnting
	}
)

// SoalEmpat - untuk menjawab case 4
func (c *Calculate) SoalEmpat() {

	//  membuat closure untuk menentukan maksimal jumlah pasangan setiap type anting
	var maksPasang = func(jml *JumlahAnting) map[string]int32 {
		var (
			resultMaxRuby, resultMaxTopaz, resultMaxPermata int32
		)

		// ruby
		modulusRuby := jml.JumlahRuby % 2
		if modulusRuby == 0 {
			resultMaxRuby = jml.JumlahRuby / 2
		} else if modulusRuby == 1 {
			resultMaxRuby = (jml.JumlahRuby - 1) / 2
		}

		// topaz
		modulusTopaz := jml.JumlahTopaz % 2
		if modulusTopaz == 0 {
			resultMaxTopaz = jml.JumlahTopaz / 2
		} else if modulusTopaz == 1 {
			resultMaxTopaz = (jml.JumlahTopaz - 1) / 2
		}

		// permata
		modulusPermata := jml.JumlahPermata % 2
		if modulusPermata == 0 {
			resultMaxPermata = jml.JumlahPermata / 2
		} else if modulusPermata == 1 {
			resultMaxPermata = (jml.JumlahPermata - 1) / 2
		}

		var resp = map[string]int32{
			"ruby":    resultMaxRuby,
			"topaz":   resultMaxTopaz,
			"permata": resultMaxPermata,
		}

		return resp

	}

	// menentukan maksimal penjualan per type anting
	fmt.Println("JAWABAN SOAL TEST NO 4")
	fmt.Println("======================")

	resultPasangan := maksPasang(&c.JumlahAnting)

	var sliceType = []string{
		"ruby",
		"topaz",
		"permata",
	}

	for _, v := range sliceType {
		fmt.Println("Maks Pasangan "+v+" = ", resultPasangan[v])
	}

	var maksPenjualan int32
	for _, v := range sliceType {

		var myPrice int32

		if v == "ruby" {
			myPrice = c.HargaAnting.HargaRuby
		} else if v == "topaz" {
			myPrice = c.HargaAnting.HargaTopaz
		} else {
			myPrice = c.HargaAnting.HargaPermata
		}

		fmt.Println("Maks Penjualan Anting type "+v+" = ", resultPasangan[v]*myPrice)

		maksPenjualan = maksPenjualan + (resultPasangan[v] * myPrice)
	}

	fmt.Println("MAKSIMAL PENJUALAN ADALAH ", maksPenjualan)
	fmt.Println("======================")
}
