package config

type Hot struct {
	Platforms []HotPlatform `json:"platforms" yaml:"platforms"` // 平台列表
}

type HotPlatform struct {
	Key     string `json:"key" yaml:"key"`         // 平台标识，对应 service 里的抓取函数
	Name    string `json:"name" yaml:"name"`       // 平台显示名称
	Enabled bool   `json:"enabled" yaml:"enabled"` // 是否启用该平台
	Icon    string `json:"icon" yaml:"icon"`       // 平台图标链接
}
