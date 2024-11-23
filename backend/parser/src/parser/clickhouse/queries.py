TEST = "SELECT 1"

CLEAR_LOCATIONS_TABLE = """
ALTER TABLE db.location_restrictions DELETE WHERE true
"""

CLEAR_EVENTS_TABLE = """
ALTER TABLE db.events DELETE WHERE true
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
