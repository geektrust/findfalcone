package main

type Planet struct {
	Name     string `json:"name"`
	Distance int    `json:"distance"`
}

type Planets []Planet
