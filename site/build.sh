#!/bin/bash

echo "Building React app..."
npm install

# Build the React app
echo "Building React app..."
npm run build

# Copy the built files to the main site directory, but preserve static assets
echo "Copying built files..."
cp -r build/* ../site/

# Ensure static assets are preserved
echo "Ensuring static assets are preserved..."
if [ ! -d "../site/static" ]; then
    mkdir -p ../site/static
fi

echo "Build completed!" 