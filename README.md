# log

logger

# git

go get https://github.com/inkochetkov/log

# exmple

### writes to file

    prodLog := logger.New(logger.ProdLog, "", "")
    prodLog.Info("Test")
    prodLog.Warn("Test")
    prodLog.Error("Test", errors.New(" 0_0"))
    prodLog.Close()

### console output

    devLog:= logger.New(logger.DevLog, "", "")
    devLog.Info("Test2")
    devLog.Warn("Test2")
    devLog.Error("Test2", errors.New(" 0_0"))
