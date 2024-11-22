from pydantic_settings import BaseSettings, SettingsConfigDict


class Settings(BaseSettings):
    ch_host: str
    ch_port: int
    ch_user: str
    ch_password: str

    remote_file: str

    log_level: str = "INFO"

    model_config = SettingsConfigDict(env_prefix="parser_", env_file=".env")


settings = Settings()
