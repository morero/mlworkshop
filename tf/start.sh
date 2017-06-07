#!/bin/bash
docker run -d -p 8888:8888 --name ml-workshop --env KERAS_BACKEND=tensorflow -v $PWD/notebooks:/home/jovyan/work ml-workshop
