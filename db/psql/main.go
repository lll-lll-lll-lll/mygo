package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "postgres://mygo-postgres:mygo-postgres@localhost/mygo-postgresdb?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	getTableName(db)

}

func getTableName(db *sql.DB) {
	rows, err := db.Query(`
	SELECT table_name
	FROM information_schema.tables
	WHERE table_schema = 'public'
	ORDER BY table_name;
`)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		var clmName string
		err := rows.Scan(&clmName)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(clmName)
	}

}

// テーブルのカラムにどんな情報が含まれているかチェックする
func typeCheck(db *sql.DB) {
	rows, err := db.Query(`
		SELECT column_name, data_type, character_maximum_length
		FROM information_schema.columns
		WHERE table_name = 'users';`)
	if err != nil {
		log.Fatal(err)
	}
	clms, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(clms)
	types, err := rows.ColumnTypes()
	if err != nil {
		log.Fatal(err)
	}
	for _, t := range types {
		fmt.Println(t.ScanType().Name())
		fmt.Println(t.Nullable())
		fmt.Println(t.Name())
		fmt.Println()
		fmt.Println("")
	}
}

func getColumnType(db *sql.DB) {
	rows, err := db.Query(`
		SELECT column_name, data_type
		FROM information_schema.columns
		WHERE table_name = 'profiles';
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	// ここのデータはクエリのselect内の文字列に関すること
	// types, err := rows.ColumnTypes()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, t := range types {
	// 	fmt.Println(t.Name())
	// 	fmt.Println(t.ScanType())
	// 	fmt.Println("end")
	// }
	columns := make([]struct {
		ColumnName string
		DataType   string
	}, 0)

	// 実際のカラムごとのタイプを知れる
	for rows.Next() {
		var column struct {
			ColumnName string
			DataType   string
		}

		err := rows.Scan(&column.ColumnName, &column.DataType)
		if err != nil {
			log.Fatal(err)
		}

		columns = append(columns, column)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	for _, column := range columns {
		fmt.Println("Column Name:", column.ColumnName)
		fmt.Println("Data Type:", column.DataType)
		fmt.Println("------")
	}

}

// func columnCheck(db *sql.DB) {
// 	tableName := "profiles"
// 	rows, err := db.Query(fmt.Sprintf("select user_name from %s;", tableName))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	types, err := rows.Columns()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	for _, t := range types {
// 		fmt.Println(t)
// 		fmt.Println("")
// 	}

// }
