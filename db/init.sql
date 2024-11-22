CREATE TABLE events (
    code String NOT NULL,
    gender Enum('male' = 1, 'female' = 2) NOT NULL,
    sport String NOT NULL,
    additional_info String,
    place String NOT NULL,
    n_participants UInt32 NOT NULL,
    stage String NOT NULL,
    start_date Datetime NOT NULL,
    end_date Datetime NOT NULL,
)
ENGINE = MergeTree()
ORDER BY (gender, start_date, code);

CREATE TABLE event_restrictions (
    code String NOT NULL,
    gender Enum('male' = 1, 'female' = 2) NOT NULL,
    left_bound Datetime NOT NULL,
    right_bound Datetime NOT NULL,
    extra_mapping String,
)
ENGINE = MergeTree()
ORDER BY (left_bound, right_bound, code);
