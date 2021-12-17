package main

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
func (g *Godic) GetUserProjectById(id string) (UserProject, error) {
	result := UserProject{}

	if err := g.client("v1/user_projects/"+id+".json", nil, &result); err != nil {
		return UserProject{}, err
	}

	return result, nil
}

// GetUserProjectList プロジェクト一覧を取得
func (g *Godic) GetUserProjectList() ([]UserProject, error) {
	results := []UserProject{}

	if err := g.client("v1/user_projects.json", nil, &results); err != nil {
		return nil, err
	}

	return results, nil
}
