package requests

import "formation-go/lib/types"

type PersonRequest struct {
	Civility types.CIVILITY `json:"civility"`
	Name     string         `json:"name"`
	Age      int            `json:"age"`
}
