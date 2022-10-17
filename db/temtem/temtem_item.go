package temtem

import (
	"context"

	log "github.com/sirupsen/logrus"
	"github.com/uptrace/bun"
	"gitlab.com/wiky.lyu/temgo/db"
)

func FindTemtemItems(query string, categorys []string, page, pageSize int) ([]*TemtemItem, int, error) {
	items := make([]*TemtemItem, 0)

	q := db.PG().NewSelect().Model(&items).Relation(`Category`).Relation(`Category.Parent`)
	if query != "" {
		q = q.Where(`"temtem_item"."name" ILIKE ?`, "%"+query+"%")
	}
	if len(categorys) > 0 {
		q = q.Where(`"temtem_item"."category" IN (?)`, bun.In(categorys))
	}

	total, err := q.Limit(pageSize).Offset((page - 1) * pageSize).Order(`temtem_item.sort ASC`).ScanAndCount(context.Background())
	if err != nil {
		log.Errorf("DB Error: %v", err)
		return nil, 0, err
	}
	return items, total, nil
}
