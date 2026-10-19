package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"os"
	"strings"
)

func motValide(mot string) bool {
	nbLettres := 0
	for _, c := range mot {
		if c >= 'a' && c <= 'z' {
			nbLettres++
		} else if c != '-' {
			return false
		}
	}
	return nbLettres > 0
}