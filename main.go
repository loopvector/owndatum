package main

import (
	"github.com/loopvector/owndatum/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
