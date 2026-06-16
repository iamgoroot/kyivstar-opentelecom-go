#!/usr/bin/env bash
set -eu

cd "$(dirname "$0")"

for dir in */; do
	name=${dir%/}
	printf '=== %s ===\n' "$name"
	if go run -C "$dir" . 2>&1; then
		echo "PASS"
	else
		echo "FAIL (set KS_CLIENT_ID, KS_CLIENT_SECRET, KS_SERVER_URL?)"
	fi
done
