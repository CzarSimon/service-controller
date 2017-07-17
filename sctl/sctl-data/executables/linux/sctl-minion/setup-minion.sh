exec_folder="/usr/local/sbin/sctl-minion"

mv sctl-minion.service /etc/systemd/sctl-minion.service
mv sctl-minion.service $init_folder/sctl-minion

mkdir $exec_folder
mv sctl-minion $exec_folder/sctl-minion
mv token-db $exec_folder/token-db

systemctl enable sctl-minion.service
systemctl start sctl-minion.service
