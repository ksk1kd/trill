package utilities

func SliceRemove(slice []string, target string) []string {

	result := []string{}

	for _, value := range slice {
		if value != target {
			result = append(result, value)
		}
	}

	return result
}
