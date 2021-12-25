package blockposition

import (
	"bufio"
	"go.uber.org/zap"
	"os"
	"path"
	"strconv"
	"strings"
)

type filePosition struct {
	minPos      uint64
	currentPos  uint64
	posFileDir  string
	posFileName string
	logger      *zap.Logger
}

func NewFilePosition(logger *zap.Logger, path, name string, minPos uint64) BlockPositionHolder {
	bp := new(filePosition)
	bp.posFileDir = path
	bp.posFileName = name
	bp.logger = logger
	bp.minPos = minPos
	return bp
}

func (fp *filePosition) Update(n uint64) error {
	for {
		p := strconv.FormatUint(n, 10)
		f, err := os.OpenFile(path.Join(fp.posFileDir, fp.posFileName), os.O_WRONLY, 0757)
		if err != nil {
			return err
		}
		_, err = f.Seek(0, 0)
		if err != nil {
			return err
		}
		w := bufio.NewWriter(f)
		_, err = w.Write([]byte(p))
		if err != nil {
			return err
		}
		err = w.Flush()
		if err != nil {
			return err
		}
		err = f.Close()
		if err != nil {
			return err
		}
		return nil
	}
}

func (fp *filePosition) Exists() bool {
	_, err := os.Open(path.Join(fp.posFileDir, fp.posFileName))
	if err != nil {
		return false
	}
	return true
}

func (fp *filePosition) Create() error {
	if _, err := os.Stat(fp.posFileDir); err != nil {
		err = os.MkdirAll(fp.posFileDir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	fil, err := os.Create(path.Join(fp.posFileDir, fp.posFileName))
	if err != nil {
		return err
	}
	defer fil.Close()
	// new blockposition file created, minimum blockposition is returned
	_, err = fil.Write([]byte(strconv.FormatUint(fp.minPos, 10)))
	if err != nil {
		return err
	}
	return nil
}

func (fp *filePosition) Read() (uint64, error) {
	fil, err := os.Open(path.Join(fp.posFileDir, fp.posFileName))
	if err != nil {
		return 0, err
	}
	scanner := bufio.NewScanner(fil)
	var block string
	for scanner.Scan() {
		block = strings.Trim(scanner.Text(), "\n")
	}
	blockI, err := strconv.ParseUint(block, 0, 64)
	if err != nil {
		return 0, err
	}
	return blockI, nil
}

func (fp *filePosition) Incr() error {
	curr, err := fp.Read()
	if err != nil {
		return err
	}
	return fp.Update(curr+1)
}
