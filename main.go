package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

type adres struct {
	host string
}

func main() {

	file, _ := os.OpenFile("dosya.txt", os.O_RDONLY, 0755)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		son := adres{scanner.Text()}

		fmt.Println("Taranan Host ip: ", son.host)
		for i := 1; i < 7000; i++ {

			hostIp := fmt.Sprintf("%s:%d", son.host, i)
			_, err := net.DialTimeout("tcp", hostIp, 10*time.Millisecond)

			if err != nil {

			} else {
				fmt.Printf("Açık Portlar %d\n", i)
			}

		}

	}

}
