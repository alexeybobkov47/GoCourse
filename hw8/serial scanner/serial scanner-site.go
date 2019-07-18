package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	for _, url := range os.Args[1:] {
		fetch(url)
	}
	//"args": ["https://ya.ru","https://mail.ru","https://gmail.com","https://dtf.ru","https://youtube.com","https://habrahabr.ru","https://vk.com","https://twitch.com","https://google.com","https://twitter.com","https://geekbrains.ru/","https://github.com/alexeybobkov47"]
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	//bytes, err := ioutil.ReadAll(resp.Body)
	nbytes, err := io.Copy(ioutil.Discard, resp.Body) // to devNull
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%.2fs %7d %s\n", time.Since(start).Seconds(), nbytes, url)
}

/*
0.23s   13729 https://ya.ru
0.16s  293893 https://mail.ru
0.59s   70419 https://gmail.com
0.13s  174769 https://dtf.ru
1.24s  430454 https://youtube.com
0.71s  240005 https://habrahabr.ru
0.25s    9309 https://vk.com
1.38s   60007 https://twitch.com
0.19s   14353 https://google.com
0.48s  331317 https://twitter.com
0.11s   47868 https://geekbrains.ru/
0.88s  125150 https://github.com/alexeybobkov47
6.34s elapsed
*/
