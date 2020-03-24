#!/bin/bash

repo_uri=$1
if [[ -z "$repo_uri" ]]; then
  echo "need a repository to push to"
  exit 1
fi

echo "building executable"
bash go-executable-build.bash basketball ./main

echo "building image"
docker build --rm -t "basketball" --build-arg ex_path=linux/amd64 --build-arg env=prod .

echo "logging in to aws"
$(aws ecr get-login --no-include-email --region us-west-2 --profile tharivol)

echo "tagging image"
docker tag basketball:latest $repo_uri

echo "pushing image"
docker push $repo_uri:latest

echo "deploying"
aws ecs update-service --cluster sandbox-cluster --service basketball-service --force-new-deployment --profile tharivol

echo "Finito!"