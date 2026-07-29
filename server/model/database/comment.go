package database

import (
	"server/global"

	"github.com/gofrs/uuid"
)

// Comment 评论表
type Comment struct {
	global.MODEL
	ArticleID string `json:"article_id"` // 文章 ID
	PID       *uint  `json:"p_id"`       // 父评论 ID
	// 1. global.MODEL 里的主键就是 uint 类型，这里正好是同类型。
	// 2. 为了实现多对多关系，这里使用了 *uint 类型，表示可以为空，uint 空值只能是 0，可能会有歧义
	PComment *Comment  `json:"-" gorm:"foreignKey:PID"`
	Children []Comment `json:"children" gorm:"foreignKey:PID"`                  // 子评论
	UserUUID uuid.UUID `json:"user_uuid" gorm:"type:char(36)"`                  // 用户 uuid
	User     User      `json:"user" gorm:"foreignKey:UserUUID;references:UUID"` // 关联的用户
	Content  string    `json:"content"`                                         // 内容
}

// TODO 创建和删除评论时需要更新文章评论数
