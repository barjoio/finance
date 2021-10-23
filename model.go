package main

type Payment struct {
	ID         int
	Year       int
	Month      int
	Day        int
	Hour       int
	Minute     int
	CategoryID int
	Category   Category
	Value      int
}

type Category struct {
	ID        int
	Name      string
	Projected int
}
