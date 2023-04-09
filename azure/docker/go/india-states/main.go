package main

import (
	"fmt"
)

type State struct {
	Name      string
	Districts []string
}

var states = []State{
	{
		Name: "Andhra Pradesh",
		Districts: []string{
			"Anantapur",
			"Chittoor",
			"East Godavari",
			"Guntur",
			"Krishna",
			"Kurnool",
			"Prakasam",
			"Srikakulam",
			"Sri Potti Sriramulu Nellore",
			"Visakhapatnam",
			"Vizianagaram",
			"West Godavari",
			"YSR Kadapa",
		},
	},
	{
		Name: "Arunachal Pradesh",
		Districts: []string{
			"Tawang",
			"West Kameng",
			"East Kameng",
			"Papum Pare",
			"Kurung Kumey",
			"Kra Daadi",
			"Lower Subansiri",
			"Upper Subansiri",
			"West Siang",
			"East Siang",
			"Siang",
			"Upper Siang",
			"Lower Siang",
			"Lower Dibang Valley",
			"Dibang Valley",
			"Anjaw",
			"Lohit",
			"Namsai",
			"Changlang",
			"Tirap",
			"Longding",
		},
	},
	// Add more states here...
}

func main() {
	for i, state := range states {
		fmt.Printf("%d. %s\n", i+1, state.Name)
	}

	var input int
	fmt.Print("Enter a state number to view its districts: ")
	_, err := fmt.Scan(&input)
	if err != nil || input < 1 || input > len(states) {
		fmt.Println("Invalid input")
		return
	}

	state := states[input-1]
	fmt.Printf("Districts of %s:\n", state.Name)
	for _, district := range state.Districts {
		fmt.Println(district)
	}
}
