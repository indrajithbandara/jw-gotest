#!/bin/bash
#author cuijun 2016-11-22
#kill tomcat pid
pidservicelist=`ps -ef|grep apache-tomcat-7-service|grep -v "grep"|awk '{print $2}'`
pidcontrlllist=`ps -ef|grep apache-tomcat-7-controller|grep -v "grep"|awk '{print $2}'`
echo "the tomcat process is --->$pidcontrlllist"
echo "the tomcat process is --->$pidservicelist"

if [ "$pidcontrlllist" = "" ]
   then
       echo "no tomcat pid alive!"
else
  echo "tomcat Id list :$pidcontrlllist"
  kill -9 $pidcontrlllist
  #killall tomcat
  echo "service stop success"
fi

if [ "$pidservicelist" = "" ]
   then
       echo "no tomcat pid alive!"
else
  echo "tomcat Id list :$pidservicelist"
  kill -9 $pidservicelist
  #killall tomcat
  echo "service stop success"
fi
echo "starting controller and service tomcat services............."

cd /opt/apache-tomcat-7-controller/bin
./startup.sh &
cd /opt/apache-tomcat-7-service/bin
./startup.sh &

echo " all controller and service tomcat has been started."
