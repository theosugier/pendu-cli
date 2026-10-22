package main

import (
	"fmt"
	"strings"
)

// ---------- couleurs ANSI ----------

const (
	reset = "\033[0m"
	gras  = "\033[1m"
	rouge = "\033[31m"
	vert  = "\033[32m"
	jaune = "\033[33m"
	cyan  = "\033[36m"
	gris  = "\033[90m"
)

// colorie un texte avec une couleur
func colorier(couleur, texte string) string {
	return couleur + texte + reset
}

// met un texte en vert (bonne lettre)
func Vert(texte string) string { return colorier(vert, texte) }

// met un texte en rouge (erreur)
func Rouge(texte string) string { return colorier(rouge, texte) }

// ---------- le dessin du pendu ----------

// un morceau de dessin : un texte placé à une ligne et une colonne, avec une couleur
type morceau struct {
	ligne   int
	col     int
	texte   string
	couleur string
}

// les 10 morceaux du pendu, dans l'ordre où ils apparaissent
var morceaux = [][]morceau{
	{{6, 0, "=========", jaune}}, // 1 : le sol
	{{1, 6, "|", jaune}, {2, 6, "|", jaune}, {3, 6, "|", jaune}, {4, 6, "|", jaune}, {5, 6, "|", jaune}}, // 2 : le poteau
	{{0, 2, "+---+", jaune}}, // 3 : la barre du haut
	{{1, 2, "|", gris}},      // 4 : la corde
	{{2, 2, "O", rouge}},     // 5 : la tête
	{{3, 2, "|", rouge}},     // 6 : le corps
	{{3, 1, "/", rouge}},     // 7 : bras gauche
	{{3, 3, "\\", rouge}},    // 8 : bras droit
	{{4, 1, "/", rouge}},     // 9 : jambe gauche
	{{4, 3, "\\", rouge}},    // 10 : jambe droite
}

// calcule combien de morceaux on montre, selon les erreurs et le nombre de vies
// (marche avec n'importe quel nombre de vies)
func nbMorceaux(erreurs, maxErreurs int) int {
	if maxErreurs <= 0 || erreurs <= 0 {
		return 0
	}
	if erreurs >= maxErreurs {
		return len(morceaux)
	}
	// arrondi vers le haut pour que chaque erreur se voie
	return (erreurs*len(morceaux) + maxErreurs - 1) / maxErreurs
}

