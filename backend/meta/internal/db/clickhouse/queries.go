package clickhouse

const countryQuery = `
    SELECT DISTINCT country FROM db.general_view
`

const sportsQuery = `
    SELECT DISTINCT sport FROM db.general_view
`

const regionQuery = `
    SELECT DISTINCT region FROM db.general_view WHERE country = @country
`

const localityQuery = `
    SELECT DISTINCT locality FROM db.general_view WHERE country = @country AND region = @region
`

const ageQuery = `
    SELECT DISTINCT left_bound, right_bound, gender, extra_mapping FROM db.general_view WHERE code = @code
`

const locationQuery = `
    SELECT DISTINCT country, region, locality FROM db.general_view WHERE code = @code
`
