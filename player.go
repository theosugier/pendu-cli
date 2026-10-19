package main

import (
	"bufio"
	"fmt"
	"strings"
	"os"
)

func askPseudo() string {
	fmt.Print("Pseudo:")
	lecteur := bufio.NewReader(os.Stdin)
	ligne, _ := lecteur.ReadString('\n')	
	pseudo := strings.TrimSpace(ligne)
	return pseudo
} 
