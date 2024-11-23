import re

import pandas as pd

from parser import util

PAGE_NUM = r"Стр\.\s*\d+\s*из\s*\d+"
SPORT_KIND = r"([а-яА-Я\-\s]+)Основной\s+состав\s"
RESERVE_SECTION = r"Молодежный\s+\(резервный\)\s+состав"
CODE = r"\d{10,}"
ROW_START = rf"\s(?={CODE}\s)"
COMPETITORS_NUMBER = r"\s+(?=\d+$)"
COMPETITION_TITLE_BEFORE = r"\s+"
COMPETITION_TITLE_AFTER = r"(?=\s+[а-я])"
DATE = r"\d{2}\.\d{2}\.\d{4}"
DATES_BEFORE = rf"\s(?={DATE}\s{DATE})"
DATES_AFTER = rf"(?<={DATE}\s{DATE})\s"
DISCIPLINE_BEFORE = r"\s?(?=[А-Я]|$)"
REGION_NAME = r"([^,]+),\s+"


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


MALE_SYNONYMS = [
    "мужчины",
    "юниоры",
    "мальчики",
    "юноши",
]

FEMALE_SYNONYMS = [
    "женщины",
    "юниорки",
    "девочки",
    "девушки",
]

age_patterns = [
    (re.compile(r"от\s*(\d+)\s*лет"), lambda m: (int(m.group(1)), 0)),
    (re.compile(r"до\s*(\d+)\s*лет"), lambda m: (0, int(m.group(1)))),
    (
        re.compile(r"(\d+)\s*-\s*(\d+)\s*лет"),
        lambda m: (int(m.group(1)), int(m.group(2))),
    ),
    (re.compile(r"(\d+)\s*лет"), lambda m: (int(m.group(1)), int(m.group(1)))),
]

group_names = {
    "мужчины": "male",
    "юноши": "male",
    "юниоры": "male",
    "мальчики": "male",
    "женщины": "female",
    "девушки": "female",
    "юниорки": "female",
    "девочки": "female",
}


def parse_restrictions(line: str):
    parsing_failed = False
    entries = []
    phrases = [phrase.strip() for phrase in line.split(",")]
    for phrase in phrases:
        phrase = phrase.strip()
        groups_in_phrase = []
        age_ranges_in_phrase = []

        for group in group_names.keys():
            if re.search(rf"\b{group}\b", phrase):
                groups_in_phrase.append(group)
                phrase = re.sub(rf"\b{group}\b", "", phrase)
        phrase = phrase.strip()

        for pattern, func in age_patterns:
            match = pattern.search(phrase)
            if match:
                lower_bound, upper_bound = func(match)
                age_ranges_in_phrase.append((lower_bound, upper_bound))
                phrase = pattern.sub("", phrase, count=1)
                phrase = phrase.strip()
                break

        if phrase:
            parsing_failed = True
            break

        if groups_in_phrase:
            if age_ranges_in_phrase:
                for group in groups_in_phrase:
                    entries.append(
                        (
                            group,
                            group_names[group],
                            age_ranges_in_phrase[0][0],
                            age_ranges_in_phrase[0][1],
                        )
                    )
            else:
                for group in groups_in_phrase:
                    if group in ("мужчины", "женщины"):
                        lower_bound = 18
                    else:
                        lower_bound = 0
                    entries.append((group, group_names[group], lower_bound, 0))
        else:
            if age_ranges_in_phrase:
                parsing_failed = True
                break
            else:
                parsing_failed = True
                break

    if parsing_failed:
        print("Failed:", line)
        return []
    return entries


def parse_group(df: pd.DataFrame) -> pd.DataFrame:
    df["Parsed Group"] = df["Group"].apply(parse_restrictions)
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
