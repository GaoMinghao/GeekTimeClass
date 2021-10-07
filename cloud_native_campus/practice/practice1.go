package main

func replace(input []string) []string {
	keyMap := make(map[string]string)
	keyMap["stupid"] = "smart"
	keyMap["weak"] = "strong"
	for index := range input {
		value, exist := keyMap[input[index]]
		if exist {
			input[index] = value
		}
	}
	return input
}
