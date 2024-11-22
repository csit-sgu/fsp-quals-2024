from datetime import datetime

import pandas as pd
from clickhouse_driver import Client

UPLOAD_EVENTS = """
INSERT INTO db.events (
    code, sport, title, additional_info, n_participants,
    stage, start_date, end_date
) VALUES
"""

UPLOAD_LOCATIONS = """
INSERT INTO db.location_restrictions (
    code, country, region, locality
) VALUES
"""


class ClickHouse:
    def __init__(self, host: str, port: int, user: str, password: str) -> None:
        self.client = Client(
            host=host, user=user, password=password, port=port
        )

    def upload(self, df: pd.DataFrame):
        df["Stage"] = ""
        df["Date Start"] = df["Date Start"].apply(
            lambda x: datetime.strptime(x, "%d.%m.%Y")
        )
        df["Date End"] = df["Date End"].apply(
            lambda x: datetime.strptime(x, "%d.%m.%Y")
        )
        df["Competitors"] = df["Competitors"].apply(lambda x: int(x))

        data = pd.concat(
            (
                df["ID"],
                df["Sport"],
                df["Title"],
                df["Raw Discipline"],
                df["Competitors"],
                df["Stage"],
                df["Date Start"],
                df["Date End"],
            ),
            axis=1,
        ).values.tolist()
        self.client.execute(UPLOAD_EVENTS, data)

        df = df.explode("Locality", ignore_index=True)
        new_cols = pd.DataFrame(
            df["Locality"].to_list(),
            columns=["Region", "Locality"],
        )
        df = df.drop("Locality", axis=1)
        df = pd.concat((df, new_cols), axis=1)

        data = pd.concat(
            (df["ID"], df["Country"], df["Region"], df["Locality"]), axis=1
        ).values.tolist()
        self.client.execute(UPLOAD_LOCATIONS, data)
