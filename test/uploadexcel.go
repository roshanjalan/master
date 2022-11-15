package test

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"github.com/Roshan/New/graph/model"
	_ "github.com/jackc/pgx"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/xuri/excelize/v2"
)

const (
	host2     = "localhost"
	port2     = 5432
	user2     = "postgres"
	password2 = "abc123"
	dbname2   = "postgres"
)


func UploadExcel(ctx context.Context) (*model.ExcelOutput, error) {
	psqlinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host2, port2, user2, password2, dbname2)
	db, err := sql.Open("postgres", psqlinfo)
	log.Println(db, err)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	var stat string
	//exceldata := &model.ExcelOutput{}

	f, err := excelize.OpenFile("E://roshan.xlsx")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rows, err := f.GetRows("student")
	if err != nil {
		panic(err)
	}

	for row, col := range rows {
		if row == 0 {
			continue
		}
		if len(col) > 1 {
			if col[0] == "" {
				log.Println("for insert")
				insertStmt := `insert into roshan_db.Student ("firstname","lastname","gender","roll","bloodgroup","dob","phno")values($1,$2,$3,$4,$5,$6,$7)`
				_, e := db.Query(insertStmt, col[1], col[2], col[3], col[4], col[5], col[6], col[7])
				log.Println(e)
				if e != nil {
					stat = "Failed to insert"
				}
				stat = "Success"
			} else if col[0] != "" {
				log.Println("for update")
				updateStmt := `update roshan_db.Student set firstname = $1, lastname = $2, gender = $3, roll = $4, bloodgroup = $5, dob = $6, phno = $7 where id = $8`
				log.Println("The query for update - ", updateStmt)
				_, er := db.Exec(updateStmt, col[1], col[2], col[3], col[4], col[5], col[6], col[7], col[0])
				log.Println("Error if any is-", er)
				if er != nil {
					stat = "Failed to update"
				}
				stat = "Success"

			}
		 } else {
		 	deleteStmt := `delete from roshan_db.Student where id =$1`
			_, er := db.Exec(deleteStmt, col[0])
		 	if er != nil {
		 		stat = "Failed to delete"
		 	}
		 	stat = "Success"
		 }

	}
	return &model.ExcelOutput{Status: &stat}, nil
	
}