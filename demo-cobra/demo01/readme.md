- [Goland: cobra简介](https://www.cnblogs.com/sparkdev/p/10856077.html)

`cobra init --pkg-name=demo.io/demo-cobra/demo01`

`cobra add image`
`cobra add container`
`cobra add version`

- 完整命令单
```bash
  509  cobra add version
  510  go run main version
  511  go run main.go version
  512  go run main.go image hello
  513  go run main.go image hello
  514  go run main.go image times -t= 3 world
  515  go run main.go image times -t=3 world
  516  go run main.go image times -t=3 
  517  go run main.go -h
  518  go run main.go network
  519  go run main.go --help
  520  go run main.go --helo
  521  go run main.go --hello
  522  go run main.go
```

>如果使用的go mod, 不需要配置GOPATH变量，或者项目不要在GOPATH里面，否则会报错