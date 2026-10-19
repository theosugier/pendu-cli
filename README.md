# pendu-cli

Jeu du pendu en ligne de commande, en Go (bibliotheque standard uniquement). TP Ynov B1, binome Theo Sugier & Amine Kouaidi.

## Lancer

```
go run . jouer       # nouvelle partie
go run . reprendre   # reprend la partie sauvegardee (save.json)
go run . score       # historique des parties (scores.csv)
```

Pendant une partie, `!save` sauvegarde et quitte.

## Fichiers

| Fichier | Role |
|---|---|
| `main.go` | FT-01 commandes CLI |
| `player.go` | FT-02 pseudo |
| `words.go` | FT-03 lecture de mot.txt + tirage |
| `game.go` | FT-04 / FT-05 boucle de jeu |
| `display.go` | affichage d'un tour |
| `scores.go` | FT-06 scores.csv |
| `save.go` | FT-07 save.json |
| `mot.txt` | liste des mots (un par ligne) |
