package main

import (
	"fmt"
	"log"

	"golang.org/x/tools/go/packages"
)

func main() {
	pkgs, err := packages.Load(&packages.Config{
		Mode: packages.NeedName | packages.NeedTypes | packages.NeedTypesInfo | packages.NeedModule | packages.NeedImports,
	}, "./lib/generics")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pkgs)
}
