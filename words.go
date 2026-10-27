package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// vérifie que le mot a seulement des lettres minuscules et des tirets
func motValide(mot string) bool {
	lettres := 0
	for _, c := range mot {
		if c >= 'a' && c <= 'z' {
			lettres++
		} else if c != '-' {
			return false
		}
	}
	// il faut au moins une lettre
	return lettres > 0
}

// lit le fichier et rend la liste des mots valides
func LireMots(nomFichier string) ([]string, error) {
	contenu, err := os.ReadFile(nomFichier)
	if err != nil {
		return nil, err
	}

	var mots []string
	for _, ligne := range strings.Split(string(contenu), "\n") {
		// on enlève les espaces et on met en minuscules
		mot := strings.ToLower(strings.TrimSpace(ligne))
		if motValide(mot) {
			mots = append(mots, mot)
		}
	}

	if len(mots) == 0 {
		return nil, fmt.Errorf("aucun mot valide")
	}
	return mots, nil
}

// choisit une lettre au hasard dans le mot (pas un tiret)
func ChoisirLettre(mot string) string {
	// on enlève les tirets du mot
	lettres := strings.ReplaceAll(mot, "-", "")
	// on prend une lettre au hasard
	i := rand.Intn(len(lettres))
	return string(lettres[i])
}

// choisit un mot au hasard et la lettre à montrer au début
func NouveauMot() (string, string, error) {
	mots, err := LireMots("mot.txt")
	if err != nil {
		return "", "", err
	}
	mot := mots[rand.Intn(len(mots))]
	return mot, ChoisirLettre(mot), nil
}
