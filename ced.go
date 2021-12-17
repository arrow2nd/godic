package main

import (
	"net/url"
)

// CED 英和辞書の見出し語情報
type CED struct {
	Id     uint64 `json:"id"`
	Title  string `json:"title"`
	Digest string `json:"digest"`
}

// GetCEDLookup 英和辞書の見出し語を前方一致検索
func (g *Godic) GetCEDLookup(params url.Values) ([]CED, error) {
	results := []CED{}

	if err := g.client("v1/ced/lookup.json", params, &results); err != nil {
		return nil, err
	}

	return results, nil
}

// CEDPronunciation 発音情報
type CEDPronunciation struct {
	Type   string   `json:"type"`
	Text   string   `json:"text"`
	Labels []string `json:"labels"`
}

// CEDTranslation 翻訳結果
type CEDTranslation struct {
	Note      string   `json:"note"`
	Labels    []string `json:"labels"`
	Etymology uint64   `json:"etymology"`
	Pos       string   `json:"pos"`
	Text      string   `json:"text"`
}

// CEDEntry 英和辞書のエントリー情報
type CEDEntry struct {
	CED
	Pronunciations []CEDPronunciation `json:"pronunciations"`
	Translations   []CEDTranslation   `json:"translations"`
	Note           string             `json:"note"`
}

// GetCEDEntries 英和辞書のエントリーを取得
func (g *Godic) GetCEDEntries(id string) (CEDEntry, error) {
	result := CEDEntry{}

	if err := g.client("v1/ced/entries/"+id+".json", nil, &result); err != nil {
		return CEDEntry{}, err
	}

	return result, nil
}
