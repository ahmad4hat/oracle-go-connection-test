package main 

import (
	"database/sql" // offical database sql package
	"fmt"

	_ "github.com/godror/godror" // not using goracle.v2 because its deprecated
)

type rowElement struct {
	grade string
	losal string
	hisal string
}

func (r rowElement) printall() {
	fmt.Println(r.grade, r.losal, r.hisal)
}

func main() {

	db, err := sql.Open("godror", "scott/tiger@localhost:1521/ORCLCDB.localdomain") // ("gordor "=database driver ,"scott/tiger.........."=conection string )
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close() // defer delays this line until the end

	rows, err := db.Query("SELECT * FROM SALGRADE ")
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		return
	}

	//var hecx string

	fmt.Println("garde ** losal ** hisal ")
	var allsalarary []rowElement

	for rows.Next() {
		var (
			grade string
			losal string
			hisal string
		)
		rows.Scan(&grade, &losal, &hisal)

		a := rowElement{
			grade: grade,
			losal: losal,
			hisal: hisal,
		}
		allsalarary = append(allsalarary, a)
	}

	for _, salarayRange := range allsalarary {
		salarayRange.printall()
	}

	defer rows.Close()

}
