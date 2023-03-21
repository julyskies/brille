package utilities

import (
	"os"
	"runtime"
	"strconv"

	"github.com/julyskies/brille/v2/constants"
)

func GetThreads() int {
	threadsNumberENV := os.Getenv(constants.ENV_THREADS)
	threadsNumber := runtime.NumCPU()
	if threadsNumberENV != "" {
		parsed, parsingError := strconv.Atoi(threadsNumberENV)
		if parsingError == nil {
			threadsNumber = Clamp(parsed, threadsNumber, 1)
		}
	}
	return threadsNumber
}
