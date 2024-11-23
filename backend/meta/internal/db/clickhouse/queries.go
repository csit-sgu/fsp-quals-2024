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
	from db.events %s
	order by start_date desc
)
select %s
from ordered o
inner join db.general_view using (code)
WHERE %s
order by o.page_index asc
`

const subInsertQuery = `
INSERT INTO db.subscriptions (
    confirmation, email, is_active, code, gender, age, sport, additional_info,
    country, region, event_type, event_scale, start_date, end_date
) VALUES (
    @confirmation, @email, false, @code, @gender, @age, @sport, @additional_info,
    @country, @region, @event_type, @event_scale, @start_date, @end_date
)
`

const filterCounterQuery = `
with ordered as (
	select distinct code, row_number() over (order by start_date desc) as page_index, start_date
	from db.events %s
	order by start_date desc
)
select count() as count
from ordered o
`

const codeQuery = `
SELECT code, title, additional_info FROM db.events WHERE code in (@codes)
`

const subCountQuery = `
SELECT count(*) FROM db.subscriptions
WHERE confirmation = @confirmation
`

const subActivateQuery = `
ALTER TABLE db.subscriptions
UPDATE is_active = true WHERE confirmation = @confirmation
`
