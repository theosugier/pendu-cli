package main

import (
	"encoding/json"
	"os"
)

// écrit la partie dans save.json
func Sauvegarder(partie Sauvegarde) error {
	// on met la partie en JSON
	data, err := json.Marshal(partie)
	if err != nil {
		return err
	}
	// on écrit dans le fichier
	return os.WriteFile("save.json", data, 0644)
}

// lit save.json puis le supprime
func ChargerSauvegarde() (Sauvegarde, error) {
	var partie Sauvegarde

	// on lit le fichier
	data, err := os.ReadFile("save.json")
	if err != nil {
		return partie, err
	}
	// on remet le JSON dans la partie
	err = json.Unmarshal(data, &partie)
	if err != nil {
		return partie, err
	}

	// on supprime le fichier
	os.Remove("save.json")
	return partie, nil
}
