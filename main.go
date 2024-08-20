package main

import "github.com/joeCavZero/library-rest-api-in-golang/internal/util/logkit"

func main() {
	logger := logkit.New("main")
	logger.Debug("this is a debug lol")
	logger.Info("this is a info lol")
	logger.Warn("this is a warn lol")
	logger.Error("this is a error lol")
	logger.Debugf("this is a %s lol", "debugf")

}
