package main

import "fmt"
func main(){
	var months = [...]string{
		"januari",
		"februari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}
	var slice1 = months[4:7]
	fmt.Println(slice1) //[mei juni juli]
	fmt.Println(len(slice1)) //3
	fmt.Println(cap(slice1)) //8

	//months[5] = "diubah"
	//fmt.Println(slice1)

	//slice1[0] = "mei lagi"
	//fmt.Println(months)

	var slice2 = months[10:] //[november desember]
	
	fmt.Println(slice2)

	var slice3 = append(slice2, "Eko") //[november desember eko]
	fmt.Println(slice3)

	slice3[1] = "bukan desember"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months) // [januari februari maret april mei juni juli agustus september oktober november desember]
	//dia tidak berubah karena jika kapasitas sudah penuh maka dia akan membuat array baru
	//jika slice2 diubh 2:4 maka hasilnya [januari februari maret /bukan desember Eko/ juni juli agustus september oktober november desember]
	//jika append masih sanggup menampung data maka array masih dengan data yang sama tetapi jika append tidak sanggup menampung data maka data array akan berubah (array baru)
	// jika diubah 11: data tidak berubah

	newSlice := make([]string, 2, 5)
	newSlice[0] = "roma"
	newSlice[1] = "yana"
	fmt.Println(newSlice)//[roma yana]
	fmt.Println(len(newSlice))// 2
	fmt.Println(cap(newSlice))// 5

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)//[roma yana]

	iniArray := [...]int{1, 2, 3, 4, 5}// [..]=[3]
	iniSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}