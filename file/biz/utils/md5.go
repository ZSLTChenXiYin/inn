package utils

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
	"sync"
)

func calculateChunkMD5(file *os.File, offset int64, size int64, wait_group *sync.WaitGroup, result_channel chan<- []byte) {
	defer wait_group.Done()

	hasher := md5.New()
	buffer := make([]byte, 64*1024) // 64KB 缓冲区

	file.Seek(offset, io.SeekStart)
	for remaining := size; remaining > 0; {
		number := int64(len(buffer))
		if remaining < number {
			number = remaining
		}

		read_number, err := file.Read(buffer[:number])
		if err != nil && err != io.EOF {
			return
		}

		hasher.Write(buffer[:read_number])
		remaining -= int64(read_number)
	}

	result_channel <- hasher.Sum(nil)
}

func GenerateFileMD5(file *os.File, file_size int64) (string, error) {
	wait_group := &sync.WaitGroup{}

	chunk_size := int64(64 * 1024 * 1024) // 64MB

	chunk_number := file_size / chunk_size
	if file_size%chunk_size != 0 {
		chunk_number++
	}

	result_channel := make(chan []byte, chunk_number)

	for index := int64(0); index < chunk_number; index++ {
		offset := index * chunk_size
		size := chunk_size
		if index == chunk_number-1 {
			size = file_size - offset
		}

		wait_group.Add(1)
		go calculateChunkMD5(file, offset, size, wait_group, result_channel)
	}

	wait_group.Wait()
	close(result_channel)

	hasher := md5.New()
	for result := range result_channel {
		hasher.Write(result)
	}

	md5 := hasher.Sum(nil)

	return hex.EncodeToString(md5), nil
}
