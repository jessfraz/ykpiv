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
	file \
	flex \
	gengetopt \
	git \
	gcc \
	go \
	help2man \
	libc-dev \
	libgcc \
	libtool \
	libusb-dev \
	openssl-dev \
	pcsc-lite-dev

RUN go get github.com/xlab/cgogen \
	&& go get github.com/golang/lint/golint

#Install pcsclite
RUN git clone https://github.com/LudovicRousseau/PCSC.git /usr/src/pcsc \
	&& ( \
		cd /usr/src/pcsc \
		&& ./bootstrap \
		&& ./configure --prefix=/usr --disable-libsystemd \
		&& make \
		&& make install \
	)

# Install ccid
RUN git clone https://github.com/LudovicRousseau/CCID.git /usr/src/ccid \
	&& ( \
		cd /usr/src/ccid \
		&& git submodule init \
		&& git submodule update \
		&& ./bootstrap \
		&& ./configure --prefix=/usr \
		&& make \
		&& make install \
	)

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

ENV PKG_CONFIG_PATH /usr/lib/pkgconfig:/usr/src/yubico-piv-tool/ykpiv
COPY . /go/src/github.com/jessfraz/ykpiv
WORKDIR /go/src/github.com/jessfraz/ykpiv
