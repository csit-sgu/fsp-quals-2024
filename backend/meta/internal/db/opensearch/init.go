package opensearch

import (
	"app/internal/config"

	"github.com/opensearch-project/opensearch-go"
)

func InitOpenSearchClient(
	c config.OpenSearchConfig,
) (*opensearch.Client, error) {
	return opensearch.NewClient(opensearch.Config{
		Addresses: []string{
			c.Host,
		},
		Username: c.Username,
		Password: c.Password,
	})
}
