package config

import (
	"fmt"
	"time"

	"github.com/schollz/progressbar/v3"
)

func StartingApps() {
	fmt.Println()
	fmt.Println("Keasikan Ngoding Taunya Udah Jam 3 Pagi")
	fmt.Println("Me :")
	PrintSurprisedPikachuMeme()
	fmt.Println()
	fmt.Println("(Wait ya sembari liat karya ini WKWK)")
	Progress()
	fmt.Println()
	fmt.Println("Developed By Moch Nurfaizal")
}
func Progress() {
	bar := progressbar.Default(100)
	for i := 0; i < 100; i++ {
		bar.Add(1)
		time.Sleep(10 * time.Millisecond)
	}
}

func PrintSurprisedPikachuMeme() {
	fmt.Println("⠀⠀⠀⣠⣾⢿⣻⣟⣿⣻⣿⣿⣿⣿⣿⣻⣟⣿⣻⣟⣿⣻⣟⣿⣻⣟⣿⣻⣟⣿⣻⣟⣿⣻⣿⢏⠆⡀⠀⠀⠀⠀⠀⠀⡀⢀⠀⡄⢤⣠")
	fmt.Println("⣴⣶⣿⣻⣽⣻⢷⣻⣾⢿⣿⣿⣿⣿⣷⣟⣾⣳⣟⣾⣳⣟⣾⣳⣟⣾⣳⣟⣾⣳⣟⣾⣳⣟⣾⣏⠖⠀⠀⠀⠀⠀⠌⠐⡀⢎⡼⣸⢧⣷")
	fmt.Println("⣿⣳⣯⢷⣯⣟⣯⣷⢯⣿⣿⣿⣿⣿⣿⣾⣳⣯⣟⣾⣽⢾⣯⢷⣻⢷⣯⣟⣾⡽⣾⢯⣷⣻⣽⣞⡌⠀⠀⠀⠀⠐⡀⠡⡘⢆⡳⣭⢿⣿")
	fmt.Println("⣿⣳⣯⢿⡾⣽⣞⣿⣻⣞⡿⣽⣻⣟⣿⣷⢯⣷⢯⣷⢯⡿⣞⡿⣯⢿⡾⣽⣞⡿⣽⣻⡾⣽⣞⣿⠰⢀⠠⠀⠐⡀⠄⢡⠘⣬⠳⣭⣟⣿")
	fmt.Println("⣿⣳⣯⢿⣽⢯⣿⣿⣿⡾⣙⢧⡛⣼⣳⢿⡿⣽⣻⢾⣻⣽⢿⣽⣻⢯⣟⣷⢯⣟⣯⡷⣟⡿⣞⡿⣏⠤⡐⠠⢁⠠⡈⣄⠫⡔⣹⠲⣯⣿")
	fmt.Println("⣿⣳⣯⢿⣾⣻⡿⣿⣿⡇⠀⠈⠑⠧⣯⣟⣿⢯⣟⡿⣽⣾⣻⢾⣽⣻⣽⢾⣻⣽⣳⢿⣯⣟⡿⣽⣿⣳⣼⣱⣮⣵⣱⣦⣳⡼⣴⣯⢷⣿")
	fmt.Println("⣿⣳⣯⢿⣞⡟⠰⣄⢻⡇⠀⠀⠀⠀⠀⠙⠻⣿⣽⣻⢷⣯⣟⣯⡷⣿⣽⣻⣽⣞⣯⡿⣾⣽⣻⢷⣻⣿⣿⣿⣿⣿⡿⠿⠛⠉⠁⠀⢸⣿")
	fmt.Println("⣿⣳⣯⢿⣾⠡⡓⣌⢖⣿⡀⠀⠀⠀⠀⠀⠀⠀⠙⠫⣿⡾⣽⣞⣿⣳⣯⢷⣻⡾⣽⣻⢷⣯⣟⣯⢿⣿⠿⠟⠋⠁⠀⠀⠀⠀⠀⠀⣾⣿")
	fmt.Println("⣿⣳⣯⢿⢣⠱⡘⡬⣞⣾⣷⡄⠀⠀⠀⠀⠀⠀⠀⠀⠈⠻⣟⠾⠳⠿⠾⠿⠽⠟⠷⠻⠿⠾⠽⠚⠉⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣸⣿⣿")
	fmt.Println("⣿⣳⣯⣟⣦⢳⣱⣣⣟⣾⣿⣿⣦⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣴⣿⣿⣿")
	fmt.Println("⣿⣳⣿⣿⣾⣿⣷⣿⣿⣿⣿⣿⣿⣷⣄⠀⢀⠄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⠞⡹⢎⡷⣿")
	fmt.Println("⣿⢯⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡦⠊⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠦⣀⣠⣶⢋⡴⢣⡞⣷⣿")
	fmt.Println("⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠁⠀⠀⠀⢀⣠⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣤⣀⠀⠀⠀⠀⠀⠘⣿⣎⡷⣜⣧⣟⣷⣿")
	fmt.Println("⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣻⠇⠀⠀⠀⠀⣯⣬⣿⡆⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣯⣼⣿⠄⠀⠀⠀⠀⠀⠸⣯⣿⣽⣾⣽⣾⣿")
	fmt.Println("⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣯⡟⠀⠀⠀⠀⠀⠙⠛⠛⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠙⠛⠋⠀⠀⠀⠀⠀⠀⠀⢻⣿⣿⣿⣿⣿⣿")
	fmt.Println("⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⢯⠃⣠⣤⣤⣀⠀⠀⠀⠀⠀⠀⠀⠘⠛⠃⠀⠀⠀⠀⠀⠀⠀⠀⢀⣠⣤⣤⡀⠀⠀⠘⣿⣿⣿⣿⣿⣿")
	fmt.Println("⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⡿⣽⣻⢰⡟⣶⢳⢯⠆⠀⠀⠀⠀⠀⠀⢀⣀⣀⣀⡀⠀⠀⠀⠀⠀⠀⣾⢧⣛⣶⢻⡆⠀⠀⢹⡿⣽⢾⣿⣿")
	fmt.Println("⣿⣟⣯⢿⡹⢯⡝⢾⡱⢏⡳⢣⢽⠀⠙⠉⠛⠋⠀⠀⠀⠀⠀⠀⢰⡋⡔⠣⡜⠹⡄⠀⠀⠀⠀⠀⠙⠯⠽⠺⠙⠀⠀⠀⠈⣟⡼⣛⡾⣿")
	fmt.Println("⣿⣽⡚⢦⡙⢆⢎⡱⢌⠣⡜⡡⢎⣇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠳⣌⣱⣌⡳⠃⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢣⡜⣱⢟⣿")
	fmt.Println("⣿⡶⣏⢧⡙⣎⢦⡑⣎⡱⢬⡱⢎⣞⣆⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠸⣜⣣⢿⣿")
	fmt.Println("⣿⣽⣻⢮⡽⣜⡦⣝⡦⣝⡶⣹⣞⣼⣻⡆⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣷⢯⣿⣿")
	fmt.Println("⢷⣫⣟⣯⣿⣽⣻⣽⣿⣽⣻⣷⣿⣾⣿⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢿⣿⣿⣿")
}
