package glob

import (
	"fmt"
	"strconv"
	"strings"
)

const NullCharacter = rune(0)

// MatchingEngine is an interface which supports the Matches function
type MatchingEngine interface {
	// Matches returns true if the given string matches the compiled pattern
	Matches(str string) bool
}

// Configurable options for the compile function
type compileOption struct {
	wildcardChar    rune
	matchOneChar    rune
	caseInsensitive bool
	handleEscapes   bool
}

// CompileOption are options for the compile function
type CompileOption interface {
	// Apply applies the option to a compile option struct
	apply(*compileOption)
}

// Struct to hold the wildcard char option
type withWildcardChar struct {
	wildcardChar rune
}

func (o *withWildcardChar) apply(co *compileOption) {
	co.wildcardChar = o.wildcardChar
}

// WithWildcardChar allows setting the wildcard char as an option
func WithWildcardChar(wildcardChar rune) CompileOption {
	return &withWildcardChar{wildcardChar: wildcardChar}
}

// WithoutWildcard disables wildcards completely
func WithoutWildcard() CompileOption {
	return &withWildcardChar{wildcardChar: NullCharacter}
}

// Struct to hold the match one char option
type withMatchOneChar struct {
	matchOneChar rune
}

func (o *withMatchOneChar) apply(co *compileOption) {
	co.matchOneChar = o.matchOneChar
}

// WithMatchOneChar allows setting the match one char as an option
func WithMatchOneChar(matchOneChar rune) CompileOption {
	return &withMatchOneChar{matchOneChar: matchOneChar}
}

// WithoutMatchOne disables match one completely
func WithoutMatchOne() CompileOption {
	return &withMatchOneChar{matchOneChar: NullCharacter}
}

// Struct to hold the case insensitive option
type withCaseInsensitive struct {
	caseInsensitive bool
}

func (o *withCaseInsensitive) apply(co *compileOption) {
	co.caseInsensitive = o.caseInsensitive
}

// WithCaseInsensitive allows setting the case insensitive as an option
func WithCaseInsensitive(caseInsensitive bool) CompileOption {
	return &withCaseInsensitive{caseInsensitive: caseInsensitive}
}

type withHandleEscapes struct {
	handleEscapes bool
}

func (o *withHandleEscapes) apply(co *compileOption) {
	co.handleEscapes = o.handleEscapes
}

func WithHandleEscapes(handleEscapes bool) CompileOption {
	return &withHandleEscapes{handleEscapes: handleEscapes}
}

