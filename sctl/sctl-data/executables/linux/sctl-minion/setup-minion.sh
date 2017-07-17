setup_service() {
  exec_folder="/usr/local/sbin/sctl-minion"
  mv sctl-minion.service /etc/systemd/system/sctl-minion.service

  mkdir $exec_folder
  mv sctl-minion $exec_folder/sctl-minion
  mv token-db $exec_folder/token-db
}

start_service() {
  ufw allow 9105
  systemctl enable sctl-minion.service
  systemctl start sctl-minion.service
}

open_swarm_ports() {
  ufw allow 2377/tcp
  ufw allow 7946/tcp
  ufw allow 7946/tcp
  ufw allow 4789/udp
}

# Setup + start tasks
setup_service
start_service
open_swarm_ports

# Removing setup-script
rm setup-minion.sh
