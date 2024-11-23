package clickhouse

import (
	"context"

	"app/internal/log"
	"app/internal/model"

	"github.com/ClickHouse/clickhouse-go/v2"
)

func (c *ClickhouseClient) GetCountries(
	ctx context.Context,
) (countries []string, err error) {
	var rows []model.Country
	if err := c.conn.Select(ctx, &rows, countryQuery); err != nil {
		log.S.Error(
			"Failed to execute country query",
			log.L().Add("query", countryQuery).Add("error", err),
		)
		return nil, err
	}

	for i := range rows {
		countries = append(countries, rows[i].Country)
	}

	log.S.Debug(
		"Countries were retrived successfully",
		log.L().Add("count", len(countries)),
	)

	return countries, nil
}

func (c *ClickhouseClient) GetRegions(
	ctx context.Context,
	country string,
) (regions []string, err error) {
	var rows []model.Region
	if err := c.conn.Select(ctx, &rows, regionQuery, clickhouse.Named("country", country)); err != nil {
		log.S.Error(
			"Failed to execute region query",
			log.L().Add("query", countryQuery).Add("error", err),
		)
		return nil, err
	}

	log.S.Debug(
		"Regions were retrived successfully",
		log.L().Add("count", len(regions)),
	)

	for i := range rows {
		regions = append(regions, rows[i].Region)
	}

	return regions, nil
}

func (c *ClickhouseClient) GetLocalities(
	ctx context.Context,
	country string,
	region string,
) (localities []string, err error) {
	var rows []model.Locality
	if err := c.conn.Select(ctx, &rows, localityQuery, clickhouse.Named("country", country), clickhouse.Named("region", region)); err != nil {
		log.S.Error(
			"Failed to execute locality query",
			log.L().Add("query", localityQuery).Add("error", err),
		)
		return nil, err
	}

	log.S.Debug(
		"Localities were retrieved successfully",
		log.L().Add("count", len(rows)),
	)

	for i := range rows {
		localities = append(localities, rows[i].Locality)
	}

	return localities, nil
}

func (c ClickhouseClient) getLocationData(
	ctx context.Context,
	code string,
) (locationData []model.LocationData, err error) {
	if err := c.conn.Select(ctx, &locationData, locationQuery, clickhouse.Named("code", code)); err != nil {
		log.S.Error(
			"Failed to execute age query",
			log.L().Add("query", locationQuery).Add("error", err),
		)
		return nil, err
	}

	return locationData, nil
}
