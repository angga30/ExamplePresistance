package lovalf

import ("database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)


type Connection interface{
	GetDb() *DbPool
}


type DbPool struct {
	Db *sql.DB
}



func (db DbPool) GetDb() *DbPool {
	return &db
}

func  SetDb() *DbPool {
	db, err := sql.Open("driver-name", "database=test1")
	if err != nil{
		log.Fatal(err)
	}

	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(50)
	return &DbPool{Db:db}
}



