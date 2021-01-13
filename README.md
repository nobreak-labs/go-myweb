go-myweb Web Application
========================

- Default Port: 8080

- Change Service Port
```shell
./myweb -port=8088
```

- Change Message
```
MESSAGE="Hello Myweb" ./myweb
```shell

- Http Response
```html
[Request Headers]
	GET / HTTP/1.1
	User-Agent: [curl/7.64.1]
	Accept: [*/*]
[Client Informations]
	Host: localhost:8080
	RemoteAddr: [::1]:61842
[Host Informations]
	Hostname: Ryan-MBP.local
	IP: 192.168.0.12
	DNS: [1.1.1.1 1.0.0.1], []
	Uptime: 15D, 4H, 19M
	Env Message: Hello world
```

