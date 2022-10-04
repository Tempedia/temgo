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

/* 获取temtem所在区域 */
func FindTemtemLocationsByTemtem(temname string) ([]*TemtemLocation, error) {
	locations := make([]*TemtemLocation, 0)

	if err := db.PG().NewSelect().Model(&locations).
		Where(`"name" IN (SELECT "location" FROM "temtem_location_area",jsonb_array_elements("temtems") j WHERE j->>'name' = ? or j->>'name' ILIKE ?)`, temname, temname+" (%").
		Scan(context.Background()); err != nil {
		log.Errorf("DB Error: %v", err)
		return nil, err
	}
	return locations, nil
}
