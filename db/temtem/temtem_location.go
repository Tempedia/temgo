package temtem

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/wiky.lyu/temgo/db"
)

func FindTemtemLocations(query string, page, pageSize int) ([]*TemtemLocation, int, error) {
	locations := make([]*TemtemLocation, 0)

	q := db.PG().NewSelect().Model(&locations).Order(`name`)
	if query != "" {
		q = q.Where(`"name" ILIKE ?`, "%"+query+"%")
	}

	total, err := q.Limit(pageSize).Offset((page - 1) * pageSize).ScanAndCount(context.Background())
	if err != nil {
		log.Errorf("DB Error: %v", err)
		return nil, 0, err
	}
	return locations, total, nil
}

func FindTemtemLocationAreasByLocation(location string) ([]*TemtemLocationArea, error) {
	areas := make([]*TemtemLocationArea, 0)

	if err := db.PG().NewSelect().Model(&areas).Where(`"location"=?`, location).Order(`id`).Scan(context.Background()); err != nil {
		log.Errorf("DB Error: %v", err)
		return nil, err
	}
	return areas, nil
}
