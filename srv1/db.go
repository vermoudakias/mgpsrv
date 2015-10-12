package main

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

var dbDriver    string = "postgres"
var defaultUser string = "postgres"
var defaultHost string = "localhost"
var defaultPass string = ""

func SimpleDbNew(cfg map[string]string) *sql.DB {
	var v string
	var found bool
	dsn := "sslmode=disable"
	if v, found = cfg["user"]; found == false {
		v = defaultUser
	}
	if len(v) > 0 {
		dsn = fmt.Sprintf("%s user=%s", dsn, v)
	}
	if v, found = cfg["password"]; found == false {
		v = defaultPass
	}
	if len(v) > 0 {
		dsn = fmt.Sprintf("%s password=%s", dsn, v)
	}
	if v, found = cfg["host"]; found == false {
		v = defaultHost
	}
	if len(v) > 0 {
		dsn = fmt.Sprintf("%s host=%s", dsn, v)
	}
	if v, found = cfg["dbname"]; found == false {
		log.Fatal("No database name given")
	}
	dsn = fmt.Sprintf("%s dbname=%s", dsn, v)
	fmt.Printf("DB: driver <%s>, data source <%s>\n", dbDriver, dsn)
	db, err := sql.Open(dbDriver, dsn)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}

func DbTest(table string) {
	var dbCfg = map[string]string{
		"user":   "postgres",
		"dbname": "test_db",
	}
	db := SimpleDbNew(dbCfg)
	qs := fmt.Sprintf("SELECT * FROM %s", table)
	rows, err := db.Query(qs)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	cols, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("DB: Columns of table %s: %s\n", table, cols)
	m := make(map[string]int)
	for i, c := range cols {
		m[c] = i
	}
	vals := make([]interface{}, len(cols))
	for i, _ := range vals {
		var val interface{}
		vals[i] = &val
	}
	var count int
	for rows.Next() {
		//var r_id, r_active, r_dlr, r_max_dlr int
		//var r_name, r_type string
		//if err := rows.Scan(&r_id, &r_name, &r_active, &r_type, &r_dlr, &r_max_dlr); err != nil {
		//	log.Fatal(err)
		//}
		//fmt.Printf("DB: Gateway %s, id %d, type %s\n", r_name, r_id, r_type)
		if err := rows.Scan(vals...); err != nil {
			log.Fatal(err)
		}
		count += 1
		s := "DB:"
		for i, colName := range cols {
			var v = *(vals[i].(*interface{}))
			var v_type = reflect.TypeOf(v)
			switch v.(type) {
			case int64:
				s = fmt.Sprintf("%s %s=%d,", s, colName, v)
			case []uint8:
				s = fmt.Sprintf("%s %s=%s,", s, colName, v)
			case time.Time:
				s = fmt.Sprintf("%s %s=%s,", s, colName, (v.(time.Time)).Format(time.RFC3339))
			default:
				s = fmt.Sprintf("%s %s=%s [%s],", s, colName, v, v_type)
			}
		}
		s = strings.TrimRight(s, ",")
		fmt.Println(s)
	}
	fmt.Printf("DB: Found %d rows\n", count)
}

//func DbGet(db *sql.DB, table string, id string) []interface{} {
func DbGet(db *sql.DB, table string, id string) {
	var qs string
	if len(id) > 0 {
		qs = fmt.Sprintf("SELECT * FROM %s where id = %s", table, id)
	} else {
		qs = fmt.Sprintf("SELECT * FROM %s", table)
	}
	rows, err := db.Query(qs)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	cols, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("DB: Columns of table %s: %s\n", table, cols)
	m := make(map[string]int)
	for i, c := range cols {
		m[c] = i
	}
	vals := make([]interface{}, len(cols))
	for i, _ := range vals {
		var val interface{}
		vals[i] = &val
	}
	for rows.Next() {
		if err := rows.Scan(vals...); err != nil {
			log.Fatal(err)
		}
		s := "DB:"
		for i, colName := range cols {
			var v = *(vals[i].(*interface{}))
			var v_type = reflect.TypeOf(v)
			s = fmt.Sprintf("%s %s=%s(%s),", s, colName, v, v_type)
		}
		s = strings.TrimRight(s, ",")
		fmt.Println(s)
	}
}

