package submovement

import "erp/resources/mariadb"

func CreateOrUpdateTable(db *mariadb.Connection) error {
	// Create table taxyear
	_,err := db.Create("CREATE TABLE IF NOT EXISTS `accounting_submovement` (`id` INT NOT NULL AUTO_INCREMENT, `debit_credit` TINYINT NOT NULL, `value` DECIMAL(10,2) NOT NULL, `movement_id` INT NOT NULL, `account_id` INT NOT NULL, PRIMARY KEY(`id`) )")
	return err
}


func New(db *mariadb.Connection,debitCredit bool, value float64, movementId int64, accountId int64) error {
	// Insert new submovement
	_,err := db.Insert("INSERT INTO `accounting_submovement` (`debit_credit`,`value`,`movement_id`,`account_id`) VALUES (?,?,?,?)",debitCredit,value,movementId,accountId)
	return err
}

func UpdateDebitCredit(db *mariadb.Connection,id int64, debitCredit bool) error {
	// Update Debit Credit 
	_,err := db.Update("UPDATE `accounting_submovement` SET `debit_credit` = ? WHERE `id` = ?",debitCredit,id)
	return err
}

func UpdateValue(db *mariadb.Connection,id int64, value float64) error {
	// Update Value 
	_,err := db.Update("UPDATE `accounting_submovement` SET `value` = ? WHERE `id` = ?",value,id)
	return err
}

func UpdateMovementId(db *mariadb.Connection,id int64, movementId int64) error {
	// Update Movement Id 
	_,err := db.Update("UPDATE `accounting_submovement` SET `movement_id` = ? WHERE `id` = ?",movementId,id)
	return err
}

func UpdateAccountId(db *mariadb.Connection,id int64, accountId int64) error {
	// Update Movement Id 
	_,err := db.Update("UPDATE `accounting_submovement` SET `account_id` = ? WHERE `id` = ?",accountId,id)
	return err
}

func Delete(db *mariadb.Connection, id int64) error {
	// Delete submovement
	_,err := db.Delete("DELETE FROM `accounting_submovement` WHERE `id` = ?",id)
	return err
}

func DeleteAllByMovementId(db *mariadb.Connection, movementId int64) error {
	// Delete all submovement
	_,err := db.Delete("DELETE FROM `accounting_submovement` WHERE `movement_id` = ?",movementId)
	return err
}