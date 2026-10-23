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

// lit save.json puis le supprime
func ChargerSauvegarde() (Sauvegarde, error) {
	var partie Sauvegarde

	// on lit le fichier
	data, err := os.ReadFile("save.json")
	if err != nil {
		// %w garde l'erreur d'origine pour errors.Is dans main
		return partie, fmt.Errorf("lecture de save.json impossible : %w", err)
	}
	// on remet le JSON dans la partie
	err = json.Unmarshal(data, &partie)
	if err != nil {
		return partie, fmt.Errorf("save.json est illisible : %w", err)
	}

	// on supprime le fichier
	os.Remove("save.json")
	return partie, nil
}
