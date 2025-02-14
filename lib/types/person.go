package types

type CIVILITY = string

const (
	CIVILITY_MR  CIVILITY = "mr"
	CIVILITY_MRS CIVILITY = "mrs"
)

type Person struct {
	Id        int
	FirstName string
	LastName  string
	Civility  CIVILITY
	Age       int
}
