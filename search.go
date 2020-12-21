package shakesearch

import (
	"github.com/sahilm/fuzzy"

	"github.com/MicahParks/shakesearch/models"
)

// Search does a fuzzy search of the complete works of Shakespeare and returns a slice of results with a length less
// than or equal to the maxMatches. A maxMatches with a value of -1 will return all matches.
func (s *ShakeSearcher) Search(maxMatches int, query string) (results []*models.Result) {

	// Find all fuzzy matching strings.
	matches := fuzzy.Find(query, s.uniqueSlice)

	// Figure out how many results to return.
	var matchCount int
	if maxMatches <= len(matches) {
		matchCount = maxMatches
	} else {
		matchCount = len(matches)
	}

	// Allocate the required memory for the return slice so it's faster to add all the strings.
	results = make([]*models.Result, matchCount)

	// Add each match to the set of strings only once. Keep track of the indexes.
	for index, match := range matches {

		// Make sure to only return the maximum number of matches.
		if maxMatches <= 0 || index == maxMatches {
			break
		}

		// Get a copy of the result to manipulate and return.
		result := *s.unique[match.Str]

		// Allocate the memory for the matching indexes slice so inserting them is faster.
		result.MatchedIndexes = make([]int64, len(match.MatchedIndexes))

		// Turn the matched indexes into the correct integer format.
		for i, matchedIndex := range match.MatchedIndexes {
			result.MatchedIndexes[i] = int64(matchedIndex)
		}

		// Use the modified copy in the result to return.
		results[index] = &result
	}

	return results
}
