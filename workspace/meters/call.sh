#!/bin/bash

REPEAT=${1}

function call_post_article() {
  curl --location --request POST 'http://localhost:3000/articles' \
--header 'Content-Type: application/json' \
--data-raw '{
	"title" : "my title",
	"content" : "my content"
}'
}

function call_members() {
  curl -XGET http://localhost:3000/members
}

for (( c=1; c<=${REPEAT}; c++ ))
do
  if [ $(echo "${c} % 2" | bc) -eq 0 ]; then
    call_post_article
  else
    call_members
  fi
done