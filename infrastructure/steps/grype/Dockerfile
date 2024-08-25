FROM BASE-IMAGE

ENV MAKE_VERSION=4.4.1-r1

RUN apk update && \
    apk add --no-cache \
    make=${MAKE_VERSION}

RUN curl -sSfL https://go.dev/dl/go1.23.0.linux-amd64.tar.gz -o /tmp/go.tar.gz && \
    tar -C /usr/local -xzf /tmp/go.tar.gz && \
    rm /tmp/go.tar.gz && \
    export PATH=$PATH:/usr/local/go/bin

RUN curl -sSfL https://raw.githubusercontent.com/anchore/grype/main/install.sh | sh -s -- -b /usr/local/bin

ENV PATH="/usr/local/go/bin:$PATH"

RUN mkdir -p /greport

COPY greport/* /greport/.

WORKDIR /greport

CMD ["make"]