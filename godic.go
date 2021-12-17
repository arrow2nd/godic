package godic

// Godic クライアント
type Godic struct {
	apiKey string
}

// NewClient 作成
func NewClient(apiKey string) *Godic {
	return &Godic{
		apiKey: apiKey,
	}
}
