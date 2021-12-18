package godic

import "net/url"

// CreateTranslateParam GetTranslateのURLパラメータを作成
func CreateTranslateParam(text string) url.Values {
	v := url.Values{}
	v.Add("text", text)
	return v
}

// CreateCEDLookupParam GetCEDLookupのURLパラメータを作成
func CreateCEDLookupParam(query string) url.Values {
	v := url.Values{}
	v.Add("query", query)
	return v
}
