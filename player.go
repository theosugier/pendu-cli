package main

import (
	"bufio"
	"fmt"
	"strings"
	"os"
)

func askPseudo() string {
	lecteur := bufio.NewReader(os.Stdin)
	for {  
		fmt.Print("Pseudo:")
		ligne, _ := lecteur.ReadString('\n')	
		pseudo := strings.TrimSpace(ligne)
		if pseudo != "" {
			return pseudo
		}
		fmt.Println("Erreur : le pseudo ne peut pas être vide.")
	}
} 
