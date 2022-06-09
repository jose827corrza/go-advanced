package main

import (
	"fmt"
	"sync"
	"time"
)

type Database struct{}

func (Database) CreateSingleConnection() {
	fmt.Println("Creating singleton for Database")
	time.Sleep(2 * time.Second)
	fmt.Println("Connection done")
}

var db *Database
var lock sync.Mutex

func GetDatabaseInstance() *Database {
	lock.Lock()
	defer lock.Unlock()
	if db == nil {
		fmt.Println("CreatingDB connection")
		db = &Database{}
		db.CreateSingleConnection()
	} else {
		fmt.Println("Connection already created")
	}
	return db
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			GetDatabaseInstance()
		}()
	}
	wg.Wait()
}
