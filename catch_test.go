package secretguard

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEnsureSafePanic(t *testing.T) {
	var secret *Secret
	defer func() {
		if r := recover(); r != nil {
			err := secret.Use(func(raw []byte) error {
				t.Fatalf("secret accessible after panic")
				return nil
			})
			require.Error(t, err)
		}
	}()

	func() {
		defer EnsureSafePanic()

		secret = NewSecretFromBytes([]byte("Hello world"))
		panic("test panic")
	}()
}
