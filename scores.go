package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

func EnregistrerScore(pseudo string, statut string, essais int, mot string, lettres []string) error {
	_, errStat := os.Stat("scores.csv")
	f, err := os.OpenFile("scores.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("impossible d'ouvrir scores.csv : %w", err)
	}
}
