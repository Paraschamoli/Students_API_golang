package sqlite

import (
	"database/sql"

	"github.com/Paraschamoli/students_API/internal/config"
	_"github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*Sqlite,error){
db,err:=sql.Open("sqlite3",cfg.StoragePath)
if err!=nil{
	return nil,err
}
}