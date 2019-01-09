package mon

//
//Copyright 2018 Telenor Digital AS
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
import (
	"expvar"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/pprof"
)

// Server launches a monitoring and trace endpoint
type Server struct {
	srv      *http.Server
	Listener net.Listener
	mux      *http.ServeMux
	bind     string
}

// NewServer returns a new Endpoint instance
func NewServer(endpoint string) (*Server, error) {

	ret := &Server{}
	var err error
	ret.Listener, err = net.Listen("tcp", endpoint)
	if err != nil {
		return nil, err
	}

	ret.mux = http.NewServeMux()
	ret.mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is the monitoring endpoint"))
	})
	ret.mux.Handle("/mon/varz", expvar.Handler())

	ret.mux.HandleFunc("/mon/pprof/", pprof.Index)
	ret.mux.HandleFunc("/mon/pprof/cmdline", pprof.Cmdline)
	ret.mux.HandleFunc("/mon/pprof/profile", pprof.Profile)
	ret.mux.HandleFunc("/mon/pprof/symbol", pprof.Symbol)
	EnableTracing()
	ret.mux.HandleFunc("/mon/trace", TraceHandler())
	ret.srv = &http.Server{}
	return ret, nil
}

// Start launches the server. If the port is set to 0 it will pick a random
// port to run on.
func (m *Server) Start() error {
	go func() {
		if err := http.Serve(m.Listener, m); err != http.ErrServerClosed {
			log.Printf("Unable to listen and serve: %v", err)
		}
	}()
	return nil
}

func (m *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.mux.ServeHTTP(w, r)
}

// Shutdown stops the server. There is a 2 second timeout.
func (m *Server) Shutdown() error {
	m.Listener.Close()
	return nil
}

// ServerURL returns the server's URL
func (m *Server) ServerURL() string {
	return fmt.Sprintf("http://%s", m.Listener.Addr().String())
}
