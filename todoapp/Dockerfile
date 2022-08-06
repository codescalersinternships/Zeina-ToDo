FROM node:14-alpine
WORKDIR /usr/src/app
# COPY rollup.config.js ./
# COPY package*.json ./
COPY . .
RUN npm install
# COPY ./src ./src
# COPY ./public ./public
# RUN npm run-script build
EXPOSE 5656
ENV HOST=0.0.0.0
CMD [ "npm", "run","dev" ]