package service

import (
	"server/global"
	"server/model/database"
	"server/model/request"
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
	err := global.DB.Where("id IN ?", ids.Ids).Delete(&database.Murmur{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (m *MurmurService) MurmurList() ([]database.Murmur, error) {
	var murmurList []database.Murmur
	err := global.DB.Order("created_at desc").Find(&murmurList).Error
	if err != nil {
		return []database.Murmur{}, err
	}
	return murmurList, nil
}
