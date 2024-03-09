## backend

#### Curl

```shell
curl -vvv -X POST http://localhost:8080/upload \
  -F "upload[]=@$(pwd)/data/TestDataset/CAMO/Imgs/camourflage_00018.jpg" \
  -F "upload[]=@$(pwd)/data/TestDataset/CAMO/GT/camourflage_00018.png" \
  -H "Content-Type: multipart/form-data"
```

```curl
Note: Unnecessary use of -X or --request, POST is already inferred.
* Host localhost:8080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
*   Trying [::1]:8080...
* Connected to localhost (::1) port 8080
> POST /upload HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/8.6.0
> Accept: */*
> Content-Length: 193079
> Content-Type: multipart/form-data; boundary=------------------------PQmMsrU5zYdsyJFx31aIbg
>
* We are completely uploaded and fine
< HTTP/1.1 200 OK
< Content-Type: text/plain; charset=utf-8
< Date: Sat, 09 Mar 2024 03:38:41 GMT
< Content-Length: 17
<
* Connection #0 to host localhost left intact
2 files uploaded!%
```