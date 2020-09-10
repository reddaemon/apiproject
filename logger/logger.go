package logger

import (
	"github.com/reddaemon/apiproject/config"
	"go.uber.org/zap"
)

// GetLogger function which get log
func GetLogger(cfg *config.Config) (*zap.Logger, error) {
	var err error
	var l *zap.Logger
	if !cfg.Debug {
		l = zap.NewNop()
		return l, nil
	}

	switch cfg.Environment {
	case "production":
		l, err = zap.NewProduction()
		defer func() {

			err := l.Sync()
			if err != nil {
				l.Fatal(err.Error())
			}
		}()
		if err != nil {
			return nil, err
		}
	case "dev":
		l, err = zap.NewDevelopment()
		defer func() {

			err := l.Sync()
			if err != nil {
				l.Fatal(err.Error())
			}
		}()
		if err != nil {
			return nil, err
		}

	default:
		l = zap.NewExample()
	}
	return l, nil
}
