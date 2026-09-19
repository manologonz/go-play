package main

import (
	"errors"
	"fmt"
	"log"

	"charm.land/huh/v2"
)

var (
	burger       string
	toppings     []string
	sauceLevel   int
	name         string
	instructions string
	discount     bool
)

func StartSimpleHuhForm() {

	// Options:
	//naahh
	burguerSelect := huh.NewSelect[string]().Title("Choose your burger").Options(
		huh.NewOption("Charmburguer Classic", "classic"),
		huh.NewOption("Chickwich", "chickwich"),
		huh.NewOption("Fishburger", "fishburger"),
		huh.NewOption("Meh Burger", "mehburger"),
	).Value(&burger)

	toppingsSelect := huh.NewMultiSelect[string]().Title("Toppings").Options(
		huh.NewOption("Lettuce", "lettuce").Selected(true),
		huh.NewOption("Tomatoes", "tomatoes").Selected(true),
		huh.NewOption("Jalapeños", "jalapeños"),
		huh.NewOption("Cheese", "cheese"),
		huh.NewOption("Vegan Cheese", "vegan cheese"),
		huh.NewOption("Nutella", "nutella"),
	).Limit(4).Value(&toppings)

	sauceOption := huh.NewSelect[int]().Title("How much Sauce do yo want?").Options(
		huh.NewOption("None", 0),
		huh.NewOption("A little", 1),
		huh.NewOption("A lot", 2),
	).Value(&sauceLevel)

	nameOption := huh.NewInput().Title("What's your name?").Value(&name).Validate(func(str string) error {
		if str == "Frank" {
			return errors.New("Sorry, we don't serve customers named Frank.")
		}
		return nil
	})

	instructionsOption := huh.NewText().Title("Special Instructions").CharLimit(400).Value(&instructions)

	discountOption := huh.NewConfirm().Title("Would you like 15% off?").Value(&discount)

	// Group
	burguerOptions := huh.NewGroup(
		burguerSelect,
		toppingsSelect,
		sauceOption,
	)

	otherOptions := huh.NewGroup(
		nameOption,
		instructionsOption,
		discountOption,
	)

	// Form
	form := huh.NewForm(burguerOptions, otherOptions)

	err := form.Run()

	if err != nil {
		log.Fatal(err)
	}

	if !discount {
		fmt.Println("What? You didn't take the discout")
	}
}
