package domain

import "strings"

type Word struct {
	ID        string    `bson:"_id"`
	EntSeq    string    `bson:"ent_seq"`
	Kanji     []Kanji   `bson:"kanji"`
	Readings  []Reading `bson:"readings"`
	Senses    []Sense   `bson:"senses"`
	JLPT      int       `bson:"jlpt"`
	IsCommon  bool      `bson:"is_common"`
	Source    string    `bson:"source"`
	CreatedBy string    `bson:"created_by"`
}

type Kanji struct {
	Text     string `bson:"text"`
	Info     string `bson:"info"`
	Priority int    `bson:"priority"`
}

type Reading struct {
	Text     string   `bson:"text"`
	Status   string   `bson:"status"`
	Info     []string `bson:"info"`
	Priority int      `bson:"priority"`
}

type Sense struct {
	POS    []string `bson:"pos"`
	Gloss  []Gloss  `bson:"gloss"`
	Source string   `bson:"source"`
}

type Gloss struct {
	Text string `bson:"text"`
	Lang string `bson:"lang"`
}

func (w *Word) SearchText() (title string, text string) {
	if len(w.Kanji) > 0 {
		title = w.Kanji[0].Text
	} else if len(w.Readings) > 0 {
		title = w.Readings[0].Text
	}

	var parts []string
	for _, k := range w.Kanji {
		parts = append(parts, k.Text)
	}
	for _, r := range w.Readings {
		parts = append(parts, r.Text)
	}
	for _, s := range w.Senses {
		for _, g := range s.Gloss {
			parts = append(parts, g.Text)
		}
	}

	text = strings.Join(parts, " ")
	return
}
