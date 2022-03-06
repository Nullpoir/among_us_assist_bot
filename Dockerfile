FROM node:12

COPY package*.json ./

RUN yarn install

COPY . .

CMD [ "yarn", "start" ]


