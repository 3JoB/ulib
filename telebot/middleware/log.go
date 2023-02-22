package middleware

import (
	"path"

	tele "github.com/3JoB/telebot"
	"github.com/3JoB/unsafeConvert"
	"github.com/goccy/go-json"
	"gopkg.in/natefinch/lumberjack.v2"
	"github.com/sirupsen/logrus"
)

var (
	logger *logrus.Logger
)

const (
	LOG_SUFFIX = ".log"
	LOG_SIZE   = 60
	LOG_BACKUP = 10
	LOG_DATE   = 7
)

func setOutPut(log *logrus.Logger, log_file_path string) {
	log.SetOutput(&lumberjack.Logger{
		Filename:   log_file_path,
		MaxSize:    LOG_SIZE,
		MaxBackups: LOG_BACKUP,
		MaxAge:     LOG_DATE,
		Compress:   true,
	})
}

func initLogger(l *LogSettings) {
	if l.Path == "" {
		l.Path = "./log/"
	}
	if l.FileName == "" {
		l.FileName = "telebot"
	}
	log_file_path := path.Join(l.Path, l.FileName+LOG_SUFFIX)
	logger = logrus.New()
	setOutPut(logger, log_file_path)
	logger.SetLevel(logrus.InfoLevel)
	logger.SetNoLock()
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
}

type LogSettings struct {
	Path     string
	FileName string
}

func Logger(l *LogSettings) tele.MiddlewareFunc {
	initLogger(l)
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			data, _ := json.MarshalIndent(c.Update(), "", "  ")
			logger.Println(unsafeConvert.String(data))
			return next(c)
		}
	}
}
