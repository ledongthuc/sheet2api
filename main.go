package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/imkira/go-ttlmap"
	"github.com/ironstar-io/chizerolog"
	"github.com/ledongthuc/goterators"
	"github.com/ledongthuc/sheet2api/cache"
	"github.com/ledongthuc/sheet2api/configs"
	"github.com/ledongthuc/sheet2api/core"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/xuri/excelize/v2"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339})

	var configFilePath string
	flag.StringVar(&configFilePath, "config-file-path", "config.yaml", "the path of config file is loaded into the system")
	flag.Parse()

	log.Info().Msgf("Config file '%s'", configFilePath)
	cs, err := configs.LoadConfigFile(configFilePath)
	if err != nil {
		log.Warn().Msgf("Fail to load config file '%s': %v", configFilePath, err)
		log.Warn().Msg("Use default configs")
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(chizerolog.LoggerMiddleware(&log.Logger))
	r.Get("/{file_name}/{sheet_name}", func(w http.ResponseWriter, r *http.Request) {
		fileName := chi.URLParam(r, "file_name")
		sheetName := chi.URLParam(r, "sheet_name")

		fileConfig, _, err := goterators.Find(cs.Files, func(item configs.File) bool { return item.URLReplacedName == fileName })
		if err != nil {
			http.Error(w, fmt.Sprintf("'%s' doesn't exist", fileName), http.StatusNotFound)
			return
		}

		cacheKey := fmt.Sprintf("/%s/%s", fileName, sheetName)
		if fileConfig.IsEnableCache() {
			cacheValue, err := cache.M.Get(cacheKey)
			if err == nil {
				log.Info().Msgf("'%s' hit cache", cacheKey)
				w.Write(cacheValue.Value().([]byte))
				return
			}
		}

		if err != nil && (errors.Is(err, ttlmap.ErrNotExist) || errors.Is(err, ttlmap.ErrDrained)) {
			http.Error(w, fmt.Sprintf("Fail to load cache: %s", err.Error()), http.StatusInternalServerError)
			return
		}

		result, err := core.GetRows(fileConfig.FilePath, sheetName)
		if err != nil {
			if errors.As(err, &excelize.ErrSheetNotExist{SheetName: sheetName}) {
				http.Error(w, fmt.Sprintf("Sheet '%s' doesn't exist", sheetName), http.StatusNotFound)
			} else {
				http.Error(w, fmt.Sprintf("Fail to get row: %s", err.Error()), http.StatusInternalServerError)
			}
			return
		}

		j, err := json.Marshal(result)
		if err != nil {
			http.Error(w, fmt.Sprintf("unmarshal response: %s", err.Error()), http.StatusInternalServerError)
			return
		}

		if fileConfig.IsEnableCache() {
			item := ttlmap.NewItem(j, ttlmap.WithTTL(time.Duration(fileConfig.CacheInSecond)*time.Second))
			if err := cache.M.Set(cacheKey, item, nil); err != nil {
				http.Error(w, fmt.Sprintf("Fail to save cache: %s", err.Error()), http.StatusInternalServerError)
				return
			}
		}

		w.Write(j)
	})

	server := fmt.Sprintf("%s:%s", cs.HostIP, cs.HostPort)
	log.Info().Msgf("Start server: %s", server)
	log.Fatal().Err(http.ListenAndServe(server, r))
}
