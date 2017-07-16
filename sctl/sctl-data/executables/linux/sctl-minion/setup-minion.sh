init_folder = "/etc/init.d/sctl-minion"
exec_folder = "/usr/bin/sctl-minion"

mkdir $init_folder
mv ./sctl-minion.service $init_folder/

mkdir $exec_folder
mv ./sctl-minion $exec_folder/
mv ./token-db $exec_folder/

systemctl start sctl-minion
