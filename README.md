Jeu du pendu en ligne de commande

TP Ynov B1 en Go fait en binome par Theo Sugier et Amine Kouaidi.

Le programme choisit un mot au hasard dans mot.txt et le joueur doit le deviner lettre par lettre. Il 6 essais. Il peut aussi proposer le mot entier (2 essais perdus si c'est faux) ou taper !save pour sauvegarder la partie et quitter (pas encore fait).

Prerequis
Go 1.22 ou plus recent.

Installation
git clone https://github.com/theosugier/pendu-cli.git
cd pendu-cli

Utilisation
go run . jouer       lance une nouvelle partie
go run . reprendre   reprend la partie sauvegardee (pas encore fait)
go run . score       affiche l'historique des parties

Fichiers
mot.txt : la liste des mots, un par ligne
save.json : la partie sauvegardee (cree par !save, supprime apres reprendre)
scores.csv : l'historique des parties terminees
