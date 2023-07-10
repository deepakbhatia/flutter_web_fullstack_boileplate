FROM ghcr.io/cirruslabs/flutter:3.7.12 AS build

WORKDIR /build

# download Flutter SDK from Flutter Github repo
#RUN git clone https://github.com/flutter/flutter.git /usr/local/flutter
#
## Set flutter environment path
#ENV PATH="/usr/local/flutter/bin:/usr/local/flutter/bin/cache/dart-sdk/bin:${PATH}"
#
## Run flutter doctor
#RUN flutter doctor

RUN flutter channel stable
RUN flutter upgrade
RUN flutter config --enable-web
RUN sudo apt-get update
RUN sudo apt-get install -y nginx
RUN sudo apt-get install -y chromium-browser
RUN sudo rm -rf /var/lib/apt/lists/*

COPY . .



