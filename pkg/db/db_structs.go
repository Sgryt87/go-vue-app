package db

import (
	"time"
	"database/sql"
)

type Migrations struct{
	Id int64	`xorm:"id"`
	Name string	`xorm:"name"`
	RunOn time.Time	`xorm:"run_on"`
}

type Users struct{
	Id int64	`xorm:"id"`
	First sql.NullString	`xorm:"first"`
	Last sql.NullString	`xorm:"last"`
	Email string	`xorm:"email"`
	Password string	`xorm:"password"`
}