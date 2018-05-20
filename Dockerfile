#!/bin/bash

FROM golang

MAINTAINER Devign92 <banziha104@gmail.com>


ADD main /
CMD ["/main"]
EXPOSE 80