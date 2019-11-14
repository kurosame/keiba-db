package service

// UniqStr is make stringSlice unique
func UniqStr(stringSlice []string) []string {
	keys := make(map[string]bool)
	uniq := []string{}
	for _, s := range stringSlice {
		if _, v := keys[s]; !v {
			keys[s] = true
			uniq = append(uniq, s)
		}
	}

	return uniq
}
