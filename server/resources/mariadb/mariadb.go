package mariadb

///////////////
//  Imports  //
///////////////

import (
	// Golang
	"database/sql"
	"fmt"
	_ "reflect"

	// External
	_ "github.com/go-sql-driver/mysql"
)

///////////////
//  Structs  //
///////////////

type Connection struct {
	*sql.DB
}

/////////////////
//  Functions  //
/////////////////

func Config(user string, pass string, name string) (*Connection,error) {
	newConn, err := sql.Open("mysql",fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s",user,pass,name))
	if err != nil { return nil,err }
	db := Connection{newConn}
	return &db,nil
}

func ConfigExtended(user string, pass string, address string, port string, name string) (*Connection,error) {
	newConn, err := sql.Open("mysql",fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",user,pass,address,port,name))
	if err != nil { return nil,err }
	db := Connection{newConn}
	return &db,nil
}

func (conn *Connection) Select(query string,args ...interface{}) ([]map[string]interface{},error){
	// Set output
	var output []map[string]interface{}
	// Get rows from query
	rows, err := conn.Query(query,args...)
	if err != nil { return nil,err }
	defer rows.Close()
	// Get types
	columnTypes, err := rows.ColumnTypes()
	if err != nil { return nil,err }

	// Get Null reflection
	scanArgs := make([]interface{}, len(columnTypes))
	for i, v := range columnTypes {
		switch v.DatabaseTypeName() {
			case "VARCHAR", "TEXT", "UUID","CHAR","MEDIUMTEXT","LONGTEXT":
				var a sql.NullString
				scanArgs[i] = &a
			case "TIMESTAMP","DATE","TIME","DATETIME":
				var a sql.NullString
				scanArgs[i] = &a
			case "BINARY", "BLOB", "MEDIUMBLOB", "LONGBLOB","TINYBLOB","VARBINARY":
				var a sql.RawBytes
				scanArgs[i] = &a
			case "TINYINT","BOOL","BOOLEAN":
				var a sql.NullBool
				scanArgs[i] = &a
			case "SMALLINT","MEDIUMINT","INT","INTEGER","BIGINT","INT1","INT2","INT3","INT4","INT8":
				var a sql.NullInt64
				scanArgs[i] = &a
			case "DECIMAL","NUMERIC","NUMBER","FIXED","FLOAT","DOUBLE","DOUBLE PRECISION":
				var a sql.NullFloat64
				scanArgs[i] = &a
			default:
				var a sql.NullString
				scanArgs[i] = &a
		}
	}

	for rows.Next() {

		err := rows.Scan(scanArgs...)
		if err != nil { return nil,err }

		row := map[string]interface{}{}

		for i, v := range columnTypes {
			if z, ok := (scanArgs[i]).(*sql.NullBool); ok  {
				row[v.Name()] = z.Bool
				continue;
			}
			if z, ok := (scanArgs[i]).(*sql.NullString); ok  {
				row[v.Name()] = z.String
				continue;
			}
			if z, ok := (scanArgs[i]).(*sql.RawBytes); ok  {
				b := make([]byte, len(*z))
				copy(b, *z)
				row[v.Name()] = b
				continue;
			}
			if z, ok := (scanArgs[i]).(*sql.NullInt64); ok  {
				row[v.Name()] = z.Int64
				continue;
			}
			if z, ok := (scanArgs[i]).(*sql.NullFloat64); ok  {
				row[v.Name()] = z.Float64
				continue;
			}
			row[v.Name()] = scanArgs[i]
		}

		output = append(output, row)
	}

	return output,nil
}

func (conn *Connection) Insert(query string, args ...interface{}) (int64,error){
	res, err := conn.Exec(query, args...)
	if err != nil { return 0,err }
	id, err := res.LastInsertId()
	if err != nil { return 0,err}
	return id,nil
}

func (conn *Connection) numRows(query string, args ...interface{}) (int64,error){
	res, err := conn.Exec(query, args...)
	if err != nil { return 0,err } 
	numRows, err := res.RowsAffected()
	if err != nil { return 0,err } 
	return numRows,nil
}

func (conn *Connection) Update(query string, args ...interface{}) (int64,error){
	return conn.numRows(query,args...)
}

func (conn *Connection) Delete(query string, args ...interface{}) (int64,error){
	return conn.numRows(query,args...)
}

func (conn *Connection) Create(query string, args ...interface{}) (int64,error){
	return conn.numRows(query,args...)
}

func (conn *Connection) Close() error {
	err := conn.Close()
	return err
}
