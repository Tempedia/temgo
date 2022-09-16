package staff

import (
	"context"
	"database/sql"

	log "github.com/sirupsen/logrus"
	"gitlab.com/wiky.lyu/temgo/db"
	"gitlab.com/wiky.lyu/temgo/x"
)

func GetStaffByUsername(username string) (*Staff, error) {
	staff := Staff{}

	if err := db.PG().NewSelect().Model(&staff).Where(`"username"=?`, username).Limit(1).Scan(context.Background()); err != nil {
		if err != sql.ErrNoRows {
			log.Errorf("DB Error: %v", err)
			return nil, err
		}
		return nil, nil
	}
	return &staff, nil
}

func GetSuperStaff() (*Staff, error) {
	staff := Staff{}

	if err := db.PG().NewSelect().Model(&staff).Where(`"is_superuser"=?`, true).Limit(1).Scan(context.Background()); err != nil {
		if err != sql.ErrNoRows {
			log.Errorf("DB Error: %v", err)
			return nil, err
		}
		return nil, nil
	}
	return &staff, nil
}

func CreateSuperStaff(username, password, name, phone, email string) (*Staff, error) {
	tx, err := db.Begin()
	if err != nil {
		log.Errorf("DB Error: %v", err)
		return nil, err
	}
	defer tx.Rollback()
	staff := Staff{}
	if err := tx.NewSelect().Model(&staff).Where(`"is_superuser"=?`, true).Limit(1).Scan(context.Background()); err != nil {
		if err != sql.ErrNoRows {
			log.Errorf("DB Error: %v", err)
			return nil, err
		}
	} else {
		/* 超级用户已经存在，不创建直接返回 */
		return &staff, nil
	}

	staff = Staff{
		Username:    username,
		Name:        name,
		Phone:       phone,
		Email:       email,
		Salt:        x.RandomString(8),
		PType:       randomPType(),
		IsSuperuser: true,
	}
	staff.Password = encrypt(staff.Salt, password, staff.PType)
	if _, err := tx.NewInsert().Model(&staff).Exec(context.Background()); err != nil {
		log.Errorf("DB Error: %v", err)
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		log.Errorf("DB Error: %v", err)
		return nil, err
	}
	return &staff, nil
}
