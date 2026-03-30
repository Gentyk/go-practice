// /////////////////////
// ///// Задача 2
// /////////////////////
package avito1

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Есть функция, работающая неопределенно долго и возвращающая число.
// Ее тело нельзя изменять (представим, что внутри сетевой запрос).
func unpredictableFunc() int64 {
	rnd := rand.Int63n(5000)
	time.Sleep(time.Duration(rnd) * time.Millisecond)

	return rnd
}

// Нужно изменить функцию обертку, которая будет работать с заданным таймаутом (например, 1 секунду).
// Если долгая функция отработала за это время, то возвращаем результат.
// Если нет, возвращаем ошибку. Результат работы тогда не важен.
// Еще нужно измерить, сколько выполнялась эта функция (просто вывести лог).
// Сигнатуру функцию обертки менять можно.

func predictableFunc1(timeout int) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	var wg sync.WaitGroup

	result := make(chan int64)
	wg.Add(1)
	start := time.Now()

	go func() {
		defer wg.Done()
		defer close(result)
		data := unpredictableFunc()
		select {
		case result <- data:
		case <-ctx.Done():
		}
	}()

	select {
	case data := <-result:
		fmt.Printf("time since %v", time.Since(start))
		return data, nil
	case <-ctx.Done():
		wg.Wait()
		return 0, fmt.Errorf("timeout")
	}
}

func predictableFunc2(timeout int) (int64, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup

	result := make(chan int64)
	wg.Add(1)
	start := time.Now()

	go func() {
		defer wg.Done()
		defer close(result)
		data := unpredictableFunc()
		select {
		case result <- data:
		case <-ctx.Done():
		}
	}()

	select {
	case data := <-result:
		fmt.Printf("time since %v", time.Since(start))
		return data, nil
	case <-time.After(time.Second * time.Duration(timeout)):
		cancel()
		wg.Wait()
		return 0, fmt.Errorf("timeout")
	}
}
