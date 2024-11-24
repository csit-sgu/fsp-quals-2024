package opensearch

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"app/internal/app/appcontext"
	"app/internal/log"
	"app/internal/model"

	"github.com/opensearch-project/opensearch-go/opensearchapi"
)

func generateBody(
	title string,
	additionalInfo string,
) io.Reader {
	body := fmt.Sprintf(`
    {
        "query": {
            "bool": {
                "should": [
                    {
                        "fuzzy": {
                            "title": "%s"
                        }
                    },
                    {
                        "fuzzy": {
                            "additional_info": "%s"
                        }
                    }
                ]
            }
        }
    }
    `, title, additionalInfo)

	return strings.NewReader(body)
}

func ApplyFuzzySearch(
	ctx context.Context,
	l log.LogObject,
	title string,
	additionalInfo string,
	e []*model.Event,
) (events []*model.Event, err error) {
	os := appcontext.Ctx.OpenSearch
	client := appcontext.Ctx.OpenSearch.Client

	codes := []string{}
	for _, event := range e {
		codes = append(codes, event.Code)
	}

	body := generateBody(title, additionalInfo)

	request := opensearchapi.SearchRequest{
		Index: []string{os.Index},
		Body:  body,
	}

	response, err := request.Do(ctx, client)
	if err != nil {
		log.S.Error(
			"Failed to execute fuzzy search query",
			log.L().Add("query", title).Error(err),
		)
		return nil, err
	}
	defer response.Body.Close()

	var result map[string]interface{}

	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		log.S.Error(
			"Failed to decode fuzzy search response",
			log.L().Add("query", title).Error(err),
		)
		return nil, err
	}

	log.S.Debug("Fuzzy search query was executed successfully", l)

	return e, nil
}
