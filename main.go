package main

import (
	"fmt"
	"os"
)

func main() {
	//defining variables
	var usrName = os.Getenv("USER")
	var ext string
	var interval int
	var dir = "/home/" + usrName + "/Desktop"

	//print greeting
	fmt.Printf("Привет, %s!\nЭто go-project, написанный just for fun\n", usrName)
	fmt.Println("Это программа будет создавать файлы на рабочем столе с заданым расширением и интервалом.\n" +
		"Введите расширение (без точки) и интервал")
	_, err := fmt.Scan(&ext, &interval)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	fmt.Printf("Создаю файлы '.%s' в дикертории '%s' каждые %d секунд\n", ext, dir, interval)

	//main loop

}
