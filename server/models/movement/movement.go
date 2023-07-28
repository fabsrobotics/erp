package movement

import (
	// Internal
	"erp/models/submovement"
	"erp/models/taxyear"
	"erp/resources/mariadb"
	"fmt"
)

func CreateOrUpdateTable(db *mariadb.Connection) error {
	// Create table taxyear
	_,err := db.Create("CREATE TABLE IF NOT EXISTS `accounting_movement` (`id` INT NOT NULL AUTO_INCREMENT, `date` DATE NOT NULL, `description` VARCHAR(255) NOT NULL, `tax_year` INT NOT NULL, PRIMARY KEY(`id`) )")
	return err
}

func New(db *mariadb.Connection, date string, description string, year int64) (int64,error) {
	// Insert new movement
	id,err := db.Insert("INSERT INTO `accounting_movement` (`date`,`description`,`tax_year`) VALUES (?,?,?)",date,description,year)
	return id,err
}

func UpdateDate(db *mariadb.Connection, id int64, newDate string) error {
	// Update date movement
	_,err := db.Update("UPDATE `accounting_movement` SET `date` = ? WHERE `id` = ?",newDate, id)
	return err
}

func UpdateDescription(db *mariadb.Connection, id int64, newDescription string) error {
	// Update description movement
	_,err := db.Update("UPDATE `accounting_movement` SET `description` = ? WHERE `id` = ?",newDescription, id)
	return err
}

func UpdateYear(db *mariadb.Connection, id int64, newYear int64) error {
	// Check New year exist
	exists,err := taxyear.CheckYearExistence(db,newYear)
	if err != nil { return err }
	if !exists { return fmt.Errorf("Year doesn't exists") }
	// Update year
	_,err = db.Update("UPDATE `accounting_movement` SET `year`= ? WHERE `id`= ?",newYear,id)
	return err
}

func Delete(db *mariadb.Connection, id int64) error {
	// Delete submovements associated
	err := submovement.DeleteAllByMovementId(db,id)
	if err != nil { return err }
	// Delete movement
	_,err = db.Delete("DELETE FROM `accounting_movement` WHERE `id` = ?",id)
	return err
}

func ChangeYearOfAll(db *mariadb.Connection, oldYear int64, newYear int64) error {
	// Update all movements of oldyear
	_,err := db.Update("UPDATE `accounting_movement` SET `year` = ? WHERE `year` = ?",newYear,oldYear)
	return err
}

func DeleteAllByYear(db *mariadb.Connection, year int64) error {
	// Select all movements
	data,err := db.Select("SELECT `id` FROM `accounting_movement` WHERE `year` = ?",year)
	if err != nil { return err }
	// Delete all submovements
	for _,movement := range(data){
		err = submovement.DeleteAllByMovementId(db,movement["id"].(int64))
		if err != nil { return err}
	}
	// Delete all movements of year
	_,err = db.Delete("DELETE FROM `accounting_movement` WHERE `year` = ?",year)
	return err
}

func DeleteAllByAccountId(db *mariadb.Connection, accountId int64) error {
	// Select all movements
	data,err := db.Select("SELECT `id` FROM `accounting_movement` WHERE `account_id` = ?",accountId)
	if err != nil { return err }
	// Delete all submovements
	for _,movement := range(data){
		err = submovement.DeleteAllByMovementId(db,movement["id"].(int64))
		if err != nil { return err}
	}
	// Delete all movements of year
	_,err = db.Delete("DELETE FROM `accounting_movement` WHERE `account_id` = ?",accountId)
	return err
}