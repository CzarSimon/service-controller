sctl_cli="github.com/CzarSimon/service-controller/sctl"
sctl_api="github.com/CzarSimon/service-controller/sctl-api"
sctl_minion="github.com/CzarSimon/service-controller/sctl-minion"
sctl_data="$GOPATH/bin/sctl-data"
start_path=$PWD

build() {
  component=$1
  echo "Fetching $component"
  go get -u $component
  echo "Building $component"
  go install $component
}

setup_sctl_data() {
  echo "Setting up sctl-data"
  mkdir $sctl_data
  mkdir $sctl_data/executables
  rsync -a $GOPATH/src/$sctl_cli/sctl-data/executables/ $sctl_data/executables
}

build_minion () {
  echo "Fetching sctl-minion"
  go get -u $sctl_minion
  echo "Building sctl-minion"
  cd $GOPATH/src/$sctl_minion

  GOOS=darwin GOARCH=amd64 go build
  mv sctl-minion $GOPATH/bin/sctl-data/executables/darwin/sctl-minion

  GOOS=linux GOARCH=amd64 go build
  mv sctl-minion $GOPATH/bin/sctl-data/executables/linux/sctl-minion
  cd $start_path
}

start_api() {
  echo "starting api-server"
  os_folder="$sctl_data/executables/$1"
  cd $os_folder
  python format-plist.py sctl-api
  cp sctl-api/com.api.sctl.plist $HOME/Library/LaunchAgents/
  launchctl unload $HOME/Library/LaunchAgents/com.api.sctl.plist
  launchctl load $HOME/Library/LaunchAgents/com.api.sctl.plist
  cd $start_path
}

format_minion_plist() {
  echo "formating minion plist"
  cd $sctl_data/executables/darwin
  python format-plist.py sctl-minion
  cd $start_path
}

build $sctl_cli
build $sctl_api
setup_sctl_data
build_minion
start_api "darwin"
format_minion_plist
