#!/bin/bash
#author : cuijun 20161223
#pwd
seldate=`date -d "1 weeks ago" +%Y-%m-%d`
echo " the current date is ->$seldate"
cd /E/IDEidea/IdeaProjects
projectname=jw-source
echo " the project name is ->$projectname"
pwd
totalcodelinenumber=`find $projectname -name "*.java"|xargs cat|grep -v ^$|wc -l`
twoannocodeline=`find $projectname -name "*.java"|xargs cat|grep -v -e ^$ -e ^\s*\/\/.*$|wc -l`
starannocodeline=`find $projectname -name "*.java"|xargs cat|grep -v -e ^$ -e ^\s*\/**\*\*/.*$|wc -l`
annocodeline1=$((totalcodelinenumber - twoannocodeline))
annocodeline2=$((totalcodelinenumber - starannocodeline))
annocodelinesum=$((annocodeline1 + annocodeline2))
echo " the totalcodelinenumber is ->$totalcodelinenumber  and  the twoannocodeline is ->$twoannocodeline and  the starannocodeline is ->$starannocodeline "
echo " the annocodeline1 is ->$annocodeline1  and  the annocodeline2 is ->$annocodeline2 , the sum is ->$annocodelinesum"

#annonumbner=`find $projectname -name "*.java"|xargs cat|grep -v ^$|wc -l`    #Output:36335
# exclude the lines begin with //  
#noannonumbner=find . -name "*.java"|xargs cat|grep -v -e ^$ -e ^\s*\/\/.*$|wc -l    #Output:36064 
#noannonumbner=`find $projectname -name *.java | xargs cat | grep -v -e ^$ -e ^\s*\/\/.*$ | wc -l`
#echo $annonumbner
#echo $noannonumbner

percent_1=$(printf "%d%%" $((annocodelinesum*100/totalcodelinenumber)))
# no 0.1
#percent_2=`awk 'BEGIN{printf "%.1f%%\n",('$noannonumbner'/'$annonumbner')*100}'`
# save 0.1
echo " the annotation code line coverage % is ->$percent_1"
#echo $percent_2
