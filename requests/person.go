package requests

type PersonRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
