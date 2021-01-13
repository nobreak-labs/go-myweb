package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"

	"github.com/miekg/dns"
	"github.com/shirou/gopsutil/host"
)

const defaultEnvMessage = "Hello world"

//GetOutboundIP function gets Outbound IP
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	checkErr(err)
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

//GetHostname function gets hostname
func GetHostname() string {
	h, err := os.Hostname()
	checkErr(err)
	return h
}

//GetUptime function gets system uptime
func GetUptime() string {
	t, err := host.Uptime()
	checkErr(err)

	d := t / (60 * 60 * 24)
	h := (t - (d * 60 * 60 * 24)) / (60 * 60)
	m := ((t - (d * 60 * 60 * 24)) - (h * 60 * 60)) / 60

	upt := fmt.Sprintf("%dD, %dH, %dM", d, h, m)

	return upt
}

//GetDNSServer function gets DNS Server & Search Domain
func GetDNSServer() string {
	r, err := dns.ClientConfigFromFile("/etc/resolv.conf")
	checkErr(err)

	resolver := fmt.Sprintf("%s, %s", r.Servers, r.Search)
	return resolver
}

//checkErr is cheking errors
func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

//RootHandler function
func RootHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatalln(err)
	}

	//Logging: request of client address
	log.Printf("Request %v --%v--> %v%v\n", r.RemoteAddr, r.Method, r.Host, r.URL)

	//Printing: http request header (method, url, protocol)
	fmt.Fprintf(w, "[Request Headers]")
	fmt.Fprintf(w, "\n\t%v %v %v\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "\t%v: %v\n", k, v)
	}

	//Printing: request host & remote address
	fmt.Fprintf(w, "[Client Informations]")
	fmt.Fprintf(w, "\n\tHost: %v", r.Host)
	fmt.Fprintf(w, "\n\tRemoteAddr: %v\n", r.RemoteAddr)

	//Printing: Host Information
	fmt.Fprintf(w, "[Host Informations]")
	fmt.Fprintf(w, "\n\tHostname: %v", GetHostname())
	fmt.Fprintf(w, "\n\tIP: %v", GetOutboundIP())
	fmt.Fprintf(w, "\n\tDNS: %v", GetDNSServer())
	fmt.Fprintf(w, "\n\tUptime: %v", GetUptime())
	fmt.Fprintf(w, "\n\tEnv Message: %v\n", setEnvMessage())
}

//NotFoundHandler for /404
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Oops! 404 Not Found!")
}

//setEnvMessage
func setEnvMessage() string {
	m := os.Getenv("MESSAGE")
	if m != "" {
		return m
	}
	return defaultEnvMessage
}

func main() {
	p := flag.Int("port", 8080, "Set service port")
	flag.Parse()

	log.Println("Start MyWeb Application")
	log.Println("Listen: ", "0.0.0.0:"+strconv.Itoa(*p))

	http.HandleFunc("/", RootHandler)
	http.HandleFunc("/404", NotFoundHandler)
	log.Fatal(
		http.ListenAndServe("0.0.0.0:"+strconv.Itoa(*p), nil),
	)
}
