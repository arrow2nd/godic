package codic

import "encoding/json"

// Owner オーナー情報
type Owner struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

// UserProject プロジェクト情報
type UserProject struct {
	Owner
	Description string `json:"description"`
	CreatedOn   string `json:"created_on"`
	WordsCount  uint   `json:"words_count"`
	OwnerInfo   Owner  `json:"owner"`
}

// GetUserProjectById IDからプロジェクトを取得
func (c *Codic) GetUserProjectById(id string) (UserProject, error) {
	bytes, err := c.client("v1/user_projects/"+id+".json", nil)
	if err != nil {
		return UserProject{}, err
	}

	result := UserProject{}
	if err := json.Unmarshal(bytes, &result); err != nil {
		return UserProject{}, err
	}

	return result, nil
}
