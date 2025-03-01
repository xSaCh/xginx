package main

import (
	"fmt"
	"net/http"

	xginx "github.com/xSaCh/xginx/pkg"
)

const PORT = 5000

func main() {
	lb := xginx.NewLoadBalancer()
	fmt.Printf("Running xginx on %d\n", PORT)
	http.ListenAndServe(fmt.Sprintf(":%d", PORT), http.HandlerFunc(lb.Router))
}
