TEST = "SELECT 1"

CLEAR_LOCATIONS_TABLE = """
TRUNCATE TABLE db.location_restrictions
"""

CLEAR_EVENTS_TABLE = """
TRUNCATE TABLE db.events
"""

CLEAR_AGE_RESTRICTIONS_TABLE = """
TRUNCATE TABLE db.age_restrictions
"""

INSERT_EVENTS = """
INSERT INTO db.events (
    code, sport, title, additional_info, n_participants,
    stage, start_date, end_date
) VALUES
"""

INSERT_LOCATIONS = """
INSERT INTO db.location_restrictions (
    code, country, region, locality
) VALUES
"""

INSERT_AGE_RESTRICTIONS = """
INSERT INTO db.age_restrictions (
    code, gender, left_bound, right_bound, extra_mapping
) VALUES
"""
