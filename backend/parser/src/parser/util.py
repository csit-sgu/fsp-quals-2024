from typing import Callable

import pandas as pd


def flat_apply(
    df: pd.DataFrame,
    column: str,
    func: Callable,
    *,
    columns: list[str] | None = None,
) -> pd.DataFrame:
    new_df = pd.DataFrame(
        df[column].apply(func).to_list(),
        columns=columns,
    )
    df = df.drop(column, axis=1)
    return pd.concat((df, new_df), axis=1)
