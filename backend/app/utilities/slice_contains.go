package utilities

func SliceContains(slice []string, target string) bool {

	for _, value := range slice {
		if value == target {
			return true
		}
	}

	return false
}
