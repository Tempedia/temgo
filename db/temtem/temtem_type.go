package temtem

import (
	"context"
	"database/sql"

	log "github.com/sirupsen/logrus"
	"gitlab.com/wiky.lyu/temgo/db"
)

// 获取temtem属性列表
func FindTemtemTypes() ([]*TemtemType, error) {
	types := make([]*TemtemType, 0)
	if err := db.PG().NewSelect().Model(&types).Order(`sort ASC`).Scan(context.Background()); err != nil {
		log.Errorf("DB Error: %v", err)
		return nil, err
	}
	return types, nil
}

// 获取属性详情
func GetTemtemType(name string) (*TemtemType, error) {
	t := TemtemType{Name: name}
	if err := db.PG().NewSelect().Model(&t).WherePK().Scan(context.Background()); err != nil {
		if err != sql.ErrNoRows {
			log.Errorf("DB Error: %v", err)
			return nil, err
		}
		return nil, nil
	}
	return &t, nil
}
