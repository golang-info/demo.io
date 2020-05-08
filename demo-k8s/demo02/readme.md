../../../../pkg/mod/k8s.io/client-go@v11.0.0+incompatible/rest/request.go:598:31: not enough arguments in call to watch.NewStreamWatcher
        have (*versioned.Decoder)
        want (watch.Decoder, watch.Reporter)

https://www.henryxieblogs.com/2019/04/errorrequestgo59831-not-enough.html

go get -u k8s.io/client-go@v11.0.0+incompatible

k8s.io/client-go v0.0.0-20190425172711-65184652c889
go mod tidy

go get k8s.io/client-go@v0.0.0-20190425172711-65184652c889
go get k8s.io/metrics@kubernetes-1.14.1

在main.go所在的路径运行，不会报错了
