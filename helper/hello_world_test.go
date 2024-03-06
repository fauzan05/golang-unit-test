package helper

import (
	"fmt"
	"runtime"
	"testing"

	// "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// jika ingin hanya me-running 1 method saja, maka perintahnya : go test -v -run=TestHelloWorld
// jika ingin me-running dari directory root ke directory package : go test -v ./...

func TestMain(m *testing.M) {
	fmt.Println("Dieksekusi Pertama Kali")

	m.Run()

	fmt.Println("Dieksekusi Terakhir Kali")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Cannot run on mac OS")
	}

	result := HelloWorld("Fauzan")
	require.Equal(t, "Hello World Fauzan", result)
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Fauzan")
	if result != "Hello World Fauzan" {
		// unit test gagal
		// t.Fail()
		// t.FailNow()
		// t.Fatal()
		t.Error("seharusnya : Hello World Fauzan")
	}

	fmt.Println("the code has been executed")
}

func TestHelloWorldWithAssert(t *testing.T) {
	result := HelloWorld("Fauzan")
	// assert.Equal(t, "Hello World Fauzan", result) // seperti Fail()
	require.Equal(t, "Hello World Fauzan", result) // seperti FailNow()
	fmt.Println("berhasil dieksekusi")
}

// jika hanya ingin menjalankan salah satu subtest saja, gunakan : go test -run TestSubTest/resultFauzan
func TestSubTest(t *testing.T) {
	t.Run("resultFauzan", func(t *testing.T) {
		result := HelloWorld("Fauzan")
		require.Equal(t, "Hello World Fauzan", result)
		fmt.Println("method resultFauzan berhasil dieksekusi")
	})

	t.Run("resultRudi", func(t *testing.T) {
		result := HelloWorld("Rudi")
		require.Equal(t, "Hello World Rudi", result)
		fmt.Println("method resultRudi berhasil dieksekusi")
	})
}

// unit test konsep table test

func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "HelloWorld(Fauzan)",
			request:  "Fauzan",
			expected: "Hello World Fauzan",
		},
		{
			name:     "HelloWorld(Rudi)",
			request:  "Rudi",
			expected: "Hello World Rudi",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

// benchmark

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("fauzan")
	}
}

func BenchmarkHelloWorldSub(b *testing.B) {
	b.Run("testing-pertama", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			result := HelloWorld("Pertama")
			require.Equal(b, "Hello World Pertama", result)
		}
	})

	b.Run("testing-kedua", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			result := HelloWorld("Kedua")
			require.Equal(b, "Hello World Kedua", result)
		}
	})
}

func BenchmarkHelloWorldTable(b *testing.B) {
	benchmarks := []struct{
		name     string
		request  string
		expected string
	}{
		{
			name: "test-fauzan",
			request: "Fauzan",
			expected: "Hello World Fauzan",
		},
	}
	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				result := HelloWorld(benchmark.request)
				require.Equal(b, benchmark.expected, result)
			}
		})
	}
}
