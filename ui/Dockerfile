FROM node:12.16.1-alpine As builder
RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app
COPY package.json package-lock.json /usr/src/app/
RUN npm install
COPY . /usr/src/app
COPY ../../shared/nginx.conf /usr/src/app/nginx.conf
RUN npm run build --prod

FROM nginx:1.15.8-alpine
COPY nginx.conf /etc/nginx/conf.d/default.conf
COPY /certificates/prod/portfolio-prod.pem /etc/ssl/certs/
COPY /certificates/prod/portfolio-prod.key /etc/ssl/private
COPY /certificates/prod/cloudflare.crt /etc/ssl/certs
COPY --from=builder /usr/src/app/dist/portfolio /usr/share/nginx/html/