package config

import (
	"fmt"
	"strings"
)

// ConnectionString generates a PostgreSQL connection string for pgxpool.
func (c *Bot) ConnectionString() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", c.Database.PSQLHOST, c.Database.PSQLPORT,
		c.Database.PSQLUSER, c.Database.PSQLPASS, c.Database.PSQLDB))

	if _, ok := c.Database.AdditionalParams["sslmode"]; !ok {
		b.WriteString(" sslmode=disable")
	}

	for key, value := range c.Database.AdditionalParams {
		b.WriteString(fmt.Sprintf(" %s=%s", key, value))
	}

	// Include pool configurations directly in the connection string
	if c.Database.DBMaxOpenConns > 0 {
		b.WriteString(fmt.Sprintf(" pool_max_conns=%d", c.Database.DBMaxOpenConns))
	}
	if c.Database.MaxIdleConns > 0 {
		b.WriteString(fmt.Sprintf(" pool_min_conns=%d", c.Database.MaxIdleConns))
	}
	if c.Database.ConnectionMaxLifetime.GoDuration() > 0 {
		b.WriteString(fmt.Sprintf(" pool_max_conn_lifetime=%s", c.Database.ConnectionMaxLifetime.GoDuration()))
	}

	return b.String()
}
