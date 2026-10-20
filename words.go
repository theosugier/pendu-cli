package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// regarde si le mot a que des lettres minuscule et des tirets
func motValide(mot string) bool {
	lettres := 0
	for _, c := range mot {
		if c >= 'a' && c <= 'z' {
			lettres++
		} else if c != '-' {
			return false
		}
	}
	return lettres > 0
}

// ouvre le fichier et recupere tout les mots qui sont bon
func LireMots(nomFichier string) ([]string, error) {
	contenu, err := os.ReadFile(nomFichier)
	if err != nil {
		return nil, err
	}

}
