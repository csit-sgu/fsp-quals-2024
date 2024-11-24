package os

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"app/internal/config"

	"github.com/opensearch-project/opensearch-go"
)

func InitOpenSearchClient(
	c config.OpenSearchConfig,
) (*opensearch.Client, error) {
	return opensearch.NewClient(opensearch.Config{
		Addresses: []string{
			fmt.Sprintf("https://%s:%s", c.Host, c.Port),
		},
		Username: c.Username,
		Password: c.Password,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	})
}
