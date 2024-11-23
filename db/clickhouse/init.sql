CREATE TABLE db.events (
    code String NOT NULL,
    sport String NOT NULL,
    title String NOT NULL,
    additional_info String,
    n_participants UInt32 NOT NULL,
    stage String NOT NULL,
    start_date Date NOT NULL,
    end_date Date NOT NULL,
)
ENGINE = MergeTree()
ORDER BY (sport, start_date, code);

CREATE TABLE db.age_restrictions (
    code String NOT NULL,
    gender Enum('male' = 1, 'female' = 2) NOT NULL,
    left_bound UInt32 NOT NULL,
    right_bound UInt32 NOT NULL,
    extra_mapping String,
)
ENGINE = MergeTree()
ORDER BY (left_bound, right_bound, code);

CREATE TABLE db.location_restrictions (
    code String NOT NULL,
    country String NOT NULL,
    region String NOT NULL,
    locality String NOT NULL,
)
ENGINE = MergeTree()
ORDER BY (country, region, locality, code);

CREATE VIEW db.general_view
AS
SELECT
	e.code AS code,
    start_date,
    country,
    region,
    title,
    additional_info,
    n_participants,
    stage,
    end_date,
    sport,
    left_bound,
    right_bound,
    gender,
    locality,
    extra_mapping,
    dense_rank() OVER ( ORDER BY e.code) AS page_index
FROM
    db.events AS e
LEFT JOIN
    db.location_restrictions AS l
ON e.code = l.code
LEFT JOIN
    db.age_restrictions AS a
ON e.code = a.code
