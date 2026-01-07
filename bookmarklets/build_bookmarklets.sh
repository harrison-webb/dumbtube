#!/bin/bash

for file in *.js; do
    echo "minifying $file..."
    # get filename without .js extension
    name="${file%.js}"

    # Read the file, remove newlines and extra spaces
    # tr -d '\n' -> Remove newlines
    # sed 's/  */ /g') -> Squeeze multiple spaces
    minified=$(cat "$file" | tr -d '\n' | sed 's/  */ /g')

    output="${name}.txt"

    echo "javascript:${minified}" > "$output"

    echo "Created: $output"
done

echo ""
echo "Done. Copy the contents of each .txt file into the url of a dedicated bookmark"
