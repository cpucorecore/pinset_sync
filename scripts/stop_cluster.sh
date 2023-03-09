pid=`ps -ef|grep ipfs-cluster-ser |grep -v grep |awk '{print $2}'`
if [ -z $pid ]; then
	echo already stoped
else
	echo pid $pid alive, to kill it
	kill -9 $pid
	echo kill finish
fi