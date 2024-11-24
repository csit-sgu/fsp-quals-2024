package os

import (
	"context"
	"fmt"
	"strings"

	"app/internal/log"
	"app/internal/model"

	"github.com/opensearch-project/opensearch-go"
	"github.com/opensearch-project/opensearch-go/opensearchapi"
)

type OpenSearch struct {
	Client *opensearch.Client
	Index  string
}

func (c *OpenSearch) convertToQuery(index []model.IndexData) string {
	query := ""
	for i := range index {
		log.S.Debug("Index data", log.L().Add("data", index[i]))
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
	for i := range indexData {
		log.S.Debug("Index data", log.L().Add("data", indexData[i]))

		req := opensearchapi.IndexRequest{
			Index: c.Index,
			Body: strings.NewReader(fmt.Sprintf(
				"{\"title\":\"%s\",\"additional_info\":\"%s\"}",
				indexData[i].Title,
				indexData[i].AdditionalInfo,
			)),
			DocumentID: indexData[i].Code,
		}


		resp, err := req.Do(ctx, c.Client)

		if err != nil {
			log.S.Warn(
				"Failed to index data",
				log.L().
					Add("index", c.Index).
					Add("code", indexData[i].Code).
					Error(err),
			)
            continue
		}

        resp.Body.Close()
	}

	log.S.Debug("Index request was completed", l)
	return nil
}
