package session

// The logger for the session.
import "github.com/str20tbl/str20tbl/logger"

var sessionLog logger.MultiLogger

func InitSession(coreLogger logger.MultiLogger) {
	sessionLog = coreLogger.New("section", "session")
}
