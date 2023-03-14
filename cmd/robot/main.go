package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/alecthomas/kingpin/v2"
	"github.com/jackc/pgx/v4"
	"github.com/mrOwner/robot/internal/app"
	cc "github.com/mrOwner/robot/internal/collect_candles"
	cg "github.com/mrOwner/robot/pkg/candles_grabber"
	"github.com/mrOwner/robot/util"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	_ app.LifeCycle = (*cc.Collector)(nil)

	configPath = kingpin.Flag("config", "The path to the directory with a config file.").Default(".").ExistingDir()

	grab = kingpin.Command("grab", "Download new candles and save them to the database.")
	file = grab.Arg("file", "The path to a file with FIGIs.").Default("./configs/FIGIs.txt").ExistingFile()
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.DurationFieldUnit = time.Second

	app := app.New()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	switch kingpin.Parse() {
	case "grab":
		config, err := util.LoadConfig(*configPath)
		if err != nil {
			log.Fatal().Err(err).Msg("cannot load config")
		}

		token := config.SandboxToken
		if config.Enviroment == util.Production {
			token = config.Token
		}

		conn, err := pgx.Connect(ctx, config.DBSource)
		if err != nil {
			log.Fatal().Err(err).Msg("unable to connect to database")
		}
		defer conn.Close(context.Background())

		grabber := cg.NewTinkoffApiReader(token, config.URLhistoricalCandles)
		collector := cc.New(*file, grabber, conn)

		app.AddSynchronous("grabber", collector)
	default:
		log.Fatal().Msg("need choose the command")
	}

	if err := app.Start(ctx); err != nil {
		log.Fatal().Err(err).Msg("error on start")
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	<-c

	if err := app.Stop(context.Background()); err != nil {
		log.Fatal().Err(err).Msg("unexpected ending")
	}
}
