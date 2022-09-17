package temtem

import (
	"context"

	log "github.com/sirupsen/logrus"
	"github.com/uptrace/bun"
	"gitlab.com/wiky.lyu/temgo/db"
)

func FindTemtems(query string, page, pageSize int) ([]*Temtem, int, error) {
	temtems := make([]*Temtem, 0)

	q := db.PG().NewSelect().Model(&temtems).Order(`no ASC`)
	if query != "" {
		q = q.WhereGroup(` AND `, func(q *bun.SelectQuery) *bun.SelectQuery {
			q = q.Where(`"name" ILIKE ?`, "%"+query+"%")
			return q
		})
	}

	total, err := q.Count(context.Background())
	if err != nil {
		log.Errorf("DB Error: %v", err)
		return nil, 0, err
	}

	if err := q.Limit(pageSize).Offset((page - 1) * pageSize).Scan(context.Background()); err != nil {
		log.Errorf("DB Error: %v", err)
		return nil, 0, err
	}
	return temtems, total, nil
}
