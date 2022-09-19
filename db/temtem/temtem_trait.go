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
