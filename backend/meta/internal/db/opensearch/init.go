package opensearch

import (
	"app/internal/config"
	"fmt"

	"github.com/opensearch-project/opensearch-go"
)

func InitOpenSearchClient(
	c config.OpenSearchConfig,
) (*opensearch.Client, error) {
	return opensearch.NewClient(opensearch.Config{
		Addresses: []string{
            fmt.Sprintf("http://%s:%s", c.Host, c.Port),
		},
		Username: c.Username,
		Password: c.Password,
	})
}
