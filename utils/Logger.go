package utils

import (
	"os"

	"github.com/sirupsen/logrus"
)

func CustomLogger() *logrus.Logger {
	logger := logrus.New()

	// 파일 핸들러 생성 (선택 사항)
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logger.Fatal("Error opening log file: ", err)
	}

	// 로그 포맷 설정
	logger.SetFormatter(&logrus.JSONFormatter{})

	// 로그 레벨 설정
	logger.SetLevel(logrus.InfoLevel)

	// 출력 설정 (파일 또는 표준 출력)
	logger.SetOutput(file)

	return logger
}
