package main
import "fmt"
func main(){

	var bil, jumlah int
	fmt.Scan(&bil)
	
	for bil > 0 {
		jumlah += bil % 10
		bil = bil / 10

	
	}
	fmt.Println(jumlah)
}

		