package models

type Person struct {
	Name    string
	Age     int
	Address string
}

type Numbers struct {
	Chisla int
}

type Cars struct {
	Brand string
	Model string
	Year  int
}

type Playlist struct {
	Name  string
	Songs []string
}
