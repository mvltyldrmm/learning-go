package main

import (
	"fmt"
	"runtime"
	"time"
)

// Eş zamanlılık ve paralellik

// eş zamanlılık = birbirnden bağımsız çalıştırmalıdır.
// paralelde bir den çok çekirdek gereklidir.

// pc de aynı anda birden çok iş yapılır. music,code,video,game,chat,mail,etc. vs
// thread lar oluşur her iş için. işlemci kaynak ayırır bu işlemler için. Ve işleri paralel olarak yapar.
// önce thread 1. thread 1 de %10 olduktan sonra thread 2 ye geçer gibi. Tüm threadlara sınırlı kaynak ayırır.
// buda aynı anda çalışıyor gibi görünür. Ama aslında gibi öyle değil.
// döngü halinde sınırlı zaman üretir.

// EŞ zamanlılık :
// ana bir thread vardır.
// kod yazıoyurm ve arkada sıkıştırma var
// tek thread da yaparsam kod yazma bitmeden sıkıştırma olmaz.
// o yüzden farklı thread lara ihtiyaç vardır.
// prof kodlamada farklı thread lar ile bunlar çözülür.
// uzun sürecek kod farklı bir thread olur.


func main(){
	//EŞ ZAMANLILIK
	//goroutine
	//go tarafından çalışma zamanında yürütülen hafif bir iş parçacığıdır. (thread)

	//go fonksiyonu ile goroutine yaratılır. go f(x,y,z,d) şeklinde yaratılır.

	go xFunc() //bu islem bir thread oluşturur. arka tarafda islemleri yapabilir. 
	//bu thread sayesinde binlerce arka tarafda islem gerceklesebilir. calistirir ve direk bitirir.
	time.Sleep(100 * time.Millisecond)

	//go kelimesi ile baslatılan bağımsız calısan maaliyeti cok uzun bir fonksiyondur. Binlerde gorouting kullanılabilir
	//isletim seviyesinde bir thread değildir. Hafiftir, kaynakları hafiftir. 

	//PARALEL İŞLEMLER

	runtime.GOMAXPROCS(1)
	go YFunc()
	time.Sleep(100* time.Millisecond)
	//tamamen farklı bir thread. Herhangi bir zamanda tamamlanabilir.

}

func xFunc(){
	for l := byte('a'); l<= byte('z'); l++{
		fmt.Println(string(l))
	}
}

func YFunc(){
	for l := byte('a'); l<= byte('z'); l++{
		go fmt.Println(string(l))
	}
}