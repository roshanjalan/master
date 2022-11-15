package main
 
import (
	"encoding/csv"
	"log"
	"os"
)
 
func main() {
	rows := [][]string{
		{"Name", "City", "Language"},
		{"Aman", "Kolkata", "Python"},
		{"Roshan", "Kolkata", "Golang"},
		{"Ravi", "Kolkata", "Aws"},
	}
 
	csvfile, err := os.Create("New.CSV File")
 
	if err != nil {
		log.Fatalf("failed to creating csv file: %s", err)
	}
 
	csvwriter := csv.NewWriter(csvfile)
 
	for _, row := range rows {
		_ = csvwriter.Write(row)
	}
 
	csvwriter.Flush()
 
	csvfile.Close()
}

// package main

// import (
// 	"log"
// 	"os"
// 	"fmt"
// )
// func readCurrentFile() {
// 	file, err := os.Open(".")
// 	if err != nil {
// 		log.Fatalf("failed to opening file: %s", err)
// 	}
// 	defer file.Close()

// 	list, _ := file.Readdirnames(0) // 0 to read all files and folders
// 	for _, name := range list {
// 		fmt.Println(name)
// 	}
// }

// func main() {
// 	readCurrentFile()
// }

// package main
 
// import (
//     "fmt"
//     "strconv"
//     )

// type rectangle struct {
// 	length  int
// 	breadth int
// 	color   string
 
// 	geometry struct {
// 		area      int
// 		perimeter int
// 	}
// }
 
// func main() {
// 	var r rectangle
// 	fmt.Println("Enter the length :")
// 	fmt.Scanln(&r.length)
// 	fmt.Println("Enter the breadth :")
// 	fmt.Scanln(&r.breadth)
// 	fmt.Println("Enter the color :")
// 	fmt.Scanln(&r.color)
//     if _, err := strconv.Atoi(r.color); err == nil {
//         fmt.Println("color format is not in integer, provide character")
//     } else {
// 	r.geometry.area = r.length * r.breadth
// 	r.geometry.perimeter = 2 * (r.length + r.breadth)
//     fmt.Println("length = ", r.length)
//     fmt.Println("breadth = ",r.breadth)
//     fmt.Println("color = ", r.color)
// 	fmt.Println("area = ", r.geometry.area)
// 	fmt.Println("Perimeter = ", r.geometry.perimeter)
// }
// }

