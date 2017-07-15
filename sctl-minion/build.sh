build_and_move () {
  # os = $1
  # architecture = $2
  GOOS=$1 GOARCH=$2 go build
  mv sctl-minion ../sctl/executables/$1/sctl-minion
}

rm sctl-minion
rm token-db

build_and_move "darwin" "amd64"
build_and_move "linux" "amd64"
