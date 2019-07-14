package main
// 1. Параллельный работает логично быстрее, "опрашивая" сайты все сразу.
// 2. Еще по работе бывает нужно мониторить страницу сайта на появление новых документов. Вот как вариант, чтобы сканировалась страница, переодически, и сравнивала с предыдущим опросом на изменения. Или что то вроде этого.
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	//"args": ["https://ya.ru","https://mail.ru","https://gmail.com","https://dtf.ru","https://youtube.com","https://habrahabr.ru","https://vk.com","https://twitch.com","https://google.com","https://twitter.com","https://geekbrains.ru/","https://github.com/alexeybobkov47"]
	for range os.Args[1:] {
		fmt.Print(<-ch) // receive from channel
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body) // to devNull
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	ch <- fmt.Sprintf("%.2fs %7d %s\n", time.Since(start).Seconds(), nbytes, url)
}

/*
0.22s  293945 https://mail.ru
0.22s  193514 https://dtf.ru
0.23s   13619 https://ya.ru
0.28s   47869 https://geekbrains.ru/
0.34s   14325 https://google.com
0.35s    9308 https://vk.com
0.50s  331317 https://twitter.com
0.67s  239834 https://habrahabr.ru
0.68s   70367 https://gmail.com
0.92s  125150 https://github.com/alexeybobkov47
1.25s   60007 https://twitch.com
1.44s  409944 https://youtube.com
1.44s elapsed
*/
