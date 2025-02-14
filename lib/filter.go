package lib

func filter[T any](slice []T, testFunction func(T) bool) *[]T {
	result := []T{} // Use empty slice instead of null slice to get [] in json response (instead of null)
	for _, element := range slice {
		if testFunction(element) {
			result = append(result, element)
		}
	}
	return &result
}
