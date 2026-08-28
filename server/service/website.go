package service

import (
	"server/global"
	"server/model/database"
)

type WebsiteService struct {
}

func (websiteService *WebsiteService) WebsiteFooterLink() []database.FooterLink {
	var footerLinks []database.FooterLink
	global.DB.Find(&footerLinks)
	return footerLinks
}

func (websiteService *WebsiteService) WebsiteCreateFooterLink(req database.FooterLink) error {
	return global.DB.Save(&req).Error
}

func (websiteService *WebsiteService) WebsiteDeleteFooterLink(req database.FooterLink) error {
	return global.DB.Delete(&req).Error
}
