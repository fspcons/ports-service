package file

import (
	"bufio"
	"context"
	"github.com/fspcons/ports-service/src/domain"
	"go.uber.org/zap"
	"io"
	"os"
	"sync"
)

type fileChecker struct {
	portFilePath string
	logger       *zap.Logger
}

// CheckOnFile if the port record exists on the json file
func (ref *fileChecker) CheckOnFile(_ context.Context, _ *domain.Port) error {
	f, err := os.Open(ref.portFilePath)
	if err != nil {
		ref.logger.Error("failed to open ports file", zap.Error(err))
		return err
	}
	defer func() {
		if err := f.Close(); err != nil {
			ref.logger.Warn("failed to close ports file", zap.Error(err))
		}
	}()

	//sync pools to reuse the memory and decrease the pressure on the GC
	linesPool := sync.Pool{New: func() interface{} {
		lines := make([]byte, 100*1024) //chunk size
		return lines
	}}

	r := bufio.NewReader(f)
	for {
		buf := linesPool.Get().([]byte) //reading the file by chunks
		n, err := r.Read(buf)
		buf = buf[:n]
		_ = buf //avoid lint errors
		if n == 0 {
			if err == io.EOF {
				break
			}
			if err != nil {
				ref.logger.Error("failed to read a file chunk into the buffer", zap.Error(err))
				break
			}
			return err
		}
	}

	//TODO From here forward I'd most likely run a Regex check or a similar method, trying to find the port ID as the KEY
	// among the json file records. I believe this would perform better than trying to unmarshall the json records from the file.
	// In case I find it I'd check to see if the user provided data matches and would return a NIL error.
	// Otherwise, I'd return some error saying that the Port record provided is invalid or something of the sorts.
	// Additionally if I could assume the keys on the file are alphabetical I could use a binary search method instead of a linear one,
	// and I could also suggest a file partitioning system based on the keys so I could have smaller files and a smarter search mechanic in that case.

	return nil
}
