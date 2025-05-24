package main

import (
	"fmt"
	"os"

	"blockchain-project/api"
	"blockchain-project/cli"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "cli" {
		cli.Run()
	} else {
		fmt.Println("Starting HTTP API on :8081")
		r := api.SetupRouter()
		r.Run(":8081")
	}
}
