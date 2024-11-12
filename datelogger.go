package datelogger

import (
	"os"
	"path/filepath"
	"time"
)

// DateLogger 自动按照日期分文件写入日志
type DateLogger struct {
	path        string //存放的目录名
	logFile     *os.File
	logFileName string //当前的日志文件名
}

// NewDateLogger 分配一个按日期存放 文件的日志
func NewDateLogger(strPath string) *DateLogger {
	d := &DateLogger{
		path: strPath,
	}
	return d
}

// Close 关闭底层文件
func (d *DateLogger) Close() error {
	if d.logFile != nil {
		err := d.logFile.Close()
		d.logFile = nil
		return err
	} else {
		return nil
	}
}

// Write 被logrus 调用来写日志，自动分文件
func (d *DateLogger) Write(p []byte) (n int, err error) {
	if err := d.checkLogFile(); err != nil {
		return -1, err
	}
	return d.logFile.Write(p)
}
func (d *DateLogger) checkLogFile() error {
	strPath := filepath.Join(d.path, time.Now().Format("2006-01-02")+".txt")
	if strPath != d.logFileName {
		if d.logFile != nil {
			if err := d.logFile.Close(); err != nil {
				return err
			}
		}
		d.logFileName = strPath
		//确保创建目录
		if err := os.MkdirAll(filepath.Dir(d.logFileName), os.ModePerm); err != nil {
			return err
		}
		flog, err := os.OpenFile(d.logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return err
		}
		d.logFile = flog
	}
	return nil
}
