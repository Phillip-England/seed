package main

import (
	"fmt"

	"github.com/Phillip-England/mood"
)

func main() {

	app := mood.New()

	app.SetDefault(func(app *mood.Mood) error {
		fmt.Println("seed - generate skeleton projects with ease")
		fmt.Println("run 'seed plant' to get started")
		return nil
	})

	app.At("plant", func(app *mood.Mood) error {

		if app.HasArg("server") {
			fmt.Println("build server")
			return nil
		}

		return nil
	})

	err := app.Run()
	if err != nil {
		panic(err)
	}

}
