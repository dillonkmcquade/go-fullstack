FROM fedora:latest
SHELL ["/bin/sh", "-c"]
COPY . /go-fullstack
RUN dnf install util-linux make -y
WORKDIR /go-fullstack
RUN ["make", "prep"]
CMD ["/bin/sh"]
