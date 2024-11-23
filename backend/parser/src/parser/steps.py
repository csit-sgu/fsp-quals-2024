import re

import pandas as pd

from parser import util
from parser.log import logger

PAGE_NUM = re.compile(r"Стр\.\s*\d+\s*из\s*\d+")
SPORT_KIND = re.compile(r"([а-яА-Я\-\s]+)Основной\s+состав\s")
RESERVE_SECTION = re.compile(r"Молодежный\s+\(резервный\)\s+состав")
CODE = r"\d{10,}"
ROW_START = re.compile(rf"\s(?={CODE}\s)")
COMPETITORS_NUMBER = re.compile(r"\s+(?=\d+$)")
COMPETITION_TITLE_BEFORE = re.compile(r"\s+")
COMPETITION_TITLE_AFTER = re.compile(r"(?=\s+[а-я])")
DATE = r"\d{2}\.\d{2}\.\d{4}"
DATES_BEFORE = re.compile(rf"\s(?={DATE}\s{DATE})")
DATES_AFTER = re.compile(rf"(?<={DATE}\s{DATE})\s")
DISCIPLINE_BEFORE = re.compile(r"\s?(?=[А-Я]|$)")
REGION_NAME = re.compile(r"([^,]+),\s+")

AGE_PATTERNS = [
    (re.compile(r"от\s+(\d+)\s+лет"), lambda m: (int(m.group(1)), 0)),
    (re.compile(r"до\s+(\d+)\s+лет"), lambda m: (0, int(m.group(1)))),
    (
        re.compile(r"(\d+)\s+-\s+(\d+)\s+лет"),
        lambda m: (int(m.group(1)), int(m.group(2))),
    ),
    (re.compile(r"(\d+)\s+лет"), lambda m: (int(m.group(1)), int(m.group(1)))),
]

GROUPS = {
    "мужчины": "male",
    "юноши": "male",
    "юниоры": "male",
    "мальчики": "male",
    "женщины": "female",
    "девушки": "female",
    "юниорки": "female",
    "девочки": "female",
}


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
        columns=["Code", "Raw"],
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
        (df["Code"], cleaned_title, cleaned_rest, df["Competitors"]),
        axis=1,
    )


def dates(df: pd.DataFrame) -> pd.DataFrame:
    df = util.flat_apply(
        df,
        "Raw",
        lambda s: re.split(DATES_BEFORE, s, maxsplit=1),
        columns=["Raw Group And Discipline", "Raw"],
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


def extract_group(df: pd.DataFrame) -> pd.DataFrame:
    df = util.flat_apply(
        df,
        "Raw Group And Discipline",
        lambda s: re.split(DISCIPLINE_BEFORE, s, maxsplit=1),
        columns=["Raw Group", "Raw Discipline"],
    )
    df["Raw Group"] = df["Raw Group"].apply(lambda s: re.sub(r"\s+", " ", s))
    return df


def age_range(token: str) -> tuple[int, int] | None:
    for pattern, func in AGE_PATTERNS:
        match = pattern.search(token)
        if match:
            return func(match)
    return None


def append_age_group(
    entries: list[tuple[str, str, int, int]],
    group_buffer: list[str],
    ages_buffer: list[tuple[int, int]],
) -> list[tuple[str, str, int, int]]:
    # if this is a group like "женщины", "мужчины, женщины"
    if (
        len(ages_buffer) == 0
        and (len(group_buffer) == 1 or len(group_buffer) == 2)
        and ("женщины" in group_buffer or "мужчины" in group_buffer)
    ):
        ages_buffer = [(18, 0)]

    return entries + [
        (group, GROUPS[group], *ages)
        for group in group_buffer
        for ages in ages_buffer
    ]


def parse_restrictions(line: str) -> list[tuple[str, str, int, int]]:
    tokens = [token.strip() for token in line.split(",")]
    entries = []
    group_buffer = []
    ages_buffer = []
    for token in tokens:
        group_found = None
        for group in GROUPS:
            group_found = group_found or re.match(rf"({group})", token)
        if group_found is not None:
            if len(ages_buffer) > 0:
                # that's a new age group, unload what we have
                entries = append_age_group(entries, group_buffer, ages_buffer)
                ages_buffer = []
                group_buffer = []

            group = group_found.group(1)
            group_buffer.append(group)
            # also check if there is an age restriction
            ages = age_range(token)
            if ages is not None:
                ages_buffer.append(ages)
        else:  # this is not a group, probably an age range
            ages = age_range(token)
            if ages is not None:
                ages_buffer.append(ages)
            else:
                logger.debug(
                    f"Failed to parse: {repr(line)}. At token: {token}"
                )

    if len(group_buffer) > 0:
        entries = append_age_group(entries, group_buffer, ages_buffer)

    if len(entries) == 0:
        logger.debug(f"Failed to parse: {repr(line)}")
        entries = append_age_group(entries, ["женщины", "мужчины"], [])

    return entries


def parse_group(df: pd.DataFrame) -> pd.DataFrame:
    df["Group"] = df["Raw Group"].apply(parse_restrictions)
    return df


def country(df: pd.DataFrame) -> pd.DataFrame:
    df["Raw Place"] = df["Raw Place"].apply(str.strip)
    df = df[df["Raw Place"] != "ПО НАЗНАЧЕНИЮ"]

    df = util.flat_apply(
        df,
        "Raw Place",
        lambda s: re.split(r"\n+", s, maxsplit=1),
        columns=["Country", "Raw Region"],
    )
    df = df.dropna()
    df["Country"] = df["Country"].apply(lambda s: re.sub(r"\s+", " ", s))
    return df


def parse_region(s: str) -> list[list[str]]:
    lines = filter(lambda x: x != "", map(str.strip, re.split(r"\n+", s)))

    res = []
    for line in lines:
        parts = list(
            filter(lambda x: x != "", re.split(REGION_NAME, line, maxsplit=1))
        )
        if len(parts) == 0:
            res.append(["", line])
        elif len(parts) == 1:
            res.append(["", parts[0]])
        elif len(parts) == 2:
            res.append(parts)
        else:
            res.append(parts[:2])

    return res


def locality(df: pd.DataFrame) -> pd.DataFrame:
    df["Locality"] = df["Raw Region"].apply(parse_region)
    df = df.drop("Raw Region", axis=1)
    return df
