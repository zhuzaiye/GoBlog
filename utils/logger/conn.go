// File:    conn
// Version: 1.0.0
// Creator: JoeLang
// Date:    2020/8/31 23:10
// DESC:

package logger

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

type connLogger struct {
	sync.Mutex
	innerWriter    io.WriteCloser
	ReconnectOnMsg bool   `json:"reconnect_on_msg"`
	Reconnect      bool   `json:"reconnect"`
	Net            string `json:"net"`
	Addr           string `json:"addr"`
	Level          string `json:"level"`
	LogLevel       int
	// illNetFlag 网络异常标记
	illNetFlag     bool
}

func (c *connLogger) Init(jsonCfg string) error {
	if len(jsonCfg) == 0 {
		return nil
	}
	err := json.Unmarshal([]byte(jsonCfg), c)
	if err != nil {
		return err
	}
	if lvl, ok := LevelMap[c.Level]; ok {
		c.LogLevel = lvl
	}
	if c.innerWriter != nil {
		_ = c.innerWriter.Close()
		c.innerWriter = nil
	}
	return nil
}

func (c *connLogger) LogWrite(when time.Time, msgText interface{}, level int) error {
	if level > c.LogLevel {
		return nil
	}
	msg, ok := msgText.(*logInfo)
	if !ok {
		return nil
	}
	if c.needToConnectOnMsg(){
		err := c.connect()
		if err != nil {
			return err
		}
		//重新连接成功
		c.illNetFlag = false
	}
	//日志频率低的服务，可以每条消息重新连接一次日志中心，可避免长时间连接，重用资源
	//如果是频繁发送日志，切勿开启该功能
	if c.ReconnectOnMsg {
		defer c.innerWriter.Close()
	}

	//当网络异常时，消息发出
	if !c.illNetFlag {
		err := c.println(when, msg)
		//网络异常，通知处理网络的go程自动重连
		if err != nil {
			c.illNetFlag = true
		}
	}
	return nil
}

func (c *connLogger) Destroy() {
	if c.innerWriter != nil {
		_ = c.innerWriter.Close()
	}
}

func (c *connLogger) connect() error {
	if c.innerWriter != nil {
		_ = c.innerWriter.Close()
		c.innerWriter = nil
	}
	addrLst := strings.Split(c.Addr, ";")
	for _, addr := range addrLst {
		conn, err := net.Dial(c.Net, addr)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "net.Dial error:%v\n", err)
			continue
			//return err
		}

		if tcpConn, ok := conn.(*net.TCPConn); ok {
			_ = tcpConn.SetKeepAlive(true)
		}
		c.innerWriter = conn
		return nil
	}
	return fmt.Errorf("hava no valid logs service addr:%v", c.Addr)
}

func (c *connLogger) needToConnectOnMsg() bool {
	if c.Reconnect {
		c.Reconnect = false
		return true
	}
	if c.innerWriter == nil {
		return true
	}
	if c.illNetFlag {
		return true
	}
	return c.ReconnectOnMsg
}

func (c *connLogger) println(when time.Time, msg *logInfo) error {
	c.Lock()
	defer c.Unlock()
	ss, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	_, err = c.innerWriter.Write(append(ss, '\n'))

	//返回err，解决日志系统网络异常后的自动重连
	return err
}

func init() {
	Register(AdapterConn, &connLogger{LogLevel: LevelTrace})
}
