FROM golang:1.18 as Builder

COPY . /src
WORKDIR /src

RUN GOPROXY=https://goproxy.cn && make build

FROM debian:stable-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
        ca-certificates  \
        netbase \
        && rm -rf /var/lib/apt/lists/ \
        && apt-get autoremove -y && apt-get autoclean -y


COPY --from=Builder /src/bin /app
COPY --from=Builder /src/conf.yml /app/

WORKDIR /app

EXPOSE 8089
EXPOSE 8085

CMD ["./easycar"]