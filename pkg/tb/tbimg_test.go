package tb

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestTbimg(t *testing.T) {
	tests := []struct {
		name string
		want map[string]string
	}{
		// TODO: Add test cases.
		{
			name: "Tbimg",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Tbimg()
			if err != nil {
				t.Errorf("Tbimg() error = %v", err)
				return
			}
			fmt.Println(got)
			rand.Seed(time.Now().Unix())
			// fmt.Println(rand.Intn(10) % 2)
			// fmt.Println(rand.Intn(10) % 2)
			// fmt.Println(rand.Intn(10) % 2)
			// fmt.Println(rand.Intn(10) % 2)
			// fmt.Println(rand.Intn(10) % 2)
			// req := new(fasthttp.Request)
			// req.SetRequestURI(got["pic"])
			// rsp := new(fasthttp.Response)
			// if err := fasthttp.Do(req, rsp); err != nil {
			// 	fmt.Println(err)
			// }
			// fmt.Println(rsp.StatusCode(), rsp.Header.String())
		})
	}
}
