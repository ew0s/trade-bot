#!/bin/bash

set -o errexit

migrations=$(find ./migrations/trade-bot -name '__*.sql')

for name in $migrations
do
	new_name=$(echo "$name" | sed "s/__/$(date "+%Y%m%d%H%M%S")_/")

	echo "moving $name -> $new_name"

	git mv "$name" "$new_name"
done
