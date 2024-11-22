package clickhouse

const countryQuery = `
    SELECT DISTINCT country FROM db.events
`

const regionQuery = `
    SELECT DISTINCT region FROM db.events WHERE country = @country
`

const localityQuery = `
    SELECT DISTINCT locality FROM db.events WHERE country = @country AND region = @region
`
