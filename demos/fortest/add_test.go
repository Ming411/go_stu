package fortest

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

// 进入对应目录  go test .

// 注意这里必须大写名称
/* func TestAdd(t *testing.T) {
	// ret := add(1, 4)
	ret := add(1, 2)
	if ret != 3 {
		t.Errorf("add(1,2) = %d, want 3", ret)
	}
} */

// 使用 go test -short
/* func TestAdd2(t *testing.T) {
	if testing.Short() {
		// 用于 跳过 那些长时间的测试
		t.Skip("short模式下跳过")
		// 后续代码不执行
	}
	ret := add(1, 2)
	if ret != 3 {
		t.Errorf("add(1,2) = %d, want 3", ret)
	}
} */

// 基于表格的驱动测试
/* func TestAdd3(t *testing.T) {
	var dataset = []struct {
		a   int
		b   int
		out int
	}{
		{1, 2, 3},
		{4, 5, 6},
	}

	for _, value := range dataset {
		re := add(value.a, value.b)
		if re != value.out {
			t.Errorf("add(%d,%d) = %d, want %d", value.a, value.b, re, value.out)
		}
	}
} */

/* 性能测试 */
const numbers = 10000

func BenchmarkStringSprintf(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var str string
		for j := 0; j < numbers; j++ {
			str = fmt.Sprintf("%s%d", str, j)
		}
	}
	b.StopTimer()
}

// 通过对比测试  builder  性能更好
func BenchmarkStringBuilder(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var builder strings.Builder
		for j := 0; j < numbers; j++ {
			builder.WriteString(strconv.Itoa(j))
		}
		_ = builder.String()
	}
	b.StopTimer()
}
