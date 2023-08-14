package main

import (
	"fmt"
	funcs "iamchurka/fuctions"
	"iamchurka/models"
)

func main() {
	person1 := models.Person{"Kevin", 25, "Shabani"}
	person2 := models.Person{"Olezha", 23, "Kamenskoe getto"}
	data1 := []models.Person{}
	data1 = append(data1, person1)
	data1 = append(data1, person2)
	fmt.Println(funcs.Task1(data1)) ///////////////////////// Первая таска типа

	numberSlice := []models.Numbers{{2}, {5}, {7}, {6}}
	fmt.Println(funcs.Task2(numberSlice)) //////////////////////////// вторая таска типа

	sort := 1995
	rides := []models.Cars{
		{
			Brand: "Toyota",
			Model: "Mark II JZ 100",
			Year:  1998,
		},
		{
			Brand: "BMW",
			Model: "M3 e46",
			Year:  2003,
		},
		{
			Brand: "Honda",
			Model: "Civic",
			Year:  1995,
		},
		{
			Brand: "Nissan",
			Model: "Cima Y32",
			Year:  1995,
		},
		{
			Brand: "Nissan",
			Model: "Laurel C33",
			Year:  1989,
		},
	}
	fmt.Println(funcs.Task3(rides, sort)) ///////////////////////// третья таска типа

	muzika := []models.Playlist{
		{
			Name:  "music for sex",
			Songs: []string{"The Prodigy - The warriors dance", "Tyler, The crator - Are we still friends?"},
		},
	}
	fmt.Println(muzika)
	funcs.Task4(muzika)
}
