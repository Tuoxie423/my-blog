package database

import "server/global"

// MurmurLike 碎碎念点赞表
// 注意：取消点赞时必须使用 Unscoped().Delete() 硬删除。
// 如果使用普通 Delete()，GORM 会软删除（只写 deleted_at），
// 该行仍然占用 uk_murmur_user 唯一索引，导致用户无法再次点赞。
type MurmurLike struct {
	global.MODEL
	MurmurID uint `json:"murmur_id" gorm:"uniqueIndex:uk_murmur_user"` // 碎碎念 ID
	UserID   uint `json:"user_id" gorm:"uniqueIndex:uk_murmur_user"`   // 用户 ID
	User     User `json:"-" gorm:"foreignKey:UserID"`
}
