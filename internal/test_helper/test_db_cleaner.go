package test_helper

import (
	"database/sql"
	"fmt"

	"github.com/jinzhu/gorm"
)

// truncate all tables

//SET FOREIGN_KEY_CHECKS=0;

// SELECT Concat('TRUNCATE TABLE ',table_schema,'.',TABLE_NAME, ';')
//FROM INFORMATION_SCHEMA.TABLES where  table_schema in (database_name)

//SET FOREIGN_KEY_CHECKS=1;

// DeleteCreatedEntities sets up GORM `onCreate` hook and return a function that can be deferred to
// remove all the entities created after the hook was set up
// You can use it as
//
// func TestSomething(t *testing.T){
//     db, _ := gorm.Open(...)
//
//     cleaner := DeleteCreatedEntities(db)
//     defer cleaner()
//
// }
func DeleteCreatedEntities(db *gorm.DB) func() {
	type entity struct {
		table   string
		keyname string
		key     interface{}
	}
	var entries []entity
	hookName := "cleanupHook"

	db.Callback().Create().After("gorm:create").Register(hookName, func(scope *gorm.Scope) {
		fmt.Printf("Inserted entities of %s with %s=%v\n", scope.TableName(), scope.PrimaryKey(), scope.PrimaryKeyValue())
		entries = append(entries, entity{table: scope.TableName(), keyname: scope.PrimaryKey(), key: scope.PrimaryKeyValue()})
	})
	return func() {
		// Remove the hook once we're done
		defer db.Callback().Create().Remove(hookName)
		// Find out if the current db object is already a transaction
		_, inTransaction := db.CommonDB().(*sql.Tx)
		tx := db
		if !inTransaction {
			tx = db.Begin()
		}
		// Loop from the end. It is important that we delete the entries in the
		// reverse order of their insertion
		for i := len(entries) - 1; i >= 0; i-- {
			entry := entries[i]
			fmt.Printf("Deleting entities from '%s' table with key %v\n", entry.table, entry.key)
			tx.Table(entry.table).Where(entry.keyname+" = ?", entry.key).Delete("")
		}

		if !inTransaction {
			tx.Commit()
		}
	}
}
