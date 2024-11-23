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

const filterQuery = `
with ordered as (
	select distinct code, row_number() over (order by start_date desc) as page_index, start_date
	from db.events
    %s
	order by start_date desc
)
select %s
from ordered o
inner join db.general_view using (code)
WHERE %s
order by o.page_index asc
`
