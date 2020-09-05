#!/usr/bin/env bash
openssl genrsa -out server.key 2048
openssl req -nodes -new -key server.key -subj "/CN=localhost" -out server.crt
openssl x509 -req -sha256 -days 365 -in server.crt -signkey server.key -out server.csr