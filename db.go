package main

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	//_ "github.com/mxk/go-sqlite"
	//_ "github.com/go-sql-driver/mysql"
	"log"
)

func dbOpen(dbFileName string) (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", dbFileName)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("Opening db file: ", dbFileName)

	db.DB()
	db.DB().Ping()
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	sqlite3Config(db.DB())
	return db, nil
}

func sqlite3Config(db *sql.DB) {
	//db.Exec("PRAGMA auto_vacuum = 0;")
	//db.Exec("PRAGMA cache_size=32768;")
	db.Exec("PRAGMA cache_size=65536;")
	db.Exec("PRAGMA count_changes = OFF;")
	db.Exec("PRAGMA cache_spill = ON;")
	//db.Exec("PRAGMA journal_size_limit = 67110000;")
	db.Exec("PRAGMA locking_mode = EXCLUSIVE;")
	//db.Exec("PRAGMA locking_mode = OFF;")
	db.Exec("PRAGMA encoding = \"UTF-8\";")
	//db.Exec("PRAGMA journal_mode = WAL;")

	db.Exec("busy_timeout=0;")
	db.Exec("legacy_file_format=OFF;")

	//db.Exec("PRAGMA mmap_size=1099511627776;")
	db.Exec("PRAGMA page_size = 40960;")

	db.Exec("PRAGMA shrink_memory;")
	db.Exec("PRAGMA synchronous=OFF;")
	//db.Exec("PRAGMA synchronous = NORMAL;")
	//db.Exec("PRAGMA temp_store = MEMORY;")
	//db.Exec("PRAGMA threads = 5;")
	//db.Exec("PRAGMA wal_autocheckpoint = 1638400;")
}

func dbInit(dbFile string) (*gorm.DB, error) {
	db, err := dbOpen(dbFile)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	//	log.Printf("%v\n", *db)

	db.CreateTable(&MeshTree{})
	return db, nil
}

const table = "mesh_trees"

func makeIndexes(db *gorm.DB) {
	db.Table(table).AddIndex("descriptor_ui", "descriptor_ui")
	db.Table(table).AddIndex("descriptor_name", "descriptor_name")

	db.Table(table).AddIndex("tree", "tree")
	db.Table(table).AddIndex("t0", "t0")
	db.Table(table).AddIndex("t0_t1", "t0", "t1")
	db.Table(table).AddIndex("t0_t1_t2", "t0", "t1", "t2")
	db.Table(table).AddIndex("t0_t1_t2_t3", "t0", "t1", "t2", "t3")
	db.Table(table).AddIndex("t0_t1_t2_t3_t4", "t0", "t1", "t2", "t3", "t4")
	db.Table(table).AddIndex("t0_t1_t2_t3_t4_t5", "t0", "t1", "t2", "t3", "t4", "t5")
	db.Table(table).AddIndex("t0_t1_t2_t3_t4_t5_t6", "t0", "t1", "t2", "t3", "t4", "t5", "t6")
	db.Table(table).AddIndex("t0_t1_t2_t3_t4_t5_t6_t7", "t0", "t1", "t2", "t3", "t4", "t5", "t6", "t7")
	db.Table(table).AddIndex("t0_t1_t2_t3_t4_t5_t6_t7_t8", "t0", "t1", "t2", "t3", "t4", "t5", "t6", "t7", "t8")
	db.Table(table).AddIndex("t0_t1_t2_t3_t4_t5_t6_t7_t8_t9", "t0", "t1", "t2", "t3", "t4", "t5", "t6", "t7", "t8", "t9")
	db.Table(table).AddIndex("t0_t1_t2_t3_t4_t5_t6_t7_t8_t9_t10", "t0", "t1", "t2", "t3", "t4", "t5", "t6", "t7", "t8", "t9", "t10")
	db.Table(table).AddIndex("t0_t1_t2_t3_t4_t5_t6_t7_t8_t9_t10_t11", "t0", "t1", "t2", "t3", "t4", "t5", "t6", "t7", "t8", "t9", "t10", "t11")

}
