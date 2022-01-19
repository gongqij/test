package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
)

const (
	logTimeFormatter      = "2006-01-02 15:04:05.000"
	fileNameTimeFormatter = "2006-01-02"
)

type Logger struct {
	logger *logrus.Logger

	filePrefix      string
	fileCreatedTime time.Time
	fileDir         string
	fileSplit       bool
	file            *os.File
}

var stdLogger = newLogger()

func newLogger() *Logger {
	return &Logger{
		logger: logrus.New(),
	}
}
func InitWithPath(dir, prefix string) {
	initLog(dir, prefix, true, logrus.InfoLevel)
}

// Init is same as InitWithPath but with a default dir "./"
func Init(prefix string) {
	initLog("./", prefix, true, logrus.InfoLevel)
}

func initLog(dir, prefix string, split bool, level logrus.Level) {

	stdLogger.logger.SetLevel(level)

	customFormatter := new(logrus.JSONFormatter)
	customFormatter.TimestampFormat = logTimeFormatter
	stdLogger.logger.SetFormatter(customFormatter)

	if err := stdLogger.createFile(dir, prefix, split); err != nil {
		panic(err)
	}
	//主要步骤
	stdLogger.logger.AddHook(&stdoutHook{})
}

func StdLogger() *Logger {
	return stdLogger
}

// stdoutHook will print log to stdout
type stdoutHook struct{}

func (hook *stdoutHook) Fire(entry *logrus.Entry) error {
	// print to stdout
	content, err := entry.String()
	if err != nil {
		return err
	}
	//输出到终端
	fmt.Print(content)
	return nil
}

func (hook *stdoutHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// createFile will make the Logger output to a file and split log file automatically if split is true
func (m *Logger) createFile(dir, prefix string, split bool) error {

	m.fileDir = dir
	m.filePrefix = prefix
	m.fileSplit = split

	if err := m.createFileAndSetOutPut(); err != nil {
		return err
	}

	if m.fileSplit {
		go m.checkSplitFile()
	}

	return nil
}

func (m *Logger) checkSplitFile() {
	t := time.NewTicker(time.Second)
	for range t.C {
		// check a new day
		if time.Now().Format(fileNameTimeFormatter) != m.fileCreatedTime.Format(fileNameTimeFormatter) {
			if err := m.createFileAndSetOutPut(); err != nil {
				m.logger.Errorf("createFileAndSetOutPut err=%s", err.Error())
			}
		}
	}
}

func (m *Logger) createFileAndSetOutPut() error {
	fileName := ""
	if m.fileSplit {
		fileName = fmt.Sprintf("%s.%s.log", path.Join(m.fileDir, m.filePrefix), time.Now().Format(fileNameTimeFormatter))
	} else {
		fileName = fmt.Sprintf("%s.log", path.Join(m.fileDir, m.filePrefix))
	}
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		return err
	}

	m.fileCreatedTime = time.Now()
	m.logger.SetOutput(file)

	// close the old file
	if m.file != nil {
		m.file.Close()
	}

	m.file = file
	return nil
}
