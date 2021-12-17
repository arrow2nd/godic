package main

import (
	"net/url"
)

type commonResult struct {
	Successful     bool   `json:"successful"`
	Text           string `json:"text"`
	TranslatedText string `json:"translated_text"`
}

// Candidate 候補
type Candidate struct {
	Text string `json:"text"`
}

// Word 単語
type Word struct {
	commonResult
	Candidates []Candidate `json:"candidates"`
}

// Translate 変換結果
type Translate struct {
	commonResult
	Words []Word `json:"words"`
}

// GetTranslate ネーミング変換結果を取得
func (g *Godic) GetTranslate(params url.Values) ([]Translate, error) {
	results := []Translate{}

	if err := g.client("v1/engine/translate.json", params, &results); err != nil {
		return nil, err
	}

	return results, nil
}
