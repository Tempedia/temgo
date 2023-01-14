package temtem

import (
	"context"
	"database/sql"

	log "github.com/sirupsen/logrus"
	"gitlab.com/wiky.lyu/temgo/db"
)

func CreateUserTeam(name string, temtems interface{}) (*TemtemUserTeam, error) {
	team := TemtemUserTeam{
		Name:    name,
		Temtems: temtems,
	}
	if _, err := db.PG().NewInsert().Model(&team).Exec(context.Background()); err != nil {
		log.Errorf("DB Error: %v", err)
		return nil, err
	}
	return &team, nil
}

func GetTemtemUserTeam(id string) (*TemtemUserTeam, error) {
	team := TemtemUserTeam{ID: id}

	if err := db.PG().NewSelect().Model(&team).WherePK().Scan(context.Background()); err != nil {
		if err != sql.ErrNoRows {
			log.Errorf("DB Error: %v", err)
			return nil, err
		}
		return nil, nil
	}
	return &team, nil
}
