package account

import (
	"erp/models/movement"
	"erp/resources/mariadb"
	"fmt"
	"time"
)

func CreateOrUpdateTable(db *mariadb.Connection) error {
	// Create table account
	_, err := db.Create("CREATE TABLE IF NOT EXISTS `accounting_account` (`id` INT NOT NULL AUTO_INCREMENT, `account` INT NOT NULL, `parent_account` INT NOT NULL, `description` VARCHAR(255) NOT NULL, `init_date` DATE NOT NULL, `end_date` DATE NULL, `nature_id` INT NOT NULL, PRIMARY KEY(`id`) )")
	return err
}

func New(db *mariadb.Connection, account int64, parent_account int64, description string, initDate string, endDate string, natureId int64) error {
	// Check if account exists in the same period
	accounts, err := db.Select("SELECT * FROM `accounting_account` WHERE `account` = ?", account)
	if err != nil {
		return err
	}
	for _, account := range accounts {
		actualInitDate, err := time.Parse("2006-01-02", initDate)
		if err != nil {
			return err
		}
		destinationInitDate, err := time.Parse("2006-01-02", account["init_date"].(string))
		if err != nil {
			return err
		}
		// Check intervals
		if actualInitDate.After(destinationInitDate) {
			if account["end_date"] == "" {
				return fmt.Errorf("actual init date is over unfinish Destination")
			} else {
				destinationEndDate, err := time.Parse("2006-01-02", account["end_date"].(string))
				if err != nil {
					return err
				}
				if actualInitDate.Before(destinationEndDate) {
					return fmt.Errorf("actual init date is before destination end date")
				}
			}
		} else if actualInitDate.Equal(destinationInitDate) {
			return fmt.Errorf("dates have same init date")
		} else {
			if endDate == "" {
				return fmt.Errorf("destination init over unfinish actual")
			} else {
				actualEndDate, err := time.Parse("2006-01-02", endDate)
				if err != nil {
					return err
				}
				if !actualEndDate.Before(destinationInitDate) {
					return fmt.Errorf("actual ends after or equal to destination init")
				}
			}
		}
	}
	// Insert new submovement
	_, err = db.Insert("INSERT INTO `accounting_account` (`account`,`parent_account`,`description`,`init_date`,`end_date`,`nature_id`) VALUES (?,?,?,?,?,?)", account, parent_account, description, initDate, endDate, natureId)
	return err
}

func UpdateAccount(db *mariadb.Connection, id int64, newAccount int64) error {
	// Check account doesn't exists from the same period
	// Get account data
	data, err := db.Select("SELECT * FROM `accounting_account` WHERE `id` = ?", id)
	if err != nil {
		return err
	}
	// Get all newAccount accounts
	accounts, err := db.Select("SELECT * FROM `accounting_account` WHERE `account` = ?", newAccount)
	if err != nil {
		return err
	}
	for _, account := range accounts {
		actualInitDate, err := time.Parse("2006-01-02", data[0]["init_date"].(string))
		if err != nil {
			return err
		}
		destinationInitDate, err := time.Parse("2006-01-02", account["init_date"].(string))
		if err != nil {
			return err
		}
		// Check intervals
		if actualInitDate.After(destinationInitDate) {
			if account["end_date"] == "" {
				return fmt.Errorf("actual init date is over unfinish Destination")
			} else {
				destinationEndDate, err := time.Parse("2006-01-02", account["end_date"].(string))
				if err != nil {
					return err
				}
				if actualInitDate.Before(destinationEndDate) {
					return fmt.Errorf("actual init date is before destination end date")
				}
			}
		} else if actualInitDate.Equal(destinationInitDate) {
			return fmt.Errorf("dates have same init date")
		} else {
			if data[0]["end_date"].(string) == "" {
				return fmt.Errorf("destination init over unfinish actual")
			} else {
				actualEndDate, err := time.Parse("2006-01-02", data[0]["end_date"].(string))
				if err != nil {
					return err
				}
				if !actualEndDate.Before(destinationInitDate) {
					return fmt.Errorf("actual ends after or equal to destination init")
				}
			}
		}
	}
	// Update account
	_, err = db.Update("UPDATE `accounting_account` SET `account` = ? WHERE `id` = ?", newAccount, id)
	return err
}

func UpdateParentAccount(db *mariadb.Connection, id int64, newParentAccount int64) error {
	// Update parent account
	_, err := db.Update("UPDATE `accounting_account` SET `parent_account` = ? WHERE `id` = ?", newParentAccount, id)
	return err
}

func UpdateDescription(db *mariadb.Connection, id int64, newDescription string) error {
	// Update description of account
	_, err := db.Update("UPDATE `accounting_account` SET `description` = ? WHERE `id` = ?", newDescription, id)
	return err
}

func UpdateInitDate(db *mariadb.Connection, id int64, newInitDate string) error {
	// Update init date of account
	_, err := db.Update("UPDATE `accounting_account` SET `init_date` = ? WHERE `id` = ?", newInitDate, id)
	return err
}

func UpdateEndDate(db *mariadb.Connection, id int64, newEndDate string) error {
	// Update end date of account
	_, err := db.Update("UPDATE `accounting_account` SET `end_date` = ? WHERE `id` = ?", newEndDate, id)
	return err
}

func UpdateNatureId(db *mariadb.Connection, id int64, newNatureId int64) error {
	// Update nature id of account
	_, err := db.Update("UPDATE `accounting_account` SET `nature_id` = ? WHERE `id` = ?", newNatureId, id)
	return err
}

func Delete(db *mariadb.Connection, id int64) error {
	// Delete all movements with this account
	err := movement.DeleteAllByAccountId(db, id)
	if err != nil {
		return err
	}
	// Delete account
	_, err = db.Delete("DELETE FROM `accounting_account` WHERE `id` = ?", id)
	return err
}

func DeleteAllByNatureId(db *mariadb.Connection, natureId int64) error {
	// Select all by nature
	data, err := db.Select("SELECT * FROM `accounting_account` WHERE `nature_id` = ?", natureId)
	if err != nil {
		return err
	}
	for _, account := range data {
		// Delete all movements with this account
		err := movement.DeleteAllByAccountId(db, account["id"].(int64))
		if err != nil {
			return err
		}
	}
	// Delete accounts
	_, err = db.Delete("DELETE FROM `accounting_account` WHERE `nature_id` = ?", natureId)
	return err
}
