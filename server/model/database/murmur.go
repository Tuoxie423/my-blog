package database

import "server/global"

type Murmur struct {
	global.MODEL        // 自带 id / created_at / updated_at
	Content      string `json:"content"` // 碎碎念内容
}
