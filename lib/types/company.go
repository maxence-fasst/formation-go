package types

type (
	Company struct {
		Name      string   `json:"name" validate:"required"`
		Siren     string   `json:"siren" validate:"required,isValidSiren"`
		Employees []Person `json:"employees"`
	}
)
