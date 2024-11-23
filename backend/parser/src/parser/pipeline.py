import httpx
import pandas as pd
from pypdf import PdfReader

from parser import steps


def download(remote: str) -> str:
    filename = "input.pdf"
    with httpx.Client(verify=False) as client:
        with client.stream("GET", remote) as response:
            response.raise_for_status()
            with open(filename, "wb") as file:
                for chunk in response.iter_bytes():
                    file.write(chunk)
    return filename


def parse(filename: str) -> pd.DataFrame:
    reader = PdfReader(filename)
    pages = reader.pages

    res = "\n".join(page.extract_text() for page in pages)
    pipeline = [
        steps.remove_page_numbers,
        steps.split_into_sport_kinds,
    ]
    for step in pipeline:
        res = step(res)

    sports = {
        key: value for key, value in zip(res[::2], res[1::2], strict=True)
    }
    pipeline = [
        steps.split_rows,
        steps.competitors_number,
        steps.competition_title,
        steps.dates,
        steps.extract_group,
        steps.parse_group,
        steps.country,
        steps.locality,
    ]
    dataset = pd.DataFrame()
    for key in sports:
        res = sports[key]
        for step in pipeline:
            res = step(res)
        res["Sport"] = key
        dataset = pd.concat((dataset, res))

    return dataset
