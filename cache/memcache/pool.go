package memcache

import (
	"context"
	"io"
	"time"

	"goat/container/pool"
	xtime "goat/time"
)

// Config memcache config.
type Config struct {
	*pool.Config

	Name         string // memcache name, for trace
	Proto        string
	Addr         string
	DialTimeout  xtime.Duration
	ReadTimeout  xtime.Duration
	WriteTimeout xtime.Duration
}

// Pool memcache connection pool struct.
type Pool struct {
	p pool.Pool
	c *Config
}

// NewPool new a memcache conn pool.
func NewPool(c *Config) (p *Pool) {
	if c.DialTimeout <= 0 || c.ReadTimeout <= 0 || c.WriteTimeout <= 0 {
		panic("must config memcache timeout")
	}
	p1 := pool.NewList(c.Config)
	cnop := DialConnectTimeout(time.Duration(c.DialTimeout))
	rdop := DialReadTimeout(time.Duration(c.ReadTimeout))
	wrop := DialWriteTimeout(time.Duration(c.WriteTimeout))
	p1.New = func(ctx context.Context) (io.Closer, error) {
		return Dial(c.Proto, c.Addr, cnop, rdop, wrop)
	}
	p = &Pool{p: p1, c: c}
	return
}

// Get gets a connection. The application must close the returned connection.
// This method always returns a valid connection so that applications can defer
// error handling to the first use of the connection. If there is an error
// getting an underlying connection, then the connection Err, Do, Send, Flush
// and Receive methods return that error.
func (p *Pool) Get(ctx context.Context) Conn {
	c, err := p.p.Get(ctx)
	if err != nil {
		return errorConnection{err}
	}
	c1, _ := c.(Conn)
	return &pooledConnection{p: p, c: c1.WithContext(ctx), ctx: ctx}
}

// Close release the resources used by the pool.
func (p *Pool) Close() error {
	return p.p.Close()
}

type pooledConnection struct {
	p   *Pool
	c   Conn
	ctx context.Context
}

func (pc *pooledConnection) Close() error {
	c := pc.c
	if _, ok := c.(errorConnection); ok {
		return nil
	}
	pc.c = errorConnection{ErrConnClosed}
	pc.p.p.Put(context.Background(), c, c.Err() != nil)
	return nil
}

func (pc *pooledConnection) Err() error {
	return pc.c.Err()
}

func (pc *pooledConnection) Set(item *Item) (err error) {
	now := time.Now()
	err = pc.c.Set(item)
	return
}

func (pc *pooledConnection) Add(item *Item) (err error) {
	now := time.Now()
	err = pc.c.Add(item)
	return
}

func (pc *pooledConnection) Replace(item *Item) (err error) {
	now := time.Now()
	err = pc.c.Replace(item)
	return
}

func (pc *pooledConnection) CompareAndSwap(item *Item) (err error) {
	now := time.Now()
	err = pc.c.CompareAndSwap(item)
	return
}

func (pc *pooledConnection) Get(key string) (r *Item, err error) {
	now := time.Now()
	r, err = pc.c.Get(key)
	return
}

func (pc *pooledConnection) GetMulti(keys []string) (res map[string]*Item, err error) {
	// if keys is empty slice returns empty map direct
	if len(keys) == 0 {
		return make(map[string]*Item), nil
	}
	now := time.Now()
	res, err = pc.c.GetMulti(keys)
	return
}

func (pc *pooledConnection) Touch(key string, timeout int32) (err error) {
	err = pc.c.Touch(key, timeout)
	return
}

func (pc *pooledConnection) Scan(item *Item, v interface{}) error {
	return pc.c.Scan(item, v)
}

func (pc *pooledConnection) WithContext(ctx context.Context) Conn {
	// TODO: set context
	pc.ctx = ctx
	return pc
}

func (pc *pooledConnection) Delete(key string) (err error) {
	err = pc.c.Delete(key)
	return
}

func (pc *pooledConnection) Increment(key string, delta uint64) (newValue uint64, err error) {
	newValue, err = pc.c.Increment(key, delta)
	return
}

func (pc *pooledConnection) Decrement(key string, delta uint64) (newValue uint64, err error) {
	newValue, err = pc.c.Decrement(key, delta)
	return
}

type errorConnection struct{ err error }

func (ec errorConnection) Err() error                                         { return ec.err }
func (ec errorConnection) Close() error                                       { return ec.err }
func (ec errorConnection) Add(item *Item) error                               { return ec.err }
func (ec errorConnection) Set(item *Item) error                               { return ec.err }
func (ec errorConnection) Replace(item *Item) error                           { return ec.err }
func (ec errorConnection) CompareAndSwap(item *Item) error                    { return ec.err }
func (ec errorConnection) Get(key string) (*Item, error)                      { return nil, ec.err }
func (ec errorConnection) GetMulti(keys []string) (map[string]*Item, error)   { return nil, ec.err }
func (ec errorConnection) Touch(key string, timeout int32) error              { return ec.err }
func (ec errorConnection) Delete(key string) error                            { return ec.err }
func (ec errorConnection) Increment(key string, delta uint64) (uint64, error) { return 0, ec.err }
func (ec errorConnection) Decrement(key string, delta uint64) (uint64, error) { return 0, ec.err }
func (ec errorConnection) Scan(item *Item, v interface{}) error               { return ec.err }
func (ec errorConnection) WithContext(ctx context.Context) Conn               { return ec }
