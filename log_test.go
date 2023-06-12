package log

import (
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLog(t *testing.T) {

	prodLog := New(ProdLog, "", "")

	go func() {
		prodLog.Name("Test1")
		prodLog.Info("1")
		time.Sleep(2 * time.Second)
		prodLog.Info("1")
	}()
	go func() {
		prodLog.Warn("2")
		prodLog.Name("Test2")
		prodLog.Error("2", errors.New("0_o"))
	}()

	file, err := os.ReadFile("./log/info.log")
	assert.NoError(t, err)

	assert.Contains(t, string(file), "0_o")

	t.Log("console output")

	devLog := New(DevLog, "", "")

	devLog.Info("Test2", nil)
	devLog.Warn("Test2", nil)
	devLog.Error("Test2", errors.New("0_o"))

	time.Sleep(3 * time.Second)
	prodLog.Close()

	err = os.RemoveAll("./log")
	assert.NoError(t, err)

	t.Log("Success test")

}
