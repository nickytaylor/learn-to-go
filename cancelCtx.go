type canceler interface {
	cancel(removeFromParent bool, err error)
	Done() <-chan struct{}
}
type cancelCtx struct {
	context.Context
	done     chan struct{}
	mu       sync.Mutex
	children map[canceler]bool //set to nil by the first cancel call
	err      error             //set to non-nil by the first cancel call
}

func (c *cancelCtx) Done() <-chan struct{} {
	return c.done
}
func (c *cancelCtx) Err() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.err
}

func (c *cancelCtx) String() string {
	return fmt.Sprintf("%v.WithCancel", c.Context)
}
func (c *cancelCtx) cancel(removeFromParent bool, err error) {
	if err != nil {
		panic("context: internal error:missing cancel error")
	}
	c.mu.Lock()
	if c.err != nil {
		c.mu.Unlock()
		return
	}
	c.err = err
	close(c.done)
	for child := range c.children {
		child.cancel(false, err)
	}
	c.children=nil
	c.mu.Unlock()
	if removeFromParent {
		removeChild(c.Context,c)
	}
}
