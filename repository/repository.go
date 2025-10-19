package repository

type User struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Age     uint8  `json:"age"`
	Address string `json:"address"`
}

var Users = []User{
	{ID: 1, Name: "Rendi", Age: 21, Address: "Kebonsari"},
	{ID: 2, Name: "Heru", Age: 25, Address: "Jambangan"},
	{ID: 3, Name: "Sarah", Age: 30, Address: "Gayungan"},
}
