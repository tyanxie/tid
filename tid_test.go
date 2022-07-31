package tid

import (
	"fmt"
	"testing"
	"time"
)

func Test_Generate(t *testing.T) {
	fmt.Println(Generate())

	begin := time.Now()
	for i := 0; i < 0xffff; i++ {
		_ = Generate()
	}
	fmt.Println("time cost: ", time.Now().Sub(begin))
}

func Benchmark_Generate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Generate()
	}
}
