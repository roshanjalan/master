package test

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
	"github.com/Roshan/New/graph/model"

	_ "github.com/jackc/pgx"
	_ "github.com/jinzhu/gorm/dialects/postgres"

)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "abc123"
	dbname   = "postgres"
)

func TableInsert(ctx context.Context, input *model.StudentInput) (*model.StudentResponse, error) {
	psqlinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlinfo)
	log.Println(db, err)
	if err != nil {
		panic(err)
	}

	defer db.Close()
	resultData := &model.StudentResponse{}
	_, e := IsValidDateWithDateObect(input.Dob)
	validPhn := checkPhno(*input.PhNo)
	input.BloodGroup = strings.ToUpper(input.BloodGroup)
	validBloodGrp := checkBlood(input.BloodGroup)
	*input.Gender = strings.ToUpper(*input.Gender)
	validGender := checkGender(*input.Gender)
	if e == nil && validPhn && validBloodGrp && validGender && input.ID == nil {

		insertStmt := `insert into roshan_db.Student ("firstname","lastname","gender","roll","bloodgroup","dob","phno")values($1,$2,$3,$4,$5,$6,$7)RETURNING ID`
		log.Println("write query")
		err := db.QueryRowContext(ctx, insertStmt, input.FirstName, input.LastName, input.Gender, input.Roll, input.BloodGroup,input.Dob, input.PhNo).Scan(&input.ID)
		if err != nil {
			resultData.ID = nil
			var output1 string =  "Failed to insert data into table"
			resultData.Status = append(resultData.Status, &output1)
			return resultData, err
		}
		resultData.ID = input.ID
		var output2 string = "Success - Inserted into table"
		resultData.Status = append(resultData.Status, &output2)
		return resultData, nil
	}else if input.ID != nil{
		updateStmt := `update roshan_db.Student set firstname = $1, lastname = $2, gender = $3, roll = $4, bloodgroup = $5, dob = $6, phno = $7 where id = $8`
		_,er := db.Exec(updateStmt,input.FirstName,input.LastName,input.Gender,input.Roll,input.BloodGroup,input.Dob,input.PhNo,input.ID)
		if er != nil{
			resultData.ID = nil
			var output1 string = "Failed to update data into table"
			resultData.Status = append(resultData.Status, &output1) 
			return resultData, er 
		}
		resultData.ID = input.ID
		var output2 string = "Success - Updated into table"
		resultData.Status = append(resultData.Status, &output2)
		return resultData,er 
	}
	log.Println(e)
	resultData.ID = nil
	if e != nil {
		var err1 string  = "Failed - date format"
		resultData.Status = append(resultData.Status, &err1)
	}

	if !validPhn {
		var err2 string = "Failed - phno format"
		resultData.Status = append(resultData.Status, &err2)
	}

	if !validBloodGrp {
		var err3 string  = "Failed - bloodgroup format"
		resultData.Status = append(resultData.Status, &err3)
	}

	if !validGender {
		var err4 string = "Failed - gender format"
		resultData.Status = append(resultData.Status, &err4)
	}

	//if (len(statusArray)>0){
	//resultData.Status = append(resultData.Status ,[statusArray])
	//}
	return resultData, nil

}

func IsValidDateWithDateObect(date string) (time.Time, error) {
	const (
		layoutISO = "2006-01-02"
		//layoutUS  = "2 January, 2006"
	)
	timeObj, err := time.Parse(layoutISO, date)
	if err != nil {
		return timeObj, err
	}
	return timeObj, err
}

func checkPhno(phno int) bool {
	s := strconv.Itoa(phno)
	if len(s) == 10 {
		return true
	}
	return false
}

func checkBlood(blood string) bool {
	bloodMap := make(map[string]string)

	bloodMap["group1"] = "A+"
	bloodMap["group2"] = "A-"
	bloodMap["group3"] = "B+"
	bloodMap["group4"] = "B-"
	bloodMap["group5"] = "AB+"
	bloodMap["group6"] = "AB-"
	bloodMap["group7"] = "O+"
	bloodMap["group8"] = "O-"

	for _, bgroup := range bloodMap {

		if blood == bgroup {
			return true
		}
	}
	return false
}

func checkGender(gender string) bool {
	genderMap := make(map[string]string)

	genderMap["type1"] = "MALE"
	genderMap["type2"] = "FEMALE"

	for _, gendtype := range genderMap {
		if gender == gendtype {
			return true
		}
	}
	return false
}