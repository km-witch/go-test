package configs

import (
	"database/sql"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB 설정
const (
	db_type2  = "mysql"
	host2     = "34.64.186.156"
	port2     = 3306
	user2     = "root"
	password2 = "2022"
	dsn2      = "root:2022@tcp(34.64.186.156:3306)/auth?parseTime=True"
)

func ConnectDB2() *gorm.DB {
	// DB에 보낼 데이터 전처리
	// DB 연결 진행
	sqlDB, err_sql := sql.Open(db_type2, dsn2)
	if err_sql != nil {
		log.Println(err_sql)
		// log.fatal(err_sql)
	}
	gormDB, err_gorm := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{
		NowFunc: func() time.Time {
			currentTime := time.Now()
			_, offset := currentTime.Zone()
			mysqlTime := currentTime.Add(time.Second * time.Duration(offset))
			return mysqlTime
		},
	})

	if err_gorm != nil {
		log.Println(err_gorm)
		// log.fatal(err_gorm)
	}
	log.Println("✅ DB2 Connected")
	return gormDB
}

// 외부 노출을 위해 대문자 DB변수 Export하며 DB변수에는 ConnectDB 실행 Return값(DB)을 할당.
var DB2 *gorm.DB = ConnectDB2()
