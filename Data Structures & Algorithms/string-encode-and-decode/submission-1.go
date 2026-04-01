type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var sb strings.Builder
	for _, str := range strs {
		sb.WriteByte(0)
		sb.WriteString(strconv.Itoa(len(str)))
		sb.WriteByte(0)
		sb.WriteString(str)
	}
	return sb.String()
}

func (s *Solution) Decode(encoded string) []string {
	var result []string
	i := 0
	for i < len(encoded) {
		// skip leading null byte
		i++
		// find the next null byte to read the length
		j := strings.IndexByte(encoded[i:], 0)
		length, _ := strconv.Atoi(encoded[i : i+j])
		// skip past the null byte, read exactly `length` bytes
		start := i + j + 1
		result = append(result, encoded[start:start+length])
		i = start + length
	}
	return result
}