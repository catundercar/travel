package main

type User struct {
	Name string `json:"name"`
	Like Like   `json:"like"`
}

type Like struct {
	Sport string `json:"sport"`
}
