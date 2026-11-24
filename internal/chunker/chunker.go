package chunker

import (
	"fmt"
	"io"
)

const DefaultChunkSize = 4 * 1024 * 1024

func ChunkProcess(r io.Reader, chunkSize int, fn func(*Chunk) error) error {
	if chunkSize <= 0 {
		return fmt.Errorf("invalid chunk size: %d", chunkSize)
	}

	buf := make([]byte, chunkSize)

	for {
		n, err := io.ReadFull(r, buf)

		if n > 0 {
			data := make([]byte, n)
			copy(data, buf[:n])

			chunk := NewChunk()
			chunk.SetChunk(data)

			if err := fn(chunk); err != nil {
				return err
			}
		}

		if err == io.EOF || err == io.ErrUnexpectedEOF {
			return nil
		}

		if err != nil {
			return err
		}
	}
}
