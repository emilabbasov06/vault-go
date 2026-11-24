package chunker

import "fmt"

type Chunk struct {
	Data []byte
}

func NewChunk() *Chunk {
	return &Chunk{}
}

func (c *Chunk) SetChunk(data []byte) error {
	if data == nil {
		return fmt.Errorf("chunk data is empty")
	}

	c.Data = data
	return nil
}
