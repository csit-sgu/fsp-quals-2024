package clickhouse

const countryQuery = `
    SELECT DISTINCT country FROM events
`

const regionQuery = `
    SELECT DISTINCT region FROM events WHERE country = $1
`

const localityQuery = `
    SELECT DISTINCT locality FROM events WHERE region = :region
`
