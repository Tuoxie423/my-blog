package task

import (
	"server/global"
	"server/service"

	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

func RegisterScheduledTasks(c *cron.Cron) error {
	if _, err := c.AddFunc("@hourly", func() {
		if err := UpdateArticleViewsSyncTask(); err != nil {
			global.Log.Error("Failed to update article views:", zap.Error(err))
		}
	}); err != nil {
		return err
	}
	if _, err := c.AddFunc("@every 30m", func() {
		if err := service.ServiceGroupApp.HotService.FetchAll(); err != nil {
			global.Log.Error("Failed to fetch hot list:", zap.Error(err))
		}
	}); err != nil {
		return err
	}
	return nil
}
