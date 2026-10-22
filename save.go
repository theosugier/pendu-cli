package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// écrit la partie dans save.jason
func Sauvegarder(partie Sauvegarde) error {
	// on transforme la partie en JSON
	data, err := json.Marshal(partie)
	if err != nil {
		return err
	}
	// on écrit dans le fichier ,ça remplace l'ancien
	return os.WriteFile("save.json", data, 0644)
}
