package main

import (
	"github.com/Turgho/GoFlowDesk/internal/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		panic(err)
	}
}
