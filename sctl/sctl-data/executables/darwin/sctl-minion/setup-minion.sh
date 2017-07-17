deamon_name="com.minion.sctl.plist"

start_minion() {
  mv $deamon_name $HOME/Library/LaunchAgents/$deamon_name
  launchctl unload $HOME/Library/LaunchAgents/$deamon_name
  launchctl load $HOME/Library/LaunchAgents/$deamon_name
}

start_minion
rm setup-minion.sh
