package main

func (s *Server) doKeepAlives() bool {
	return atomic.LoadInt32(&s.disableKeepAlives) == 0 && !s.shuttingDown()
}

func xxx() {
    panicChan := make(chan interface{}, 1)
	go func() {
		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()
		h.handler.ServeHTTP(tw, r)
		close(done)
	}()
	select {
	case p := <-panicChan:
		panic(p)
    }
}
