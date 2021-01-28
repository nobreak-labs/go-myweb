go-myweb Web Application
========================

### Default Port
```8080```

### Default Response
```
Hello World!
localhost
```

### Detailed Response
```http://X.X.X.X/?detail=V```

- detail
   	- Numeric
       	- 1
       	- 2
       	- 3
   	- Strings
       	- header
       	- client
       	- container

```
Hello World!
localhost
[Request Headers]
	GET /?detail=3 HTTP/1.1
	User-Agent: [curl/7.64.1]
	Accept: [*/*]
[Client Informations]
	Host: localhost:8080
	RemoteAddr: [::1]:62949
[Container Informations]
	Hostname: Ryan-MBP.local
	IP: 192.168.0.12
	DNS: [1.1.1.1 1.0.0.1], []
	Uptime: 3D, 2H, 58M%
```

### Change Service Port
```shell
./myweb -port=8088
```

### Change Message
```shell
MESSAGE="Hello Myweb" ./myweb
```
```
Hello Myweb!
localhost
```
