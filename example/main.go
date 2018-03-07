package main

import (
	"context"
	"github.com/rakeshnair/waitctx"
	"time"
	"log"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func(wc *waitctx.WaitCtx) {

		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-wc.Stop():
				{
					log.Println("exiting task 1")
					wc.MarkAsDone()
					return
				}
			case <-ticker.C:
				{
					log.Println("starting task 2")
					time.Sleep(time.Second * 3)
					log.Println("finished task 2")
				}
			}
		}
		return
	}(waitctx.RoutineCtx(ctx))

	go func(wc *waitctx.WaitCtx) {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-wc.Stop():
				{
					log.Println("exiting task 2")
					wc.MarkAsDone()
					return
				}
			case <-ticker.C:
				{
					log.Println("starting task 2")
					time.Sleep(time.Second * 3)
					log.Println("finished task 2")
				}
			}
		}
		return
	}(waitctx.RoutineCtx(ctx))

	t := time.Second * 5
	log.Printf("going to sleep for %dsec and then call cancel\n", t)
	time.Sleep(time.Second * 2)
	log.Println("going to cancel")
	cancel()
	waitctx.ServiceCtx(ctx).Wait()
	log.Println("program exiting")
}
