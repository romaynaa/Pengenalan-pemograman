package main 

import "fmt"

func main(){
	var kar int
	fmt.Scanf("%c", &kar)
	if 'a' <= kar && kar <= 'z' || 'A' <= kar && kar <= 'Z' {
		fmt.Println("Alphabet")
	}else{
		fmt.Println("Bukan Alphabet")
	}
}
// fmt.Scanf digunakan untuk membaca input dari pengguna dengan format tertentu dan menyimpannya kedalam variabel yang diberikan sebagai argumen