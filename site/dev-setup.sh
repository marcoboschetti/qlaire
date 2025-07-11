#!/bin/bash

echo "🚀 Setting up Qlaire React Frontend Development Environment"

# Check if Node.js is installed
if ! command -v node &> /dev/null; then
    echo "❌ Node.js is not installed. Please install Node.js 16+ first."
    exit 1
fi

# Check if npm is installed
if ! command -v npm &> /dev/null; then
    echo "❌ npm is not installed. Please install npm first."
    exit 1
fi

echo "✅ Node.js and npm found"

# Install dependencies
echo "📦 Installing dependencies..."
npm install

if [ $? -eq 0 ]; then
    echo "✅ Dependencies installed successfully"
else
    echo "❌ Failed to install dependencies"
    exit 1
fi

echo ""
echo "🎉 Development environment ready!"
echo ""
echo "To start development:"
echo "  npm start"
echo ""
echo "To build for production:"
echo "  npm run build"
echo "  ./build.sh"
echo ""
echo "The React app will run on http://localhost:3000"
echo "Make sure your Go backend is running on http://localhost:8080" 