package shakesearch

// SurroundingWorks retrieves the line numbers surrounding the given line. The number of lines returned will be less
// than or equal to maxLines.
func (s *ShakeSearcher) SurroundingWorks(lineNumber int, maxLines uint) (lines Lines) {

	// Get the length of the raw slice.
	length := len(s.raw)

	// Confirm a valid line number.
	if !lineNumberValid(length, lineNumber) {
		return lines
	}

	// Get the lower index.
	lowerIndex := lineNumber - 1 - int(maxLines)/2
	if !lineNumberValid(length, lowerIndex) {
		lowerIndex = 0
	}

	// Get the higher index.
	higherIndex := lineNumber - 1 + int(maxLines)/2
	if !lineNumberValid(length, higherIndex) {
		higherIndex = length
	}

	return s.raw[lowerIndex:higherIndex]
}

// lineNumberValid indicates if the line number is valid for the given length.
func lineNumberValid(length, lineNumber int) (valid bool) {
	valid = true
	if lineNumber-1 < 0 || lineNumber > length {
		valid = false
	}
	return valid
}
