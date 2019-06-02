#!/bin/bash
docker-compose down && docker build -t hdfs . &&  docker-compose up -d
