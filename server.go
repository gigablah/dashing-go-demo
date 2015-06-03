package main

import (
	"github.com/gigablah/dashing-go"
	_ "github.com/gigablah/dashing-go-demo/jobs"
)

func main() {
    dashing.Start()
}
