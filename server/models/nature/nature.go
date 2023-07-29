package nature

import (
	"erp/models/account"
	"erp/resources/mariadb"
)

func CreateOrUpdateTable(db *mariadb.Connection) error {
	// Create table nature
	_, err := db.Create("CREATE TABLE IF NOT EXISTS `accounting_nature` (`id` INT NOT NULL AUTO_INCREMENT, `name` VARCHAR(45) NOT NULL, `increment` TINYINT NOT NULL, PRIMARY KEY(`id`) )")
	return err
}

func New(db *mariadb.Connection, name string, increment bool) error {
	// Insert new nature
	_, err := db.Insert("INSERT INTO `accounting_nature` (`name`,`increment`) VALUES (?,?)", name, increment)
	return err
}

func UpdateName(db *mariadb.Connection, id int64, newName string) error {
	// Update nature name
	_, err := db.Update("UPDATE `accounting_nature` SET `name` = ? WHERE `id` = ?", newName, id)
	return err
}

func UpdateIncrement(db *mariadb.Connection, id int64, newIncrement bool) error {
	// Update nature increment
	_, err := db.Update("UPDATE `accounting_nature` SET `increment` = ? WHERE `id` = ?", newIncrement, id)
	return err
}

func Delete(db *mariadb.Connection, id int64) error {
	// Delete all accounts with this nature
	err := account.DeleteAllByNatureId(db, id)
	if err != nil {
		return err
	}
	// Delete nature
	_, err = db.Delete("DELETE FROM `accounting_nature` WHERE `id` = ?", id)
	return err
}
