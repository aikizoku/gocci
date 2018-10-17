package worker

import (
	"net/http"

	"github.com/aikizoku/beego/src/handler"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

// SampleHandler ... サンプルのハンドラ定義
type SampleHandler struct {
}

// Cron ... Cronから実行されるハンドラ
func (h *SampleHandler) Cron(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	log.Debugf(ctx, "call cron handler")
	handler.RenderSuccess(w)
}

// TaskQueue ... TaskQueueで実行されるハンドラ
func (h *SampleHandler) TaskQueue(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	log.Debugf(ctx, "call task queue handler")
	handler.RenderSuccess(w)
}

// NewSampleHandler ... SampleHandlerを作成する
func NewSampleHandler() *SampleHandler {
	return &SampleHandler{}
}
