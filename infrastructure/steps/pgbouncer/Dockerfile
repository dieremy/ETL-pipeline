FROM BASE-IMAGE

RUN apk add --no-cache \
    postgresql \
    pgbouncer

# Set environment variables (adjust as needed)
ARG PGBOUNCER_LISTEN_PORT
ARG PGBOUNCER_DATABASE
ARG PGBOUNCER_USER
ARG PGBOUNCER_PASSWORD
ARG PGBOUNCER_ALLOW_REMOTE
ARG PGBOUNCER_SERVER_HOST
ARG PGBOUNCER_SERVER_PORT
ARG PGBOUNCER_SERVER_DATABASE
ARG PGBOUNCER_SERVER_USER
ARG PGBOUNCER_SERVER_PASSWORD

# Set the default command
CMD ["pgbouncer"]