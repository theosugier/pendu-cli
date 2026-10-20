package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 { // verifie que dans la commande il y a bien 2 argument run + jouer ou reprendre ou score
		fmt.Println("Usage : go run . <jouer|reprendre|score>") // print les argu
		os.Exit(1)
	}
	switch os.Args[1] { // reagrde commadne taper et choisi en fonction
	case "jouer": // si on tape jouer alors
		fmt.Println("nouvelle partie")
		pseudo, err := askPseudo() // appel la fonction de player.go et renvoie pseudo + erreur si y'en a une
		if err != nil {            // si il y a une erreur on la print
			fmt.Println("Erreur :", err)
			os.Exit(1)
		}
		fmt.Println("Bonjour:", pseudo)                                                               // affiche le pseudo
		partie := Sauvegarde{Pseudo: pseudo, Mot: "porte-monnaie", Lettres: []string{"e"}, Essais: 6} // crée var partie qui prend la struct sauvegarde ( fiche vide ) et on la remplis
		jouerPartie(partie)                                                                           // appel jouerPartie pour lancer la parti
	case "reprendre": // si on tape reprendre alors
		fmt.Println("reprise de la partie")
	case "score": // si on tape score alors
		fmt.Println("scores:")
	default:
		fmt.Println("Usage : go run . <jouer|reprendre|score>") // result si mauvais saisis autre que les cases
		os.Exit(1)
	}
}
