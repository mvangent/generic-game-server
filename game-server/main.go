package main

import (
	"embed"
	"fmt"
	"log"
	"strings"

	"github.com/mvangent/greetings"
)

var (
	Version string = strings.TrimSpace(version)
	//go:embed version.txt
	version string
	//go:embed static/*
	files embed.FS
)

func print_index() {
	data, _ := files.ReadFile("static/index.html")
	fmt.Printf("%s\n", string(data))
}

func main() {
	print_index()

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Melvin")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(message)
}
