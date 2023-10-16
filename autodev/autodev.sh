#!/bin/bash -e

# Display help information
function show_help {
    echo "Usage: deploy.sh [branch] [version]"
    echo ""
    echo "Deploy the current branch to the specified branch and create a package with the specified version."
    echo ""
    echo "Arguments:"
    echo "  branch    The target branch to deploy to. Default is 'dev'."
    echo "  version   The version number for the package. Default is '1.0.0' for 'dev', '2.0.0' for 'test', and '3.0.0' for 'master'."
    echo ""
    echo "Options:"
    echo "  -h, --help    Show this help message and exit."
}

# Function to handle errors and checkout back to original branch
function handle_error {
    echo "$(tput setaf 1)$1$(tput sgr0)" # '$(tput setaf 1)' set the text to blold and red and '$(tput sgr0)' reset the text formatting to normal
    git checkout "$ORIGINAL_BRANCH"
    exit 1
}

# Check if help is requested
if [[ "$1" == "-h" || "$1" == "--help" ]]; then
    show_help
    exit 0
fi

# Get current feature branch
FEATURE_BRANCH=$(git rev-parse --abbrev-ref HEAD)

# Set target branch and version
TARGET_BRANCH=${1:-dev}
case "$TARGET_BRANCH" in
    dev) VERSION=1.0.0 ;;
    test) VERSION=2.0.0 ;;
    master) VERSION=3.0.0 ;;
    *) handle_error "Invalid branch. Use dev, test, or master." ;;
esac
VERSION=${2:-$VERSION}

function package() {
    DIR_NAME=$(basename "$(pwd)")
    TAG=$( sh release.sh -v "$VERSION" | grep -o "TAG='[^']*'" | awk -F"'" '{print $2}' | head -n 1)
    if [ "$TAG" = "" ]; then
        handle_error "Failed to create package."
    fi

    echo "================================================================================"
    echo "Please copy the following content and provide it to the deployment team"
    echo "Service: $DIR_NAME, Tag: $TAG"
    echo "================================================================================"
}

# Print current branch
echo "Current branch is $FEATURE_BRANCH"

# Check if the feature branch is valid
#if [[ "$FEATURE_BRANCH" == "$TARGET_BRANCH" || "$FEATURE_BRANCH" == "dev" || "$FEATURE_BRANCH" == "test" || "$FEATURE_BRANCH" == "pd" ]]; then
#    handle_error "Cannot run deploy on dev, test, or pd branch."
#fi
if [[ "$FEATURE_BRANCH" == "$TARGET_BRANCH" || "$FEATURE_BRANCH" == "dev" || "$FEATURE_BRANCH" == "test" || "$FEATURE_BRANCH" == "pd" ]]; then
    package
    exit 0
fi


# Save original branch
ORIGINAL_BRANCH=$(git rev-parse --abbrev-ref HEAD)

# Pull Current branch
git pull origin "$ORIGINAL_BRANCH"|| handle_error "Failed to update branch $ORIGINAL_BRANCH."

# Switch to target branch
git checkout "$TARGET_BRANCH" || handle_error "Failed to switch to target branch $TARGET_BRANCH."

# Pull latest changes
git pull origin "$TARGET_BRANCH" || handle_error "Failed to pull changes from $TARGET_BRANCH."

# Merge feature branch
git merge --no-edit "$FEATURE_BRANCH" || handle_error "Failed to merge $FEATURE_BRANCH into $TARGET_BRANCH."

# Push changes to remote repository
git push origin "$TARGET_BRANCH" || handle_error "Failed to push changes to $TARGET_BRANCH."

# Create package
package

# Switch back to original branch
git checkout "$ORIGINAL_BRANCH" || handle_error "Failed to switch back to the original branch."