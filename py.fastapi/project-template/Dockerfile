# versions: https://www.python.org/doc/versions/
FROM python:latest

WORKDIR /app

# versions: https://pypi.org/project/poetry/#history
RUN pip install "poetry==1.8.2"

# COPY pyproject.toml poetry.lock* /app/

# virtualenvs.create false: Poetryはデフォルトでプロジェクトごとに仮想環境を作成しますが、この設定はDockerコンテナ内で仮想環境を作成することを無効にする。
# RUN poetry config virtualenvs.create false \
#     && poetry install

COPY ./ /app
