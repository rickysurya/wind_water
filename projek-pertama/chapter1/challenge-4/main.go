package main

import (
	"fmt"
	"os"
	"strconv"
)

type student struct {
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

const job = "IT Dev"

type biodataGenerator interface{
	biodata() []student
}

type namaBiodata struct{
	nama []string
}


func newNamaBiodata(nama []string)biodataGenerator {
	return &namaBiodata{nama: nama}
}

func (n *namaBiodata)biodata() []student {
	almt := []string{"Jakarta", "Bogor", "Padang", "Medan", "Bandung"}
	alasan := []string{"Milik abdul", "Milik budi", "Milik caca", "Milik dede", "Milik erlan"}
	generate := make([]student, 0)
	var s student
	for key, value := range n.nama {
		s.nama = value
		s.alamat = almt[key]
		s.pekerjaan = job
		s.alasan = alasan[key]
		generate = append(generate, s)
	}
	return generate
}

func main() {
	nama := []string{"Abdul", "Budi", "Caca", "Dede", "Erlan"}
	generator := newNamaBiodata(nama)
	output := generator.biodata()
	arg, _ := strconv.Atoi(os.Args[1])

	var inama string

	for i, v := range nama {
		if arg == i {
			inama = v
		}
	}

	for i, v := range output {
		if inama == v.nama {
			fmt.Println("ID :", i)
			fmt.Println("Nama :", v.nama)
			fmt.Println("Pekerjaan :", v.pekerjaan)
			fmt.Println("Alasan :", v.alasan)
		}
	}

}