// construit le dessin ligne par ligne
func DessinPendu(erreurs, maxErreurs int) string {
	const hauteur, largeur = 7, 9

	// la grille : chaque case a un caractère et une couleur
	car := make([][]string, hauteur)
	coul := make([][]string, hauteur)
	for l := 0; l < hauteur; l++ {
		car[l] = strings.Split(strings.Repeat(" ", largeur), "")
		coul[l] = make([]string, largeur)
	}

	// on place les morceaux
	n := nbMorceaux(erreurs, maxErreurs)
	for i := 0; i < n; i++ {
		for _, m := range morceaux[i] {
			for k, c := range m.texte {
				car[m.ligne][m.col+k] = string(c)
				coul[m.ligne][m.col+k] = m.couleur
			}
		}
	}

	// si c'est perdu, la tête devient un X
	if maxErreurs > 0 && erreurs >= maxErreurs {
		car[2][2] = "X"
	}

	// on assemble la grille en texte
	var sb strings.Builder
	for l := 0; l < hauteur; l++ {
		for c := 0; c < largeur; c++ {
			if coul[l][c] != "" {
				sb.WriteString(colorier(coul[l][c], car[l][c]))
			} else {
				sb.WriteString(car[l][c])
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

// ---------- les infos du jeu ----------

// lettres proposées sans doublon, en minuscules
func lettresUniques(lettres string) []rune {
	vues := map[rune]bool{}
	liste := []rune{}
	for _, r := range strings.ToLower(lettres) {
		if !vues[r] {
			vues[r] = true
			liste = append(liste, r)
		}
	}
	return liste
}

// compte les erreurs : les lettres proposées qui ne sont pas dans le mot
func CompterErreurs(mot, lettres string) int {
	erreurs := 0
	for _, r := range lettresUniques(lettres) {
		if !strings.ContainsRune(mot, r) {
			erreurs++
		}
	}
	return erreurs
}

// vrai si toutes les lettres du mot ont été trouvées
func MotTrouve(mot, lettres string) bool {
	for _, r := range mot {
		if r != '-' && !strings.ContainsRune(strings.ToLower(lettres), r) {
			return false
		}
	}
	return true
}

// le mot avec les lettres trouvées en vert et des _ pour les autres
// si revele est vrai (perdu), les lettres manquantes s'affichent en rouge
func AfficherMot(mot, lettres string, revele bool) string {
	trouvees := strings.ToLower(lettres)
	parts := []string{}
	for _, r := range mot {
		lettre := string(r)
		switch {
		case r == '-':
			parts = append(parts, colorier(cyan, lettre))
		case strings.ContainsRune(trouvees, r):
			parts = append(parts, colorier(gras+vert, strings.ToUpper(lettre)))
		case revele:
			parts = append(parts, colorier(gras+rouge, strings.ToUpper(lettre)))
		default:
			parts = append(parts, colorier(gris, "_"))
		}
	}
	return strings.Join(parts, " ")
}

// l'alphabet : vert = bonne lettre, rouge = erreur, gris = pas encore essayée
func AfficherAlphabet(mot, lettres string) string {
	trouvees := strings.ToLower(lettres)
	parts := []string{}
	for c := 'a'; c <= 'z'; c++ {
		lettre := strings.ToUpper(string(c))
		switch {
		case !strings.ContainsRune(trouvees, c):
			parts = append(parts, colorier(gris, lettre))
		case strings.ContainsRune(mot, c):
			parts = append(parts, colorier(gras+vert, lettre))
		default:
			parts = append(parts, colorier(gras+rouge, lettre))
		}
	}
	return strings.Join(parts, " ")
}

// la barre de vies : change de couleur quand il en reste peu
func BarreVies(erreurs, maxErreurs int) string {
	restant := maxErreurs - erreurs
	if restant < 0 {
		restant = 0
	}
	perdues := maxErreurs - restant
	couleur := vert
	if restant*3 <= maxErreurs*2 {
		couleur = jaune
	}
	if restant*3 <= maxErreurs {
		couleur = rouge
	}
	barre := colorier(couleur, strings.Repeat("█", restant)) + colorier(gris, strings.Repeat("░", perdues))
	return fmt.Sprintf("Vies : %s %d/%d", barre, restant, maxErreurs)
}

// un message dans un cadre
func Cadre(message, couleur string) string {
	ligne := strings.Repeat("═", len([]rune(message))+2)
	return colorier(couleur, "╔"+ligne+"╗\n║ "+message+" ║\n╚"+ligne+"╝")
}

// ---------- l'écran complet ----------

// construit tout l'écran du jeu
func Ecran(mot, lettres string, maxErreurs int) string {
	erreurs := CompterErreurs(mot, lettres)
	perdu := erreurs >= maxErreurs
	gagne := MotTrouve(mot, lettres)

	var sb strings.Builder
	sb.WriteString(colorier(gras+cyan, "=== JEU DU PENDU ===") + "\n\n")
	sb.WriteString(DessinPendu(erreurs, maxErreurs) + "\n")
	sb.WriteString(AfficherMot(mot, lettres, perdu) + "\n\n")
	sb.WriteString(AfficherAlphabet(mot, lettres) + "\n\n")
	sb.WriteString(BarreVies(erreurs, maxErreurs) + "\n")

	if gagne {
		sb.WriteString("\n" + Cadre("Bravo, tu as gagné !", vert) + "\n")
	} else if perdu {
		sb.WriteString("\n" + Cadre("Perdu ! Le mot était : "+strings.ToUpper(mot), rouge) + "\n")
	}
	return sb.String()
}

// efface le terminal puis affiche l'écran
func AfficherEcran(mot, lettres string, maxErreurs int) {
	fmt.Print("\033[H\033[2J")
	fmt.Print(Ecran(mot, lettres, maxErreurs))
}
