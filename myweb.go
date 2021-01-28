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

//rootHandler function
func RootHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatalln(err)
	}

	//Logging: request of client address
	log.Printf("Request %v --%v--> %v%v\n", r.RemoteAddr, r.Method, r.Host, r.URL)

	query := r.URL.Query()
	detail := query.Get("detail")

	if detail == "" {
		detail = "0"
	}

	detailNum, err := strconv.Atoi(detail)
	if err != nil {
		detailNum = 0
	}

	//Printing: welcome message with hostname
	fmt.Fprintf(w, welcomeMessage())
	fmt.Fprintf(w, GetHostname())

	if detailNum >= 1 || detail == "header" {
		//Printing: http request header (method, url, protocol)
		fmt.Fprintln(w)
		fmt.Fprintf(w, "[Request Headers]")
		fmt.Fprintf(w, "\n\t%v %v %v\n", r.Method, r.URL, r.Proto)
		for k, v := range r.Header {
			fmt.Fprintf(w, "\t%v: %v\n", k, v)
		}
	}
	if detailNum >= 2 || detail == "client" {
		//Printing: request host & remote address
		fmt.Fprintf(w, "[Client Informations]")
		fmt.Fprintf(w, "\n\tHost: %v", r.Host)
		fmt.Fprintf(w, "\n\tRemoteAddr: %v\n", r.RemoteAddr)
	}
	if detailNum >= 3 || detail == "container" {
		//Printing: container information
		fmt.Fprintf(w, "[Container Informations]")
		fmt.Fprintf(w, "\n\tHostname: %v", GetHostname())
		fmt.Fprintf(w, "\n\tIP: %v", GetOutboundIP())
		fmt.Fprintf(w, "\n\tDNS: %v", GetDNSServer())
		fmt.Fprintf(w, "\n\tUptime: %v", GetUptime())
	}
}

//healthCheckHandler function
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	code := query.Get("code")

	if code == "" {
		code = "200"
	}

	codeNum, err := strconv.Atoi(code)
	if err != nil {
		codeNum = 999
	}

	if codeNum == 0 || codeNum == 200 {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Health Check: OK")
		log.Println("Health Check: OK")
	} else if codeNum >= 100 && codeNum < 199 {
		w.WriteHeader(codeNum)
		fmt.Fprintln(w, "Health Check: Informational")
		log.Println("Health Check: Informational")
	} else if codeNum >= 201 && codeNum < 299 {
		w.WriteHeader(codeNum)
		fmt.Fprintln(w, "Health Check: Success")
		log.Println("Health Check: Success")
	} else if codeNum >= 300 && codeNum < 399 {
		w.WriteHeader(codeNum)
		fmt.Fprintln(w, "Health Check: Redirection")
		log.Println("Health Check: Redirection")
	} else if codeNum >= 400 && codeNum < 499 {
		w.WriteHeader(codeNum)
		fmt.Fprintln(w, "Health Check: Client Error")
		log.Println("Health Check: Client Error")
	} else if codeNum >= 500 && codeNum < 599 {
		w.WriteHeader(codeNum)
		fmt.Fprintln(w, "Health Check: Server Error")
		log.Println("Health Check: Server Error")
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Health Check: No Valid Code")
		log.Println("Health Cehck: No Valid Code")
	}
}

//welcomeMessage
func welcomeMessage() string {
	m := os.Getenv("MESSAGE")
	if m == "" {
		m = "Hello World!"
	}
	return m + "\n"
}

func main() {
	p := flag.Int("port", 8080, "Set service port")
	flag.Parse()

	log.Println("Start MyWeb Application")
	log.Println("Listen: ", "0.0.0.0:"+strconv.Itoa(*p))

	http.HandleFunc("/", RootHandler)
	http.HandleFunc("/health", HealthCheckHandler)
	log.Fatal(
		http.ListenAndServe("0.0.0.0:"+strconv.Itoa(*p), nil),
	)
}
