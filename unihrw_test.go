package unihrw

import (
	"fmt"
	"github.com/spaolacci/murmur3"
	"math"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

type (
	mock struct {
		data []byte
	}
)

const benchmarkKey = "This is some key for benchmark"

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (m mock) Raw() []byte {
	return m.data
}

func (m mock) Data() []byte {
	return m.data
}

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func newMock(s string) mock {
	return mock{data: []byte(s)}
}

func TestDistributionMurMur32(t *testing.T) {
	var (
		i        uint64
		err      error
		size     = uint64(10)
		keys     = uint64(1000000) // 1 million
		accuracy = 0.03            // for stddev

		h = murmur3.New32() // in this test we use mmh3-32bit
	)

	t.Run("test rawers objects", func(t *testing.T) {
		var (
			nodes  = make([]mock, 0, size)
			work   = make([]mock, size)
			counts = make(map[string]uint64)
		)

		for i = 0; i < size; i++ {
			nodes = append(nodes, newMock(RandStringRunes(10)))
		}

		t1 := time.Now()
		for i = 0; i < keys; i++ {
			copy(work, nodes)
			err = HrwSort32(work, []byte(strconv.FormatUint(i, 10)), h)
			if err != nil {
				t.Errorf("hrw sort finished with error")
			}
			counts[string(work[0].Data())]++
		}
		fmt.Println(time.Since(t1))

		var chi2 float64
		mean := float64(keys) / float64(len(nodes))
		delta := mean * accuracy
		for _, count := range counts {
			d := mean - float64(count)
			chi2 += math.Pow(float64(count)-mean, 2) / mean
			if d > delta || (0-d) > delta {
				t.Errorf(
					"out of stddev range: expected %.0f +/- %.2f, got %.2f",
					mean, delta, float64(count),
				)
			}
		}
		t.Log("chi2 = ", chi2)
	})

	t.Run("test strings", func(t *testing.T) {
		var (
			nodes  = make([]string, 0, size)
			work   = make([]string, size)
			counts = make(map[string]uint64)
		)

		for i = 0; i < size; i++ {
			nodes = append(nodes, RandStringRunes(10))
		}

		for i = 0; i < keys; i++ {
			copy(work, nodes)
			err = HrwSort32(work, []byte(strconv.FormatUint(i, 10)), h)
			if err != nil {
				t.Errorf("hrw sort finished with error")
			}
			counts[string(work[0])]++
		}

		var chi2 float64
		mean := float64(keys) / float64(len(nodes))
		delta := mean * accuracy
		for _, count := range counts {
			d := mean - float64(count)
			chi2 += math.Pow(float64(count)-mean, 2) / mean
			if d > delta || (0-d) > delta {
				t.Errorf(
					"out of stddev range: expected %.0f +/- %.2f, got %.2f",
					mean, delta, float64(count),
				)
			}
		}
		t.Log("chi2 = ", chi2)
	})

	t.Run("test bytes", func(t *testing.T) {
		var (
			nodes  = make([][]byte, 0, size)
			work   = make([][]byte, size)
			counts = make(map[string]uint64)
		)

		for i = 0; i < size; i++ {
			nodes = append(nodes, []byte(RandStringRunes(10)))
		}

		for i = 0; i < keys; i++ {
			copy(work, nodes)
			err = HrwSort32(work, []byte(strconv.FormatUint(i, 10)), h)
			if err != nil {
				t.Errorf("hrw sort finished with error")
			}
			counts[string(work[0])]++
		}

		var chi2 float64
		mean := float64(keys) / float64(len(nodes))
		delta := mean * accuracy
		for _, count := range counts {
			d := mean - float64(count)
			chi2 += math.Pow(float64(count)-mean, 2) / mean
			if d > delta || (0-d) > delta {
				t.Errorf(
					"out of stddev range: expected %.0f +/- %.2f, got %.2f",
					mean, delta, float64(count),
				)
			}
		}
		t.Log("chi2 = ", chi2)
	})

	t.Run("test ints", func(t *testing.T) {
		var (
			nodes  = make([]int, 0, size)
			work   = make([]int, size)
			counts = make(map[int]uint64)
		)

		for i = 0; i < size; i++ {
			nodes = append(nodes, int(i))
		}

		for i = 0; i < keys; i++ {
			copy(work, nodes)
			err = HrwSort32(work, []byte(strconv.FormatUint(i, 10)), h)
			if err != nil {
				t.Errorf("hrw sort finished with error")
			}
			counts[work[0]]++
		}

		var chi2 float64
		mean := float64(keys) / float64(len(nodes))
		delta := mean * accuracy
		for _, count := range counts {
			d := mean - float64(count)
			chi2 += math.Pow(float64(count)-mean, 2) / mean
			if d > delta || (0-d) > delta {
				t.Errorf(
					"out of stddev range: expected %.0f +/- %.2f, got %.2f",
					mean, delta, float64(count),
				)
			}
		}
		t.Log("chi2 = ", chi2)
	})
}

func TestDistributionMurMur64(t *testing.T) {
	var (
		i        uint64
		err      error
		size     = uint64(10)
		keys     = uint64(1000000) // 1 million
		accuracy = 0.03            // for stddev

		h = murmur3.New64() // in this test we use mmh3-32bit
	)

	t.Run("test rawers objects", func(t *testing.T) {
		var (
			nodes  = make([]mock, 0, size)
			work   = make([]mock, size)
			counts = make(map[string]uint64)
		)

		for i = 0; i < size; i++ {
			nodes = append(nodes, newMock(RandStringRunes(10)))
		}

		t1 := time.Now()
		for i = 0; i < keys; i++ {
			copy(work, nodes)
			err = HrwSort64(work, []byte(strconv.FormatUint(i, 10)), h)
			if err != nil {
				t.Errorf("hrw sort finished with error")
			}
			counts[string(work[0].Data())]++
		}
		fmt.Println(time.Since(t1))

		var chi2 float64
		mean := float64(keys) / float64(len(nodes))
		delta := mean * accuracy
		for _, count := range counts {
			d := mean - float64(count)
			chi2 += math.Pow(float64(count)-mean, 2) / mean
			if d > delta || (0-d) > delta {
				t.Errorf(
					"out of stddev range: expected %.0f +/- %.2f, got %.2f",
					mean, delta, float64(count),
				)
			}
		}
		t.Log("chi2 = ", chi2)
	})

	t.Run("test strings", func(t *testing.T) {
		var (
			nodes  = make([]string, 0, size)
			work   = make([]string, size)
			counts = make(map[string]uint64)
		)

		for i = 0; i < size; i++ {
			nodes = append(nodes, RandStringRunes(10))
		}

		for i = 0; i < keys; i++ {
			copy(work, nodes)
			err = HrwSort64(work, []byte(strconv.FormatUint(i, 10)), h)
			if err != nil {
				t.Errorf("hrw sort finished with error")
			}
			counts[string(work[0])]++
		}

		var chi2 float64
		mean := float64(keys) / float64(len(nodes))
		delta := mean * accuracy
		for _, count := range counts {
			d := mean - float64(count)
			chi2 += math.Pow(float64(count)-mean, 2) / mean
			if d > delta || (0-d) > delta {
				t.Errorf(
					"out of stddev range: expected %.0f +/- %.2f, got %.2f",
					mean, delta, float64(count),
				)
			}
		}
		t.Log("chi2 = ", chi2)
	})

	t.Run("test bytes", func(t *testing.T) {
		var (
			nodes  = make([][]byte, 0, size)
			work   = make([][]byte, size)
			counts = make(map[string]uint64)
		)

		for i = 0; i < size; i++ {
			nodes = append(nodes, []byte(RandStringRunes(10)))
		}

		for i = 0; i < keys; i++ {
			copy(work, nodes)
			err = HrwSort64(work, []byte(strconv.FormatUint(i, 10)), h)
			if err != nil {
				t.Errorf("hrw sort finished with error")
			}
			counts[string(work[0])]++
		}

		var chi2 float64
		mean := float64(keys) / float64(len(nodes))
		delta := mean * accuracy
		for _, count := range counts {
			d := mean - float64(count)
			chi2 += math.Pow(float64(count)-mean, 2) / mean
			if d > delta || (0-d) > delta {
				t.Errorf(
					"out of stddev range: expected %.0f +/- %.2f, got %.2f",
					mean, delta, float64(count),
				)
			}
		}
		t.Log("chi2 = ", chi2)
	})

	t.Run("test ints", func(t *testing.T) {
		var (
			nodes  = make([]int, 0, size)
			work   = make([]int, size)
			counts = make(map[int]uint64)
		)

		for i = 0; i < size; i++ {
			nodes = append(nodes, int(i))
		}

		for i = 0; i < keys; i++ {
			copy(work, nodes)
			err = HrwSort64(work, []byte(strconv.FormatUint(i, 10)), h)
			if err != nil {
				t.Errorf("hrw sort finished with error")
			}
			counts[work[0]]++
		}

		var chi2 float64
		mean := float64(keys) / float64(len(nodes))
		delta := mean * accuracy
		for _, count := range counts {
			d := mean - float64(count)
			chi2 += math.Pow(float64(count)-mean, 2) / mean
			if d > delta || (0-d) > delta {
				t.Errorf(
					"out of stddev range: expected %.0f +/- %.2f, got %.2f",
					mean, delta, float64(count),
				)
			}
		}
		t.Log("chi2 = ", chi2)
	})
}

// benchmarks

func benchmarkMurMur32Objects(b *testing.B, n int, obj []byte) {
	servers := make([]mock, n)
	for i := uint64(0); i < uint64(len(servers)); i++ {
		servers[i] = mock{
			data: []byte("localhost:" + strconv.FormatUint(60000-i, 10)),
		}
	}

	h := murmur3.New32()

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		if err := HrwSort32(servers, obj, h); err != nil {
			b.Fatal(err)
		}
	}
}

