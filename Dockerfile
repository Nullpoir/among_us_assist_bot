FROM node:12

COPY ./app /app
WORKDIR /app

RUN yarn install

CMD [ "yarn", "start" ]
