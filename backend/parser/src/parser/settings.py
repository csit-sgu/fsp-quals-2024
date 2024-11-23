from pydantic_settings import BaseSettings, SettingsConfigDict


class Settings(BaseSettings):
    ch_host: str
    ch_port: int
    ch_user: str
    ch_password: str

    unload_timeout: int = 300
    remote_file: str
    no_download: bool = False

    log_level: str = "INFO"

    model_config = SettingsConfigDict(env_prefix="parser_", env_file=".env")


settings = Settings()
