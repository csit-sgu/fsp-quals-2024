from datetime import datetime

import pandas as pd
from clickhouse_driver import Client

from parser.clickhouse import queries
from parser.log import logger


class ClickHouse:
    def __init__(self, host: str, port: int, user: str, password: str) -> None:
        self.client = Client(
            host=host, user=user, password=password, port=port
        )
        logger.info("Testing the ClickHouse connection")
        self.client.execute(queries.TEST)
        logger.info("ClickHouse connection successful")

    def upload(self, df: pd.DataFrame):
        logger.info("Clearing the tables")
        self.client.execute(queries.CLEAR_LOCATIONS_TABLE)
        self.client.execute(queries.CLEAR_EVENTS_TABLE)
        self.client.execute(queries.CLEAR_AGE_RESTRICTIONS_TABLE)
        logger.info("Tables have been cleared")

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
                df["Code"],
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
        logger.info(f"Uploading event data ({len(data)} records)")
        self.client.execute(queries.INSERT_EVENTS, data)
        logger.info("Event data has been uploaded")

        tmp_df = df.explode("Locality", ignore_index=True)
        new_cols = pd.DataFrame(
            tmp_df["Locality"].to_list(),
            columns=["Region", "Locality"],
        )
        tmp_df = tmp_df.drop("Locality", axis=1)
        tmp_df = pd.concat((tmp_df, new_cols), axis=1)

        data = pd.concat(
            (
                tmp_df["Code"],
                tmp_df["Country"],
                tmp_df["Region"],
                tmp_df["Locality"],
            ),
            axis=1,
        ).values.tolist()
        logger.info(
            f"Inserting event location information ({len(data)} records)"
        )
        self.client.execute(queries.INSERT_LOCATIONS, data)
        logger.info("Event locations have been uploaded")

        tmp_df = df.explode("Group", ignore_index=True)
        new_cols = pd.DataFrame(
            tmp_df["Group"].to_list(),
            columns=["Original", "Gender", "Lower Bound", "Upper Bound"],
        )
        tmp_df = tmp_df.drop("Group", axis=1)
        tmp_df = pd.concat((tmp_df, new_cols), axis=1)

        data = pd.concat(
            (
                tmp_df["Code"],
                tmp_df["Gender"],
                tmp_df["Lower Bound"],
                tmp_df["Upper Bound"],
                tmp_df["Original"],
            ),
            axis=1,
        ).values.tolist()
        logger.info(
            f"Inserting age restriction information ({len(data)} records)"
        )
        self.client.execute(queries.INSERT_AGE_RESTRICTIONS, data)
        logger.info("Event age restrictions have been uploaded")
