package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

type server struct {
	addr  string
	proxy *httputil.ReverseProxy
}
type Server interface {
	Address() string
	IsAlive() bool
	Server(w http.ResponseWriter, r *http.Request)
}

func newServer(addr string) *server {
	serverURL, err := url.Parse(addr)
	//create a function to handle err
	handleError(err)
	return &server{
		addr:  addr,
		proxy: httputil.NewSingleHostReverseProxy(serverURL),
	}
}

type LoadBalancer struct {
	port            string
	roundRobincount int
	servers         []Server
}

func NewLoadBalancer(port string, servers []Server) *LoadBalancer {
	return &LoadBalancer{
		port:            port,
		roundRobincount: 0,
		servers:         servers,
	}
}

func handleError(err error) {
	if err != nil {
		fmt.Println("error ", err)
		os.Exit(1)
		// panic(err)
	}
}

func (s *server) Address() string { return s.addr }

func (s *server) IsAlive() bool { return true }

func (s *server) Server(w http.ResponseWriter, r *http.Request) {
	s.proxy.ServeHTTP(w, r)

}

func (lb *LoadBalancer) getNextAvailableServer() Server {
	server := lb.servers[lb.roundRobincount%len((lb.servers))]
	for server.IsAlive() {
		lb.roundRobincount++
		server = lb.servers[lb.roundRobincount%len((lb.servers))]
	}
	lb.roundRobincount++
	return server
}

func (lb *LoadBalancer) serverProxy(w http.ResponseWriter, r *http.Request) {
	targetServer := lb.getNextAvailableServer()
	fmt.Println("forwarding request to address ", targetServer.Address())
	targetServer.Server(w, r)
}

func main() {
	servers := []Server{
		newServer("https://www.duckduckgo.com"),
		newServer("https://www.youtube.com"),
		newServer("https://www.bing.com"),
	}
	lb := NewLoadBalancer("8000", servers)
	handleRedirect := func(w http.ResponseWriter, r *http.Request) {
		lb.serverProxy(w, r)
	}
	http.HandleFunc("/", handleRedirect)
	fmt.Println("serving request at localhost:", lb.port)
	http.ListenAndServe(":"+lb.port, nil)

}
