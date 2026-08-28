package service

import (
	"encoding/json"
	"server/config"
	"server/global"
	"server/model/other"
	"server/model/response"
	"time"

	uapi "github.com/AxT-Team/uapi-sdk-go/uapi"
	"go.uber.org/zap"
)

type HotService struct{}

// GetHotByType 抓取指定平台的热榜原始数据
func (hotService *HotService) GetHotByType(plattype string) (other.HotResponse, error) {
	var data other.HotResponse
	client := uapi.New("https://uapis.cn", "")
	params := map[string]any{
		"type": plattype,
	}
	resp, err := client.Misc().GetMiscHotboard(params)
	if err != nil {
		return other.HotResponse{}, err
	}
	b, err := json.Marshal(resp)
	if err != nil {
		return other.HotResponse{}, err
	}
	if err = json.Unmarshal(b, &data); err != nil {
		return other.HotResponse{}, err
	}
	return data, nil
}

// fetchAndCache 抓取单个平台、转成统一结构并写入 Redis
func (hotService *HotService) fetchAndCache(p config.HotPlatform) (response.PlatformHot, error) {
	data, err := hotService.GetHotByType(p.Key)
	if err != nil {
		return response.PlatformHot{}, err
	}

	ph := response.PlatformHot{
		Type: data.Type,
		Name: p.Name,
		Icon: p.Icon,
		List: data.List,
	}

	b, err := json.Marshal(ph)
	if err != nil {
		return response.PlatformHot{}, err
	}
	// Redis 只是缓存，写失败不影响返回数据
	if err := global.Redis.Set("hot:"+p.Key, b, 40*time.Minute).Err(); err != nil {
		global.Log.Error("cache hot failed:", zap.String("platform", p.Key), zap.Error(err))
	}
	return ph, nil
}

// GetHotAll 获取所有启用平台的热榜（优先读 Redis，miss 则抓取兜底）
func (hotService *HotService) GetHotAll() ([]response.PlatformHot, error) {
	var result []response.PlatformHot

	for _, p := range global.Config.Hot.Platforms {
		if !p.Enabled {
			continue
		}

		// 先读 Redis，命中直接返回
		bytes, err := global.Redis.Get("hot:" + p.Key).Bytes()
		if err == nil {
			var ph response.PlatformHot
			if json.Unmarshal(bytes, &ph) == nil {
				result = append(result, ph)
				continue
			}
		}

		// miss 或解析失败 → 兜底抓取
		ph, err := hotService.fetchAndCache(p)
		if err != nil {
			global.Log.Error("GetHotByType error:", zap.String("platform:", p.Key), zap.Error(err))
			continue
		}
		result = append(result, ph)
	}
	return result, nil
}

// FetchAll 抓取所有启用平台并写入 Redis（给 cron 定时预热用）
func (hotService *HotService) FetchAll() error {
	for _, p := range global.Config.Hot.Platforms {
		if !p.Enabled {
			continue
		}
		if _, err := hotService.fetchAndCache(p); err != nil {
			global.Log.Error("fetch hot failed:", zap.String("platform:", p.Key), zap.Error(err))
			continue
		}
	}
	return nil
}
