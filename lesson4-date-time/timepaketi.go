package main

import (
	"fmt"
	"time"
)

func main() {
	now_time := time.Now().Unix()
	fmt.Println(now_time)

	/*
		Unix time nedir ?
		1 ocak 1970 den itibaren geçen saniye sayısıdır.
	*/

	t := time.Date(2018, time.January, 1, 0, 0, 0, 0, time.UTC) //yeni tarih

	fmt.Println(t)

	now := time.Now()

	fmt.Printf("Mevcut zaman : %s\n", now)

	fmt.Printf("Ay bilgisi : %d\n", now.Month())
	fmt.Printf("Gun bilgisi : %d\n", now.Day())
	fmt.Printf("Saat bilgisi : %d\n", now.Hour())
	fmt.Printf("Dakika bilgisi : %d\n", now.Minute())
	fmt.Printf("Saniye bilgisi : %d\n", now.Second())

	tomorrow := now.AddDate(0, 0, 1)
	fmt.Printf("Tomorrow : %s\n", tomorrow)

	// tarih karsilastirma

	fmt.Printf("%s\n", now.Before(tomorrow))
	fmt.Printf("%s\n", now.After(tomorrow))
	fmt.Printf("%s\n", now.Equal(tomorrow))

	diff := tomorrow.Sub(now)

	fmt.Printf("Fark : %s\n", diff)
	fmt.Printf("Fark : %d\n", diff.Hours())
	fmt.Printf("Fark : %d\n", diff.Minutes())
	fmt.Printf("Fark : %d\n", diff.Seconds())
	fmt.Printf("Fark : %d\n", diff.Nanoseconds())
	fmt.Printf("Fark : %d\n", diff.Milliseconds())
}
