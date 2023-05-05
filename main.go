package main

// struct is like an object in javascript
// the backtick “ is a value pair which act as additional information.
// For example in this case, the ID key when encoding or decoding as json should be in key "id" instead of "ID"
type Movie struct {
	ID       string `json:"id"`
	Isbn     string `json:"isbn"`
	Title    string `json:"title"`
	Director *Director
}

type Director struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func main() {
	return
}
