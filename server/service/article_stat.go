package service

import (
	"server/global"
	"strconv"
)

func (articleService *ArticleService) NewArticleView() CountDB {
	return CountDB{
		Index: "article_views",
	}
}

type CountDB struct {
	Index string
}

// Incr 在原有基础上加一并返回最新值
func (c CountDB) Incr(id string) (int, error) {
	num, err := global.Redis.HIncrBy(c.Index, id, 1).Result()
	return int(num), err
}

// Set 在原有基础上加一
func (c CountDB) Set(id string) error {
	_, err := c.Incr(id)
	return err
}

// GetInfo 取出数据
func (c CountDB) GetInfo() map[string]int {
	var Info = map[string]int{}
	maps := global.Redis.HGetAll(c.Index).Val()
	for id, val := range maps {
		num, _ := strconv.Atoi(val)
		Info[id] = num
	}
	return Info
}

// Clear 清除数据
func (c CountDB) Clear() {
	global.Redis.Del(c.Index)
}
