package starter_test

import (
	"github.com/stretchr/testify/assert"
	"io"
	"net/http/httptest"
	starter "tddproject"
	"testing"
)

func TestSayHello(t *testing.T) {
	greeting := starter.SayHello("Mevlut")
	assert.Equal(t, "Hello Mevlut. Welcome!", greeting)
	greeting_two := starter.SayHello("Test Name")
	assert.Equal(t, "Hello Test Name. Welcome!", greeting_two)
}

func TestOddOrEven(t *testing.T) {
	t.Run("Check Non Negative Numbers", func(t *testing.T) { //sub test
		assert.Equal(t, "45 is an odd number", starter.OddOrEven(45))
		assert.Equal(t, "42 is an even number", starter.OddOrEven(42))
		assert.Equal(t, "0 is an even number", starter.OddOrEven(0))
	})
	t.Run("Check Negative Numbers", func(t *testing.T) {
		assert.Equal(t, "-45 is an odd number", starter.OddOrEven(-45))
		assert.Equal(t, "-42 is an even number", starter.OddOrEven(-42))
	})
}

func TestCheckhealth(t *testing.T) {
	t.Run("Check health status", func(t *testing.T) {
		req := httptest.NewRequest("GET", "http://mysite.com/example", nil)
		writer := httptest.NewRecorder()
		starter.Checkhealth(writer, req)
		response := writer.Result()
		body, err := io.ReadAll(response.Body)

		assert.Equal(t, "health check passed", string(body))
		assert.Equal(t, 200, response.StatusCode)
		assert.Equal(t,
			"text/plain; charset=utf-8",
			response.Header.Get("Content-Type"))

		assert.Equal(t, nil, err)
	})
}

func TestRegisterExamp(t *testing.T) {
	t.Run("Register check", func(t *testing.T) {
		assert.Equal(t, "mtest is registered", starter.RegisterExample("Mevlut", "test", "123456"))
		assert.Equal(t, "Name, Surname or password is empty", starter.RegisterExample("", "", ""))
		assert.Equal(t, "1kaya is registered", starter.RegisterExample("1", "Kaya", "123456"))
	})
}
