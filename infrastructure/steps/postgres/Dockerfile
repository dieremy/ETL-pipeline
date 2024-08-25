FROM BASE-IMAGE

ARG PSQL_DATABASE
ARG PSQL_USER
ARG PSQL_PASSWORD
ARG PSQL_ROOT_PASSWORD

RUN apk update && \
    apk add --no-cache \
    postgresql \
    postgresql-contrib

# Add data directory and set ownership
RUN mkdir -p /var/lib/postgresql/data

RUN chown postgres:postgres /var/lib/postgresql/data

ENTRYPOINT ["postgres", "-D", "/var/lib/postgresql/data"]