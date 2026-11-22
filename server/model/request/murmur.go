package request

type MurmurCreate struct {
	Content string `json:"content" binding:"required"`
}

type MurmurDelete struct {
	Ids []uint `json:"ids" binding:"required"`
}

// MurmurLike 点赞/取消点赞
// UserID 不通过请求体传入，由 JWT 解析后填充，避免伪造他人身份。
type MurmurLike struct {
	UserID   uint `json:"-"`
	MurmurID uint `json:"murmur_id" binding:"required"`
}
