FROM python:3.11-alpine3.20 AS builder

ARG WD_NAME=/app
WORKDIR $WD_NAME

ENV UV_HOME="/opt/uv"
ENV UV_HTTP_TIMEOUT=300
ENV UV_UNMANAGED_INSTALL="$UV_HOME"
ENV UV_SYSTEM_PYTHON=true
ENV PATH="$UV_HOME:$PATH"

RUN apk update && apk add curl
RUN curl -LsSf https://astral.sh/uv/install.sh | sh

COPY pyproject.toml .
COPY uv.lock .
COPY .python-version .
RUN uv sync --frozen --no-dev

FROM python:3.11-alpine3.20 AS runtime

ARG WD_NAME=/app
WORKDIR $WD_NAME

ENV PYTHONPATH="src:.venv/lib/python3.11/site-packages"
COPY --from=builder $WD_NAME/.venv .venv
COPY src src

ENTRYPOINT ["python", "-c", "from parser.main import main; main()"]
