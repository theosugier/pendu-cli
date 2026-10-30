package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// enregistre le score à la fin de la partie
func EnregistrerScore(pseudo string, statut string, essais int, mot string, lettres []string) error {
	// on lit l'ancien fichier
	ancien, err := os.ReadFile("scores.csv")
	texte := ""
	if err != nil {
		// le fichier n'existe pas encore, on part de zéro
		texte = ""
	} else {
		// le fichier existe, on garde son contenu
		texte = string(ancien)
	}

	// la date d'aujourd'hui (ce format est toujours le même en Go)
	date := time.Now().Format("2006-01-02 15:04")

	// on prépare la nouvelle ligne avec des virgules
	ligne := fmt.Sprintf("%s,%s,%d,%s,%s,%s\n", pseudo, statut, essais, mot, strings.Join(lettres, ""), date)

	// on ajoute la ligne à la fin et on écrit le fichier
	texte = texte + ligne
	return os.WriteFile("scores.csv", []byte(texte), 0644)
}

// affiche tous les scores
func AfficherScores() error {
	contenu, err := os.ReadFile("scores.csv")
	if err != nil {
		// pas de fichier = pas encore de parties
		fmt.Println("aucune partie")
		return nil
	}

	fmt.Println("Pseudo statut essais mot lettres date")

	// on coupe le fichier ligne par ligne
	lignes := strings.Split(string(contenu), "\n")
	for _, ligne := range lignes {
		// on coupe la ligne à chaque virgule
		l := strings.Split(ligne, ",")
		if len(l) < 6 {
			// ligne vide ou abîmée, on la saute
			continue
		} else {
			fmt.Println(l[0], l[1], l[2], l[3], l[4], l[5])
		}
	}
	return nil
}
