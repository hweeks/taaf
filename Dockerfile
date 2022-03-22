FROM golang:rc-buster

EXPOSE 3005

ARG GIT_COMMIT=unspecified
ARG USER=app
ARG GROUP=app
ARG UID=1337
ARG GID=1337

LABEL git_commit=$GIT_COMMIT

USER root

RUN curl -fsSL https://deb.nodesource.com/setup_16.x | bash - &&\
  apt-get install -y nodejs && npm i -g yarn

RUN apt-get install -y python3 && apt-get install -y python3-pip

RUN groupadd --gid $GID $GROUP &&\
  useradd --uid $UID --gid $GID $USER

RUN mkdir -p /app/be

COPY . /app/be

WORKDIR /app/be

RUN chown -R root:root /app && cd /app/be && yarn install && \
  yarn deps

CMD yarn dev-inner