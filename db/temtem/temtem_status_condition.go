package temtem

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/wiky.lyu/temgo/db"
)

func FindTemtemStatusConditions(query, group string, page, pageSize int) ([]*TemtemStatusCondition, int, error) {
	conditions := make([]*TemtemStatusCondition, 0)

	q := db.PG().NewSelect().Model(&conditions).Order(`temtem_status_condition.name`)

	if query != "" {
		q = q.Where(`"temtem_status_condition"."name" ILIKE ?`, "%"+query+"%")
	}
	if group != "" {
		q = q.Where(`"temtem_status_condition"."group"=?`, group)
	}

	total, err := q.Limit(pageSize).Offset((page - 1) * pageSize).ScanAndCount(context.Background())
	if err != nil {
		log.Errorf("DB Error: %v", err)
		return nil, 0, err
	}
	return conditions, total, nil
}

func FindTemtemTechniquesByStatusCondition(condition string) ([]*TemtemTechnique, error) {
	techniques := make([]*TemtemTechnique, 0)

	if err := db.PG().NewSelect().Model(&techniques).
		Where(`"name" IN (SELECT UNNEST("techniques") FROM "temtem_status_condition" WHERE "name"=?)`, condition).
		Scan(context.Background()); err != nil {
		log.Errorf("DB Error: %v", err)
		return nil, err
	}
	return techniques, nil
}

func FindTemtemTraitsByStatusCondition(condition string) ([]*TemtemTrait, error) {
	traits := make([]*TemtemTrait, 0)

	if err := db.PG().NewSelect().Model(&traits).
		Where(`"name" IN (SELECT UNNEST("traits") FROM "temtem_status_condition" WHERE "name"=?)`, condition).
		Scan(context.Background()); err != nil {
		log.Errorf("DB Error: %v", err)
		return nil, err
	}
	return traits, nil
}
