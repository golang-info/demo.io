package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

/*
liuyanchao@ycliu912-mac:~$ curl -v http://localhost:8080/welcome
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 8080 (#0)
> GET /welcome HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: *\/*
>
< HTTP/1.1 200 OK
< Msgid:
< Date: Fri, 06 Mar 2020 02:59:53 GMT
< Content-Length: 12
< Content-Type: text/plain; charset=utf-8
<
* Connection #0 to host localhost left intact
Hello, world* Closing connection 0
*/


//HelloWorld hello world handler
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	msgID := " "
	if m := r.Context().Value("msgId"); m != nil {
		if value, ok := m.(string); ok {
			msgID = value
		}
	}
	w.Header().Add("msgId", msgID)
	w.Write([]byte("Hello, world"))
}

func injectMsgID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		msgID := uuid.New().String()
		fmt.Println("msgID ", msgID)
		ctx := context.WithValue(r.Context(), "msgID", msgID)
		fmt.Println("ctx ", ctx)
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}

func main() {
	helloWorldHandler := http.HandlerFunc(HelloWorld)
	http.Handle("/welcome", injectMsgID(helloWorldHandler))
	http.ListenAndServe(":8080", nil)
}
