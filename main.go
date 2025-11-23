package main

import (
	"fmt"
	"math"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	//defining variables
	var usrName = os.Getenv("USER")
	var ext string
	var interval, maxFiles int
	var dir = "/home/" + usrName + "/Desktop"

	//print greeting
	fmt.Printf("Привет, %s!\nЭто go-project, написанный just for fun\n", usrName)
	fmt.Println("Это программа будет создавать файлы на рабочем столе с заданым расширением и интервалом.\n" +
		"Введите расширение (без точки), интервал, макс. количество файлов")

	//getting data
	_, err := fmt.Scan(&ext, &interval, &maxFiles)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	fmt.Printf("Создаю %d файла(ов) '.%s' в дикертории '%s' каждые %d секунд\n", maxFiles, ext, dir, interval)

	//main loop
	wg.Add(2)
	go func(n int, interval int) {
		//defining number of segments of bar
		segment := float64(100) / float64(n)
		baseSegment := segment
		//bar output
		for segment <= 100 {
			//cleaning thr output
			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			commandErr := cmd.Run()
			if commandErr != nil {
				fmt.Print("Запускайте вне IDE!\n")
			}
			//setting parameters
			ceilSeg := int(math.Ceil(segment))
			currentBlocks := strings.Repeat("█", ceilSeg)
			currentDots := strings.Repeat(".", 100-ceilSeg)

			//printing the bar
			fmt.Printf(
				"╔════════════════════════════════════════════════════════════════════════════════════════════════════════╗\n"+
					"║                                                                                                        ║\n"+
					"║                                             Создаю файлы….                                             ║\n"+
					"║                                                                                                        ║\n"+
					"║ ["+currentBlocks+currentDots+"] ║\n"+
					"║                                                                                                        ║\n"+
					"║                                        Выполнено: (%03d / 100)%%                                         ║\n"+
					"║                                                                                                        ║\n"+
					"╚════════════════════════════════════════════════════════════════════════════════════════════════════════╝\n", int(segment))
			time.Sleep(time.Duration(interval) * time.Second)
			segment = segment + baseSegment
		}
		wg.Done()
	}(maxFiles, interval)
	go fileCreator(ext, dir, maxFiles, interval)
	wg.Wait()
	fmt.Println("Завершено исполнение горутин")
}

func fileCreator(ext string, dir string, maxFiles int, interval int) {
	//creating files within a loop
	for i := 0; i < maxFiles; i++ {
		//size := i + 1
		//matrix := make([][]int, size)
		//for j := 0; j <= size; i++ {
		//	matrix[j] = make([]int, size)
		//}
		//fmt.Println(matrix)
		path := dir + "/" + strconv.Itoa(i) + "file." + ext
		file, errFileCreate := os.Create(path)
		if errFileCreate != nil {
			fmt.Print(errFileCreate)
			os.Exit(1)
		}
		text := "This is file #" + strconv.Itoa(i)
		_, errFileWriting := file.WriteString(text)
		if errFileWriting != nil {
			fmt.Print(errFileWriting)
			os.Exit(1)
		}
		errFileClosing := file.Close()
		if errFileClosing != nil {
			fmt.Print(errFileClosing)
			os.Exit(1)
		}
		time.Sleep(time.Duration(interval) * time.Second)
	}
	wg.Done()
}
