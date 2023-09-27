package tools

import (
	"encoding/binary"
	"github.com/Yihen-Liu/go-service/log"
	"io"
	"math/rand"

	"net/http"
	"strconv"
	"strings"
	"time"
)

var client = &http.Client{}

func Get(url string, timeout int) ([]byte, error) {
	// url := "http://106.3.133.179:46657/tri_block_info?height=104360"

	client.Timeout = time.Second * time.Duration(timeout)
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func Post(url string, timeout int, payload *strings.Reader) ([]byte, error) {

	client := &http.Client{}
	client.Timeout = time.Second * time.Duration(timeout)
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func SliceContain[Item comparable](items []Item, item Item) bool {
	for i := 0; i < len(items); i++ {
		if items[i] == item {
			return true
		}
	}
	return false
}

func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

func BytesToInt64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}
func Hex2int64(hexStr string) (int64, error) {
	hexStr = strings.TrimPrefix(hexStr, "0x")
	if hexStr == "" {
		return 0, nil
	}
	return strconv.ParseInt(hexStr, 16, 64)

}

// Str2Int64
//
//	@Description: 将字符串转为int64
//	@param str
//	@return int64
func Str2Int64(str string) int64 {

	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		log.Error(err)
	}
	return num
}

// RandSleep 随机休眠(毫秒)
func RandSleep(min, max int) {
	randNum := rand.Intn(max-min) + min
	time.Sleep(time.Duration(randNum) * time.Millisecond)
}

type Int64S []uint64

func (a Int64S) Len() int           { return len(a) }
func (a Int64S) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Int64S) Less(i, j int) bool { return a[i] < a[j] }
