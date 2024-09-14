package migration

import (
	"context"
	"go_starter/internal/domain/user"

	"gorm.io/gorm"
)

func Run(ctx context.Context, db *gorm.DB) error {
	if err := db.AutoMigrate(&user.User{}); err != nil {
		return err
	}
	return nil
}
