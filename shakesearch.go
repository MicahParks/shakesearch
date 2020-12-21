package shakesearch

import (
	"io/ioutil"
	"strings"

	"github.com/MicahParks/shakesearch/models"
)

// Line represents a line in Shakespeare's works.
type Line struct {
	Number uint
	Text   string
}

// Lines is an array of Line. It is used as typing information for the Go template.
type Lines []Line

// ShakeSearcher is the data structure that holds the unique lines of Shakespeare's complete works in a convent format
// to search. The format is a map of all lines and a slice of the line numbers they belong to.
type ShakeSearcher struct {
	raw         []Line
	unique      map[string]*models.Result
	uniqueSlice []string
}

// NewShakeSearcher loads the give file as a slice of string with trimmed space.
func NewShakeSearcher(filePath string) (shakeSearcher *ShakeSearcher, err error) {

	// Create the ShakeSearcher.
	shakeSearcher = &ShakeSearcher{
		raw:         make([]Line, 0),
		unique:      make(map[string]*models.Result),
		uniqueSlice: make([]string, 0),
	}

	// Read the file into memory.
	var fileData []byte
	if fileData, err = ioutil.ReadFile(filePath); err != nil {
		return nil, err
	}

	// Split the file data by newlines.
	lines := strings.Split(string(fileData), "\n")

	// Create the required data structures.
	shakeSearcher.createDataStructures(lines)

	return shakeSearcher, nil
}

// createDataStructures creates the required data structures for the ShakeSearcher given the slice of lines from
// Shakespeare's complete works.
func (s *ShakeSearcher) createDataStructures(lines []string) {

	// Allocate the memory for the raw lines up front so it inserts faster.
	s.raw = make([]Line, len(lines))

	// Iterate through the given lines and create the desired data structures.
	for index, line := range lines {

		// Add the line to the raw lines.
		s.raw[index] = Line{
			Number: uint(index) + 1,
			Text:   line,
		}

		// Add the line to the unique set if it isn't just whitespace.
		if line = strings.TrimSpace(line); line != "" {

			// Get the existing UniqueLine data structure's pointer from the map.
			uniqueLine, ok := s.unique[line]

			// If this is the first time this line has been seen, create the data structure and assign its pointer to
			// the map.
			if !ok {
				uniqueLine = &models.Result{
					Line:        line,
					LineNumbers: make([]int64, 0),
				}
				s.unique[line] = uniqueLine
			}

			// Add the line number for this line to the UniqueLine data structure in the map.
			uniqueLine.LineNumbers = append(uniqueLine.LineNumbers, int64(index+1))
		}
	}

	// Create a slice that will hold all the unique lines. Allocate the memory up front so it inserts faster.
	s.uniqueSlice = make([]string, len(s.unique))

	// Insert all of the unique lines to the slice.
	index := uint(0)
	for line := range s.unique {
		s.uniqueSlice[index] = line
		index++
	}
}
