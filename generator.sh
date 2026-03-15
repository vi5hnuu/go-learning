#!/bin/sh

START=$1
MAX=$2
DIR=$3

if [ -z "$START" ] || [ -z "$MAX" ] || [ -z "$DIR" ]; then
  echo "Usage: $0 <start> <max_files> <directory>"
  exit 1
fi

mkdir -p "$DIR"

i=$START
END=$((START + MAX))
while [ $i -le $END ]
do
  FILE="$DIR/example-$i.go"

  cat > "$FILE" <<EOF
package main

import (
    "fmt"
)

func main() {
    fmt.Println("example-$i")
}
EOF

  i=$((i + 1))
done

echo "Generated $MAX files in $DIR"