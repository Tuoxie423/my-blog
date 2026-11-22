package service

import (
	"errors"
	"server/global"
	"server/model/database"
	"server/model/request"
	"server/model/response"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MurmurService struct{}

func (m *MurmurService) MurmurCreate(mur database.Murmur) error {
	err := global.DB.Create(&mur).Error
	if err != nil {
		return err
	}
	return nil
}

func (m *MurmurService) MurmurDelete(ids request.MurmurDelete) error {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		// 碎碎念本身是软删除
		if err := tx.Where("id IN ?", ids.Ids).Delete(&database.Murmur{}).Error; err != nil {
			return err
		}
		// 同步硬删除赞同记录，避免留下无法再被访问的孤儿数据
		return tx.Unscoped().
			Where("murmur_id IN ?", ids.Ids).
			Delete(&database.MurmurLike{}).Error
	})
}

func (m *MurmurService) MurmurList() ([]database.Murmur, error) {
	var murmurList []database.Murmur
	err := global.DB.Order("created_at desc").Find(&murmurList).Error
	if err != nil {
		return []database.Murmur{}, err
	}
	if len(murmurList) == 0 {
		return murmurList, nil
	}

	ids := make([]uint, 0, len(murmurList))
	for _, v := range murmurList {
		ids = append(ids, v.ID)
	}

	// 一次 GROUP BY 查出所有碎碎念的赞同数，避免逐条查询（N+1）
	var rows []struct {
		MurmurID  uint
		LikeCount int64
	}
	err = global.DB.Model(&database.MurmurLike{}).
		Select("murmur_id, count(*) AS like_count").
		Where("murmur_id IN ?", ids).
		Group("murmur_id").
		Scan(&rows).Error
	if err != nil {
		return []database.Murmur{}, err
	}

	countMap := make(map[uint]int64, len(rows))
	for _, row := range rows {
		countMap[row.MurmurID] = row.LikeCount
	}
	for i := range murmurList {
		murmurList[i].Likes = countMap[murmurList[i].ID]
	}

	return murmurList, nil
}

// MurmurLike 切换赞同状态。
// 返回操作后的最终状态和赞同数，前端以此为准。
func (m *MurmurService) MurmurLike(req request.MurmurLike) (response.MurmurLikeResult, error) {
	var result response.MurmurLikeResult

	err := global.DB.Transaction(func(tx *gorm.DB) error {
		// 确认碎碎念存在（软删除的记录会被 GORM 自动排除）
		var murmur database.Murmur
		if err := tx.First(&murmur, req.MurmurID).Error; err != nil {
			return err
		}

		// 查询当前用户是否已赞同
		var like database.MurmurLike
		err := tx.Where("murmur_id = ? AND user_id = ?", req.MurmurID, req.UserID).
			First(&like).Error

		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			// 未赞同：新建。OnConflict 兜住并发重复插入，避免唯一键报错
			if err := tx.Clauses(clause.OnConflict{DoNothing: true}).
				Create(&database.MurmurLike{
					MurmurID: req.MurmurID,
					UserID:   req.UserID,
				}).Error; err != nil {
				return err
			}
		case err != nil:
			return err
		default:
			// 已赞同：取消。必须硬删除，否则软删除行会占用唯一索引
			if err := tx.Unscoped().Delete(&like).Error; err != nil {
				return err
			}
		}

		// 读回最终状态（使用新变量，避免旧结构体主键影响查询条件）
		var after database.MurmurLike
		err = tx.Where("murmur_id = ? AND user_id = ?", req.MurmurID, req.UserID).
			First(&after).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		result.IsLiked = err == nil

		// 统计最新赞同数
		return tx.Model(&database.MurmurLike{}).
			Where("murmur_id = ?", req.MurmurID).
			Count(&result.Likes).Error
	})

	return result, err
}

// MurmurLiked 返回当前用户点过赞的碎碎念 ID 列表。
// 公开列表接口拿不到用户身份，由前端调用该接口后与列表合并。
func (m *MurmurService) MurmurLiked(userID uint) ([]uint, error) {
	ids := make([]uint, 0)
	err := global.DB.Model(&database.MurmurLike{}).
		Where("user_id = ?", userID).
		Pluck("murmur_id", &ids).Error
	if err != nil {
		return nil, err
	}
	return ids, nil
}
