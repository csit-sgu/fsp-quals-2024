import re

import pandas as pd
from pypdf import PdfReader

from parser import util

PAGE_NUM = r"Стр\.\s*\d+\s*из\s*\d+"
ID = r"\d{10,}"
ROW_START = rf"\n(?={ID}\s+)"
COMPETITORS_NUMBER = r"\s+(?=\d+$)"
COMPETITION_TITLE_BEFORE = r"\s+"
LOWERCASE_RUS = r"[а-я]"
COMPETITION_TITLE_AFTER = rf"(?=\s+{LOWERCASE_RUS})"
DATE = r"\d{2}\.\d{2}\.\d{4}"
DATES_BEFORE = rf"\n(?={DATE}\n{DATE})"
DATES_AFTER = rf"(?<={DATE}\n{DATE})\n"
UPPERCASE_RUS = r"[А-Я]"
DISCIPLINE_BEFORE = rf"\n?(?={UPPERCASE_RUS}|$)"
REGION_AFTER = r"(?:,\s)|$"


def remove_page_numbers(s: str) -> str:
    return re.sub(PAGE_NUM, "", s)


def split_rows(s: str) -> pd.DataFrame:
    return pd.DataFrame(
        map(str.strip, re.split(ROW_START, s)), columns=["Raw"]
    )


def competitors_number(df: pd.DataFrame) -> pd.DataFrame:
    return pd.DataFrame(
        df["Raw"]
        .apply(lambda s: re.split(COMPETITORS_NUMBER, s, maxsplit=1))
        .to_list(),
        columns=["Raw", "Competitors"],
    )


def competition_title(df: pd.DataFrame) -> pd.DataFrame:
    df = util.flat_apply(
        df,
        "Raw",
        lambda s: re.split(COMPETITION_TITLE_BEFORE, s, maxsplit=1),
        columns=["ID", "Raw"],
    )

    df = util.flat_apply(
        df,
        "Raw",
        lambda s: re.split(COMPETITION_TITLE_AFTER, s, maxsplit=1),
        columns=["Title", "Raw"],
    )

    cleaned_title = pd.DataFrame(
        df["Title"].apply(lambda s: re.sub(r"\s+", " ", s.strip()))
    )
    cleaned_rest = pd.DataFrame(df["Raw"].apply(lambda s: s.strip()))

    df = df.drop("Raw", axis=1)
    return pd.concat(
        (df["ID"], cleaned_title, cleaned_rest, df["Competitors"]),
        axis=1,
    )


def dates(df: pd.DataFrame) -> pd.DataFrame:
    df = util.flat_apply(
        df,
        "Raw",
        lambda s: re.split(DATES_BEFORE, s, maxsplit=1),
        columns=["Raw Group", "Raw"],
    )

    df = util.flat_apply(
        df,
        "Raw",
        lambda s: re.split(DATES_AFTER, s, maxsplit=1),
        columns=["Raw Dates", "Raw Place"],
    )

    return util.flat_apply(
        df, "Raw Dates", str.split, columns=["Date Start", "Date End"]
    )


def group(df: pd.DataFrame) -> pd.DataFrame:
    return util.flat_apply(
        df,
        "Raw Group",
        lambda s: re.split(DISCIPLINE_BEFORE, s, maxsplit=1),
        columns=["Group", "Raw Discipline"],
    )


def country(df: pd.DataFrame) -> pd.DataFrame:
    return util.flat_apply(
        df,
        "Raw Place",
        lambda s: re.split("\n+", s, maxsplit=1),
        columns=["Country", "Raw Region"],
    )


def region(df: pd.DataFrame) -> pd.DataFrame:
    return util.flat_apply(
        df,
        "Raw Region",
        lambda s: re.split(REGION_AFTER, s, maxsplit=1),
        columns=["Region", "City"],
    )


def main():
    reader = PdfReader("input.pdf")
    page = reader.pages[159]

    res = page.extract_text()
    pipeline = [
        remove_page_numbers,
        split_rows,
        competitors_number,
        competition_title,
        dates,
        group,
        country,
        region,
    ]
    for step in pipeline:
        res = step(res)
    df = res

    print(df.info())
    for col in df:
        print(df[col].head())

    for raw in df["Country"]:
        print(repr(raw))

    for raw in df["Group"]:
        print(repr(raw))
