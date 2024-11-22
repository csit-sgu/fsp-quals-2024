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


def remove_page_numbers(s: str) -> str:
    return re.sub(PAGE_NUM, "", s)


def split_rows(s: str) -> pd.DataFrame:
    return pd.DataFrame(map(str.strip, re.split(ROW_START, s)))


def competitors_number(df: pd.DataFrame) -> pd.DataFrame:
    rows = (
        df[0]
        .apply(lambda s: re.split(COMPETITORS_NUMBER, s, maxsplit=1))
        .to_list()
    )
    return pd.DataFrame(rows)


def competition_title(df: pd.DataFrame) -> pd.DataFrame:
    comp_number = pd.DataFrame(
        df[0]
        .apply(lambda s: re.split(COMPETITION_TITLE_BEFORE, s, maxsplit=1))
        .to_list()
    )
    df = df.drop(0, axis=1)
    df = pd.concat((comp_number, df), axis=1, ignore_index=True)

    comp_name = pd.DataFrame(
        df[1]
        .apply(lambda s: re.split(COMPETITION_TITLE_AFTER, s, maxsplit=1))
        .to_list()
    )
    df = df.drop(1, axis=1)
    df = pd.concat((df[0], comp_name, df[2]), axis=1, ignore_index=True)

    cleaned_name = pd.DataFrame(
        df[1]
        .apply(lambda s: re.sub(r"\s+", " ", s.strip()))
        .to_list()
    )
    cleaned_rest = pd.DataFrame(
        df[2]
        .apply(lambda s: s.strip())
        .to_list()
    )

    df = df.drop(2, axis=1)
    return pd.concat((
        df[0],
        cleaned_name,
        cleaned_rest,
        df[3]
    ), axis=1, ignore_index=True)


def main():
    reader = PdfReader("input.pdf")
    page = reader.pages[1]

    text = page.extract_text()
    res = remove_page_numbers(text)
    res = split_rows(res)
    res = competitors_number(res)
    res = competition_title(res)

    for col in res:
        print(res[col].head())
