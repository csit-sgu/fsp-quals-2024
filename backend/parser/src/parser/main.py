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
UPPERCASE_RUS = r"[А-Я]"
DISCIPLINE_BEFORE = rf"\n?(?={UPPERCASE_RUS}|$)"


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
        df["Raw Dates"].apply(str.split).to_list(),
        columns=["Date Start", "Date End"],
    )
    df = df.drop("Raw Dates", axis=1)
    df = pd.concat((df, clean_dates), axis=1)

    return df

def group(df: pd.DataFrame) -> pd.DataFrame:
    groups = pd.DataFrame(
        df["Raw Group"]
        .apply(lambda s: re.split(DISCIPLINE_BEFORE, s, maxsplit=1))
        .to_list(),
        columns=["Group", "Raw Discipline"],
    )
    df = df.drop("Raw Group", axis=1)
    return pd.concat((df, groups), axis=1)

def main():
    reader = PdfReader("input.pdf")
    page = reader.pages[111]

    text = page.extract_text()
    df = remove_page_numbers(text)
    df = split_rows(df)
    df = competitors_number(df)
    df = competition_title(df)
    df = dates(df)
    df = group(df)

    print(df.info())
    for col in df:
        print(df[col].head())

    for raw in df["Raw"]:
        print(repr(raw))

    for raw in df["Group"]:
        print(repr(raw))
