package main

func Manipulate(text []string) []string {
	// check first elements length
	if len(text[0]) == 1 {
		return text
	}

	perm := []string{}
	uniquePerm := make(map[string]bool)

	for i, v := range text[0] {
		remains := text[0][:i] + text[0][i+1:]
		subs := Manipulate([]string{remains})

		for _, s := range subs {
			permutation := string(v) + s
			if !uniquePerm[permutation] {
				perm = append(perm, permutation)
				uniquePerm[permutation] = true
			}
		}
	}
	return perm
}
