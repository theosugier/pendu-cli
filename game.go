package main

import (
	"fmt"
	"strings"
)

func ProposerMot(proposition string, mot string, essais int) (int, bool, error) {
		proposition = strings.ToLower(strings.TrimSpace(proposition))

}

type Sauvegarde struct { // fiche d'une partie, utile pour jouer + sauvegarder en json
	Pseudo  string   `json:"pseudo"`
	Mot     string   `json:"mot"`
	Lettres []string `json:"lettres"`
	Essais  int      `json:"essais"`
}