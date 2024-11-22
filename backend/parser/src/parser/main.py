import re

import pandas as pd
from pypdf import PdfReader

from parser import util

PAGE_NUM = r"Стр\.\s*\d+\s*из\s*\d+"
SPORT_KIND = r"([А-Я\s]+)Основной\s+состав\s"
RESERVE_SECTION = r"Молодежный\s+\(резервный\)\s+состав"
ID = r"\d{10,}"
ROW_START = rf"\s(?={ID}\s)"
COMPETITORS_NUMBER = r"\s+(?=\d+$)"
COMPETITION_TITLE_BEFORE = r"\s+"
LOWERCASE_RUS = r"[а-я]"
COMPETITION_TITLE_AFTER = rf"(?=\s+{LOWERCASE_RUS})"
DATE = r"\d{2}\.\d{2}\.\d{4}"
DATES_BEFORE = rf"\s(?={DATE}\s{DATE})"
DATES_AFTER = rf"(?<={DATE}\s{DATE})\s"
UPPERCASE_RUS = r"[А-Я]"
DISCIPLINE_BEFORE = rf"\s?(?={UPPERCASE_RUS}|$)"
CITY_BEFORE = r"(?:,\s)|$"


def remove_page_numbers(s: str) -> str:
    return re.sub(PAGE_NUM, "", s)


def split_into_sport_kinds(s: str) -> list[str]:
    s = re.sub(RESERVE_SECTION, "", s)
    return list(map(str.strip, re.split(SPORT_KIND, s)))[1:]


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
        df["Title"].apply(lambda s: re.sub(r"\s+", " ", s))
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
    df = util.flat_apply(
        df,
        "Raw Group",
        lambda s: re.split(DISCIPLINE_BEFORE, s, maxsplit=1),
        columns=["Group", "Raw Discipline"],
    )
    df["Group"] = df["Group"].apply(lambda s: re.sub(r"\s+", " ", s))
    return df


def country(df: pd.DataFrame) -> pd.DataFrame:
    df["Raw Place"] = df["Raw Place"].apply(str.strip)
    df = df[df["Raw Place"] != "ПО НАЗНАЧЕНИЮ"]

    df = util.flat_apply(
        df,
        "Raw Place",
        lambda s: re.split("\n+", s, maxsplit=1),
        columns=["Country", "Raw Region"],
    )
    df["Country"] = df["Country"].apply(lambda s: re.sub(r"\s+", " ", s))
    return df


def region(df: pd.DataFrame) -> pd.DataFrame:
    df = util.flat_apply(
        df,
        "Raw Region",
        lambda s: re.split(CITY_BEFORE, s, maxsplit=1),
        columns=["Region", "City"],
    )
    df["Region"] = df["Region"].apply(lambda s: re.sub(r"\s+", " ", s))

    df["City Copy"] = df["City"].copy()
    df["City"] = df.apply(
        lambda row: row["Region"] if row["City"] == "" else row["City"],
        axis=1,
    )
    df["Region"] = df.apply(
        lambda row: "" if row["City Copy"] == "" else row["Region"],
        axis=1,
    )
    df = df.drop("City Copy", axis=1)
    return df


def main():
    reader = PdfReader("input.pdf")
    # pages = reader.pages[86:95]
    pages = reader.pages[:50]

    res = "\n".join(page.extract_text() for page in pages)
    pipeline = [
        remove_page_numbers,
        split_into_sport_kinds,
    ]
    for step in pipeline:
        res = step(res)

    parsed_sports = {}
    sports = {
        key: value for key, value in zip(res[::2], res[1::2], strict=True)
    }
    pipeline = [
        split_rows,
        competitors_number,
        competition_title,
        dates,
        group,
        country,
        # region,
    ]
    for key in sports:
        res = sports[key]
        for step in pipeline:
            res = step(res)
        parsed_sports[key] = res

    for key in parsed_sports:
        df = parsed_sports[key]
        print(key)
        print(df.head())
        print()
