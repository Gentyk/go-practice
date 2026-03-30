////////////////////
///////Задача 1
////////////////////
// Есть master сервис и 2 реплики. Дан интерфейс Client, с методом Get
// Метод Get может выполняться долго
// Нужно реализовать метод GetHedged, который:
//   - делает запрос в master
//   - если ответ не приходит за hedgedDelay, сдеалть еще 2 запроса в replicas
//   - продолжать ждать ответ от master, пока не закончится maxDelay

package avito1

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Client interface {
	Get(ctx context.Context, key string) (result string, err error)
}

type Result struct {
	r   string
	err error
}

func GetHedged(
	ctx context.Context,
	master Client,
	replicas [2]Client,
	key string,
	hedgedDelay time.Duration,
	maxDelay time.Duration,
) (result string, err error) {
	// implement
	localCtx, cancel := context.WithTimeout(ctx, maxDelay)
	defer cancel()
	var wg sync.WaitGroup
	results := make(chan Result, 3)

	wg.Add(1)
	go func() {
		defer wg.Done()
		answ, err := master.Get(localCtx, key)
		select {
		case results <- Result{answ, err}:
		case <-localCtx.Done():
		}
	}()

	select {
	case val := <-results:
		close(results)
		return val.r, val.err
	case <-time.After(hedgedDelay):
		for _, r := range replicas {
			wg.Add(1)
			go func() {
				defer wg.Done()
				answ, err := r.Get(localCtx, key)
				select {
				case results <- Result{answ, err}:
				case <-localCtx.Done():
				}
			}()
		}
	case <-localCtx.Done():
		close(results)
		return "", fmt.Errorf("maxDelay timeout")
	}

	select {
	case val := <-results:
		cancel()
		wg.Wait()
		close(results)
		return val.r, val.err
	case <-localCtx.Done():
		wg.Wait()
		close(results)
		return "", fmt.Errorf("maxDelay timeout")
	}
}

func GetHedged2(
	ctx context.Context,
	master Client,
	replicas [2]Client,
	key string,
	hedgedDelay time.Duration,
	maxDelay time.Duration,
) (result string, err error) {
	// implement
	localCtx, cancel := context.WithTimeout(ctx, maxDelay)
	defer cancel()
	var wg sync.WaitGroup
	results := make(chan Result, 3)

	wg.Add(1)
	go func() {
		defer wg.Done()
		answ, err := master.Get(ctx, key)
		select {
		case results <- Result{answ, err}:
		case <-localCtx.Done():
		}
	}()

	t := time.AfterFunc(hedgedDelay, func() {
		for _, r := range replicas {
			wg.Add(1)
			go func() {
				defer wg.Done()
				answ, err := r.Get(ctx, key)
				select {
				case results <- Result{answ, err}:
				case <-localCtx.Done():
				}
			}()
		}
	})
	defer t.Stop()

	select {
	case val := <-results:
		go func() {
			cancel()
			wg.Wait()
			close(results)
		}()
		return val.r, val.err
	case <-localCtx.Done():
		wg.Wait()
		close(results)
		return "", fmt.Errorf("maxDelay timeout")
	}
}
