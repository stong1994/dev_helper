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
	echo "  -v, --version    Show version."
}

# Function to handle errors and checkout back to original branch
function handle_error {
	echo "$(tput setaf 1)$1$(tput sgr0)" # '$(tput setaf 1)' set the text to blold and red and '$(tput sgr0)' reset the text formatting to normal
	git checkout "$FEATURE_BRANCH"
	exit 1
}

# if has code uncommited, commit it with aicommit if exists
function commit_uncommitted_changes {
	# Check if there are uncommitted changes
	# git diff-index --quiet HEAD -- | handle_error "execute command failed: git diff --cached --quiet"

	#git diff --cached --quiet | handle_error "execute command failed: git diff --cached --quiet"
	#echo $status | handle_error "execute command failed: echo $status"
	#changes=$? | handle_error "execute command failed: echo $?"
	#if [ "$changes" != 0 ]; then
	if ! git diff-index --quiet HEAD --; then
		echo "code has uncommited. try to commit"
		# Check if aicommit exists
		if command -v aicommit &>/dev/null; then
			# Use aicommit to generate commit command
			commit_command=$(aicommit --platform lingyi --quiet)
			# Execute the generated command
			echo "will autocommit: $commit_command"
			eval "$commit_command"
		else
			# If aicommit does not exist, report it to user
			handle_error "aicommit hasn't been installed. Please commit changes before running deploy or install aicommit."
		fi
	else
		echo "code has commited."
	fi
}

# Check if help is requested
if [[ "$1" == "-h" || "$1" == "--help" ]]; then
	show_help
	exit 0
fi

if [[ "$1" == "-v" || "$1" == "--version" ]]; then
	echo "1.0.1"
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
	TAG=$(sh release.sh -v "$VERSION" | grep -o "TAG='[^']*'" | awk -F"'" '{print $2}' | head -n 1)
	if [ "$TAG" = "" ]; then
		handle_error "Failed to create package."
	fi

	echo "================================================================================"
	echo "Please copy the following content and provide it to the deployment team"
	echo "Service: $DIR_NAME, Tag: $TAG"
	echo "================================================================================"
}

commit_uncommitted_changes

# Check if the feature branch exists on the remote
if git ls-remote --heads origin | grep -q "$FEATURE_BRANCH"; then
	# Pull Current branch
	git pull origin "$FEATURE_BRANCH" || handle_error "Failed to update branch $FEATURE_BRANCH."
fi
# update the remote featuere branch
git push origin "$FEATURE_BRANCH"

# Check if the feature branch is valid
#if [[ "$FEATURE_BRANCH" == "$TARGET_BRANCH" || "$FEATURE_BRANCH" == "dev" || "$FEATURE_BRANCH" == "test" || "$FEATURE_BRANCH" == "pd" ]]; then
#    handle_error "Cannot run deploy on dev, test, or pd branch."
#fi
#if [[ "$FEATURE_BRANCH" == "$TARGET_BRANCH" || "$FEATURE_BRANCH" == "dev" || "$FEATURE_BRANCH" == "test" || "$FEATURE_BRANCH" == "pd" ]]; then
if [[ "$FEATURE_BRANCH" == "$TARGET_BRANCH" && "$TARGET_BRANCH" == "dev" ]]; then
	package
	exit 0
fi

# Check if the target branch exists on the local and switch to it
if ! git ls-remote --heads origin | grep -q "$TARGET_BRANCH"; then
	echo "The target branch $TARGET_BRANCH does not exist on the local. Creating it..."
	git checkout -b "$TARGET_BRANCH" || handle_error "Failed to switch to target branch $TARGET_BRANCH."
else
	git checkout "$TARGET_BRANCH" || handle_error "Failed to switch to target branch $TARGET_BRANCH."
fi

# Pull latest changes
git pull origin "$TARGET_BRANCH" || handle_error "Failed to pull changes from $TARGET_BRANCH."

# Merge feature branch
git merge --no-edit "$FEATURE_BRANCH" || handle_error "Failed to merge $FEATURE_BRANCH into $TARGET_BRANCH."

# Push changes to remote repository
git push origin "$TARGET_BRANCH" || handle_error "Failed to push changes to $TARGET_BRANCH."

# Create package
package

# Switch back to original branch
git checkout "$FEATURE_BRANCH" || handle_error "Failed to switch back to the original branch."
