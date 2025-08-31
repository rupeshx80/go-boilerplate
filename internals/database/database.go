package database

import (
	"context"
	"fmt"
	"net"
	"net/url"
	"strconv"
	"time"

	pgxzero "github.com/jackc/pgx-zerolog"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/tracelog"
	"github.com/newrelic/go-agent/v3/integrations/nrpgx5"
	"github.com/rs/zerolog"
	"github.com/sriniously/go-boilerplate/internal/config"
	loggerConfig "github.com/sriniously/go-boilerplate/internal/logger"
)

type Database struct {
	Pool *pgxpool.Pool
	log *zerolog.Logger
}

const DatabasePingTimeout = 10