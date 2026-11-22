package task

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"server/global"
	"server/model/elasticsearch"
	"server/service"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/scriptlanguage"
)

// UpdateArticleViewsSyncTask 将 Redis 中的文章浏览量（增量），同步到 Elasticsearch
func UpdateArticleViewsSyncTask() error {
	// 获取redis中的缓存数据
	articleView := service.ServiceGroupApp.ArticleService.NewArticleView()

	viewsInfo := articleView.GetInfo()
	var errs []error
	for id, num := range viewsInfo {
		// 无变化就跳过，顺手清理掉这个字段
		if num == 0 {
			if err := global.Redis.HDel(articleView.Index, id).Err(); err != nil {
				errs = append(errs, fmt.Errorf("清理零值缓存 %s 失败: %w", id, err))
			}
			continue
		}

		// 更新数据 之前的数据+缓存中的数据；字段缺失时按 0 处理，避免脚本报错
		source := "ctx._source.views = (ctx._source.views ?: 0) + params.n"
		script := types.Script{
			Source: &source,
			Lang:   &scriptlanguage.Painless,
			Params: map[string]json.RawMessage{"n": json.RawMessage(strconv.Itoa(num))},
		}
		_, err := global.ESClient.Update(elasticsearch.ArticleIndex(), id).Script(&script).Do(context.TODO())
		if err != nil {
			if dropOnError(err) {
				// 文章已不存在或数据无法修复：记录后丢弃，避免永久阻塞
				errs = append(errs, fmt.Errorf("同步文章 %s 浏览量失败，已丢弃该缓存: %w", id, err))
				_ = global.Redis.HDel(articleView.Index, id).Err()
				continue
			}
			// 可恢复的错误（ES 不可用、索引不存在等）：保留缓存，等待下次重试
			errs = append(errs, fmt.Errorf("同步文章 %s 浏览量失败，保留待重试: %w", id, err))
			continue
		}

		// 同步成功后立刻删除该字段，避免下次任务重复累加
		if err := global.Redis.HDel(articleView.Index, id).Err(); err != nil {
			errs = append(errs, fmt.Errorf("同步成功但清理缓存 %s 失败: %w", id, err))
		}
	}

	return errors.Join(errs...)
}

// dropOnError 判断失败的缓存是否应该直接丢弃（不再重试）。
// 只有确定无法通过重试恢复的错误才丢弃，其余一律保留等待下次重试。
func dropOnError(err error) bool {
	var esErr *types.ElasticsearchError
	if !errors.As(err, &esErr) {
		return false // 网络错误、超时等，保留重试
	}

	// 404：只有"文档不存在"才丢弃；索引不存在等情况要保留等待恢复
	if esErr.Status == http.StatusNotFound {
		return esErr.ErrorCause.Type == "document_missing_exception"
	}

	// 4xx 中无法通过重试解决的请求错误直接丢弃；5xx、429、401、403 等保留重试
	switch esErr.Status {
	case http.StatusBadRequest, http.StatusUnsupportedMediaType, http.StatusUnprocessableEntity:
		return true
	}
	return false
}
