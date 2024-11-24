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

const filterQuery = `
with country_view as (
    select distinct country, region, locality, code from db.location_restrictions %s
),
age_view as (
    select distinct left_bound, right_bound, gender, code, extra_mapping from db.age_restrictions %s
),
ordered as (
	select distinct code, row_number() over (order by start_date desc) as page_index, start_date, title, additional_info, event_type, event_scale, n_participants, end_date, sport
	from db.events %s
	order by start_date desc
)
select %s
from ordered o
inner join country_view cv on cv.code = o.code
inner join age_view av on av.code = o.code %s
order by page_index asc
`

const filterCounterQuery = `
with country_view as (
    select distinct country, region, locality, code from db.location_restrictions %s
),
age_view as (
    select distinct left_bound, right_bound, gender, code, extra_mapping from db.age_restrictions %s
),
ordered as (
	select distinct code, row_number() over (order by start_date desc) as page_index, start_date, title, additional_info, event_type, event_scale, n_participants, end_date, sport
	from db.events %s
	order by start_date desc
),
available as (
    select distinct o.code
    from ordered o
    inner join country_view cv on cv.code = o.code
    inner join age_view av on av.code = o.code
    order by page_index asc
)
select count() as count from available;
`

const codeQuery = `
SELECT code, title, additional_info FROM db.events WHERE code in (@codes)
`

const subFindByMail = `
SELECT * FROM db.subscriptions WHERE email = @email AND NOT is_active
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

const subCountQuery = `
SELECT count(*) FROM db.subscriptions
WHERE confirmation = @confirmation
`

const subActivateQuery = `
ALTER TABLE db.subscriptions
UPDATE is_active = true WHERE confirmation = @confirmation
`
