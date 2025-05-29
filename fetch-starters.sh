#!/bin/bash

# Check if both organization and repository name are provided
if [ $# -ne 2 ]; then
    echo "Usage: $0 <github-org> <repo-name>"
    exit 1
fi

ORG=$1
REPO=$2

# Create directory if it doesn't exist
if [ ! -d "$REPO" ]; then
    mkdir "$REPO"
fi

# Clone the repository
echo "Cloning repository: $ORG/$REPO"
git clone "https://github.com/$ORG/$REPO.git" temp_repo

# Move files from temp directory to target directory, excluding students directory, .git, and .gitignore
echo "Moving files to $REPO directory..."
find temp_repo -mindepth 1 -maxdepth 1 -not -name ".git" -not -name ".gitignore" -not -name "students" -exec cp -r {} "$REPO/" \;

# Clean up temporary directory
rm -rf temp_repo

echo "Done! Files have been fetched to the $REPO directory."