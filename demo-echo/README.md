1. go version: go1.11.4
2. echo version: v4.1.10
3. k8s version: v1.14.6
3.1 error:
# k8s.io/client-go/rest
# https://github.com/kubernetes/client-go/blob/master/INSTALL.md#go-modules
go get k8s.io/client-go@v11.0.0              # replace v11.0.0 with the required version (or use kubernetes-1.x.y tags if desired)
go get k8s.io/api@kubernetes-1.14.0          # replace kubernetes-1.14.0 with the required version
go get k8s.io/apimachinery@kubernetes-1.14.0 # replace kubernetes-1.14.0 with the required version

../../../../../go/pkg/mod/k8s.io/client-go@v11.0.0+incompatible/rest/request.go:598:31: not enough arguments in call to watch.NewStreamWatcher
	have (*versioned.Decoder)
	want (watch.Decoder, watch.Reporter)

go get k8s.io/apimachinery@kubernetes-1.14.1


4. client-go version: v11.0
5. eureka
6. config
7. swagger-ui
8. unit test
9. 