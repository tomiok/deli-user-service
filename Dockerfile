FROM golang

ARG app_env
ARG vault_token

ENV ENV $app_env

COPY . /go/src/github.com/pedidosya/search-federator
WORKDIR /go/src/github.com/pedidosya/search-federator

# added vendor services will need to be included here
#RUN go get -u github.com/golang/dep/...
#RUN dep ensure -vendor-only
RUN make build

RUN echo ${vault_token} > /root/.vault_token

#RUN echo $app_env
CMD ./search-federator -E ${ENV}

EXPOSE 8080
