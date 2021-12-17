package codic

// Codic クライアント
type Codic struct {
	apiKey string
}

// NewClient 作成
func NewClient(apiKey string) *Codic {
	return &Codic{
		apiKey: apiKey,
	}
}
