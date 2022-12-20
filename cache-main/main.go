package main

import (
	"context"
	"errors"
	"fmt"
	"sync/atomic"
	"time"
)

var count int64

func main() {

	cache := NewLogCache(&CacheConf{
		FetchSize:     3000,
		SignalSize:    3000,
		BroadcastSize: 9000,
		CacheSize:     100000,
		FlexDefault:   100,
	})

	//go func() {
	//	for {
	//		time.Sleep(10 * time.Second)
	//		cache.NoticeWorker()
	//	}
	//}()
	//初始化消费者
	for i := 0; i < 20; i++ {
		go func(i int) {
			var fetchedData = make([]LogRecord, cache.Conf.FetchSize)
			for {
				//fetchedData = fetchedData[0:0]
				err := cache.FetchData(context.TODO(), &fetchedData)
				if err != nil && errors.Is(err, ExistErr) {
					fmt.Println(err.Error())
					return
				}
				//fetchedData, err := cache.FetchDataV2(context.TODO())
				//if err != nil && errors.Is(err, ExistErr) {
				//	return
				//}
				time.Sleep(300 * time.Millisecond)
				for _, datum := range fetchedData {
					atomic.AddInt64(&count, datum.Id)
				}
			}
		}(i)
	}

	//生产者
	now := time.Now()
	for i := 0; i < 400000; i++ {
		err := cache.PutData(context.TODO(), []LogRecord{
			{Id: 1}, {Id: 2}, {Id: 3}, {Id: 7}, {Id: 8}, {Id: 9},
		})
		if (i > 100000 && i < 200000) || i > 300000 {
			time.Sleep(100 * time.Millisecond)
		}
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	now2 := time.Now()

	time.Sleep(time.Second * 5)

	cache.ExistWorkerBroadCase()
	time.Sleep(time.Second * 15)

	fmt.Println("count: ", count) //21 * 100000
	fmt.Println("taste: ", now2.Sub(now).Milliseconds())

	return

	//cond := sync.NewCond(&sync.Mutex{})

	//go func() {
	//	time.Sleep(3 * time.Second)
	//	if len(data) > 3000 {
	//		cond.Broadcast()
	//	} else {
	//		cond.Signal()
	//	}
	//}()

	//cancel, cancelFunc := context.WithCancel(context.TODO())

	//
	//for i := 0; i < 20; i++ {
	//	go read(context.TODO(), "reader" + strconv.Itoa(i), cond)
	//}
	//now := time.Now()
	//
	//for i := 0; i < 2000; i++ {
	//	write("writer", cond)
	//}
	//now2 := time.Now()
	//
	//time.Sleep(time.Second * 10)
	////cancelFunc()
	//existFlag1 = 1
	//cond.Broadcast()
	//
	//
	//time.Sleep(time.Second * 5)
	//
	//fmt.Println("data:", len(data))
	//fmt.Println("count: ", count) //21 * 100000
	//fmt.Println("taste: ", now2.Sub(now).Milliseconds())
}
