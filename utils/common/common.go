package common

import (
	//"crypto/md5"
	"crypto/sha1"
	//"encoding/hex"
	"fmt"
	"io"
	//"math"
	//"math/big"
	"math/rand"
	//crand "crypto/rand"
	"net"
	"strconv"
	"strings"
	"time"
)


//获取随机数  数字和文字
func GetRandomBoth(n int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

//sha1加密
func Sha1En(data string) string {
	t := sha1.New()///产生一个散列值得方式
	_,_=io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

func IpStringToInt(ipstring string) int {
	if net.ParseIP(ipstring)==nil {
		return 0
	}
	ipSegs := strings.Split(ipstring, ".")
	var ipInt int = 0
	var pos uint = 24
	for _, ipSeg := range ipSegs {
		tempInt, _ := strconv.Atoi(ipSeg)
		tempInt = tempInt << pos
		ipInt = ipInt | tempInt
		pos -= 8
	}
	return ipInt
}