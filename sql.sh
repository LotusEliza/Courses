#!/bin/bash
PROJECT=site-dev
docker exec -it $PROJECT ./cockroach sql --insecure
