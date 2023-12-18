# azure

docker cp ./ <コンテナ名またはコンテナ ID>:/
docker run -it mcr.microsoft.com/azure-cli
docker cp ../vuoy_monitor_smaple 5478:/

親ディレクトリのファイルをコピー

//az cli をビルド
docker build -f azure/Dockerfile -t az-cli .

//az cli を dind と run
docker run --link dind:docker -v /var/run/docker.sock:/var/run/docker.sock -it az-cli

az commnads

```
az group create --name myResourceGroup --location eastus
az acr create --resource-group myResourceGroup --name my-first-acr  --sku Basic


az acr create --resource-group myResourceGroup --name maemuralaboacr  --sku Basic
```

docker tag myapp:latest maemuralab/myrepo:myapp
docker push maemuralab/myrepo:myapp

influxdb query

```
influx query --raw 'from(bucket:"vuoy_monitor") |> range(start:-1mo)'

```

```
from(bucket: "vuoy_monitor")
  |> range(start: -2m)
  |> filter(fn: (r) => r["_measurement"] == "vuoy_surroundings")
  |> filter(fn: (r) => r["_field"] == "AirPressure" or r["_field"] == "Moisuture" or r["_field"] == "Tempreture" or r["_field"] == "Rssi")
  |> count()
```

```flux
influx delete --org  iot --bucket vuoy_monitor --start '1970-01-01T00:00:00Z' --stop $(date +"%Y-%m-%dT%H:%M:%SZ")
```

```
echo 'from(bucket: "vuoy_monitor") |> range(start: -10m) |> filter(fn: (r) => r["_measurement"] == "vuoy_surroundings") |> filter(fn: (r) => r["_field"] == "AirPressure" or r["_field"] == "Moisuture" or r["_field"] == "Tempreture" or r["_field"] == "Rssi") |> count()' | docker exec -i vuoy_monitor_smaple-influxdb-1 influx query --org $INFLUXDB_ORG --token $INFLUXDB_TOKEN

echo 'from(bucket: "vuoy_monitor") |> range(start: -1d) |> filter(fn: (r) => r["_measurement"] == "vuoy_surroundings") |> filter(fn: (r) => r["_field"] == "AirPressure" or r["_field"] == "Moisuture" or r["_field"] == "Tempreture" or r["_field"] == "Rssi") |> count()' | docker exec -i  vuoy_monitor_smaple-influxdb-1 influx query --org $INFLUXDB_ORG --token $INFLUXDB_TOKEN
```

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

# influx delete --org iot --bucket vuoy_monitor --start '1970-01-01T00:00:00Z' --stop $(date +"%Y-%m-%dT%H:%M:%SZ")

# ^[[A^[[2~^[[2~^[[2~^[[H^[[H^C

# ^C

# influx query --raw 'from(bucket: "vuoy_monitor")

|> range(start: -30d)
|> filter(fn: (r) => r["_measurement"] == "vuoy_surroundings")
|> filter(fn: (r) => r["_field"] == "AirPressure")
'> > > >

# influx query --raw 'from(bucket: "vuoy_monitor")

|> range(start: -30d)
|> filter(fn: (r) => r["_measurement"] == "vuoy_surroundings")
|> filter(fn: (r) => r["_field"] == "AirPressure")
'> > > >
#group,false,false,true,true,false,false,true,true,true
#datatype,string,long,dateTime:RFC3339,dateTime:RFC3339,dateTime:RFC3339,double,string,string,string
#default,\_result,,,,,,,,
,result,table,\_start,\_stop,\_time,\_value,\_field,\_measurement,user
,,0,2023-04-23T15:48:33.472179383Z,2023-05-23T15:48:33.472179383Z,2023-05-23T10:00:00Z,1011.9,AirPressure,vuoy_surroundings,bar
