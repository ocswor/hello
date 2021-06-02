#本项目是根据个人习惯对[key](https://github.com/rakyll/hey) 进行了改造


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