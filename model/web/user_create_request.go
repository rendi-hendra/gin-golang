package web

type UserCreateRequest struct {
	Name    string `json:"name" binding:"required"`
	Age     uint8  `json:"age" binding:"required,gte=1,lte=120"`
	Address string `json:"address" binding:"required"`
}
