package test

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"github.com/Roshan/New/graph/model"
	_ "github.com/jackc/pgx"
	_ "github.com/jinzhu/gorm/dialects/postgres"

)

const (
	host1     = "localhost"
	port1     = 5432
	user1     = "postgres"
	password1 = "abc123"
	dbname1   = "postgres"
)

func FetchData(ctx context.Context, input *model.StudentData) ([]*model.StudentOutput, error) {
	psqlinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host1, port1, user1, password1, dbname1)
	db, err := sql.Open("postgres", psqlinfo)
	log.Println(db, err)
	if err != nil {
		panic(err)
	}
defer db.Close()
if input.Inp != nil{
fetchRow := `SELECT * FROM roshan_db.student WHERE firstname LIKE '%`+*input.Inp+`%'`
row, err := db.Query(fetchRow)

rowEntry := []*model.StudentOutput{}
if err != nil {
	return nil, fmt.Errorf("RowFromQuery %v: %v",input.Inp,err)
}

for row.Next() {
	resultRow := &model.StudentOutput{}
	 err := row.Scan(&resultRow.ID, &resultRow.FirstName, &resultRow.LastName, &resultRow.Gender, &resultRow.Roll, &resultRow.BloodGroup, &resultRow.Dob, &resultRow.PhNo)
	 if err != nil {
	return nil,  fmt.Errorf("RowFromQuery %v: %v",input.Inp,err)
}
if resultRow.LastName == nil {
	a := ""
	resultRow.LastName = &a 
}
if resultRow.Gender == nil {
	b := ""
	resultRow.Gender = &b
}
rowEntry = append(rowEntry, resultRow)
}
return rowEntry, nil 
}
fetchRow := `SELECT * FROM roshan_db.student`
row, err := db.Query(fetchRow)

rowEntry := []*model.StudentOutput{}
if err != nil {
	return nil, fmt.Errorf("RowFromQuery %v: %v",input.Inp,err)
}

for row.Next() {
	resultRow := &model.StudentOutput{}
	 err := row.Scan(&resultRow.ID, &resultRow.FirstName, &resultRow.LastName, &resultRow.Gender, &resultRow.Roll, &resultRow.BloodGroup, &resultRow.Dob, &resultRow.PhNo)
	 if err != nil {
	return nil,  fmt.Errorf("RowFromQuery %v: %v",input.Inp,err)
}
if resultRow.LastName == nil {
	a := ""
	resultRow.LastName = &a 
}
if resultRow.Gender == nil {
	b := ""
	resultRow.Gender = &b
}
rowEntry = append(rowEntry, resultRow)
}
return rowEntry, nil 
}