package main

import (
	"fmt"
	"os"
)

type localeMeta struct{
	language string
	country  string
}

func main() {
	supportedLocale:=[]localeMeta{
		{"en-US", "United States"},
		{"fr-FR", "France"},
		{"es-ES", "Spain"},
	}

	if len(os.Args) < 3 {
		fmt.Println("Please provide a locale and country.")
		return
	}

	locale := localeMeta{
		language: os.Args[1],
		country:  os.Args[2],
	}

	for _, loc := range supportedLocale {
		if loc == locale {
			fmt.Printf("Locale: %s, Language: %s, Country: %s\n", loc, locale.language, locale.country)
		}
	}
}