package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func askPseudo() (string, error) {
	lecteur := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Pseudo:")
		ligne, err := lecteur.ReadString('\n')
		if err != nil {
			return "", fmt.Errorf("lecture du pseudo : %w", err)
		}
		pseudo := strings.TrimSpace(ligne)
		if pseudo != "" {
			return pseudo, err
		}
		fmt.Println("Erreur : le pseudo ne peut pas être vide.")
	}
}
