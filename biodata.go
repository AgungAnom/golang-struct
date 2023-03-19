package main

import (
	"fmt"
	"os"
	"strconv"
)


type Biodata struct{
absen int
nama string
alamat string
pekerjaan string
alasan string

}


func (b Biodata) printBio(){
	fmt.Println("Absen 		:", b.absen)
	fmt.Println("Nama 		:", b.nama)
	fmt.Println("Alamat		:", b.alamat)
	fmt.Println("Pekerjaan 	:", b.pekerjaan)
	fmt.Println("Alasan		:", b.alasan)
}

func main(){
var bio = []Biodata{
	{absen:1,nama:"Agung", alamat:"Tabanan, Bali",pekerjaan: "Full Stack Developer", alasan: "Belajar Golang"},
	{absen:2,nama:"Anom", alamat:"Tabanan, Bali",pekerjaan: "Front End Developer", alasan: "Belajar Back End"},
	{absen:3,nama:"Semara", alamat:"Tabanan, Bali",pekerjaan: "UI/UX", alasan: "Menambah ilmu di bidang lain"},
	{absen:4,nama:"Putra", alamat:"Tabanan, Bali",pekerjaan: "Back End Developer", alasan: "Belajar Bahasa Baru"},
}
input := os.Args[1]
inputInt,_ := strconv.Atoi(input)
state := 0


for _, data := range bio {
	if data.absen == inputInt {
		state = 1
		data.printBio()
	} 
}


if state == 0 {
	fmt.Println("Peserta dengan nomor absen", inputInt, "tidak ditemukan!")
}

}


