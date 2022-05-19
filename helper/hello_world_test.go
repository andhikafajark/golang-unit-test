package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("Before Unit Test")

	m.Run()

	fmt.Println("After Unit Test")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Master")

	if result != "Hello Master" {
		t.FailNow()
	}

	fmt.Println("TestHelloWorld Done")
}

func TestHelloWorldError(t *testing.T) {
	result := HelloWorld("Master")

	if result != "Hello Master" {
		t.Error("Result must be 'Hello Master'")
	}

	fmt.Println("TestHelloWorldError Done")
}

func TestHelloWorldFatal(t *testing.T) {
	result := HelloWorld("Master")

	if result != "Hello Master" {
		t.Fatal("Result must be 'Hello Master'")
	}

	fmt.Println("TestHelloWorldFatal() Done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Master")

	assert.Equal(t, "Hello Master", result, "Result must be 'Hello Master'")

	fmt.Println("TestHelloWorldAssert Done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Master")

	require.Equal(t, "Hello Master", result, "Result must be 'Hello Master'")

	fmt.Println("TestHelloWorldAssert Done")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Can not run on MacOS")
	}

	result := HelloWorld("Master")

	require.Equal(t, "Hello Master", result, "Result must be 'Hello Master'")
}

func TestSubTest(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "HelloWorld(Master)",
			request:  "Master",
			expected: "Hello Master",
		},
		{
			name:     "HelloWorld(Zero)",
			request:  "Zero",
			expected: "Hello Zero",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)

			require.Equal(t, test.expected, result, "Result must be '"+test.expected+"'")
		})
	}
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Master")
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Master", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Master")
		}
	})
	b.Run("Zero", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Zero")
		}
	})
}
