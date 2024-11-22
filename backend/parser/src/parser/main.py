from numpy.ma.core import diff
import re

import pandas as pd
from pypdf import PdfReader

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
    comp_id = pd.DataFrame(
        df["Raw"]
        .apply(lambda s: re.split(COMPETITION_TITLE_BEFORE, s, maxsplit=1))
        .to_list(),
        columns=["ID", "Raw"],
    )
    df = df.drop("Raw", axis=1)
    df = pd.concat((comp_id, df), axis=1)

    comp_name = pd.DataFrame(
        df["Raw"]
        .apply(lambda s: re.split(COMPETITION_TITLE_AFTER, s, maxsplit=1))
        .to_list(),
        columns=["Title", "Raw"],
    )
    df = df.drop("Raw", axis=1)
    df = pd.concat(
        (df["ID"], comp_name, df["Competitors"]),
        axis=1,
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
    dates_before = pd.DataFrame(
        df["Raw"]
        .apply(lambda s: re.split(DATES_BEFORE, s, maxsplit=1))
        .to_list(),
        columns=["Raw Group", "Raw"],
    )
    df = df.drop("Raw", axis=1)
    df = pd.concat((df, dates_before), axis=1)
    
    dates_after = pd.DataFrame(
        df["Raw"]
        .apply(lambda s: re.split(DATES_AFTER, s, maxsplit=1))
        .to_list(),
        columns=["Raw Dates", "Raw"],
    )
    df = df.drop("Raw", axis=1)
    df = pd.concat((df, dates_after), axis=1)

    clean_dates = pd.DataFrame(
        df["Raw Dates"]
        .apply(str.split)
        .to_list(),
        columns=["Date Start", "Date End"]
    )
    df = df.drop("Raw Dates", axis=1)
    df = pd.concat((df, clean_dates), axis=1)

    return df    


def main():
    reader = PdfReader("input.pdf")
    page = reader.pages[111]

    text = page.extract_text()
    res = remove_page_numbers(text)
    res = split_rows(res)
    res = competitors_number(res)
    res = competition_title(res)
    res = dates(res)

    print(res.info())
    for col in res:
        print(res[col].head())

    for raw in res["Raw"]:
        print(repr(raw))
