package secretguard

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSealedSecretCanBeUsed(t *testing.T) {
	sensitiveData := []byte("hello world")
	expectedResult := someOperation(sensitiveData)

	secret := NewSecretFromBytes(sensitiveData)
	var result int
	err := secret.Use(func(raw []byte) error {
		result = someOperation(raw)
		return nil
	})
	require.NoError(t, err)
	assert.Equal(t, expectedResult, result)
}

func TestSealedSecretCanBeUsedMultipleTimes(t *testing.T) {
	sensitiveData := []byte("hello world")
	expectedResult := someOperation(sensitiveData)

	secret := NewSecretFromBytes(sensitiveData)
	for range 10 {
		var result int
		err := secret.Use(func(raw []byte) error {
			result = someOperation(raw)
			return nil
		})
		require.NoError(t, err)
		assert.Equal(t, expectedResult, result)
	}
}

func TestErrorOnUse(t *testing.T) {
	sensitiveData := []byte("hello world")
	secret := NewSecretFromBytes(sensitiveData)
	err := secret.Use(func(raw []byte) error {
		return errors.New("test error")
	})
	assert.Error(t, err)
}

func TestDataIsWipedAfterSealing(t *testing.T) {
	sensitiveData := []byte("hello world")
	_ = NewSecretFromBytes(sensitiveData)
	assertSliceIsWiped(t, sensitiveData)
}

func TestInternalPanicIsSafe(t *testing.T) {
	sensitiveData := []byte("hello world")

	secret := NewSecretFromBytes(sensitiveData)
	otherSecret := NewSecretFromBytes([]byte("other secret"))
	t.Run("when inner func panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Run("other secret should not be accessible", func(t *testing.T) {
					err := otherSecret.Use(func(otherSecretData []byte) error {
						t.Fatalf("secret accessible after panic")
						return nil
					})
					require.Error(t, err)
				})
			}
		}()
		err := secret.Use(func(raw []byte) error {
			panic("something happened")
		})
		require.NoError(t, err)
	})

}

func someOperation(raw []byte) int {
	sum := 0
	for _, b := range raw {
		sum += int(b)
	}

	return sum
}

func assertSliceIsWiped(t *testing.T, slice []byte) {
	t.Helper()
	for _, b := range slice {
		if b != 0 {
			t.Fail()
		}
	}
}
