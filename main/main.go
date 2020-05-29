package main
import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)
var cache = map[int]Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
func main(){
	wg := &sync.WaitGroup{}
	m := &sync.RWMutex{}
	cacheCh := make(chan Book)
	dbCh := make(chan Book)
	for i:=0;i<10 ;i++  {
		id := rnd.Intn(10) + 1
		/*if b, ok := queryCache(id); ok {
			fmt.Println("from cache")
			fmt.Println(b)
			continue
		}
		if b, ok := queryDatabase(id); ok {
			fmt.Println("from database")
			fmt.Println(b)
			continue
			}
		//fmt.Printf("Book not found with id: '%v'", id)
		//time.Sleep(150 * time.Millisecond)
		*/
		wg.Add(2)
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex, ch chan<- Book) {
			if b, ok := queryCache(id,m); ok {
				/*fmt.Println("from cache")
				fmt.Println(b)*/
				ch <- b
			}
			wg.Done()
		}(id,wg,m,cacheCh)

		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex, ch chan<- Book) {
			if b, ok := queryDatabase(id,m); ok {
				ch <- b
				/*fmt.Println("from database")
				fmt.Println(b)*/
			}
			wg.Done()
		}(id,wg,m,dbCh)
		// create one goroutine per query to handle response
		go func(cacheCh, dbCh <-chan Book) {
			select {
			case b := <- cacheCh:
				fmt.Println("from cache")
				fmt.Println(b)
				<-dbCh
			case b := <- dbCh:
				fmt.Println("from database")
				fmt.Println(b)
			}
		}(cacheCh,dbCh)
		time.Sleep(150 * time.Millisecond)
	}
	wg.Wait()
}




func queryCache(id int, m *sync.RWMutex) (Book, bool) {
	m.RLock()
	b, ok := cache[id]
	m.RUnlock()
	return b, ok
}

func queryDatabase(id int, m *sync.RWMutex) (Book, bool){
	time.Sleep(100 * time.Millisecond)
	for _, b := range books {
		m.Lock()
		cache[id] = b
		m.Unlock()
		if b.ID == id {
			return b, true
		}
	}
	return Book{}, false
}
