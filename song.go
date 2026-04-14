package main

import (
	"fmt"
	"time"
)

// Function: Ek-ek akshar (character) print karne ke liye
func typeWrite(text string, delay time.Duration) {
	for _, char := range text {
		fmt.Printf("%c", char)
		time.Sleep(delay)
	}
	fmt.Println()
}

func main() {
	// ANSI Color Codes (Terminal ko rangeen banane ke liye)
	colorReset := "\033[0m"
	colorRed := "\033[31m"
	colorYellow := "\033[33m"
	colorCyan := "\033[36m"
	colorBold := "\033[1m"

	fmt.Println(colorCyan + colorBold + "🎵 Playing: Kesariya (Brahmastra) 🎵" + colorReset)
	fmt.Println("------------------------------------------")
	time.Sleep(1 * time.Second)

	// Song Data: Line aur uska color/emoji
	lyrics := []struct {
		text  string
		color string
	}{
		{"Mujhko itna bata de koi... 🧐", colorReset},
		{"Kaise tujhse dil na lagaye koi... ❤️", colorRed},
		{"", colorReset},
		{"Rabba ne tujhko banane mein... ✨", colorYellow},
		{"Kardi hai husn ki khaali tijoriyaan... 💎", colorYellow},
		{"", colorReset},
		{"Kajal ki siyahi se likhi hai tune... 🖋️", colorCyan},
		{"Jaane kitno ki love storiyan... 📖", colorCyan},
		{"", colorReset},
		{"Kesariya tera ishq hai piya... 🧡", colorYellow},
		{"Rang jaaun jo main haath lagaun... 🎨", colorYellow},
		{"Din beete saara teri fikr mein... ☀️", colorRed},
		{"Rain saari teri khair manaun... 🌙", colorRed},
	}

	// Loop jo lyrics ko stylist tarike se print karega
	for _, line := range lyrics {
		if line.text == "" {
			fmt.Println() // Empty line ke liye
			time.Sleep(500 * time.Millisecond)
			continue
		}

		// Line print ho rahi hai thode style ke saath
		fmt.Print(line.color)
		typeWrite(line.text, 50*time.Millisecond) // Har akshar 50ms mein aayega

		time.Sleep(1 * time.Second) // Har line ke baad 1 second ka pause
	}

	fmt.Println(colorCyan + "\n------------------------------------------")
	fmt.Println("✨ Program Finished! Hope you liked it! ✨" + colorReset)
}
