from parser import pipeline
from parser.clickhouse.core import ClickHouse
from parser.log import logger
from parser.settings import settings


def main():
    logger.info("Connecting to ClickHouse")
    ch_client = ClickHouse(
        settings.ch_host,
        settings.ch_port,
        settings.ch_user,
        settings.ch_password,
    )
    logger.info("Successfully connected to ClickHouse")

    logger.info("Downloading the remote file...")
    # local_file = pipeline.download(settings.remote_file)
    logger.info(
        "PDF file has been downloaded. Parsing file's contents. "
        + "This may take some time"
    )
    local_file = "input.pdf"
    sports = pipeline.parse(local_file)
    logger.info("Data has been parsed. Uploading to ClickHouse")
    ch_client.upload(sports)
    logger.info("Data has been uploaded")
