package main

import (
	"fmt"
	"sync"
	"time"
)

// Go -> GoRoutine
// GoRoutine -> es un hilo de ejecución ligero virtual
// Channel -> para comunicar cosas entre GoRoutine
func decirHola(canal chan<- string) {
	time.Sleep(1 * time.Second)
	canal <- "Hola desde la GoRoutine"
}

func imprimirMensaje(canal <-chan string) {
	fmt.Println("Esperando mensaje...")
	msg := <-canal
	fmt.Println("Mensaje recibido:", msg)
}

func main() {
	canal := make(chan string)
	go decirHola(canal)
	imprimirMensaje(canal)

	canal2 := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			canal2 <- i
		}
		close(canal2)
	}()

	for num := range canal2 {
		fmt.Println("Número recibido:", num)
	}

	// Mutex
	var contador int
	var mu sync.RWMutex

	// writer
	go func() {
		for i := 0; i < 10; i++ {
			mu.Lock()
			contador++
			mu.Unlock()
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// Reader
	for i := 0; i < 10; i++ {
		go func(id int) {
			for j := 0; j < 5; j++ {
				mu.RLock()
				fmt.Println("Contador:", contador)
				mu.RUnlock()
				time.Sleep(1 * time.Second)
			}
		}(i)
	}

	time.Sleep(10 * time.Second)
}
