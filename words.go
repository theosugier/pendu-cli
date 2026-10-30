package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"strings"
)

// regarde si le mot est bon
func motValide(mot string) bool {
	nbLettres := 0
	for _, c := range mot {
		if c >= 'a' && c <= 'z' {
			// sa c une lettre
			nbLettres++
		} else if c == '-' {
			// le tire est ok
		} else {
			// pas bon sa
			return false
		}
	}
	// faut au moin une lettre
	if nbLettres > 0 {
		return true
	} else {
		return false
	}
}

// lit le fichier et done les mots
func LireMots(nomFichier string) ([]string, error) {
	contenu, err := os.ReadFile(nomFichier)
	if err != nil {
		return nil, err
	}

	var mots []string
	lignes := strings.Split(string(contenu), "\n")
	for _, ligne := range lignes {
		// on enleve les espaces et les magiscules
		mot := strings.TrimSpace(ligne)
		mot = strings.ToLower(mot)
		if motValide(mot) {
			// mot bon donc on le met dans la liste
			mots = append(mots, mot)
		}
	}

	if len(mots) == 0 {
		return nil, fmt.Errorf("aucun mot valide")
	} else {
		return mots, nil
	}
}

// prend une letre au hasard
func ChoisirLettre(mot string) string {
	// on garde just les lettres sans les tires
	lettres := ""
	for _, c := range mot {
		if c != '-' {
			lettres = lettres + string(c)
		}
	}
	// letre au hasard
	i := rand.IntN(len(lettres))
	return string(lettres[i])
}

// prend un mot au hasard et une letre pour commencer
func NouveauMot() (string, string, error) {
	mots, err := LireMots("mot.txt")
	if err != nil {
		return "", "", err
	}
	// un mot au hasard dans la liste
	i := rand.IntN(len(mots))
	mot := mots[i]
	lettre := ChoisirLettre(mot)
	return mot, lettre, nil
}
