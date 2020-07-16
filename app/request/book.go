package request

type CreateBookRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UpdateBookRequest struct {
	ID   string `json:"mobile"`
	Name string `json:"password"`
}
