package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mateuspmello/testes/jfilho"
)

const dbFileName = "game.db.json"

func main() {
	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")

	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}

	store := jfilho.NewFileSystemPlayerStore(db)

	game := jfilho.CLI{store, os.Stdin}
	game.PlayPoker()
}
