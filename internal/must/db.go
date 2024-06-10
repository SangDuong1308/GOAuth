package must

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDb(conStr string) *gorm.DB {
	sqlDb, err := sql.Open("mysql", conStr)

	if err != nil {
		panic(err)
	}

	gormDb, err := gorm.Open(
		mysql.New(mysql.Config{
			Conn: sqlDb,
		}),
		&gorm.Config{})
	if err != nil {
		panic(err)
	}

	return gormDb
}
