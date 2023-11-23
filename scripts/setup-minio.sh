#!/bin/bash

#This should not be run as a standalone script but only from the docker-compose

set -u

#install minio client
if [[ -z "$(which mc)" ]]; then
	wget -O mc https://dl.min.io/client/mc/release/linux-amd64/mc
	chmod +x mc
	mv mc /usr/local/bin/mc
fi

mc alias set local http://minio:9000 minioadmin minioadmin
mc mb --ignore-existing local/werr
mc admin user svcacct add --access-key "minioadminaccesskey" --secret-key "minioadminsecretkey" local minioadmin