// Compile compiles the glob pattern and returns a MatchingEngine or error if there was one
func Compile(globPattern string, options ...CompileOption) (MatchingEngine, error) {

	// Build our compile options with default values
	co := compileOption{
		wildcardChar:  rune('*'),
		matchOneChar:  rune('?'),
		handleEscapes: true,
	}

	// Process any options passed in on the compile options
	for _, option := range options {
		option.apply(&co)
	}

	// The 'raw' variables hold the pattern before we have processed it.  We just copy it here so we can directly
	//  access the characters and deal with the case sensitivity in one place.  After this we don't have to
	//  worry about case sensitivity again
	lengthRaw := len(globPattern)
	var lowerCaseRaw []rune
	var upperCaseRaw []rune

	if co.caseInsensitive {
		// Case in-sensitive match so convert the pattern to upper and lower and assign it accordingly
		lowerCaseRaw = []rune(strings.ToLower(globPattern))
		upperCaseRaw = []rune(strings.ToUpper(globPattern))
	} else {
		// Case sensitive match so just get the backing character array and assign it to the upper and lower
		lowerCaseRaw = []rune(globPattern)
		upperCaseRaw = []rune(globPattern)
	}

	// These will hold our strings after we process the escapes, multiple wildcards and anything else
	lowerCase := make([]rune, lengthRaw)
	upperCase := make([]rune, lengthRaw)

	// Identifies the positions in the character array with wildcard characters
	wildcard := make([]bool, lengthRaw)

	// Identifies the positions in the character array match one characters
	matchOne := make([]bool, lengthRaw)

	// Flag to identify if we have processed an escape character.  If handleEscape is false then this
	// will never be set to true
	var inEscape bool

	// Walk through the raw character arrays and copy them into the upperCase and lowerCase arrays.
	//  While we walk through we will collapse multiple globs and handle escaped characters.

	// Count the wildcards for shortcut handling later
	wildcardCount := 0

	// Index points to where we are in the lowerCase and upperCase arrays.
	index := 0
	for i := 0; i < lengthRaw; i++ {

		// Pull the characters/runes out
		lowerC := lowerCaseRaw[i]
		upperC := upperCaseRaw[i]

		// Handle escaping
		if co.handleEscapes && lowerC == '\\' && !inEscape {
			// Mark the flag and eat the escape
			inEscape = true
			continue
		}

		// Check if we have hit a wildcard character
		if co.wildcardChar != NullCharacter && lowerC == co.wildcardChar {

			// If we are in an escape then ignore the wildcard and let it fall through as a simple character
			if !inEscape {
				// Check if this is a duplicate
				if index != 0 && wildcard[index-1] {
					// Yes, they have entered two wildcard patterns in a row so ignore the second wildcard
					continue
				}

				// We need to mark this location as being a wildcard
				wildcard[index] = true
				wildcardCount++
			}
		} else if co.matchOneChar != NullCharacter && lowerC == co.matchOneChar {
			// As before, if we are in an escape then let the exactly one fall through as a regular character.
			if !inEscape {
				// Mark the location as being a matchOne
				matchOne[index] = true
			}
		} else if inEscape {
			switch lowerC {
			case 'r':
				// Return
				lowerC = '\r'
				upperC = '\r'
			case 'n':
				// New Line
				lowerC = '\n'
				upperC = '\n'
			case 't':
				// Tab
				lowerC = '\t'
				upperC = '\t'
			case '\\':
				// Backslash
				lowerC = '\\'
				upperC = '\\'
			case 'u':
				// Unicode
				if i+4 < lengthRaw {
					s := string(lowerCaseRaw[i+1 : i+5])
					// Try to convert it
					u, err := strconv.ParseUint(s, 16, 64)
					if err != nil {
						return nil, fmt.Errorf("bad unicode character: %s", s)
					}
					upperC = rune(u)
					lowerC = rune(u)
					// Move i ahead 4 for the characters we ate
					i += 4
				} else {
					return nil, fmt.Errorf("bad Unicode char at end of input")
				}
			default:
				return nil, fmt.Errorf("unknown escape sequence : \\%U", lowerC)
			}
		}

		// Copy the character over
		lowerCase[index] = lowerC
		upperCase[index] = upperC
		index++

		// Turn off escaping
		inEscape = false
	}

	// Index holds our length
	length := index

	// At this point we are done compiling the wildcard pattern into the lowerCase and upperCase arrays.
	//  But we can inspect the resulting patterns and make some simple optimizations.

	// If the pattern is empty ... they gave us an empty string so short cut it.
	if length == 0 {
		// specifically: ''
		return emptyOnlyEngineInst, nil
	}

	// If our pattern is just one character long, and its a wildcard, then this will match everything
	//  so return the Match Everything Engine.
	if length == 1 && wildcard[0] {
		// specifically: '%'
		return everythingEngineInst, nil
	}

	// If there are no wildcards then this is a simple equal to engine
	if wildcardCount == 0 {
		// ex: 'foo'
		return &equalToEngine{lowerCase: lowerCase, upperCase: upperCase, matchOne: matchOne, length: length}, nil
	}

	// If there is only one wildcard and it is at either the start or end then return
	//  an EndsWith or StartsWith engine accordingly
	if wildcardCount == 1 {
		if wildcard[0] {
			// ex: '%foo'
			return &endsWithEngine{lowerCase: lowerCase, upperCase: upperCase, matchOne: matchOne, length: length}, nil
		}
		if wildcard[length-1] {
			// ex: 'foo%'
			return &startsWithEngine{lowerCase: lowerCase, upperCase: upperCase, matchOne: matchOne, length: length}, nil

		}
	}

	// If there are two wildcards and they are at the start AND end then this is a contains
	if wildcardCount == 2 && wildcard[0] && wildcard[length-1] {
		// ex: '%foo%'
		return &containsEngine{lowerCase: lowerCase, upperCase: upperCase, matchOne: matchOne, length: length}, nil
	}

	// No other shortcuts so fall back to the glob engine
	return &globEngine{lowerCase: lowerCase, upperCase: upperCase, wildcard: wildcard, matchOne: matchOne, length: length}, nil
}
