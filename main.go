package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main() {
	var input string
	fmt.Println("Vai a questo link, mandalo a lorenzo e chiedigli come si legge, scrivi la risposta.")
	fmt.Println("Chiedigli di dirti in italiano lettera per lettera cosa dice il messaggio, senza spazi e tutto minuscolo")
	fmt.Println("sorpresa.jpg")
	//I'm sorry, the file is no longer available
	fmt.Scanln(&input)
	if input != "lamultiani" {
		fmt.Println("Mi spiace, pensaci bene. CTRL + C per chiudere.")
		return
	} else {
		fmt.Println("Dì bravo a lorenzo che sa la sua lingua")
		fmt.Println("Un giorno poi mi chiederai cosa significa")
		fmt.Println("			 ____	 __     __     _____	 __     __    _______     ____    ____")
		fmt.Println("			/    |  |  |   |  |   / ____/   |  |   |  |  |  ___   \\   |  |    |  |")
		fmt.Println("		       /  /| |  |  |   |  |  / /	|  |   |  |  | |___|  |   |  |    |  |")
		fmt.Println("		      /  / | |  |  |   |  |  | |  ___   |  |   |  |  |  ___  /	  |  |    |  |")
		fmt.Println("		     /  /__| |  |  |   |  |  | | |_ _|  |  |   |  |  |  | \\  \\    |  |    |__|")
		fmt.Println("	            /  ____  |  |  \\___/  |  | \\__/ |   |  |___|  |  |  |  \\  \\   |  |    ____")
		fmt.Println("	           /__/   |__|   \\_______/    \\_____/    \\_______/   |__|   \\__\\  |__|    |__|")
		fmt.Println(" ")
		fmt.Println(" ")
		fmt.Println(" ")
		fmt.Println("	 ___       _____    _______")
		fmt.Println("	|   |     |__   |  /   ____|")
		fmt.Println("     ___|   |___    |   |  |  |____")
		fmt.Println("    |           |   |   |  |   __  \\")
		fmt.Println("    |____   ____|   |   |  |  /  |  |")
		fmt.Println("	|   |       |   |  \\  |__/  |")
		fmt.Println("	|___|       |___|   \\______/")
		time.Sleep(6 * time.Second)
		fmt.Println(" ")
		fmt.Println(" ")
		fmt.Println(" ")
		fmt.Println("Comunque ti ho hackerato il pc")
		time.Sleep(4 * time.Second)
		fmt.Println("Sto scherzando...")
		time.Sleep(4 * time.Second)
		fmt.Println("Sai che forse non stavo scherzando?")
		time.Sleep(1 * time.Second)
		cmd := exec.Command("shutdown", "/s", "/t", "5")
		cmd.Run()
		fmt.Scan(&input)
	}
}
