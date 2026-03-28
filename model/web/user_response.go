package web

type UserResponse struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Age     uint8  `json:"age"`
	Address string `json:"address"`
}
