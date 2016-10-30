FROM golang:1.7.1-alpine

COPY . $GOPATH/src/github.com/Berlin-opendb-hack/mdbga
WORKDIR $GOPATH/src/github.com/Berlin-opendb-hack/mdbga

ENV BANK_SCHEME = ''
ENV BANK_HOST = ''
ENV BANK_PATH = ''
ENV TRANSFER_FEE_PERCENTAGE=''
ENV BLOCKCHAIN_SCHEME=''
ENV BLOCKCHAIN_HOST=''
ENV BLOCKCHAIN_PATH=''

johanness-macbook-pro-2.local

RUN go build -v && \
    mkdir -p /mdbga && \
    cp mdbga /mdbga/mdbga && \
    rm -R $GOPATH/src/github.com/Berlin-opendb-hack/mdbga
EXPOSE 8881
WORKDIR /mdbga
CMD ["/mdbga/mdbga"]


