package funcs

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
