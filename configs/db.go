package configs

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB 설정
const (
	db_type  = "mysql"
	host     = "34.64.186.156"
	port     = 3306
	user     = "root"
	password = "2022"
	dsn      = "root:2022@tcp(34.64.186.156:3306)/test?parseTime=True"
)

func ConnectDB() *gorm.DB {
	// DB에 보낼 데이터 전처리
	// DB 연결 진행
	sqlDB, err_sql := sql.Open(db_type, dsn)
	if err_sql != nil {
		fmt.Println(err_sql)
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

	// // AutoMirgate는 Model에 적용된 스키마를 기준으로 새 데이터베이스 적용시, 테이블이 없으면 생성하도록 함.
	// gormDB.AutoMigrate(&model.Collection{})
	// gormDB.AutoMigrate(&model.ProductGroup{})
	// gormDB.AutoMigrate(&model.Nft{})
	// gormDB.AutoMigrate(&model.Sale{})
	// gormDB.AutoMigrate(&model.Saleslog{})
	// gormDB.AutoMigrate(&model.Block{})
	// gormDB.AutoMigrate(&model.Obj{})
	// gormDB.AutoMigrate(&model.Obj_msg{})

	if err_gorm != nil {
		fmt.Println(err_gorm)
	}
	fmt.Println("✅ DB Connected")
	return gormDB
}

// 외부 노출을 위해 대문자 DB변수 Export하며 DB변수에는 ConnectDB 실행 Return값(DB)을 할당.
var DB *gorm.DB = ConnectDB()