func benchmarkMurMur32Strings(b *testing.B, n int, obj []byte) {
	servers := make([]string, n)
	for i := uint64(0); i < uint64(len(servers)); i++ {
		servers[i] = "localhost:" + strconv.FormatUint(60000-i, 10)
	}

	h := murmur3.New32()

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		if err := HrwSort32(servers, obj, h); err != nil {
			b.Fatal(err)
		}
	}
}

func benchmarkMurMur32Bytes(b *testing.B, n int, obj []byte) {
	servers := make([][]byte, n)
	for i := uint64(0); i < uint64(len(servers)); i++ {
		servers[i] = []byte("localhost:" + strconv.FormatUint(60000-i, 10))
	}

	h := murmur3.New32()

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		if err := HrwSort32(servers, obj, h); err != nil {
			b.Fatal(err)
		}
	}
}

func benchmarkMurMur32Ints(b *testing.B, n int, obj []byte) {
	servers := make([]int, n)
	for i := uint64(0); i < uint64(len(servers)); i++ {
		servers[i] = int(i)
	}

	h := murmur3.New32()

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		if err := HrwSort32(servers, obj, h); err != nil {
			b.Fatal(err)
		}
	}
}

func benchmarkMurMur64Objects(b *testing.B, n int, obj []byte) {
	servers := make([]mock, n)
	for i := uint64(0); i < uint64(len(servers)); i++ {
		servers[i] = mock{
			data: []byte("localhost:" + strconv.FormatUint(60000-i, 10)),
		}
	}

	h := murmur3.New32()

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		if err := HrwSort32(servers, obj, h); err != nil {
			b.Fatal(err)
		}
	}
}

