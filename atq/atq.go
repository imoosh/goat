package atq

import (
	"context"
	"github.com/hibiken/asynq"
)

type AsyncTaskQ struct {
	cli *asynq.Client
	srv *asynq.Server
	mux *asynq.ServeMux
}

type Config struct {
	Addr string
	Auth string
}

type HandleFunc func(context.Context, *asynq.Task) error

func NewAsyncTaskQ(c Config) (acp *AsyncTaskQ) {
	opt := asynq.RedisClientOpt{Addr: c.Addr, Password: c.Auth}
	cfg := asynq.Config{Logger: logger{}, LogLevel: defaultLoggerLevel}
	return &AsyncTaskQ{
		cli: asynq.NewClient(opt),
		srv: asynq.NewServer(opt, cfg),
		mux: asynq.NewServeMux(),
	}
}

func (atq *AsyncTaskQ) Run() {
	atq.srv.Run(atq.mux)
}

func (atq *AsyncTaskQ) Register(pattern string, handler HandleFunc) {
	atq.mux.HandleFunc(pattern, handler)
}

/*
Enqueue(task, asynq.ProcessAt(t))
Enqueue(task, asynq.ProcessIn(t))
*/
func (atq *AsyncTaskQ) Enqueue(task *asynq.Task, opts ...asynq.Option) (err error) {
	_, err = atq.cli.Enqueue(task, opts...)
	return
}
