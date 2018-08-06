#! /bin/bash

result=`curl -s 'https://api.github.com/repos/developer-learning/night-reading-go/contributors' | jq '.[].login' | sed -e 's/"//g'`

for element in $result
do
	if [ `grep -c $element CONTRIBUTORS` -eq '0' ]; then
		echo 'add contributors.'
		echo $element >> CONTRIBUTORS
		echo $element
    	all-contributors add $element code
	fi	
done

all-contributors check

echo 'add contributors completed.'
echo 'contributors generate...'
all-contributors generate
echo 'contributors generate completed.'
