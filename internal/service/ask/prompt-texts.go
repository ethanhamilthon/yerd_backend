package ask

type Targets struct {
	English string
	German  string
}

var (
	englishText Targets = Targets{
		English: `Explain to me what "[[]]" means in English (if it is another language then translate to English). First explain in general what this word/phrase means. Then make 3 sentences in English, and a translation in English. And explain exactly in the context of each sentence. The answer is needed without Markdown markup.`,
		German:  `Explain to me what "[[]]" means in German (if it is another language then translate to German). First explain in general what this word/phrase means. Then make 3 sentences in German, and a translation in English. And explain exactly in the context of each sentence. The answer is needed without Markdown markup.`,
	}
	russianText Targets = Targets{
		English: `Объясните мне, что означает "[[]]" на английском языке (если это другой язык, то переведите на английский). Сначала объясните в общих чертах, что означает это слово/фраза. Затем составьте 3 предложения на английском языке и перевод на русский. И объясните точно в контексте каждого предложения. Ответ нужен без разметки Markdown.`,
		German:  `Объясните мне, что означает "[[]]" на немецком языке (если это другой язык, то переведите на немецкий). Сначала объясните в общих чертах, что означает это слово/фраза. Затем составьте 3 предложения на немецком языке и перевод на русский. И объясните точно в контексте каждого предложения. Ответ нужен без разметки Markdown.`,
	}
	frenchText Targets = Targets{
		English: `Expliquez-moi ce que signifie "[[]]" en anglais (s'il s'agit d'une autre langue, traduisez-la en anglais). Expliquez d'abord en général ce que ce mot/phrase signifie. Ensuite, faites 3 phrases en anglais, et une traduction en français. Et expliquez exactement dans le contexte de chaque phrase. La réponse est nécessaire sans balisage Markdown.`,
		German:  `Expliquez-moi ce que signifie "[[]]" en allemand (s'il s'agit d'une autre langue, traduisez-la en allemand). Expliquez d'abord en général ce que ce mot/phrase signifie. Ensuite, faites 3 phrases en allemand, et une traduction en français. Et expliquez exactement dans le contexte de chaque phrase. La réponse est nécessaire sans balisage Markdown.`,
	}
	turkishText Targets = Targets{
		English: `Bana "[[]]" ifadesinin İngilizcede ne anlama geldiğini açıklayın (eğer başka bir dil ise İngilizceye çevirin). Önce açıklayın Genel olarak bu kelimenin/cümlenin ne anlama geldiğini. Daha sonra İngilizce 3 cümle ve Türkçe bir çeviri yapın. Ve tam olarak açıklayın her cümle bağlamında. Cevap Markdown işaretlemesi olmadan gereklidir.`,
		German:  `Bana "[[]]" ifadesinin Almanca'da ne anlama geldiğini açıklayın (eğer başka bir dilde ise Almanca'ya çevirin). Önce açıklayın Genel olarak bu kelimenin/cümlenin ne anlama geldiğini. Daha sonra Almanca 3 cümle ve Türkçe bir çeviri yapın. Ve tam olarak açıklayın her cümle bağlamında. Cevap Markdown işaretlemesi olmadan gereklidir.`,
	}
	chineseText Targets = Targets{
		English: `请解释一下"[[]]"在英语中是什么意思（如果是其他语言，请翻译成英语）。首先解释这个词/短语的大致意思。然后用英语造 3 个句子，并翻译成中文。并准确解释每个句子的上下文。答案不需要 Markdown 标记。`,
		German:  `请解释一下"[[]]"在德语中是什么意思（如果是其他语言，请翻译成德语）。首先解释这个词/短语的大致意思。然后用德语造 3 个句子，并翻译成中文。并准确解释每个句子的上下文。答案不需要 Markdown 标记。`,
	}
)

func getPropmtWithLanguage(Lang string, TargetLang string) string {
	userLang := englishText
	switch Lang {
	case "english":
		userLang = englishText
	case "russian":
		userLang = russianText
	case "french":
		userLang = frenchText
	case "turkish":
		userLang = turkishText
	case "chinese":
		userLang = chineseText
	}

	result := userLang.English
	switch TargetLang {
	case "english":
		result = userLang.English
	case "german":
		result = userLang.German
	}

	return result
}
