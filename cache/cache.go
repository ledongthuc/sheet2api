package cache

import (
	"sync"

	"github.com/imkira/go-ttlmap"
	"github.com/rs/zerolog/log"
)

var M *ttlmap.Map

func init() {
	sync.OnceFunc(func() {
		log.Info().Msgf("Init cache")
		M = ttlmap.New(nil)
	})()
}
