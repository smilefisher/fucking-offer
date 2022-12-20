package main

import (
	"context"
	"errors"
	"reflect"
	"sync"
	"sync/atomic"
)

type LogCache struct {
	Data      []LogRecord
	Cond      *sync.Cond
	ExistFlag *int32 //消费者退出
	Conf      *CacheConf
}

// 构造cache
func NewLogCache(Conf *CacheConf) *LogCache {
	var f int32 = 0
	return &LogCache{
		Data:      make([]LogRecord, 0, Conf.CacheSize),
		Cond:      sync.NewCond(&sync.Mutex{}),
		ExistFlag: &f,
		Conf:      Conf,
	}
}

type LogRecord struct {
	Id int64 `json:"id"`
}

type CacheConf struct {
	FetchSize     int
	SignalSize    int
	BroadcastSize int //倍数
	CacheSize     int
}

var fullError = errors.New("cache full")

func (l *LogCache) PutData(ctx context.Context, data []LogRecord) error {

	l.Cond.L.Lock()
	defer l.Cond.L.Unlock()
	if len(l.Data)+len(data) > l.Conf.CacheSize {
		return fullError
	}
	l.Data = append(l.Data, data...)
	length := len(l.Data)

	of := reflect.TypeOf(l.Data)
	of.FieldByName("data")

	if length >= l.Conf.BroadcastSize {
		l.Cond.Broadcast()
	} else if length >= l.Conf.SignalSize {
		l.Cond.Signal()
	}
	return nil
}

var ExistErr = errors.New("exist error")

func (l *LogCache) FetchData(ctx context.Context, record *[]LogRecord) (err error) {
	*record = (*record)[0:0]
	l.Cond.L.Lock()
	for len(l.Data) == 0 {
		l.Cond.Wait()
		if atomic.LoadInt32(l.ExistFlag) != 0 {
			l.Cond.L.Unlock()
			return ExistErr
		}
	}
	//fmt.Printf("%p \n ",l.Data)
	length := len(l.Data)
	if length > l.Conf.FetchSize {

		*record = append(*record, l.Data[(length-l.Conf.FetchSize):]...)
		l.Data = l.Data[:length-l.Conf.FetchSize]
	} else {
		*record = append(*record, l.Data...)
		l.Data = l.Data[0:0]
	}
	l.Cond.L.Unlock()
	return nil
}

func (l *LogCache) ExistWorkerBroadCase() {
	atomic.StoreInt32(l.ExistFlag, 1)
	l.Cond.Broadcast()
}

// 通知worker获取数据
func (l *LogCache) NoticeWorker() {
	if len(l.Data) > l.Conf.BroadcastSize {
		l.Cond.Broadcast()
	} else {
		l.Cond.Signal()
	}
}
