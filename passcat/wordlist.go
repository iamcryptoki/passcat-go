package passcat

import (
	"fmt"
)

type worder interface {
	WordAt(int) string
}

type wordList struct {
	words map[int]string
}

func (w *wordList) WordAt(i int) string {
	return w.words[i]
}

func getWordList(lang string) (worder, error) {
	switch lang {
	case "de":
		return wordListGerman, nil
	case "eff":
		return wordListEFF, nil
	case "es":
		return wordListSpanish, nil
	case "fr":
		return wordListFrench, nil
	case "it":
		return wordListItalian, nil
	case "jp":
		return wordListJapanese, nil
	case "nl":
		return wordListDutch, nil
	case "no":
		return wordListNorwegian, nil
	case "pl":
		return wordListPolish, nil
	case "ro":
		return wordListRomanian, nil
	case "ru":
		return wordListRussian, nil
	case "se":
		return wordListSwedish, nil
	case "tr":
		return wordListTurkish, nil
	default:
		return nil, fmt.Errorf("Wordlist not found (%s).", lang)
	}
}
