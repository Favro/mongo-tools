package main

import (
	"github.com/10gen/mongoplay"
	"github.com/jessevdk/go-flags"

	"os"
)

func main() {
	opts := mongoplay.Options{}
	var parser = flags.NewParser(&opts, flags.Default)
	parser.AddCommand("play", "Play captured traffic against a mongodb instance", "",
		&mongoplay.PlayCommand{GlobalOpts: &opts})
	parser.AddCommand("record", "Convert network traffic into mongodb queries", "",
		&mongoplay.RecordCommand{GlobalOpts: &opts})
	parser.AddCommand("stat", "Generate statistics on captured traffic", "",
		&mongoplay.StatCommand{GlobalOpts: &opts})

	_, err := parser.Parse()
	if err != nil {
		os.Exit(1)
	}
}
