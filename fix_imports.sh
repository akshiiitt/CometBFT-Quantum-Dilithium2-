#!/bin/bash

# Script to update all import paths from old module to new module
# This fixes the module path mismatch issue

OLD_MODULE="github.com/cometbft/cometbft"
NEW_MODULE="github.com/akshiiitt/CometBFT-Quantum-Dilithium2-"

echo "Updating import paths from $OLD_MODULE to $NEW_MODULE"
echo "This may take a few minutes..."

# Find all .go files and update the import paths
find . -name "*.go" -type f -exec sed -i "s|$OLD_MODULE|$NEW_MODULE|g" {} +

echo "Import paths updated successfully!"
echo "Running go mod tidy to clean up dependencies..."

go mod tidy

echo "Done! All import paths have been updated."
