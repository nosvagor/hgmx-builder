package include

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var (
	fileHashes      = make(map[string]string)
	fileHashesMutex sync.RWMutex
)

func getFileHash(filePath string) string {
	fileHashesMutex.RLock()
	hash, exists := fileHashes[filePath]
	fileHashesMutex.RUnlock()

	if exists {
		return hash
	}

	fullPath := filepath.Join("views/static", filePath)
	file, err := os.Open(fullPath)
	if err != nil {
		panic(err)
	} else {
		defer file.Close()
		hashMD5 := md5.New()
		if _, err := io.Copy(hashMD5, file); err != nil {
			panic(err)
		} else {
			hash = fmt.Sprintf("?v=%x", hashMD5.Sum(nil))
		}
	}

	fileHashesMutex.Lock()
	fileHashes[filePath] = hash
	fileHashesMutex.Unlock()
	return hash
}

func Bust() string {
	return fmt.Sprintf("?v=%d", time.Now().UnixNano())
}
