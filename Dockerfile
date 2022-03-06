FROM node:12

# Add Tini
ENV TINI_VERSION v0.15.0
ADD https://github.com/krallin/tini/releases/download/${TINI_VERSION}/tini /tini
RUN chmod +x /tini
ENTRYPOINT ["/tini", "--"]

# application config
RUN mkdir /app

WORKDIR /app
COPY ./app/package*.json /app/
RUN yarn install
COPY ./app /app

CMD [ "yarn", "start" ]
