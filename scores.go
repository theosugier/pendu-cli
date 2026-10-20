package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

// enregistre le score du joueur a la fin de la partie dans scores.csv
func EnregistrerScore(pseudo string, statut string, essais int, mot string, lettres []string) error {
	// regarde si le fichier existe deja puis l'ouvre sinon il le creer
	_, errStat := os.Stat("scores.csv")
	f, err := os.OpenFile("scores.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("impossible d'ouvrir scores.csv : %w", err)
	}
}
