package opensearch

import (
	"app/internal/app/appcontext"
	"app/internal/log"
	"app/internal/model"
	"context"
	"fmt"
	"io"
	"strings"

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
                            "title": "соревнование"
                        }
                    },
                    {
                        "fuzzy": {
                            "additional_info": "класс"
                        }
                    }
                ]
            }
        }
    }
    `)

    return strings.NewReader(body)

}

func ApplyFuzzySearch(
	ctx context.Context,
	l log.LogObject,
	searchString string,
	e []*model.Event,
) (events []*model.Event, err error) {
    os := appcontext.Ctx.OpenSearch
    client := appcontext.Ctx.OpenSearch.Client

    codes := []string{}
    for _, event := range e {
        codes = append(codes, event.Code)
    }

    body := generateBody(searchString, searchString)

    request := opensearchapi.SearchRequest{
        Index: []string{os.Index},
        Body: body,
    }

    response, err := request.Do(ctx, client)

    if err != nil {
        log.S.Error(
            "Failed to execute fuzzy search query",
            log.L().Add("query", searchString).Error(err),
        )
        return nil, err
    }
    defer response.Body.Close()

    log.S.Debug("Fuzzy search query was executed successfully", l.Add("count", response.Body))




	return e, nil
}
