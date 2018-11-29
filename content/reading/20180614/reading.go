package main

type Hijacker interface {
	// After a call to Hijack, the original Request.Body must not be used.
	Hijack() (net.Conn, *bufio.ReadWriter, error)
}

func xxx() {
    // ...
    if proto := c.tlsState.NegotiatedProtocol; validNPN(proto) {
	if fn := c.server.TLSNextProto[proto]; fn != nil {
		h := initNPNRequest{tlsConn, serverHandler{c.server}}
		fn(c.server, tlsConn, h)
	}
	return
}
}

// onceSetNextProtoDefaults configures HTTP/2, if the user hasn't
// configured otherwise. (by setting srv.TLSNextProto non-nil)
// It must only be called via srv.nextProtoOnce (use srv.setupHTTP2_*).
func (srv *Server) onceSetNextProtoDefaults() {
	if strings.Contains(os.Getenv("GODEBUG"), "http2server=0") {
		return
	}
	// Enable HTTP/2 by default if the user hasn't otherwise
	// configured their TLSNextProto map.
	if srv.TLSNextProto == nil {
		conf := &http2Server{
			NewWriteScheduler: func() http2WriteScheduler { return http2NewPriorityWriteScheduler(nil) },
		}
		srv.nextProtoErr = http2ConfigureServer(srv, conf)
	}
}
