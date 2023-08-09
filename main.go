package main

import (
	"fmt"
	"iamchurka/models"
)

func Task1(data1 []models.Person) map[string]int {
	r := make(map[string]int)
	for _, v := range data1 {
		r[v.Name] = v.Age
	}
	return r
}

func Task2(numberSlice []models.Numbers) models.Numbers {
	result := models.Numbers{}
	for _, v := range numberSlice {
		result.Chisla += v.Chisla
	}
	return result
}

func Task3(rides []models.Cars, sort int) []string {
	r := []string{}
	for _, v := range rides {
		if v.Year == sort {
			r = append(r, v.Brand, v.Model)
		}
	}
	return r
}

func Task4(muzika []models.Playlist) {
	addSong := "Tyler, The Creator - sweet, i thought you wanted to dance"
	for _, v := range muzika {
		v.Songs = append(v.Songs, addSong)
		fmt.Println(v.Name, v.Songs)
	}
}

func main() {
	person1 := models.Person{"Kevin", 25, "Shabani"}
	person2 := models.Person{"Olezha", 23, "Kamenskoe getto"}
	data1 := []models.Person{}
	data1 = append(data1, person1)
	data1 = append(data1, person2)
	fmt.Println(Task1(data1)) ///////////////////////// Первая таска типа

	numberSlice := []models.Numbers{{2}, {5}, {7}, {6}}
	fmt.Println(Task2(numberSlice)) //////////////////////////// вторая таска типа

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
	fmt.Println(Task3(rides, sort)) ///////////////////////// третья таска типа

	muzika := []models.Playlist{
		{
			Name:  "music for sex",
			Songs: []string{"The Prodigy - The warriors dance", "Tyler, The crator - Are we still friends?"},
		},
	}
	fmt.Println(muzika)
	Task4(muzika)
}
