package other

// HotResponse 聚合接口统一返回结构（各平台一样）
type HotResponse struct {
	Type string `json:"type"` // 平台标识（bilibili/weibo...）
	List []struct {
		Index    int    `json:"index"`
		Title    string `json:"title"`
		URL      string `json:"url"`
		HotValue string `json:"hot_value"`
	} `json:"list"`
}
