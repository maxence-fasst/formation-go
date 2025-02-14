package responses

type (
	ReferentielItem struct {
		Value               string            `json:"value"`
		Label               string            `json:"label"`
		ParameterIdentifier string            `json:"-"`
		Metadata            map[string]string `json:"metadata"`
	}
)
