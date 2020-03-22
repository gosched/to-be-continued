package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
)

func downloadAlbum(targetDir, albumURL, extension string, numberOfFiles int) {
	urls := []string{}

	temp := strings.Split(albumURL, "/")
	partURL := temp[len(temp)-1]

	albumURL += "/" + partURL + "-"

	for i := 1; i <= numberOfFiles; i++ {
		if 1 <= i && i < 10 {
			temp := albumURL + "000" + strconv.Itoa(i) + extension
			urls = append(urls, temp)
		}

		if 10 <= i && i < 100 {
			temp := albumURL + "00" + strconv.Itoa(i) + extension
			urls = append(urls, temp)
		}

		if 100 <= i && i < 1000 {
			temp := albumURL + "0" + strconv.Itoa(i) + extension
			urls = append(urls, temp)
		}
	}

	err := os.Mkdir(targetDir, 0777)
	if err != nil {
		log.Println(err)
	}

	currentDir, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	err = os.Chdir(targetDir)
	if err != nil {
		log.Println(err)
	}

	wg := &sync.WaitGroup{}
	limiter := make(chan bool, 4)
	for i, url := range urls {
		wg.Add(1)
		limiter <- true
		go downloadFile(i+1, url, limiter, wg)
	}
	wg.Wait()

	err = os.Chdir(currentDir)
	if err != nil {
		log.Println(err)
	}

}

func downloadFiles(targetDir, albumURL, extension string, begin, end int) {
	urls := []string{}

	for i := begin; i <= end; i++ {
		if 1 <= i && i < 10 {
			temp := albumURL + "000" + strconv.Itoa(i) + extension
			urls = append(urls, temp)
		}

		if 10 <= i && i < 100 {
			temp := albumURL + "00" + strconv.Itoa(i) + extension
			urls = append(urls, temp)
		}

		if 100 <= i && i < 1000 {
			temp := albumURL + "0" + strconv.Itoa(i) + extension
			urls = append(urls, temp)
		}
	}

	err := os.Mkdir(targetDir, 0777)
	if err != nil {
		log.Println(err)
	}

	currentDir, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	err = os.Chdir(targetDir)
	if err != nil {
		log.Println(err)
	}

	wg := &sync.WaitGroup{}
	limiter := make(chan bool, 4)
	for i, url := range urls {
		// fmt.Println(i, url)
		wg.Add(1)
		limiter <- true
		go downloadFile(i+1, url, limiter, wg)
	}
	wg.Wait()

	err = os.Chdir(currentDir)
	if err != nil {
		log.Println(err)
	}
}

func downloadFile(index int, url string, limiter chan bool, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("start to download :", index)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	temp := strings.Split(url, "/")
	name := temp[len(temp)-1]

	file, err := os.OpenFile(name, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("download finished:", index)

	<-limiter
}
