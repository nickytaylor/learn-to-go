package main

import "fmt"

//const (
//	c0 = iota
//	c1 = iota
//	c2 = iota
//)
//
//var a = "hello,world"
//var b = a[0]
//var c = "hello"
//
//type User struct {
//	name string
//	age  int
//}
//
//var andes = User{
//	name: "andes",
//	age:  18,
//}
//var p = &andes
//var e = [...]int{1, 2, 3}
//
//type T struct {
//	a int
//}
//
//func (t T) Get() int {
//	return t.a
//}
//func doinout(f func(int, int) int, a, b int) int {
//	return f(a, b)
//}
//
////func main() {
////	//fmt.Println(p.name + " p.age")
////	//e[0] = 2
////	//fmt.Println(e)
////	////for i := 0; i < len(d); i++ {
////	////	fmt.Println(d[i])
////	////
////	////}
////	////for i, v := range d {
////	////	fmt.Println(i, v)
////	////}
////	//doinout(func(x, y int) int {
////	//	return x + y
////	//}, 1, 2)
////	t := T{
////		a: 1,
////	}
////	t.Get()
////
////}
//
//type Int int
//
//func (a Int) Max(b Int) Int {
//	if a >= b {
//		return a
//	} else {
//		return b
//	}
//}
//func (i *Int) Set(a Int) {
//	*i = a
//}
//func (i Int) Print() {
//	fmt.Printf("value=%d\n", i)
//}
//
//var wg sync.WaitGroup
//var urls = []string{
//	"http://www.baidu.com",
//	"http://www.baidu.com",
//	"http://www.baidu.com",
//}
//
//func GenerateIntA() chan int {
//	ch := make(chan int, 10)
//	go func() {
//		for {
//			ch <- rand.Int()
//		}
//	}()
//	return ch
//}
//func GenerateIntB() chan int {
//	ch := make(chan int, 10)
//	go func() {
//		for {
//			ch <- rand.Int()
//		}
//	}()
//	return ch
//}
//func GenerateInt() chan int {
//	ch := make(chan int, 20)
//	go func() {
//		for {
//			select {
//			case ch <- <-GenerateIntA():
//			case ch <- <-GenerateIntB():
//			}
//		}
//	}()
//	return ch
//}

const (
	NUMBER = 10
)

type task struct {
	begin  int
	end    int
	result chan<- int
}

func (t *task) do() {
	sum := 0
	for i := t.begin; i <= t.end; i++ {
		sum += i

	}
	t.result <- sum
}
func InitTask(taskchan chan<- task, r chan int, p int) {
	qu := p / 10
	mod := p % 10
	high := qu * 10
	for j := 0; j < qu; j++ {
		b := 10*j + 1
		e := 10 * (j + 1)
		tsk := task{
			begin:  b,
			end:    e,
			result: r,
		}
		taskchan <- tsk
	}
	if mod != 0 {
		tsk := task{
			begin:  high + 1,
			end:    p,
			result: r,
		}
		taskchan <- tsk
	}
	close(taskchan)
}
func DistributeTask(taskchan <-chan task, workers int, done chan struct{}) {
	for i := 0; i < workers; i++ {
		go ProcessTask(taskchan, done)
	}
}
func ProcessTask(taskchan <-chan task, done chan struct{}) {
	for t := range taskchan {
		t.do()
	}
	done <- struct{}{}
}
func CloseResult(done chan struct{}, resultchan chan int, workers int) {
	for i := 0; i < workers; i++ {
		<-done
	}
	close(done)
	close(resultchan)
}
func ProcessResult(resultchan chan int) int {
	sum := 0
	for r := range resultchan {
		sum += r
	}
	return sum
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
	workers := NUMBER
	taskchan := make(chan task, 10)
	resultchan := make(chan int, 10)
	done := make(chan struct{}, 10)
	go InitTask(taskchan, resultchan, 100)
	DistributeTask(taskchan, workers, done)
	go CloseResult(done, resultchan, workers)
	sum := ProcessResult(resultchan)
	fmt.Println(sum)
}
