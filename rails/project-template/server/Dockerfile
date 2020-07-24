FROM ruby:latest

COPY . .
WORKDIR /workspace/server

RUN apt-get install -y libpq-dev
RUN gem install bundler
RUN gem update

ENV BUNDLE_PATH=/bundle
