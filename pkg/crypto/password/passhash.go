package password

var _ PassHash = (*passHash)(nil)

type (
	PassHash interface {
		Generate(text []byte) (hash string, err error)
		Verify(text []byte, hash string) (bool, error)
		t()
	}

	passHash struct {
		pepper      string
		memory      uint32
		iterations  uint32
		parallelism uint8
		saltLength  uint32
		keyLength   uint32
	}
)

func New(pepper string, memory, iterations, saltLength, keyLength uint32, parallelism uint8) PassHash {
	return &passHash{
		pepper:      pepper,
		memory:      memory,
		iterations:  iterations,
		parallelism: parallelism,
		saltLength:  saltLength,
		keyLength:   keyLength,
	}
}

func (a *passHash) t() {}
