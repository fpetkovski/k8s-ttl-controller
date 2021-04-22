#!/usr/bin/env bash

set -o errexit
set -o pipefail

tmp_root=/tmp
envtest_root_dir=$tmp_root/envtest

k8s_version="${ENVTEST_K8S_VERSION:-1.19.2}"
goarch="$(go env GOARCH)"
goos="$(go env GOOS)"

if [[ ("$goos" != "linux" && "$goos" != "darwin") || ("$goos" == "darwin" && "$goarch" != "amd64") ]]; then
  echo "OS '$goos' with '$goarch' arch is not supported. Aborting." >&2
  return 1
fi

dest_dir="$(pwd)/test"
mkdir -p "${dest_dir}/bin"

# use the pre-existing version in the temporary folder if it matches our k8s version
if [[ -x "${dest_dir}/bin/kube-apiserver" ]]; then
  version=$("${dest_dir}"/bin/kube-apiserver --version)
  if [[ $version == *"${k8s_version}"* ]]; then
    echo "Using cached envtest tools from ${dest_dir}"
    return 0
  fi
fi

echo "fetching envtest tools@${k8s_version} (into '${dest_dir}')"
envtest_tools_archive_name="kubebuilder-tools-$k8s_version-$goos-$goarch.tar.gz"
envtest_tools_download_url="https://storage.googleapis.com/kubebuilder-tools/$envtest_tools_archive_name"

envtest_tools_archive_path="$tmp_root/$envtest_tools_archive_name"
if [ ! -f $envtest_tools_archive_path ]; then
  curl -sL ${envtest_tools_download_url} -o "$envtest_tools_archive_path"
fi

mkdir -p "${dest_dir}"
tar -C "${dest_dir}" --strip-components=1 -zvxf "$envtest_tools_archive_path"
