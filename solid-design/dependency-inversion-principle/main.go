package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/basselalaraaj/design-patterns/solid-design/dependency-inversion-principle/client"
)

type person interface {
	GetQuote() (string, error)
}

func main() {
	k := client.NewKanye()
	q := getQuote(k)

	printQuote(q)
}

func getQuote(person person) string {
	q, err := person.GetQuote()
	if err != nil {
		log.Fatalln(err)
	}

	return q
}

func printQuote(quote string) {
	q := strings.ToLower(quote)
	fmt.Printf("%s.. think about it! \n", q)
}
