
FROM node:latest as build

WORKDIR /app

COPY package*.json ./

RUN npm i

RUN npm install -g @angular/cli

COPY . .

RUN npm run build --configuration=production

FROM nginx:latest

COPY ./nginx.conf /etc/nginx/conf.d/default.conf

COPY --from=build /app/dist/frontend-angular/browser /usr/share/nginx/html

EXPOSE 80

