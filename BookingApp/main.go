// Selalu gunakan package
package main

import "fmt"

// inisialisasi program yang akan dijalankan func main() adalah program yang pertama akan dijalankan
func main(){
	// Setelah mendeklarasi kan variable, var tersebut harus digunakan
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets int =50

	remainingTickets = -1
	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	// %v untuk memberitahu go untuk mengganti dengan variable
	fmt.Printf("Wellcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var userName string
	var userTickets int
	// ask user for their name

	userName = "Tom"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)
}
