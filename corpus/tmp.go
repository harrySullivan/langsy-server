package main

import (
	"fmt"
	"gopkg.in/jdkato/prose.v2"
	"github.com/gertd/go-pluralize"
)

func main() {
	doc, _ := prose.NewDocument("The chairs are red")
	fmt.Println(doc.Tokens())
	client := pluralize.NewClient()
	fmt.Println(client.Pluralize("chair", 4, true))
}

