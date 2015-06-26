KAFKADIR="kafka_2.9.2-0.8.2.1"
SERVERCONFIG="config/server.properties"
ZOOCONFIG="config/zookeeper.properties"

if [ "$1" == "-server-config" ]; then
	SERVERCONFIG=$2
fi

if [ "$3" == "-server-config" ]; then
	SERVERCONFIG=$4
fi

if [ "$1" == "-zoo-config" ]; then
	ZOOCONFIG=$2
fi

if [ "$3" == "-zoo-config" ]; then
	ZOOCONFIG=$4
fi

$KAFKADIR/bin/zookeeper-server-start.sh $KAFKADIR/$ZOOCONFIG &
$KAFKADIR/bin/kafka-server-start.sh $KAFKADIR/$SERVERCONFIG & 

