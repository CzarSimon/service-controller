sctl_cli="github.com/CzarSimon/service-controller/sctl"
sctl_api="github.com/CzarSimon/service-controller/sctl-api"
sctl_minion="github.com/CzarSimon/service-controller/sctl-minion"
sctl_data="$GOPATH/bin/sctl-data"

build() {
  echo "Building sctl and sctl-api"
  go get -u github.com/CzarSimon/service-controller
  go install $sctl_cli
  go install $sctl_api
}

setup_sctl_data() {
  echo "Setting upd sctl-data"
  mkdir $sctl_data
  mkdir $sctl_data/executables
  setup_os_exec_files "darwin"
  setup_os_exec_files "linux"
}

setup_os_exec_files() {
  os=$1
  echo "setting up executables for $os minion"
  mkdir $sctl_data/executables/$os
  mkdir $sctl_data/executables/$os/sctl-minion
  source="$GOPATH/src/$sctl_cli/sctl-data/executables/$os/sctl-minion/"
  target="$sctl_data/executables/$os/"
  rm $target/sctl-minion
  rsync -a $source $target
}

build_minion () {
  echo "Building sctl-minion"
  cd $GOPATH/src/$sctl_minion

  GOOS=darwin GOARCH=amd64 go build
  mv sctl-minion $GOPATH/bin/sctl-data/executables/darwin/sctl-minion

  GOOS=linux GOARCH=amd64 go build
  mv sctl-minion $GOPATH/bin/sctl-data/executables/linux/sctl-minion
  cd start_path
}

start_path=$PWD
build
setup_sctl_data
build_minion
