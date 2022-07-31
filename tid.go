// Package tid 生成唯一id工具
// 生成规则：机器ip+进程pid+毫秒时间戳+自增序列
// 注意该包如果初始化失败则会直接panic
package tid

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"net"
	"os"
	"strconv"
	"sync/atomic"
	"time"
)

var (
	ip       string // 机器ip地址进过计算后得出的hash字符串
	pid      string // 进程pid
	ipAndPid string // 机器ip和进程pid拼接后的字符串

	seq uint32 = 0 // 自增序列号：[0,65535(0xffff)]
)

func init() {
	// 获取机器所有ip
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(fmt.Errorf("get interface addrs failed: %w", err))
	}
	if addrs == nil || len(addrs) == 0 {
		panic(errors.New("get interface addrs success but is nil"))
	}
	var addrBuf bytes.Buffer
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ip4 := ipnet.IP.To4(); ip4 != nil {
				addrBuf.Write(ip4)
			} else {
				addrBuf.Write(ipnet.IP)
			}
		}
	}
	addrBytes := addrBuf.Bytes()
	var addrHash uint32 = 0
	for i := 0; i < len(addrBytes); i++ {
		addrHash = uint32(addrBytes[i]) + (addrHash << 6) + (addrHash << 16) - addrHash
	}
	ip = strconv.FormatUint(uint64(addrHash), 16)

	// 进程pid
	pid = strconv.Itoa(os.Getpid())

	// 拼接
	ipAndPid = ip + pid
}

// Generate 生成唯一id
func Generate() string {
	buf := bytes.NewBufferString(ipAndPid)

	// 获取时间戳
	ts := strconv.FormatInt(time.Now().UnixMilli(), 10)
	buf.WriteString(ts)

	// 自增序列
	seqNum := make([]byte, 2)
	currSeq := atomic.AddUint32(&seq, 1)
	if seq == (0xffff + 1) {
		atomic.CompareAndSwapUint32(&seq, 0xffff+1, 0)
	}
	seqNum[0] = byte(currSeq >> 8)
	seqNum[1] = byte(currSeq)
	buf.WriteString(hex.EncodeToString(seqNum))

	return buf.String()
}
