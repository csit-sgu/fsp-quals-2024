package app

import (
	"bytes"
	"context"
	"strings"

	"app/internal/app/appcontext"
	"app/internal/config"
	"app/internal/db/clickhouse"
	"app/internal/db/os"
	"app/internal/log"

	"github.com/opensearch-project/opensearch-go/opensearchapi"
)

// Add all required onShutdown logic here.
func onStartup(ctx context.Context) error {
	chClient, err := clickhouse.InitClickhouseClient(
		config.C.Database.Clickhouse,
	)
	if err != nil {
		return err
	}

	osClient, err := os.InitOpenSearchClient(config.C.Database.OpenSearch)
	if err != nil {
		log.S.Error(
			"Failed to init OpenSearch client",
			log.L().Error(err),
		)
		return err
	}

	settings := strings.NewReader(`{
    "settings": {
        "number_of_shards": 1,
        "number_of_replicas": 0,
        "analysis": {
            "filter": {
                "russian_stop": {
                    "type": "stop",
                    "stopwords": "_russian_"
                },
                "russian_stemmer": {
                    "type": "stemmer",
                    "language": "russian"
                },
                "russian_keywords": {
                    "type": "keyword_marker",
                    "keywords": []
                }
            },
            "tokenizer": {
                "russian_ngram": {
                    "type": "ngram",
                    "min_gram": 3,
                    "max_gram": 4
                }
            },
            "analyzer": {
                "russian_analyzer": {
                    "type": "custom",
                    "tokenizer": "russian_ngram",
                    "filter": [
                        "lowercase",
                        "russian_stop",
                        "russian_stemmer",
                        "russian_keywords"
                    ]
                }
            }
        }
    },
    "mappings": {
        "properties": {
            "title": {
                "type": "text",
                "analyzer": "russian_analyzer"
            },
            "additional_info": {
                "type": "text",
                "analyzer": "russian_analyzer"
            },
            "code": {
                "type": "text",
                "analyzer": "standard"
            }
        }
    }
}`)
	resp, err := opensearchapi.IndicesCreateRequest{
		Index: config.C.Database.OpenSearch.Index,
		Body:  settings,
	}.Do(ctx, osClient)

	bodyBytes := new(bytes.Buffer)
	bodyBytes.ReadFrom(resp.Body)
	log.S.Debug("resp", log.L().Add("body", bodyBytes.String()))

	if err != nil {
		log.S.Error("Failed to create OpenSearch index", log.L().Error(err))
		return err
	} else {
		log.S.Debug(
			"OpenSearch index was created",
			log.L().Add("index", config.C.Database.OpenSearch.Index).Add("response", resp),
		)
	}

	appcontext.Ctx = &appcontext.AppContext{
		Clickhouse: chClient,
		OpenSearch: &os.OpenSearch{
			Client: osClient,
			Index:  config.C.Database.OpenSearch.Index,
		},
	}

	log.S.Debug("App startup is complete", log.L())

	return nil
}
