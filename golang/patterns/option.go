package patterns

const (
	default_port       = 80
	default_timeout    = 60000
	default_encryption = false
)

// server struct. We want explicitly to keep properties private
type server struct {
	host       string
	port       uint16
	timeout_ms uint64
	encryption bool
}

// -----------------------------
//  With all params constructor
// -----------------------------

// NewServerDiscarded
func NewServerDiscarded(host string, port uint16, ms uint64, encryption bool) *server {
	return &server{
		host:       host,
		port:       default_port,
		timeout_ms: default_timeout,
		encryption: default_encryption,
	}
}

// ---------------------
//  With Option pattern
// ---------------------

// serverOpt
type serverOpt func(*server)

// NewServer constructor accepting host name and serverOpt
func NewServer(host string, opts ...serverOpt) *server {
	server := &server{
		host,
		default_port,
		default_timeout,
		default_encryption,
	}
	for _, opt := range opts {
		opt(server)
	}
	return server
}

// WithPort determines port for listen to
func WithPort(port uint16) serverOpt {
	return func(s *server) {
		s.port = port
	}
}

// WithTimeoutMs determines millis to timeout
func WithTimeoutMs(ms uint64) serverOpt {
	return func(s *server) {
		s.timeout_ms = ms
	}
}

// WithEncryption is to enable server encryption
func WithEncryption() serverOpt {
	return func(s *server) {
		s.encryption = true
	}
}

// WithoutEncryption is to disable server encryption
func WithoutEncryption() serverOpt {
	return func(s *server) {
		s.encryption = false
	}
}
