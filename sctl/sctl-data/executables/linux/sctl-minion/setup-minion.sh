exec_folder="/usr/local/sbin/sctl-minion"

mv sctl-minion.service /etc/systemd/system/sctl-minion.service

mkdir $exec_folder
mv sctl-minion $exec_folder/sctl-minion
mv token-db $exec_folder/token-db

ufw allow 9105

systemctl enable sctl-minion.service
systemctl start sctl-minion.service

rm setup-minion.sh
