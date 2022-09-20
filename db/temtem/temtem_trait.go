package temtem

import (
	"context"
	"database/sql"

	log "github.com/sirupsen/logrus"
	"gitlab.com/wiky.lyu/temgo/db"
)

func GetTemtemTrait(name string) (*TemtemTrait, error) {
	trait := TemtemTrait{Name: name}

	if err := db.PG().NewSelect().Model(&trait).WherePK().Scan(context.Background()); err != nil {
		if err != sql.ErrNoRows {
			log.Errorf("DB Error: %v", err)
			return nil, err
		}
		return nil, nil
	}
	return &trait, nil
}

func FindTemtemTraits(query string, page, pageSize int) ([]*TemtemTrait, int, error) {
	traits := make([]*TemtemTrait, 0)

	q := db.PG().NewSelect().Model(&traits)

	if query != "" {
		q = q.Where(`"name" ILIKE ?`, "%"+query+"%")
	}

	q = q.Order(`name ASC`)
	total, err := q.Limit(pageSize).Offset((page - 1) * pageSize).ScanAndCount(context.Background())
	if err != nil {
		log.Errorf("DB Error: %v", err)
		return nil, 0, err
	}
	return traits, total, nil
}
