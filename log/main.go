package main

/**
 * @Author: LFM
 * @Date: 2023/11/23 23:45
 * @Since: 1.0.0
 * @Desc: TODO
 */
import (
	"errors"

	"github.com/go-logr/logr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	k8s_zap "sigs.k8s.io/controller-runtime/pkg/log/zap"
)

func main() {
	Start()
}

func Start() {
	logf.SetLogger(CreateLogger(false, false))
	log := logf.Log.WithName("alter")

	NewApp(log)
}

func NewApp(log logr.Logger) {
	log.Error(errors.New("xxx"), "alerts garbage collection failed")

	log.Info("alerts garbage collection failed")

	log.V(1).Info("alerts garbage collection failed")

	log.V(2).Info("alerts garbage collection failed-2")

	log.V(3).Info("alerts garbage collection failed-3")

	log.V(4).Info("alerts garbage collection failed-4")

	log.V(4).Info("alerts garbage collection failed-4")
}

func CreateLogger(debug bool, development bool) logr.Logger {
	var config zapcore.EncoderConfig
	if development {
		config = zap.NewDevelopmentEncoderConfig()
	} else {
		config = zap.NewProductionEncoderConfig()
	}
	config.EncodeTime = zapcore.ISO8601TimeEncoder

	var encoder zapcore.Encoder
	if development {
		encoder = zapcore.NewConsoleEncoder(config)
	} else {
		encoder = zapcore.NewJSONEncoder(config)
	}

	level := zap.InfoLevel
	if debug {
		level = zap.DebugLevel
	}

	return k8s_zap.New(k8s_zap.UseDevMode(development), k8s_zap.Encoder(encoder), k8s_zap.Level(level))
}
