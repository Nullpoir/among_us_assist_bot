#!/bin/bash -eu
docker ps -a | grep among_us_assist_bot | awk '{print $1}' | xargs docker rm
docker volume ls | grep among_us_assist_bot | awk '{print $2}' | xargs docker volume rm
docker image ls | grep among_us_assist_bot | awk '{print $1}' | xargs docker image rm
