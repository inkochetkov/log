package log

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLog(t *testing.T) {

	prodLog := New(ProdLog, "", "")

	prodLog.Info("Test", nil)
	prodLog.Warn("Test", 1)
	prodLog.Error("Test", errors.New("0_o"))
	prodLog.Close()

	file, err := os.ReadFile("./log/info.log")
	assert.NoError(t, err)

	assert.Contains(t, string(file), "0_o")

	err = os.RemoveAll("./log")
	assert.NoError(t, err)

	t.Log("console output")

	devLog := New(DevLog, "", "")

	devLog.Info("Test2", nil)
	devLog.Warn("Test2", nil)
	devLog.Error("Test2", errors.New("0_o"))

	t.Log("Success test")

}
