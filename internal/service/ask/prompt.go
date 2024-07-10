package ask

import (
	"regexp"
)

// promptGenarate is entry point
func promptGenarate(UserLanguage string, TargetLanguage string, Word string) string {
	return putWord(getPropmtWithLanguage(UserLanguage, TargetLanguage), Word)
}

// putWord puts the word to prompt
func putWord(Text string, Word string) string {
	regex := regexp.MustCompile(`\[\[.*?\]\]`)
	return regex.ReplaceAllString(Text, Word)
}
