package jmdict

import (
	"encoding/json"
	"io"
	"strings"

	"github.com/Tranduy1dol/learning-japanese/internal/domain"
)

type JMdictFile struct {
	Words []JMdictWord `json:"words"`
}

type JMdictWord struct {
	ID    string        `json:"id"`
	Kanji []JMdictKanji `json:"kanji"`
	Kana  []JMdictKana  `json:"kana"`
	Sense []JMdictSense `json:"sense"`
}

type JMdictKanji struct {
	Common bool     `json:"common"`
	Text   string   `json:"text"`
	Tags   []string `json:"tags"`
}

type JMdictKana struct {
	Common         bool     `json:"common"`
	Text           string   `json:"text"`
	Tags           []string `json:"tags"`
	AppliesToKanji []string `json:"appliesToKanji"`
}

type JMdictSense struct {
	PartOfSpeech []string      `json:"partOfSpeech"`
	Gloss        []JMdictGloss `json:"gloss"`
	Field        []string      `json:"field"`
	Dialect      []string      `json:"dialect"`
}

type JMdictGloss struct {
	Lang string `json:"lang"`
	Text string `json:"text"`
}

func Parse(reader io.Reader) ([]*domain.Word, error) {
	var file JMdictFile
	err := json.NewDecoder(reader).Decode(&file)
	if err != nil {
		return nil, err
	}

	words := make([]*domain.Word, 0, len(file.Words))
	for _, e := range file.Words {
		word := new(domain.Word)

		word.ID = e.ID
		word.EntSeq = e.ID
		word.Source = "jmdict"

		for _, k := range e.Kanji {
			word.Kanji = append(word.Kanji, domain.Kanji{
				Text: k.Text,
				Info: strings.Join(k.Tags, ","),
			})

			if k.Common {
				word.IsCommon = true
			}
		}

		for _, r := range e.Kana {
			word.Readings = append(word.Readings, domain.Reading{
				Text: r.Text,
				Info: r.Tags,
			})

			if r.Common {
				word.IsCommon = true
			}
		}

		for _, s := range e.Sense {
			sense := domain.Sense{
				POS:    s.PartOfSpeech,
				Source: strings.Join(s.Field, ","),
			}

			for _, g := range s.Gloss {
				sense.Gloss = append(sense.Gloss, domain.Gloss{
					Text: g.Text,
					Lang: g.Lang,
				})
			}

			word.Senses = append(word.Senses, sense)
		}

		words = append(words, word)
	}

	return words, nil
}
