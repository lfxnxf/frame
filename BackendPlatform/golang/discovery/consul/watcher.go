package consul

// "time"

// log "github.com/lfxnxf/frame/BackendPlatform/golang/logging"

type ConsulWatcher struct {
	client *Client
	// update    chan *naming.Update
	side      chan int
	close     chan bool
	address   []string
	lastIndex uint64
}

type ConsulResolver struct {
	w      *ConsulWatcher
	nodes  []string
	scheme string
	tag    string
	proto  string
}

func NewConsulResolver(nodes []string, tag string) *ConsulResolver {
	r := ConsulResolver{
		nodes:  nodes,
		scheme: "http",
		tag:    tag,
		proto:  "http",
	}
	return &r
}

func newConsulWatcher(nodes []string, scheme string) (*ConsulWatcher, error) {
	c, err := New(nodes, scheme, nil)
	if err != nil {
		return nil, err
	}
	w := ConsulWatcher{
		client: c,
		// update: make(chan *naming.Update, 1),
		side:  make(chan int, 1),
		close: make(chan bool),
	}
	return &w, nil
}

func (w *ConsulWatcher) watch(target, tag, proto string) ([]string, error) {
	return nil, nil
}

func (w *ConsulWatcher) Close() {
	close(w.side)
	close(w.close)
}
