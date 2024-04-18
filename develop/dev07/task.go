package main

import (
	"fmt"
	"sync"
	"time"
)

/*
=== Or channel ===

Реализовать функцию, которая будет объединять один или более done каналов в single канал если один из его составляющих каналов закроется.
Одним из вариантов было бы очевидно написать выражение при помощи select, которое бы реализовывало эту связь,
однако иногда неизестно общее число done каналов, с которыми вы работаете в рантайме.
В этом случае удобнее использовать вызов единственной функции, которая, приняв на вход один или более or каналов, реализовывала весь функционал.

Определение функции:
var or func(channels ...<- chan interface{}) <- chan interface{}

Пример использования функции:
sig := func(after time.Duration) <- chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
}()
return c
}

start := time.Now()
<-or (
	sig(2*time.Hour),
	sig(5*time.Minute),
	sig(1*time.Second),
	sig(1*time.Hour),
	sig(1*time.Minute),
)

fmt.Printf(“fone after %v”, time.Since(start))
*/

func mergeDoneChannels(doneChans ...<-chan struct{}) <-chan struct{} {
	merged := make(chan struct{})
	var wg sync.WaitGroup

	for _, doneChan := range doneChans {
		wg.Add(1)
		go func(ch <-chan struct{}) {
			defer wg.Done()
			<-ch
			close(merged)
		}(doneChan)
	}

	go func() {
		wg.Wait()
		close(merged)
	}()

	return merged
}

func main() {
	done1 := make(chan struct{})
	done2 := make(chan struct{})
	done3 := make(chan struct{})

	single := mergeDoneChannels(done1, done2, done3)

	// Пример использования single-канала
	go func() {
		for {
			_, ok := <-single
			if !ok {
				fmt.Println("Single channel closed")
				break
			}
		}
	}()

	// Закрытие одного из done-каналов
	time.Sleep(5 * time.Second)
	close(done2)
}
