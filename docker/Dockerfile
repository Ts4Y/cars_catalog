FROM golang:1.22 AS build-env
ARG NAME
ENV BASE_PATH "/go/src/${NAME}"

RUN apt-get update -qq \
    && apt-get install -yq \
    libc6 \
    libc6-dev \ 
    libaio1 \ 
    libaio-dev

COPY src ${BASE_PATH}
WORKDIR ${BASE_PATH}
ARG NAME

RUN --mount=type=cache,id=${NAME},target=/go/pkg/mod \
    --mount=type=cache,id=${NAME}_cache,target=/root/.cache/go-build \ 
    cd cmd && \
    go build -o ${NAME} .

FROM ubuntu:22.04
ARG NAME
ENV BASE_PATH "/go/src/${NAME}"

ENV TZ="Asia/Tashkent"
ENV DEBIAN_FRONTEND=noninteractive
RUN  apt-get update -qq \
    && apt-get install -yq \
    tzdata \
    ca-certificates && \ 
    ln -fs /usr/share/zoneinfo/Asia/Tashkent /etc/localtime && \
    dpkg-reconfigure -f noninteractive tzdata && \
    rm -rf /var/lib/apt/lists/*;

ENV BIN_PATH "/go/src/${NAME}"
ENV LOG_DIR "log"

COPY --from=build-env ${BASE_PATH}/cmd/${NAME} /app/${NAME}
COPY src/config /app/config
COPY src/vars.env /app/vars.env

WORKDIR /app
VOLUME ["/app/log"]

# EXPOSE 9000
ARG NAME
ENV NAME_BIN=$NAME

CMD /app/${NAME_BIN}