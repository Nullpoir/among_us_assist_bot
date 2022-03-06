FROM node:12

RUN mkdir /app

WORKDIR /app
COPY ./app/package*.json /app/
RUN yarn install
COPY ./app /app

CMD [ "yarn", "start" ]
