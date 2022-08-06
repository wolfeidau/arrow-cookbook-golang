FROM golang:1.19

ENV DEBIAN_FRONTEND noninteractive

# install libarrow-dev to link against with CGO
RUN apt-get update -y -q && \
    apt-get install -y -q --no-install-recommends ca-certificates lsb-release wget && \
    wget https://apache.jfrog.io/artifactory/arrow/$(lsb_release --id --short | tr 'A-Z' 'a-z')/apache-arrow-apt-source-latest-$(lsb_release --codename --short).deb && \
    apt-get install -y -q --no-install-recommends ./apache-arrow-apt-source-latest-$(lsb_release --codename --short).deb && \
    apt-get update -y -q && \
    apt-get install -y -q --no-install-recommends \
        cmake \
        libarrow-dev && \
    apt-get clean
