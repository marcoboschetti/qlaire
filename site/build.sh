#!/bin/bash


echo "Building React app..."
npm install

# Build the React app
echo "Building React app..."
npm run build

# Copy the built files to the main site directory
echo "Copying built files..."
cp -r build/* ../site/

echo "Build completed!" 