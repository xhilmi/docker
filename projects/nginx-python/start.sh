#!/bin/sh

cd /opt/app && gunicorn -c gunicorn.conf.py wsgi --daemon
nginx -g "daemon off;"
