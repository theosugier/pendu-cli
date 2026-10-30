package main

import (
	"fmt"
	"strings"
)

// ProposerMot vérifie que le joueur qui a mis mot est correct
func ProposerMot(proposition string, mot string, essais int) (int, bool, error) {
	proposition = strings.ToLower(strings.TrimSpace(proposition))

	//n'autorise pas à tenter un mot si il y a moins de deux essais
	if essais < 2 {
		return essais, false, fmt.Errorf("il faut au moins 2 essais pour proposer un mot")
	}

	//refuse le mot mis si ce n'est pas le bon
	if len(proposition) != len(mot) {
		return essais, false, fmt.Errorf("le mot doit contenir %d lettres", len(mot))
	}

	//obtention de la victoire en cas de trouveent du mot caché
	if proposition == mot {
		return essais, true, nil
	}

	//retire 2essais en cas de faute et envoie un "false"
	return essais - 2, false, nil
}

type Sauvegarde struct { // fiche d'une partie, utile pour jouer + sauvegarder en json
	Pseudo  string   `json:"pseudo"`
	Mot     string   `json:"mot"`
	Lettres []string `json:"lettres"`
	Essais  int      `json:"essais"`
}

func jouerPartie(partie Sauvegarde) {
	for partie.Essais > 0 { // tant quon a des essaies
		fmt.Println(AfficherMot(partie.Mot, strings.Join(partie.Lettres, ""), false)) // affiche le mot avec les lettres trouvées
		fmt.Println("Essais restants :", partie.Essais)
		lettres := strings.Join(partie.Lettres, " ") // affiche lettres séparer despaces sinon sa serais dnas tableau
		fmt.Println("Lettres proposées :", lettres)
		fmt.Print("> ")
		ligne, err := lecteur.ReadString('\n')
		if err != nil {
			fmt.Println("Erreur :", err)
			return
		}
		saisie := strings.ToLower(strings.TrimSpace(ligne))
		if saisie == "" {
			fmt.Println("Erreur : saisissez une lettre, un mot ou !save.")
			continue
		}
		deja := false
		for _, l := range partie.Lettres {
			if l == saisie {
				deja = true
			}
		}
		if deja {
			fmt.Println("Erreur : la lettre", saisie, "a déjà été proposée.")
			continue
		}
		if len(saisie) == 1 && (saisie[0] < 'a' || saisie[0] > 'z') {
			fmt.Println("Erreur : saisissez une lettre, un mot ou !save.")
			continue
		}
		if saisie == "!save" {
			fmt.Println("La sauvegarde n'est pas encore disponible.")
			continue
		}
		if len(saisie) > 1 {
			essais, gagne, err := ProposerMot(saisie, partie.Mot, partie.Essais)
			if err != nil {
				fmt.Println("Erreur :", err)
				continue
			}
			partie.Essais = essais
			if gagne {
				fmt.Println("Bravo", partie.Pseudo, "le mot: ", partie.Mot)
				return
			}
			fmt.Println("Rate ce n'est pas le mot.")
			continue
		}
		partie.Lettres = append(partie.Lettres, saisie)
		if strings.Contains(partie.Mot, saisie) == true {
			fmt.Println("Bien joué, la lettre", saisie, "est dans le mot.")
		} else {
			fmt.Println("Rate la lettre", saisie, "n'est pas dans le mot.")
			partie.Essais-- // on perd un essaie
		}
		if MotTrouve(partie.Mot, strings.Join(partie.Lettres, "")) {
			fmt.Println("Bravo", partie.Pseudo, "le mot: ", partie.Mot)
			return
		}
	}
	fmt.Println("Perdu", partie.Pseudo, "le mot etait", partie.Mot)
}
