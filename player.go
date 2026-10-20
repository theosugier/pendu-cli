package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var lecteur = bufio.NewReader(os.Stdin) // branche un truc qui enregistre le clavier

func askPseudo() (string, error) {
	for {
		fmt.Print("Pseudo:")                   // demande le pseudo
		ligne, err := lecteur.ReadString('\n') // lis tout se quon ecris juqua quon appuis sur entrée et renv texte  + retour ligne et renv err si besoin
		if err != nil {                        // si il y a une erreur
			return "", fmt.Errorf("lecture du pseudo : %w", err) // on retourne l'err
		}
		pseudo := strings.TrimSpace(ligne) // enleve les espace et le retour a la ligne debut + fin
		if pseudo != "" {                  // si pseudo pas vide
			return pseudo, err // on retourne pseudo + err si il y a
		}
		fmt.Println("Erreur : le pseudo ne peut pas être vide.") // sinon on dit que pseudo peux pas etre vide
	}
}
