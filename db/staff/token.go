package staff

import (
	"context"
	"database/sql"
	"time"

	log "github.com/sirupsen/logrus"
	"gitlab.com/wiky.lyu/temgo/db"
)

func GetStaffToken(id string) (*StaffToken, error) {
	token := StaffToken{ID: id}
	if err := db.PG().NewSelect().Model(&token).Relation("Staff").
		WherePK().Scan(context.Background()); err != nil {
		if err != sql.ErrNoRows {
			log.Errorf("DB Error: %v", err)
			return nil, err
		}
		return nil, nil
	}
	return &token, nil
}

func CreateStaffToken(staffID int64, expiresAt time.Time, device, ip string) (*StaffToken, error) {
	tx, err := db.Begin()
	if err != nil {
		log.Errorf("DB Error: %v", err)
		return nil, err
	}
	defer tx.Rollback()

	if _, err := tx.NewUpdate().Model((*StaffToken)(nil)).
		Where(`"staff_id"=?`, staffID).Where(`"status"=?`, StaffTokenStatusOK).
		Set(`"status"=?`, StaffTokenStatusInvalid).Exec(context.Background()); err != nil {
		log.Errorf("DB Error: %v", err)
		return nil, err
	}

	token := StaffToken{
		StaffID:   staffID,
		Device:    device,
		IP:        ip,
		Status:    StaffTokenStatusOK,
		ExpiresAt: expiresAt,
	}
	if _, err := tx.NewInsert().Model(&token).Exec(context.Background()); err != nil {
		log.Errorf("DB Error: %v", err)
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		log.Errorf("DB Error: %v", err)
		return nil, err
	}

	return &token, nil
}
