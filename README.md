# 本项目是根据个人习惯对[hey](https://github.com/rakyll/hey) 进行了改造

```
Usage: hey [options...] <url>

Options:
  -n  Number of requests to run. Default is 200.
  -c  Number of workers to run concurrently. Total number of requests cannot
      be smaller than the concurrency level. Default is 50.
  -q  Rate limit, in queries per second (QPS) per worker. Default is no rate limit.
  -z  Duration of application to send requests. When duration is reached,
      application stops and exits. If duration is specified, n is ignored.
      Examples: -z 10s -z 3m.
  -o  Output type. If none provided, a summary is printed.
      "csv" is the only supported alternative. Dumps the response
      metrics in comma-separated values format.

  -m  HTTP method, one of GET, POST, PUT, DELETE, HEAD, OPTIONS.
  -H  Custom HTTP header. You can specify as many as needed by repeating the flag.
      For example, -H "Accept: text/html" -H "Content-Type: application/xml" .
  -t  Timeout for each request in seconds. Default is 20, use 0 for infinite.
  -A  HTTP Accept header.
  -d  HTTP request body.
  -D  HTTP request body from file. For example, /home/user/file.txt or ./file.txt.
  -T  Content-type, defaults to "text/html".
  -a  Basic authentication, username:password.
  -x  HTTP Proxy address as host:port.
  -h2 Enable HTTP/2.

  -host	HTTP Host header.

  -disable-compression  Disable compression.
  -disable-keepalive    Disable keep-alive, prevents re-use of TCP
                        connections between different HTTP requests.
  -disable-redirects    Disable following of HTTP redirects
  -cpus                 Number of used cpu cores.
                        (default for current machine is 8 cores)
```
* 增加了支持curl文件 请求
  日常使用PostMan很多,但是PostMan不支持延迟,所以可以复制PostMan的curl文件,直接使用hello进行压测.
  
```
go run hello.go -n 10 -c 1 -curlfile 'D:\goproject\hello\examples\test.chrome.curl.txt'
```

* 增加了debug 模式
  主要是方便确认这次请求返回结果是否和预期一直. debug 只会请求一次 (-n -c 配置不起作用)

```
go run hello.go -curlfile 'D:\goproject\hello\examples\test.chrome.curl.txt' -debug
```