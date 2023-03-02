type query struct {
	sql    chan string
	result chan string
}

func execQuery(q query) {
	go func() {
		sql := <-q.sql
		q.result <- "result from" + sql
	}()

}

func main() {

	//c := make(chan struct{})
	//go func(i chan struct{}) {
	//	sum := 0
	//	for i := 0; i < 10000; i++ {
	//		sum += i
	//	}
	//	println(sum)
	//	c <- struct{}{}
	//}(c)
	//
	//println("NumGoroutine=", runtime.NumGoroutine())
	//<-c
	//for _, url := range urls {
	//	wg.Add(1)
	//	go func(url string) {
	//		defer wg.Done()
	//		resp, err := http.Get(url)
	//		if err == nil {
	//			println(resp.Status)
	//		}
	//	}(url)
	//}
	//wg.Wait()
	q := query{make(chan string, 1), make(chan string, 1)}
	go execQuery(q)
	q.sql <- "select * from table"
	time.Sleep(1 * time.Second)
	fmt.Println(<-q.result)
}
