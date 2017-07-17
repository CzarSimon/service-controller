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
  echo "Compiling sctl-minion for linux and darwin"
  xgo --targets=linux/amd64,darwin/amd64 github.com/CzarSimon/service-controller/sctl-minion

  darwin_file="sctl-minion-darwin-10.6-amd64"
  mv $darwin_file $sctl_data/executables/darwin/sctl-minion/sctl-minion
  rm $darwin_file

  linux_file="sctl-minion-linux-amd64"
  mv $linux_file $sctl_data/executables/linux/sctl-minion/sctl-minion
  rm $linux_file
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
