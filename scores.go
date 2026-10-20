package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// enregistre le score a la fin de la partie
func EnregistrerScore(pseudo string, statut string, essais int, mot string, lettres []string) error {
	f, err := os.OpenFile("scores.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	// on ferme le fichier a la fin de la fonction
	defer f.Close()

	date := time.Now().Format("2006-01-02 15:04")
	// on ecrit la ligne dans le fichier avec des virgules entre chaque info
	_, err = fmt.Fprintf(f, "%s,%s,%d,%s,%s,%s\n", pseudo, statut, essais, mot, strings.Join(lettres, ""), date)
	return err
}

//affiche tous les scores qui sont dans le fichier
func AfficherScores() error {
		contenu, err := os.ReadFile("scores.csv")
		if err != nil {
			// si le ficjhier existe pas y a  pas encore de parties
			fmt.Println("aucune partie")
			return nil
		}

		fmt.Println("Pseudo statut essais mot date")
		//on coupe le fichier ligne par ligne
		lignes := strings.Split(string(contenu), "\n")
		for _, ligne := range lignes {
				// on coupe la ligne a chaque virgule 
				l := strings.Split(ligne, ",")
				//si la ligne est vide ou abime on la saute
				if len(l) < 6 {
						continue
				}
				fmt.Println(l[0], l[1], l[2], l[3], l[4], l[5])
		}
		return nil

}	
	