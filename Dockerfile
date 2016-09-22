FROM alpine:edge
MAINTAINER Jessica Frazelle <jess@linux.com>

ENV PATH /go/bin:/usr/local/go/bin:$PATH
ENV GOPATH /go

RUN	apk add --no-cache \
	autoconf \
	automake \
	bash \
	build-base \
	ca-certificates \
	gengetopt \
	git \
	gcc \
	go \
	help2man \
	libc-dev \
	libgcc \
	libtool \
	openssl-dev \
	pcsc-lite-dev

RUN go get github.com/xlab/cgogen \
	&& go get github.com/golang/lint/golint

ENV YKPIV_VERSION 1.4.2

RUN git clone --depth 1 -b "yubico-piv-tool-${YKPIV_VERSION}" \
	https://github.com/Yubico/yubico-piv-tool.git /usr/src/yubico-piv-tool \
	&& ( \
		cd /usr/src/yubico-piv-tool \
		&& autoreconf --install \
		&& ./configure --prefix=/usr \
		&& make \
		&& make install \
	)

ENV PKG_CONFIG_PATH ${PKG_CONFIG_PATH}:/usr/src/yubico-piv-tool/ykpiv
COPY . /go/src/github.com/jfrazelle/ykpiv
WORKDIR /go/src/github.com/jfrazelle/ykpiv
