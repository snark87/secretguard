package secretguard

import (
	"github.com/awnumar/memguard"
)

type Secret struct {
	secret *memguard.Enclave
}

func NewSecretFromBytes(raw []byte) *Secret {
	buffer := memguard.NewBufferFromBytes(raw)

	return &Secret{
		secret: buffer.Seal(),
	}
}

func (s *Secret) Use(operation func(raw []byte) error) error {
	defer EnsureSafePanic()

	buffer, err := s.secret.Open()
	if err != nil {
		return err
	}
	defer func() {
		s.secret = buffer.Seal()
	}()

	if err := operation(buffer.Bytes()); err != nil {
		return err
	}

	return err
}
