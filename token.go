package main

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

// Generate a random token
func NewRandomToken(index int) string {
	digits := []int{}
	const length = len(alphabet)
	for index > 0 {
		digits = append(digits, index%length)
		index /= length
	}

	if len(digits) == 0 {
		digits = []int{0}
	}

	token := ""
	for _, v := range digits {
		token += string(alphabet[v%length])
	}

	return token
}
