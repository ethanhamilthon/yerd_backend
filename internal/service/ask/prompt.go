package ask

import (
	"regexp"
	"strings"
)

var prompt = `Explain to me what "[[]]" means in tttt (if it is another language then translate to tttt). First explain in general what this word/phrase means.
Then make 3 sentences in tttt, and a translation in uuuu. And explain exactly in the context of each sentence.
Write what it means in more detail, and in the examples should be 3 points, number each example,
but not the sentence itself: the sentence in tttt, its translation,
what the word/phrase "[[[]]" means in this context. The answer is needed without Markdown markup.`

// promptGenarate is entry point
func promptGenarate(UserLanguage string, TargetLanguage string, Word string) (string, string) {
	return putTargetLanguage(putUserLanguage(putWord(prompt, Word), UserLanguage), TargetLanguage), getSystemPrompt(UserLanguage)
}

// putWord puts the word to prompt
func putWord(Text string, Word string) string {
	regex := regexp.MustCompile(`\[\[.*?\]\]`)
	return regex.ReplaceAllString(Text, Word)
}

// getSystemPrompt returns system prompt, which is important to explain to the openai api output language
func getSystemPrompt(UserLanguage string) string {
	text := "YOU HAVE TO ANSWER ONLY IN [[]], I SPEAK ONLY [[]]"
	regex := regexp.MustCompile(`\[\[.*?\]\]`)
	return regex.ReplaceAllString(text, strings.ToUpper(UserLanguage))
}

// putUserLanguage puts the users language to prompt
func putUserLanguage(Text string, LanguageName string) string {
	capitalised := сapitalize(LanguageName)
	regex := regexp.MustCompile(`uuuu`)
	return regex.ReplaceAllString(Text, capitalised)
}

// putTargetLanguage puts the users target language to prompt
func putTargetLanguage(Text string, TargetLanguage string) string {
	regex := regexp.MustCompile(`tttt`)
	return regex.ReplaceAllString(Text, сapitalize(TargetLanguage))
}

// сapitalize makes first char of text uppercase
func сapitalize(str string) string {
	if len(str) == 0 {
		return str
	}
	// Преобразуем первый символ в верхний регистр и соединяем с оставшейся строкой.
	return strings.ToUpper(string(str[0])) + str[1:]
}
