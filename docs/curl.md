// works fine

curl --header "Content-Type: application/json" \
 --request POST \
 --data '{"surroundings":[{"number": 1, "timestamp": "2023-04-24T10:00:00Z", "Rssi":-3 ,"tempreture": 10.5, "moisuture": 0.5, "airPressure": 1012.2} ]}' \
 http://localhost:8080/handle

curl --header "Content-Type: application/json" -X POST --data '{"surroundings":[{"number": 1, "timestamp": "2023-05-23 23:51:54.008849999 +0900 JST m=+0.000462300", "rssi": -3, "temperature": 10.5, "moisture": 0.5, "airPressure": 1012.2},{"number": 2, "timestamp": "2023-05-23 23:52:36.405033943 +0900 JST m=+0.000520363", "rssi": -3, "temperature": 15.3, "moisture": 0.6, "airPressure": 1011.9}]}' http://localhost:8080/handle

curl -X POST 'https://slack.com/api/chat.postMessage' \
-H 'Content-Type: application/json; charset=utf-8' \
-H 'Authorization: Bearer '$SLACK_BOT_OAUTH \
-d '{"channel":"#api-test", "text":"test"}'
