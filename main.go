package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage : go run . <jouer|reprendre|score>")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "jouer":
		fmt.Println("nouvelle partie")
		pseudo, err := askPseudo()
		if err != nil {
			fmt.Println("Erreur :", err)
			os.Exit(1)
		}
		fmt.Println("Bonjour:", pseudo)
	case "reprendre":
		fmt.Println("reprise de la partie")
	case "score":
		fmt.Println("scores:")
	default:
		fmt.Println("Usage : go run . <jouer|reprendre|score>")
		os.Exit(1)
	}
}
