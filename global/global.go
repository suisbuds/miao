package global

import (
	"github.com/suisbuds/miao/pkg/logger"
	"github.com/suisbuds/miao/pkg/setting"
	"github.com/suisbuds/miao/pkg/validator"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Database
var(
	DBEngine *gorm.DB
)

// Settings
var (
	ServerSetting   *setting.ServerSetting
	AppSetting      *setting.AppSetting
	DatabaseSetting *setting.DatabaseSetting
	Logger 		*logger.Logger
	Zapper *zap.SugaredLogger
)

// Validator
var (
	Validator *validator.MiaoValidator
)