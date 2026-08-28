package response

// PlatformHot 一个平台的热榜（也直接用来解析 uapi 返回）
type PlatformHot struct {
	Type string `json:"type"` // 平台标识（uapi 直接返回）
	Name string `json:"name"` // 显示名（从 config 填）
	Icon string `json:"icon"` // 图标（从 config 填）
	List []struct {
		Index    int    `json:"index"`
		Title    string `json:"title"`
		URL      string `json:"url"`
		HotValue string `json:"hot_value"`
	} `json:"list"`
}
