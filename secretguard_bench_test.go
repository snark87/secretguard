package secretguard

import (
	"crypto/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func BenchmarkLongSecret(b *testing.B) {
	sliceSize := 1 * 1024 * 1024 // 1MB
	for range b.N {
		b.StopTimer()
		rawData := make([]byte, sliceSize)
		_, err := rand.Read(rawData)
		require.NoError(b, err)

		b.StartTimer()
		secret := NewSecretFromBytes(rawData)
		secret.Use(func(raw []byte) error {
			return nil
		})
	}
}
