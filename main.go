package main

import (
	"fmt"
	"goconcurrent/book"
	"math/rand"
	"sync"
	"time"
)

var cache = map[int]book.Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	wg := &sync.WaitGroup{}
	m := &sync.RWMutex{}
	cacheCh := make(chan book.Book)
	dbCh := make(chan book.Book)
	for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1
		wg.Add(2)
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex, ch chan<- book.Book) {
			if b, ok := queryCache(id, m); ok {
				// fmt.Println("From cache")
				// fmt.Println(b)
				ch <- b
			}
			wg.Done()
		}(id, wg, m, cacheCh)
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex, ch chan<- book.Book) {
			if b, ok := queryDatabase(id, m); ok {
				// fmt.Println(b)
				m.Lock()
				cache[id] = b
				m.Unlock()
				ch <- b
			}
			wg.Done()
		}(id, wg, m, dbCh)

		go func(cacheCh, dbCh <-chan book.Book) {
			select {
			case b := <-cacheCh:
				fmt.Println("From cache")
				fmt.Println(b)
				<-dbCh
			case b := <-dbCh:
				fmt.Println("From database")
				fmt.Println(b)
			}
		}(cacheCh, dbCh)
		time.Sleep(150 * time.Millisecond)
	}
	wg.Wait()
	// wg := &sync.WaitGroup{}
	// ch := make(chan int, 3)

	// wg.Add(2)

	// go func(ch chan int, wg *sync.WaitGroup) {
	// 	// if msg, ok := <-ch; ok {
	// 	// 	fmt.Println(msg, ok)
	// 	// }
	// 	for msg := range ch {
	// 		fmt.Println(msg)
	// 	}
	// 	wg.Done()
	// }(ch, wg)

	// go func(ch chan int, wg *sync.WaitGroup) {
	// 	for i := 0; i < 10; i++ {
	// 		ch <- i
	// 	}
	// 	close(ch)
	// 	wg.Done()
	// }(ch, wg)

	// wg.Wait()
}

func queryCache(id int, m *sync.RWMutex) (book.Book, bool) {
	m.Lock()
	b, ok := cache[id]
	m.Unlock()
	return b, ok
}

func queryDatabase(id int, m *sync.RWMutex) (book.Book, bool) {
	time.Sleep(100 * time.Millisecond)
	for _, b := range book.GetBooks() {
		if b.ID == id {
			// m.Lock()
			// cache[id] = b
			// m.Unlock()
			return b, true
		}
	}
	return book.Book{}, false
}
