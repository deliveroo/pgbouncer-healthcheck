package main

type configuration struct {
	Port                 int    `default:"8000"`
	Connstr              string `default:"host=localhost port=6543 dbname=pgbouncer sslmode=disable"`
	EnhancedCheck        bool   `default:"false" split_words:"true"`
	CheckDDAgent         bool   `envconfig:"check_ddagent" default:"false"`
	EnableDebugEndpoints bool   `default:"false" split_words:"true"`
	PGBouncerPort        int    `envconfig:"pgbouncer_port" default:"6543"`
}
