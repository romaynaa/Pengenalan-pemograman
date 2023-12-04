package main
import "fmt"
func main(){
	var nilai_mhs int
	var mengerjakan_tugas bool
	var status_tugas bool
	fmt.Scan(&nilai_mhs, &mengerjakan_tugas)
	status_tugas = (nilai_mhs > 55 && mengerjakan_tugas==true) || nilai_mhs > 90
	fmt.Println(status_tugas)
}