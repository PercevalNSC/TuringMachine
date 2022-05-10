package dtm

func OverWrite(w string, b byte, i int) string {
	result := []byte(w)
	result[i] = b
	return string(result)
}
