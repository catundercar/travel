host host.docker.internal|head -1|awk '{print $4}'|xargs -i sudo sed -i "s/host.docker.internal/{}/g" /etc/proxychains4.conf
