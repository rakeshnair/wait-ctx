package waitctx

import (
	"context"
	"log"
	"sync"
)

var once sync.Once
var wc *WaitCtx

func ServiceCtx(ctx context.Context) *WaitCtx {
	once.Do(func() {
		log.Println("waitCtx registered")
		var wg sync.WaitGroup
		wc = &WaitCtx{
			ctx: ctx,
			wg:  &wg,
		}
	})
	return wc
}

func RoutineCtx(ctx context.Context) *WaitCtx {
	wc := ServiceCtx(ctx)
	log.Println("new routine added to service waitCtx")
	wc.wg.Add(1)
	return wc
}

type WaitCtx struct {
	ctx context.Context
	wg  *sync.WaitGroup
}

func (wc *WaitCtx) Stop() <-chan struct{} {
	return wc.ctx.Done()
}

func (wc *WaitCtx) MarkAsDone() {
	wc.wg.Done()
}

func (wc *WaitCtx) Wait() {
	wc.wg.Wait()
}
