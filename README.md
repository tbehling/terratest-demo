Clone this repo:

```shell
$ git clone git@github.com:tbehling/terratest-demo --recursive
```

Install mage:

```
$ cd mage
$ go run bootstrap.go install
```

Run mage:

```
$ $(go env GOPATH)/bin/mage
Targets:
  build            
  terratest        
  terratestHttp

$ $(go env GOPATH)/bin/mage terratestHttp
TestHttp 2020-06-28T15:58:45-04:00 logger.go:66: Running 'docker run' on image 'nginx', returning stdout
TestHttp 2020-06-28T15:58:45-04:00 logger.go:66: Running command docker with args [run --detach --rm -P nginx]
TestHttp 2020-06-28T15:58:46-04:00 logger.go:66: a0e400db498bbd5211df42754eb182b1f7b9f1a140d600d9fb64a18728e43b38
TestHttp 2020-06-28T15:58:47-04:00 retry.go:72: HTTP GET to URL http://localhost:32782
TestHttp 2020-06-28T15:58:47-04:00 http_helper.go:32: Making an HTTP GET call to URL http://localhost:32782
TestHttp 2020-06-28T15:58:47-04:00 logger.go:66: Running 'docker stop' on containers '[a0e400db498bbd5211df42754eb182b1f7b9f1a140d600d9fb64a18728e43b38]'
TestHttp 2020-06-28T15:58:47-04:00 logger.go:66: Running command docker with args [stop a0e400db498bbd5211df42754eb182b1f7b9f1a140d600d9fb64a18728e43b38]
TestHttp 2020-06-28T15:58:48-04:00 logger.go:66: a0e400db498bbd5211df42754eb182b1f7b9f1a140d600d9fb64a18728e43b38
PASS
```