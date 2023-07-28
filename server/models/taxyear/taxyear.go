package taxyear

import (
	"erp/models/movement"
	"erp/resources/mariadb"
	"fmt"
	"regexp"
)

func CreateOrUpdateTable(db *mariadb.Connection){
	// Create table taxyear
	_,err := db.Create("CREATE TABLE IF NOT EXISTS `accounting_tax_year` (`year` INT NOT NULL, `init_date` DATE NOT NULL, `end_date` DATE NOT NULL, PRIMARY KEY(`year`) )")
	if err != nil { return }
}

func New(db *mariadb.Connection, year int64, startDate string, endDate string) error {
	re := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
	// Check startDate follows format
	if !re.MatchString(startDate){ return fmt.Errorf("startDate doesn't follow format YYYY-MM-DD")}
	// Check endDate follows format
	if !re.MatchString(startDate){ return fmt.Errorf("endDate doesn't follow format YYYY-MM-DD")}
	// Check Year doesn't already exists
	raw,err := db.Select("SELECT * FROM `taxyear` WHERE `year` = ?",year)
	if err != nil { return err }
	if len(raw) > 2 { return fmt.Errorf("year already exists") }
	return nil
}

func UpdateYear(db *mariadb.Connection, oldYear int64, newYear int64) error {
	// Check new year doesn't already exists
	raw,err := db.Select("SELECT * FROM `accounting_tax_year` WHERE `year` = ?",newYear)
	if err != nil { return err }
	if len(raw) > 2 { return fmt.Errorf("year already exists") }
	// Update year
	_,err = db.Update("UPDATE `accounting_tax_year` SET `year` = ? WHERE `year` = ?",newYear,oldYear)
	if err != nil { return err }
	// Update all account movements year
	err = movement.ChangeYearOfAll(db,oldYear,newYear)
	return err
}

func Delete(db *mariadb.Connection, year int64) error {
	// Delete all movements from this year
	err := movement.DeleteAllByYear(db,year)
	if err != nil { return err }
	// Delete year
	_,err = db.Delete("DELETE FROM `accounting_tax_year` WHERE `year` = ?",year)
	return err
}

func CheckYearExistence(db *mariadb.Connection, year int64) (bool,error) {
	// Check if year already exists
	raw,err := db.Select("SELECT * FROM `accounting_tax_year` WHERE `year` = ?",year)
	if err != nil { return false,err }
	if len(raw) > 2 {
		return true,nil
	} else {
		return false,nil
	}
}