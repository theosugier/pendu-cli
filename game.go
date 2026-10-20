package main

import (
	"fmt"
	"strings"
)

//ProposerMot vérifie que le joueur qui a mis mot est correct
func ProposerMot(proposition string, mot string, essais int) (int, bool, error) {
		proposition = strings.ToLower(strings.TrimSpace(proposition))

		//n'autorise pas à tenter un mot si il y a moins de deux essais
		if essais < 2 {
				return essais, false, fmt.Errorf("il faut au moins 2 essais en stock pour proposer de mettre un mot en entier si non je decline l'offre hahahahah")
		}

		//refuse le mot mis si ce n'est pas le bon 
		if len(proposition) != len(mot) {
				return essais, false, fmt.Errorf("le mot doit contenir %d lettres", len(mot))
		}

		//obtention de la victoire en cas de trouveent du mot caché
		if proposition == mot {
				return essais, true, nil
		}

		//retire 2essais en cas de faute et envoie un "false" 
		return essais -2, false, nil
}

type Sauvegarde struct { // fiche d'une partie, utile pour jouer + sauvegarder en json
	Pseudo  string   `json:"pseudo"`
	Mot     string   `json:"mot"`
	Lettres []string `json:"lettres"`
	Essais  int      `json:"essais"`
}
