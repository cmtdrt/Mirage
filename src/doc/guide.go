package doc

import (
	_ "embed"
	"fmt"
	"log"
	"os"
)

//go:embed guides/guide-en.md
var guideEN []byte

//go:embed guides/guide-fr.md
var guideFR []byte

const (
	GUIDE_OUTPUT_EN = "mirage-guide-en.md"
	GUIDE_OUTPUT_FR = "mirage-guide-fr.md"
)

func GenerateGuide(lang string) {
	var content []byte
	var filename string

	switch lang {
	case "en":
		content = guideEN
		filename = GUIDE_OUTPUT_EN
	case "fr":
		content = guideFR
		filename = GUIDE_OUTPUT_FR
	default:
		log.Fatalf("invalid guide language: %s", lang)
	}

	err := os.WriteFile(filename, content, 0644)
	if err != nil {
		log.Fatalf("failed to write guide file: %v", err)
	}
	fmt.Printf("Created guide: %s\n", filename)
}
