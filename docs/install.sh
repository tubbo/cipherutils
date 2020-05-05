#!/usr/bin/env bash
#
# Install cipherutils on the system. Note that this is _not_ a ZSH script,
# because ZSH may not actually be on the system yet. We can be pretty
# sure that Bash will be there, though, so that's why this script is
# written in Bash instead of ZSH.

# Find latest released version from GitHub API
repo="https://api.github.com/repos/tubbo/cipherutils/releases/latest"
version=$(curl --silent "$repo" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
tmpdir=$(mktemp -d)
sourcedir="cipherutils-${version:1}"
arch="amd64"

if [ -e /etc/distro ]; then
  os="linux"
else
  os="darwin"
fi

filename="cipherutils-$os-$arch.tar.gz"

# Ensure source code is in place
echo "Downloading cipherutils $version from GitHub..."
pushd "$tmpdir" > /dev/null 2>&1
curl -sLO "https://github.com/tubbo/cipherutils/releases/$version/downloads/$filename"

# Extract source code from tarball
tar -zxf "$filename"

# Install cipherutils from source to /usr/local and capture exit code
echo "Installing cipherutils..."
make install
code=$?

# Clean up the source directory
popd > /dev/null 2>&1
rm -rf "$tmpdir"

exit $code
