package core

func hashCode(s string) uint16 {
	var result uint16
	for i := 0; i < len(s); i++ {
		result = 31*result + uint16(s[i])
	}
	return result
}
