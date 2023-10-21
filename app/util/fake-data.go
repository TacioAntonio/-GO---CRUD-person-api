package util

type Person struct {
	Id        string
	Avatar    string
	Firstname string
	Lastname  string
	Email     string
	Age       int
}

var Persons = []Person{
	{
		Id:        "1",
		Avatar:    "https://cdn-icons-png.flaticon.com/512/147/147142.png",
		Firstname: "Jorge",
		Lastname:  "Milcox",
		Email:     "Jorge@gmail.com",
		Age:       1,
	},
}
