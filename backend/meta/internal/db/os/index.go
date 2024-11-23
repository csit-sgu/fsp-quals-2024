package os

import (
	"context"
	"fmt"
	"strings"

	"app/internal/log"
	"app/internal/model"

	"github.com/opensearch-project/opensearch-go"
)

type OpenSearch struct {
	Client *opensearch.Client
	Index  string
}

func (c *OpenSearch) convertToQuery(index []model.IndexData) string {
	query := ""
	for i := range index {
		query += fmt.Sprintf(
			"{\"index\":{\"_index\":\"%s\",\"_id\":\"%s\"}}\n",
			c.Index,
			index[i].Code,
		)
		query += fmt.Sprintf(
			"{\"title\":\"%s\",\"additional_info\":\"%s\"}\n",
			index[i].Title,
			index[i].AdditionalInfo,
		)
	}
	return query
}

func (c *OpenSearch) IndexData(
	ctx context.Context,
	l log.LogObject,
	indexData []model.IndexData,
) error {
	query := c.convertToQuery(indexData)

	log.S.Debug("Built query for bulk request", l.Add("query", query))
	c.Client.Bulk(strings.NewReader(query))
	log.S.Debug("Bulk request was sent successfully", l)
	return nil
}
