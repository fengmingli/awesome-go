package log

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

/**
 * @Author: LFM
 * @Date: 2022/3/27 7:56 下午
 * @Since: 1.0.0
 * @Desc: TODO
 */

func InitZeroLog() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "2006-01-02 15:04:05:001"})

}
