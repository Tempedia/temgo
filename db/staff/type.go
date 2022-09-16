package staff

import (
	"time"

	"github.com/uptrace/bun"
)

const (
	StaffStatusOK     = "OK"
	StaffStatusBanned = "Banned"
)

type Staff struct {
	bun.BaseModel `bun:"table:staff"`

	ID       int64  `bun:"id,notnull,pk" json:"id"`
	Username string `bun:"username,notnull" json:"username"`
	Name     string `bun:"name,notnull,nullzero" json:"name"`

	Salt     string `bun:"salt,notnull,nullzero" json:"-"`
	PType    string `bun:"ptype,notnull,nullzero" json:"-"`
	Password string `bun:"password,notnull,nullzero" json:"-"`

	Status      string `bun:"status,notnull,nullzero" json:"status"`
	IsSuperuser bool   `bun:"is_superuser,notnull,nullzero" json:"is_superuser"`
	Phone       string `bun:"phone,notnull,nullzero" json:"phone"`
	Email       string `bun:"email,notnull,nullzero" json:"email"`

	CreatedBy int64     `bun:"created_by,nullzero,notnull" json:"created_by"`
	CreatedAt time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `bun:"updated_at,nullzero,notnull,default:current_timestamp" json:"updated_at"`
}

const (
	StaffTokenStatusOK      = "OK"
	StaffTokenStatusInvalid = "Invalid"
)

type StaffToken struct {
	bun.BaseModel `bun:"table:staff_token"`

	ID        string    `bun:"id,pk,default:gen_random_uuid()" json:"id"`
	StaffID   int64     `bun:"staff_id,notnull" json:"staff_id"`
	Device    string    `bun:"device,notnull,nullzero" json:"device"`
	IP        string    `bun:"ip,notnull,nullzero" json:"ip"`
	ExpiresAt time.Time `bun:"expires_at,notnull,nullzero" json:"expires_at"`
	Status    string    `bun:"status,notnull,nullzero" json:"status"`

	CreatedAt time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `bun:"updated_at,nullzero,notnull,default:current_timestamp" json:"updated_at"`

	Staff *Staff `bun:"rel:belongs-to,join:staff_id=id"`
}
