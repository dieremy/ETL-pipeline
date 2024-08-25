FROM BASE-IMAGE

ENV AIRFLOW_VERSION=2.10.0

ENV AIRFLOW__CORE__SQL_ALCHEMY_CONN="postgresql://${PSQL_USER}:${PSQL_PASSWORD}@postgres:5432/${PSQL_DATABASE}"
ENV AIRFLOW__CORE__DEFAULT_DAG_FOLDER="/opt/airflow/dags"

ARG PSQL_DATABASE
ARG PSQL_USER
ARG PSQL_PASSWORD
ARG PSQL_ROOT_PASSWORD

RUN apk update && \
    apk add --no-cache \
    python3-dev \
    g++ \
    make \
    gcc \
    musl-dev \
    postgresql-dev

RUN pip3 install --upgrade pip

RUN PYTHON_VERSION="$(python -c 'import sys; print(f"{sys.version_info.major}.{sys.version_info.minor}")')" \
    CONSTRAINT_URL="https://raw.githubusercontent.com/apache/airflow/constraints-${AIRFLOW_VERSION}/constraints-${PYTHON_VERSION}.txt" \
    pip install "apache-airflow[async,postgres,google,celery]==${AIRFLOW_VERSION}" --constraint "${CONSTRAINT_URL}"

RUN mkdir -p /opt/airflow/dag

COPY DAGs/* /opt/airflow/dag/.

CMD [ "python3", "/opt/airflow/dag/DAG_operator.py"]
