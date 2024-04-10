# versions: https://www.python.org/doc/versions/
FROM python:latest

WORKDIR /app

# versions: https://pypi.org/project/poetry/#history
RUN pip install "poetry==1.8.2"

ENV PATH="/app/.venv/bin:$PATH"

COPY pyproject.toml poetry.lock* /app/
RUN poetry config virtualenvs.in-project true \
    && poetry install

COPY ./ /app
