package main

import (
	"fmt"
	"strings"
)

func afficherMot(mot string, lettres []string) {
	for i := 0; i < len(mot); i++ { // je parcours chaque caractère du mot un par un
		if mot[i] == '-' { // si cest un tiré jaffiche direct ( mot composé )
			fmt.Print("- ")
		} else { // je verif si le car a deja ete proposé
			trouvé := false
			for _, l := range lettres {
				if l == strings.ToLower(string(mot[i])) { // je compare en min pour ignore la case 
					trouvé = true
				}
			}
			if trouvé { // si la lettre a ete trv je laffiche
				fmt.Print(strings.ToLower(string(mot[i])) + " ")
			} else {
				fmt.Print("_ ") // sinon je cache avec "_"
			}
		}
	}
	fmt.Println() // retour ligne a la fin 
}
