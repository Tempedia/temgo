package temtem

import (
	"context"
	"database/sql"
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"gitlab.com/wiky.lyu/temgo/db"
	"gitlab.com/wiky.lyu/temgo/x"
)

func FindTemtems(query string, type_ []string, sort string, page, pageSize int) ([]*Temtem, int, error) {
	temtems := make([]*Temtem, 0)

	q := db.PG().NewSelect().Model(&temtems)
	if query != "" {
		q = q.WhereGroup(` AND `, func(q *bun.SelectQuery) *bun.SelectQuery {
			q = q.Where(`"name" ILIKE ?`, "%"+query+"%")
			if x.IsDigit(query) {
				q = q.WhereOr(`"no" = ?`, query)
			}
			return q
		})
	}
	if len(type_) > 0 {
		q = q.Where(`"type" @> ?`, pgdialect.Array(type_))
	}

	total, err := q.Count(context.Background())
	if err != nil {
		log.Errorf("DB Error: %v", err)
		return nil, 0, err
	}
	sortField, sortOrder := x.ParseSortParam(sort)
	if sortField != "" && sortOrder != "" {
		switch sortField {
		case "no":
			q = q.Order(fmt.Sprintf("no %s", sortOrder))
		case "name":
			q = q.Order(fmt.Sprintf("name %s", sortOrder))
		case "type1":
			q = q.OrderExpr(fmt.Sprintf("type[1] %s", sortOrder))
		case "type2":
			q = q.OrderExpr(fmt.Sprintf("type[2] %s", sortOrder))
		case "trait1":
			q = q.OrderExpr(fmt.Sprintf("traits[1] %s", sortOrder))
		case "trait2":
			q = q.OrderExpr(fmt.Sprintf("traits[2] %s", sortOrder))
		case "hp":
			q = q.OrderExpr(fmt.Sprintf(`"stats"->'HP'->'base' %s`, sortOrder))
		case "sta":
			q = q.OrderExpr(fmt.Sprintf(`"stats"->'STA'->'base' %s`, sortOrder))
		case "spd":
			q = q.OrderExpr(fmt.Sprintf(`"stats"->'SPD'->'base' %s`, sortOrder))
		case "atk":
			q = q.OrderExpr(fmt.Sprintf(`"stats"->'ATK'->'base' %s`, sortOrder))
		case "def":
			q = q.OrderExpr(fmt.Sprintf(`"stats"->'DEF'->'base' %s`, sortOrder))
		case "spatk":
			q = q.OrderExpr(fmt.Sprintf(`"stats"->'SPATK'->'base' %s`, sortOrder))
		case "spdef":
			q = q.OrderExpr(fmt.Sprintf(`"stats"->'SPDEF'->'base' %s`, sortOrder))
		case "total":
			q = q.OrderExpr(fmt.Sprintf(`(("stats"->'HP'->>'base')::int + ("stats"->'STA'->>'base')::int + ("stats"->'SPD'->>'base')::int + ("stats"->'ATK'->>'base')::int + ("stats"->'DEF'->>'base')::int + ("stats"->'SPATK'->>'base')::int + ("stats"->'SPDEF'->>'base')::int) %s`, sortOrder))
		}

	} else {
		q = q.Order(`no ASC`)
	}

	if err := q.Limit(pageSize).Offset((page - 1) * pageSize).Scan(context.Background()); err != nil {
		log.Errorf("DB Error: %v", err)
		return nil, 0, err
	}
	return temtems, total, nil
}

/* 获取temtem的进化前形态 */
func FindTemtemsEvolvesFrom(name string) ([]*Temtem, error) {
	temtems := make([]*Temtem, 0)

	if err := db.PG().NewSelect().Model(&temtems).
		Where(`evolves_to@>?::jsonb`, fmt.Sprintf(`[{"to":"%s"}]`, name)).
		Order(`no ASC`).Scan(context.Background()); err != nil {
		log.Errorf("DB Error: %v", err)
		return nil, err
	}

	return temtems, nil
}

func GetTemtemByName(name string) (*Temtem, error) {
	temtem := Temtem{Name: name}

	if err := db.PG().NewSelect().Model(&temtem).Where(`"name"=?name`).Scan(context.Background()); err != nil {
		if err != sql.ErrNoRows {
			log.Errorf("DB Error: %v", err)
			return nil, err
		}
		return nil, nil
	}
	return &temtem, nil
}
