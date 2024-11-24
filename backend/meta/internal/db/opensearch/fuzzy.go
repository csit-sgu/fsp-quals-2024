package opensearch

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"sort"
	"strings"

	"app/internal/app/appcontext"
	"app/internal/log"
	"app/internal/model"

	"github.com/opensearch-project/opensearch-go/opensearchapi"
)

type HitSource struct {
	Title          string `ch:"title"`
	AdditionalInfo string `ch:"additional_info"`
	Code           string `ch:"code"`
}

type Hit struct {
	Index  string    `json:"_index"`
	ID     string    `json:"_id"`
	Score  float64   `json:"_score"`
	Source HitSource `json:"_source"`
}

func generateBody(
	title string,
	additionalInfo string,
) (io.Reader, string) {
	body := fmt.Sprintf(`
    {
        "query": {
            "bool": {
                "should": [
                    {
                        "term": {
                            "title": "%s"
                        }
                    },
                    {
                        "term": {
                            "additional_info": "%s"
                        }
                    },
                    {
                        "fuzzy": {
                            "title": {
                                "value": "%s",
                                "fuzziness": "AUTO",
                                "prefix_length": 1,
                                "max_expansions": 10
                            }
                        }
                    }
                ]
            }
        }
    }
    `, title, additionalInfo)

	return strings.NewReader(body), body
}

type Hits struct {
	Hits []Hit `json:"hits"`
}

func ApplyFuzzySearch(
	ctx context.Context,
	l log.LogObject,
	title string,
	additionalInfo string,
	e []*model.Event,
) ([]*model.Event, error) {
	os := appcontext.Ctx.OpenSearch
	client := appcontext.Ctx.OpenSearch.Client

	codes := []string{}
	for _, event := range e {
		codes = append(codes, event.Code)
	}

	requestBody, debug := generateBody(title, additionalInfo)
	log.S.Debug("Fuzzy search query", log.L().Add("body", debug))

	request := opensearchapi.SearchRequest{
		Index: []string{os.Index},
		Body:  requestBody,
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

	var parsedBody struct {
		Hits struct {
			Hits []Hit `json:"hits"`
		} `json:"hits"`
	}

	raw, err := io.ReadAll(response.Body)
	log.S.Debug("Raw response", log.L().Add("raw", string(raw)))
	if err != nil {
		log.S.Error(
			"Failed to read fuzzy search response",
			log.L().Add("query", title).Error(err),
		)
		return nil, err
	}

	err = json.Unmarshal(raw, &parsedBody)

	hits := parsedBody.Hits.Hits

	log.S.Debug("Fuzzy search query was executed successfully", l.Add("hits", parsedBody))

	scores := map[string]float64{}

    foundCodes := []string{}

	for _, hit := range hits {
		if err != nil {
			log.S.Warn("Failed to decode code", log.L().Add("code", hit.ID).Error(err))
			continue
		}

        foundCodes = append(foundCodes, hit.Source.Code)
		scores[hit.Source.Code] = hit.Score
        log.S.Info("Hit", log.L().Add("code", hit.Source.Code).Add("score", hit.Score))
	}

	for _, event := range e {
		event.Score = scores[event.Code]
		log.S.Debug(
			"Event score",
			log.L().Add("code", event.Code).Add(
				"score",
				event.Score,
			).Add("scores", scores[event.Code]),
		)
	}

	sort.Slice(e, func(i, j int) bool {
		return e[i].Score < e[j].Score
	})

	log.S.Debug("Fuzzy search query was executed successfully", l)

	return e, nil
}
