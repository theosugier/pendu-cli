package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"os"
	"strings"
)

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

func LireMots(nomFichier string) ([]string, error) {
	contenu, err := os.ReadFile(nomFichier)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, fmt.Errorf("le fichier %s n'existe pas : %w", nomFichier, err)
		}
		return nil, fmt.Errorf("impossible de lire %s : %w", nomFichier, err)
	}

	var mots []string
	for _, ligne := range strings.Split(string(contenu), "\n") {
		mot := strings.ToLower(strings.TrimSpace(ligne))
		if motValide(mot) {
			mots = append(mots, mot)
		}
	}

	if len(mots) == 0 {
		return nil, fmt.Errorf("aucun mot valide dans %s", nomFichier)
	}
	return mots, nil
}
