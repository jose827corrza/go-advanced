package main

import (
	"fmt"
	"sync"
)

var (
	balance = 100
)

func Deposit(ammount int, wg *sync.WaitGroup, lock *sync.RWMutex) {
	defer wg.Done()
	lock.Lock() //Entre todas las GoRoutines, la primera que llegue aca va a impedir que se mod lode abajo
	//pues, hasta que se desbloquee
	b := balance
	balance = b + ammount
	lock.Unlock()
}

func Balance(lock *sync.RWMutex) int {
	lock.RLock() //La bloquea para lectura
	b := balance
	lock.RUnlock()
	return b
}
func main() {
	//Ahora se va a bloquear el programa, para saber cuando finaliza
	var wg sync.WaitGroup
	//var lock sync.Mutex //Esta linea hace la magia
	var lock sync.RWMutex //Este permite especificar si se bloquea para escritura o lectura, aparte
	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go Deposit(i*100, &wg, &lock)
	}
	wg.Wait() //<-aca se bloqueo el programa
	fmt.Println(Balance(&lock))
}

//go build --race sync/main.go
//nos va a indicar si hay condicion de carrea
