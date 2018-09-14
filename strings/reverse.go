package strings

// Reverse gives the reversed string of string_in
func Reverse(string_in string) string {
	rune_slice := []rune(string_in)

	var forward_rune rune
	forward_idx := 0
	backward_idx := len(rune_slice) - 1
	for backward_idx > forward_idx {
		forward_rune = rune_slice[forward_idx]
		rune_slice[forward_idx] = rune_slice[backward_idx]
		rune_slice[backward_idx] = forward_rune

		forward_idx++
		backward_idx--
	}
	return string(rune_slice)
}
