package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	//defining variables
	var usrName = os.Getenv("USER")
	var ext string
	var interval, maxFiles, n int
	var dir = "/home/" + usrName + "/Desktop"

	//print greeting
	fmt.Printf("Привет, %s!\nЭто go-project, написанный just for fun\n", usrName)
	fmt.Println("Это программа будет создавать файлы на рабочем столе с заданым расширением и интервалом, " +
		"а также считать до n.\n" +
		"Введите n, расширение (без точки), интервал, макс. количество файлов")

	//getting data
	_, err := fmt.Scan(&n, &ext, &interval, &maxFiles)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	fmt.Printf("Создаю %d файла(ов) '.%s' в дикертории '%s' каждые %d секунд\n", maxFiles, ext, dir, interval)

	//main loop
	wg.Add(2)
	go func(n int) {
		for i := 0; i <= n; i++ {
			fmt.Printf("%d\n", i)
			time.Sleep(2 * time.Second)
		}
		wg.Done()
	}(n)
	go fileCreator(ext, dir, maxFiles, interval)
	wg.Wait()
	fmt.Println("Завершено исполнение горутин")
}

func fileCreator(ext string, dir string, maxFiles int, interval int) {
	//creating files within a loop
	for i := 0; i < maxFiles; i++ {
		path := dir + "/" + strconv.Itoa(i) + "file." + ext
		_, err := os.Create(path)
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}
		time.Sleep(time.Duration(interval) * time.Second)
	}
	wg.Done()
}
