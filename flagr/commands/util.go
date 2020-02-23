package commands

import (
	"github.com/checkr/goflagr"
	"github.com/urfave/cli"
)

func getFlagrClient(c *cli.Context) *goflagr.APIClient {
	// For now, just punt to using localhost:13480

	cfg := goflagr.NewConfiguration()
	cfg.BasePath = "http://localhost:13480/api/v1"
	client := goflagr.NewAPIClient(cfg)
	return client
}
