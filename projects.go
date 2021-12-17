package codic

import "encoding/json"

// GetUserProjectList プロジェクト一覧を取得
func (c *Codic) GetUserProjectList() ([]UserProject, error) {
	bytes, err := c.client("v1/user_projects.json", nil)
	if err != nil {
		return nil, err
	}

	results := []UserProject{}
	if err := json.Unmarshal(bytes, &results); err != nil {
		return nil, err
	}

	return results, nil
}
