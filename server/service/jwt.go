package service

import (
	"server/global"
	"server/model/database"
	"server/utils"

	"github.com/gofrs/uuid"
	"go.uber.org/zap"
)

type JwtService struct {
}

func (j *JwtService) SetRedisJWT(jwt string, uuid uuid.UUID) error {
	dr, err := utils.ParseDuration(global.Config.Jwt.RefreshTokenExpiryTime)
	if err != nil {
		return err
	}
	return global.Redis.Set(uuid.String(), jwt, dr).Err()
}

func (jwtService *JwtService) GetRedisJWT(uuid uuid.UUID) (string, error) {
	// 从Redis获取指定uuid对应的JWT
	return global.Redis.Get(uuid.String()).Result()
}

func (jwtService *JwtService) JoinInBlacklist(jwtList database.JwtBlacklist) error {
	err := global.DB.Create(&jwtList).Error
	if err != nil {
		return err
	}
	global.BlackCache.SetDefault(jwtList.Jwt, struct{}{})
	return nil
}

func (jwtService *JwtService) IsInBlacklist(jwt string) bool {
	// 从黑名单缓存中检查JWT是否存在
	_, ok := global.BlackCache.Get(jwt)
	return ok
}

func LoadAll() {
	var data []string
	// 从数据库中获取所有的黑名单JWT
	if err := global.DB.Model(&database.JwtBlacklist{}).Pluck("jwt", &data).Error; err != nil {
		// 如果获取失败，记录错误日志
		global.Log.Error("Failed to load JWT blacklist from the database", zap.Error(err))
		return
	}
	// 将所有JWT添加到BlackCache缓存中
	for i := 0; i < len(data); i++ {
		global.BlackCache.SetDefault(data[i], struct{}{})
	}
}
