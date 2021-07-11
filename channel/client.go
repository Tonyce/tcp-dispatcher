package channel

const DELIMITER byte = '\n'

type Client interface {
	Read() (string, error)
	Write(string) error
	// Read(time.Duration) ([]byte, error)
	// Write([]byte, time.Duration) error
	Close()
	CheckValid() bool
}
