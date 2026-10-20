package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"os"
	"strings"
)

//verfifie que le mot a que des lettre minisucules et des tirets avec au moins une lettre
func motValide(mot string) bool {
	nbLettres := 0
	for _, c := range mot {
		if c >= 'a' && c <= 'z' {
			nbLettres++
		} else if c != '-' {
			return false
		}
	}
	return nbLettres > 0
}

//lit le fichier et renvoit la liste des mots valides
func LireMots(nomFichier string) ([]string, error) {
		// lit le fichier en entier et renvoit une erreur si il existe pas ou si on peut pas le lire
	contenu, err := os.ReadFile(nomFichier)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, fmt.Errorf("le fichier %s n'existe pas : %w", nomFichier, err)
		}
		return nil, fmt.Errorf("impossible de lire %s : %w", nomFichier, err)
	}

	//regarde chaque ligne et fait le tris 
	var mots []string
	for _, ligne := range strings.Split(string(contenu), "\n") {
		mot := strings.ToLower(strings.TrimSpace(ligne))
		if motValide(mot) {
			mots = append(mots, mot)
		}
	}

	// si y'a aucun mot valide on renvoit une erreur sinon la liste
	if len(mots) == 0 {
		return nil, fmt.Errorf("aucun mot valide dans %s", nomFichier)
	}
	return mots, nil
}

// choisi une lettre au hasard dans le mot sans prendre les tirets
func ChoisirLettre(mot string) string {
	for {
		c := mot[rand.IntN(len(mot))]
		if c != '-' {
			return string(c)
		}
	}
}

// prend un mot au hasard dans mot.txt et renvoit le mot avec la lettre a reveler
func NouveauMot() (string, string, error) {
	mots, err := LireMots("mot.txt")
	if err != nil {
		return "", "", fmt.Errorf("choix du mot : %w", err)
	}
	mot := mots[rand.IntN(len(mots))]
	return mot, ChoisirLettre(mot), nil
}
