package temtem

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/wiky.lyu/temgo/db"
)

// 获取temtem属性列表
func FindTemtemTypes() ([]*TemtemType, error) {
	types := make([]*TemtemType, 0)
	if err := db.PG().NewSelect().Model(&types).Scan(context.Background()); err != nil {
		log.Errorf("DB Error: %v", err)
		return nil, err
	}
	return types, nil
}
