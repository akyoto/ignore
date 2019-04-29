package ignore

// Reader is a reader that can ignore runes.
type Reader struct {
	inString          bool
	inCharacterString bool
	inParentheses     int
	escape            bool
}

// CanIgnore tells you if you can ignore the rune.
func (r *Reader) CanIgnore(letter rune) bool {
	if letter == '\\' && !r.escape {
		r.escape = true
		return true
	}

	defer func() {
		r.escape = false
	}()

	if letter == '"' && !r.escape {
		r.inString = !r.inString
		return true
	}

	if r.inString {
		return true
	}

	if letter == '\'' && !r.escape {
		r.inCharacterString = !r.inCharacterString
		return true
	}

	if r.inCharacterString {
		return true
	}

	if letter == '(' || letter == '[' || letter == '{' {
		r.inParentheses++
		return true
	}

	if letter == ')' || letter == ']' || letter == '}' {
		r.inParentheses--

		if r.inParentheses == 0 {
			return true
		}
	}

	if r.inParentheses > 0 {
		return true
	}

	return false
}
