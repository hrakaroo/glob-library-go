package glob

const startingStackSize = 16

// Used to identify when there is no save point to jump back to
const noSavePoint = -1

var emptyOnlyEngineInst = &emptyOnlyEngine{}
var everythingEngineInst = &everythingEngine{}

// emptyOnlyEngine is one which just matches the empty string
type emptyOnlyEngine struct {
}

func (e *emptyOnlyEngine) Matches(str string) bool {
	return len(str) == 0
}

// everythingEngine matches everything
type everythingEngine struct {
}

func (e *everythingEngine) Matches(str string) bool {
	return true
}

// EqualToEngine is an equality check
type equalToEngine struct {
	lowerCase string
	upperCase string
	matchOne  []bool
	length    int
}

func (e *equalToEngine) Matches(str string) bool {

	// Test if the lengths are the same and shortcut if not
	if e.length != len(str) {
		return false
	}

	// Walk down the two character arrays and verify the characters are equal.
	index := 0
	for {
		// If we reach the end the they must be equal
		if index == e.length {
			return true
		}

		if !e.matchOne[index] && str[index] != e.lowerCase[index] && str[index] != e.upperCase[index] {
			return false
		}
		index++
	}
}

type endsWithEngine struct {
	lowerCase string
	upperCase string
	matchOne  []bool
	length    int
}

func (e *endsWithEngine) Matches(str string) bool {

	// The input string must be at least as long as the glob pattern -1 (for the first wildcard)
	if len(str) < e.length-1 {
		return false
	}

	charsIndex := len(str) - 1

	for patternIndex := e.length - 1; patternIndex > 0; patternIndex-- {
		if !e.matchOne[patternIndex] && str[charsIndex] != e.lowerCase[patternIndex] &&
			str[charsIndex] != e.upperCase[patternIndex] {
			return false
		}
		charsIndex--
	}
	return true
}

type startsWithEngine struct {
	lowerCase string
	upperCase string
	matchOne  []bool
	length    int
}

func (e *startsWithEngine) Matches(str string) bool {

	// The input string must be at least as long as the glob pattern -1 (for the final wildcard)
	if len(str) < e.length-1 {
		return false
	}

	// Run down through the pattern and compare them.  If we get to the end of our pattern then have matched.
	for index := 0; index < e.length-1; index++ {
		if !e.matchOne[index] && str[index] != e.lowerCase[index] && str[index] != e.upperCase[index] {
			return false
		}
	}
	return true
}

type containsEngine struct {
	lowerCase string
	upperCase string
	matchOne  []bool
	length    int
}

func (e *containsEngine) Matches(str string) bool {

	// The input string must be at least as long as the glob pattern - 2 (for the first and last wildcard)
	if len(str) < e.length-2 {
		return false
	}

	charsIndex := 0
	patternIndex := 1 // Skip the starting wildcard
	charsSavePoint := noSavePoint

	// The only counter we really care about is the one walking the pattern.  Once the pattern index gets
	//  to the end of its pattern (-1) we are done.  The chars index is just along for the ride.

	// The difficulty in using a for loop here is that the patternIndex and charsIndex can both be reset
	//  which means we would have to reset them at value-1 to offset the for loop incrementing them.

	for {
		if patternIndex == e.length-1 {
			// We got to the end of our pattern (-1 for the final wildcard char)
			return true
		}

		// Check if we are not at the end of our string and if the character is a matchOne or matches
		if charsIndex != len(str) &&
			(e.matchOne[patternIndex] ||
				str[charsIndex] == e.lowerCase[patternIndex] ||
				str[charsIndex] == e.upperCase[patternIndex]) {

			// At this point the input string and the pattern are matching, but this may be a false
			// match in the string (meaning it starts with the pattern but is not a full pattern)
			// Consider string = "my dog fred my dog spot" and the pattern "%my dog spot%"
			// This pattern is going to *start* matching at charIndex 0, but its a false match since
			// the two will diverge before we get to the end of the pattern.
			// Therefore, if we are not already waking down a branch we need to save a branch point
			// that assumes its a false match.

			if charsSavePoint == noSavePoint {
				charsSavePoint = charsIndex + 1
			}
			// Keep walking forward
			patternIndex++
			charsIndex++
		} else if charsIndex != len(str) && charsSavePoint == noSavePoint {

			// charsSavePoint == noSavePoint means we are still pointing at the wildcard and
			//  have not yet matched, so as long as we are not at the end of our string, keep walking
			charsIndex++
		} else if charsSavePoint != noSavePoint {

			// Our chars index has reached the end of the input string before our pattern index reached
			//  the end of the pattern. Normally we would be dead here, but we have a save point so
			//  reset the charsIndex to the save point (which has already been moved forward), reset
			//  the pattern index to the start of the pattern and remove our save point.
			charsIndex = charsSavePoint
			patternIndex = 1
			charsSavePoint = noSavePoint
		} else {
			// We failed and we have no alternative branch to try
			return false
		}
	}
}

type globEngine struct {
	lowerCase string
	upperCase string
	wildcard  []bool
	matchOne  []bool
	length    int
}

func (e *globEngine) Matches(str string) bool {

	// We use a stack instead of doing recursion.
	//  We could move this stack declaration to a class field and doing so gains us about a 5% speed
	//  improvement.  However, it would also remove the thread safety of this class so for the time
	//  being I'm going to leave it here.
	stack := make([]int, startingStackSize)
	stackIndex := 0

	// Our two indexes, one to walk down the input string and the other to walk down the pattern
	charsIndex := 0
	patternIndex := 0

	for {
		// If we are not at the end of our pattern and we have hit a wildcard
		if patternIndex != e.length && e.wildcard[patternIndex] {

			// If we are not at the end of our input string
			if charsIndex != len(str) {
				// Check our stack is big enough.
				if stackIndex == cap(stack) {
					tmp := make([]int, cap(stack)*2)
					copy(tmp, stack)
					stack = tmp
				}
				// Add a path branch where we try walking the char index forward
				stack[stackIndex] = patternIndex
				stackIndex++
				stack[stackIndex] = charsIndex + 1
				stackIndex++
			}

			// Walk the wildcard index forward (non-greedy)
			patternIndex++
		}

		// If we reach the end of our character array and
		//  we are at the end of the wildcard then we are are successful
		if patternIndex == e.length && charsIndex == len(str) {
			// Boom!
			return true
		}

		// This path fails if we are either at the end of the character string, or the pattern string as we
		//  can not be at the end of both of them at the same time since that would have been caught by the
		//  above conditional.
		// If we are not at the end of either string then we fail if the characters match or if we are on a
		//  matchOne (meaning we match any single character.)
		if charsIndex == len(str) || patternIndex == e.length ||
			(!e.matchOne[patternIndex] && str[charsIndex] != e.lowerCase[patternIndex] && str[charsIndex] != e.upperCase[patternIndex]) {

			// This branch has failed to match, so check if we have a different branch we can try
			if stackIndex != 0 {
				// Yes, so reset our indexes and loop.
				stackIndex--
				charsIndex = stack[stackIndex]
				stackIndex--

				patternIndex = stack[stackIndex]
			} else {
				// No, we have exhausted all available branches so this pattern fails to branch.
				return false
			}
		} else {
			// We get here if we did not fail in any of those above conditions and we should continue
			//  walking down both strings.
			patternIndex++
			charsIndex++
		}
	}
}
