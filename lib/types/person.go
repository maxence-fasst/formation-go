package types

type CIVILITY = string

const (
	CIVILITY_MR  CIVILITY = "mr"
	CIVILITY_MRS CIVILITY = "mrs"
)

type Person struct {
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Civility  CIVILITY `json:"civility"`
	Age       int      `json:"age"`
}