func benchmarkMurMur64Strings(b *testing.B, n int, obj []byte) {
	servers := make([]string, n)
	for i := uint64(0); i < uint64(len(servers)); i++ {
		servers[i] = "localhost:" + strconv.FormatUint(60000-i, 10)
	}

	h := murmur3.New64()

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		if err := HrwSort64(servers, obj, h); err != nil {
			b.Fatal(err)
		}
	}
}

func benchmarkMurMur64Bytes(b *testing.B, n int, obj []byte) {
	servers := make([][]byte, n)
	for i := uint64(0); i < uint64(len(servers)); i++ {
		servers[i] = []byte("localhost:" + strconv.FormatUint(60000-i, 10))
	}

	h := murmur3.New64()

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		if err := HrwSort64(servers, obj, h); err != nil {
			b.Fatal(err)
		}
	}
}

func benchmarkMurMur64Ints(b *testing.B, n int, obj []byte) {
	servers := make([]int, n)
	for i := uint64(0); i < uint64(len(servers)); i++ {
		servers[i] = int(i)
	}

	h := murmur3.New64()

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		if err := HrwSort64(servers, obj, h); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMurMur32Objects(b *testing.B) {
	b.Run("10", func(b *testing.B) {
		benchmarkMurMur32Objects(b, 10, []byte(benchmarkKey))
	})
	b.Run("100", func(b *testing.B) {
		benchmarkMurMur32Objects(b, 10, []byte(benchmarkKey))
	})
	b.Run("1000", func(b *testing.B) {
		benchmarkMurMur32Objects(b, 10, []byte(benchmarkKey))
	})
}

func BenchmarkMurMur32Strings(b *testing.B) {
	b.Run("10", func(b *testing.B) {
		benchmarkMurMur32Strings(b, 10, []byte(benchmarkKey))
	})
	b.Run("100", func(b *testing.B) {
		benchmarkMurMur32Strings(b, 100, []byte(benchmarkKey))
	})
	b.Run("1000", func(b *testing.B) {
		benchmarkMurMur32Strings(b, 100, []byte(benchmarkKey))
	})
}

func BenchmarkMurMur32Bytes(b *testing.B) {
	b.Run("10", func(b *testing.B) {
		benchmarkMurMur32Bytes(b, 10, []byte(benchmarkKey))
	})
	b.Run("100", func(b *testing.B) {
		benchmarkMurMur32Bytes(b, 100, []byte(benchmarkKey))
	})
	b.Run("1000", func(b *testing.B) {
		benchmarkMurMur32Bytes(b, 100, []byte(benchmarkKey))
	})
}

func BenchmarkMurMur32Ints(b *testing.B) {
	b.Run("10", func(b *testing.B) {
		benchmarkMurMur32Ints(b, 10, []byte(benchmarkKey))
	})
	b.Run("100", func(b *testing.B) {
		benchmarkMurMur32Ints(b, 100, []byte(benchmarkKey))
	})
	b.Run("1000", func(b *testing.B) {
		benchmarkMurMur32Ints(b, 100, []byte(benchmarkKey))
	})
}

func BenchmarkMurMur64Objects(b *testing.B) {
	b.Run("10", func(b *testing.B) {
		benchmarkMurMur64Objects(b, 10, []byte(benchmarkKey))
	})
	b.Run("100", func(b *testing.B) {
		benchmarkMurMur64Objects(b, 100, []byte(benchmarkKey))
	})
	b.Run("1000", func(b *testing.B) {
		benchmarkMurMur64Objects(b, 100, []byte(benchmarkKey))
	})
}

func BenchmarkMurMur64Strings(b *testing.B) {
	b.Run("10", func(b *testing.B) {
		benchmarkMurMur64Strings(b, 10, []byte(benchmarkKey))
	})
	b.Run("100", func(b *testing.B) {
		benchmarkMurMur64Strings(b, 100, []byte(benchmarkKey))
	})
	b.Run("1000", func(b *testing.B) {
		benchmarkMurMur64Strings(b, 100, []byte(benchmarkKey))
	})
}

func BenchmarkMurMur64Bytes(b *testing.B) {
	b.Run("10", func(b *testing.B) {
		benchmarkMurMur64Bytes(b, 10, []byte(benchmarkKey))
	})
	b.Run("100", func(b *testing.B) {
		benchmarkMurMur64Bytes(b, 100, []byte(benchmarkKey))
	})
	b.Run("1000", func(b *testing.B) {
		benchmarkMurMur64Bytes(b, 100, []byte(benchmarkKey))
	})
}

func BenchmarkMurMur64Ints(b *testing.B) {
	b.Run("10", func(b *testing.B) {
		benchmarkMurMur64Ints(b, 10, []byte(benchmarkKey))
	})
	b.Run("100", func(b *testing.B) {
		benchmarkMurMur64Ints(b, 100, []byte(benchmarkKey))
	})
	b.Run("1000", func(b *testing.B) {
		benchmarkMurMur64Ints(b, 100, []byte(benchmarkKey))
	})
}
