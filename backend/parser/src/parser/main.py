import time

from parser import pipeline
from parser.clickhouse.core import ClickHouse
from parser.log import logger
from parser.settings import settings


def payload(ch_client: ClickHouse):
    if settings.no_download:
        local_file = "input.pdf"
    else:
        logger.info("Downloading the remote file...")
        local_file = pipeline.download(settings.remote_file)
        logger.info("PDF file has been downloaded")
    logger.info("Parsing file's contents. This may take some time")
    sports = pipeline.parse(local_file)
    logger.info("Data has been parsed. Uploading to ClickHouse")
    upd_codes = ch_client.upload(sports)
    logger.info("Data has been uploaded")
    if len(upd_codes) != 0:
        logger.info(f"Pushing data to Meta. Host {settings.meta_host}")
        pipeline.meta_push(settings.meta_host, upd_codes)
        logger.info("Pushed data to Meta")


def main():
    logger.info("Connecting to ClickHouse")
    ch_client = ClickHouse(
        settings.ch_host,
        settings.ch_port,
        settings.ch_user,
        settings.ch_password,
    )
    logger.info("Successfully connected to ClickHouse")

    while True:
        payload(ch_client)
        time.sleep(settings.unload_timeout)
