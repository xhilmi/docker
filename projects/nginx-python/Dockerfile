# base image
FROM alpine:latest

# argument
ARG DEBIAN_FRONTEND=noninteractive

# install nginx + python3 into image
RUN apk add --update --no-cache python3 \
    nginx curl wget nano vim git bash fish && \
    ln -sf python3 /usr/bin/python && \
    python3 -m ensurepip && \
    pip3 install --no-cache --upgrade pip setuptools

# create environment variable
ENV nginx_vhost=/etc/nginx/conf.d/default.conf
ENV nginx_conf=/etc/nginx/nginx.conf

# copy application files/folder into image
COPY ./nginx/default ${nginx_vhost}
COPY ./nginx/nginx.conf ${nginx_conf}
COPY ./app /opt/app
COPY ./start.sh /start.sh
COPY . /home

# install required python module
RUN pip install -r /opt/app/requirements.txt

# add execute permission
RUN chmod +x /start.sh

# expose ports 80 + 443
EXPOSE 80 443

# start the application using start.sh
CMD [ "./start.sh" ]