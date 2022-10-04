package temtem

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/wiky.lyu/temgo/db"
)

/* 获取temtem的升级技能 */
func FindTemtemLevelingUpTechniques(name string) ([]*TemtemLevelingUpTechnique, error) {
	techniques := make([]*TemtemLevelingUpTechnique, 0)

	if err := db.PG().NewSelect().Model(&techniques).Relation(`Technique`).
		Where(`"temtem"=?`, name).Order(`id ASC`).Scan(context.Background()); err != nil {
		log.Errorf("DB Error: %v", err)
		return nil, err
	}
	return techniques, nil
}

func FindTemtemCourseTechniques(name string) ([]*TemtemCourseTechnique, error) {
	techniques := make([]*TemtemCourseTechnique, 0)

	if err := db.PG().NewSelect().Model(&techniques).Relation(`Technique`).
		Where(`"temtem"=?`, name).Order(`course ASC`).Scan(context.Background()); err != nil {
		log.Errorf("DB Error: %v", err)
		return nil, err
	}
	return techniques, nil
}

func FindTemtemBreedingTechniques(name string) ([]*TemtemBreedingTechnique, error) {
	techniques := make([]*TemtemBreedingTechnique, 0)

	if err := db.PG().NewSelect().Model(&techniques).Relation(`Technique`).
		Where(`"temtem"=?`, name).Scan(context.Background()); err != nil {
		log.Errorf("DB Error: %v", err)
		return nil, err
	}
	return techniques, nil
}

func FindTemtemTechniques(query string, types []string, class string, page, pageSize int) ([]*TemtemTechnique, int, error) {
	techniques := make([]*TemtemTechnique, 0)
	q := db.PG().NewSelect().Model(&techniques).Order(`name`)

	if len(types) >= 1 {
		q = q.Where(`"type"=?`, types[0])
	}
	if len(types) >= 2 {
		q = q.Where(`"synergy_type"=?`, types[1])
	}
	if class != "" {
		q = q.Where(`"class"=?`, class)
	}

	total, err := q.Limit(pageSize).Offset((page - 1) * pageSize).ScanAndCount(context.Background())
	if err != nil {
		log.Errorf("DB Error: %v", err)
		return nil, 0, err
	}
	return techniques, total, nil
}
