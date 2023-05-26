package test

import (
	"errors"
	"os"
	"testing"

	"github.com/inkochetkov/log/pkg/logger"
	"github.com/stretchr/testify/assert"
)

func TestLog(t *testing.T) {

	t.Log("writes to file")

	prodLog := logger.New(logger.ProdLog, "", "")

	prodLog.Info("Test", nil)
	prodLog.Warn("Test", 1)
	prodLog.Error("Test", errors.New("0_o"))
	prodLog.Close()

	file, err := os.ReadFile("info.log")
	assert.NoError(t, err)

	assert.Contains(t, string(file), "0_o")

	t.Log("console output")

	devLog := logger.New(logger.DevLog, "", "")

	devLog.Info("Test2", nil)
	devLog.Warn("Test2", nil)
	devLog.Error("Test2", errors.New("0_o"))

	t.Log("Success test")

}
