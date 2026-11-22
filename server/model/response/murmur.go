package response

// MurmurLikeResult 赞同操作后的最终状态
// 直接返回操作结果，前端无需自行猜测加减，避免多标签页/重复点击后状态错乱。
type MurmurLikeResult struct {
	IsLiked bool  `json:"is_liked"` // 是否已赞同
	Likes   int64 `json:"likes"`    // 最新赞同数
}
