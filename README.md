#Cell Radio backend

##Prerequisites

Requires go >=1.9, dep (`go get -u github.com/golang/dep/cmd/dep`).

No need to clone sources to your GOPATH.

##Build

###Locally

```
export GOPATH=`pwd` && \
cd src/cell_radio && \
dep ensure && \
go install && \
cd -
```

###Docker image

`docker build -t cell_radio_backend:latest -f Dockerfile .`

##Run

###Locally

`./bin/cell_radio`

###In Docker container

`docker run --rm -it -p 9090:8080 -p 9091:8081 -e "SRV_HOST=0.0.0.0" --name cell_radio_backend cell_radio_backend:latest`

or

`docker run --rm -it -p 9090:8080 -p 9091:8081 --name cell_radio_backend cell_radio_backend:latest cell_radio -host 0.0.0.0`

##Test

###Manually

(if run in Docker with specified ports, change `http://localhost:8080` to `http://<docker_network_gateway_or_127.0.0.1>:<docker_port_mapped_to_8080>`

curl -iv http://localhost:8080/ht/addmany -X POST -d '{"DftrxId":"dftrx0", "Config":{"RxGain":19, "Network":0, "2g_params":{"Arfcn":540, "Ulsc":65, "Dlsc":2}, "3g_params":{"Arfcn":10788, "Ulsc":8758542, "Dlsc":369}, "4g_params":{"Arfcn":1400, "Ulsc":5, "Dlsc":0}, "Alpha":0.017, "Slot":"0", "Celltrack":0, "Watcher":0, "Antenna":0, "GpsSrc":1, "Version":2, "App":1, "GsmMode":1, "Url":"http://localhost:8081/addmany", "Sound":1.000}, "State":{"DlMode":0, "Tune":"ul", "SignalFound":0, "Hsdpa":0, "Standard":"DCS", "Usrp":1, "Capabilities":16678719, "Expiration":1607713200000, "Compression":1, "Error":0, "Remote":0}, "GpsExt":{"Status":0, "Lat":32.097828, "Lon":34.846316, "Alt":52.6, "TS":138}, "Compass":{"Hdg":0.0, "FitErr":""}, "Timestamp":0, "Points":[{"Id":552,"SettingsVersion":0,"Channel":65,"SNR":33.2,"Angles":[{"TSi":139, "TSf":0.259143384615, "RSSI_m":-28.3, "RSSI_s":0.0, "SNR_m":35.9, "SNR_s":0.0, "Master":1, "An":0.00, "Phase":0.00, "Visible":1, "Int_m":-3.54462, "Int_s":0.00000, "Int":-3.54462, "Ant":0, "RxGain":19, "V":1},{"TSi":139, "TSf":0.263758769231, "RSSI_m":-28.3, "RSSI_s":0.0, "SNR_m":37.4, "SNR_s":0.0, "Master":1, "An":0.00, "Phase":0.00, "Visible":1, "Int_m":-3.61846, "Int_s":0.00000, "Int":-3.61846, "Ant":0, "RxGain":19, "V":1},{"TSi":139, "TSf":0.268374153846, "RSSI_m":-28.4, "RSSI_s":0.0, "SNR_m":35.9, "SNR_s":0.0, "Master":1, "An":0.00, "Phase":0.00, "Visible":1, "Int_m":-3.54462, "Int_s":0.00000, "Int":-3.54462, "Ant":0, "RxGain":19, "V":1},{"TSi":139, "TSf":0.342220307692, "RSSI_m":-28.3, "RSSI_s":0.0, "SNR_m":37.4, "SNR_s":0.0, "Master":1, "An":0.00, "Phase":0.00, "Visible":1, "Int_m":-3.61846, "Int_s":0.00000, "Int":-3.61846, "Ant":0, "RxGain":19, "V":1},{"TSi":139, "TSf":0.346835692308, "RSSI_m":-28.3, "RSSI_s":0.0, "SNR_m":40.5, "SNR_s":0.0, "Master":1, "An":0.00, "Phase":0.00, "Visible":1, "Int_m":-3.61846, "Int_s":0.00000, "Int":-3.61846, "Ant":0, "RxGain":19, "V":1},{"TSi":139, "TSf":0.351451076923, "RSSI_m":-28.3, "RSSI_s":0.0, "SNR_m":33.2, "SNR_s":0.0, "Master":1, "An":0.00, "Phase":0.00, "Visible":1, "Int_m":-3.61846, "Int_s":0.00000, "Int":-3.61846, "Ant":0, "RxGain":19, "V":1}],"RSSI":-29.3,"Angle2a":0.0,"Antenna2a":22652576,"Compression":1,"Antenna":22652576,"TSi":139,"TSf":0.351451076923,"Clocks":[0, 0, 1536665497356427, 1536665497398082]}]}'

curl -iv http://localhost:8080/ht/addmany -X POST -d '{"DftrxId":"dftrx0", "Config":{"RxGain":19, "Network":0, "2g_params":{"Arfcn":540, "Ulsc":65, "Dlsc":2}, "3g_params":{"Arfcn":10788, "Ulsc":8758542, "Dlsc":369}, "4g_params":{"Arfcn":1400, "Ulsc":5, "Dlsc":0}, "Alpha":0.017, "Slot":"0", "Celltrack":0, "Watcher":0, "Antenna":0, "GpsSrc":1, "Version":2, "App":1, "GsmMode":1, "Url":"http://localhost:8081/addmany", "Sound":1.000}, "State":{"DlMode":0, "Tune":"ul", "SignalFound":0, "Hsdpa":0, "Standard":"DCS", "Usrp":1, "Capabilities":16678719, "Expiration":1607713200000, "Compression":1, "Error":0, "Remote":0}, "GpsExt":{"Status":0, "Lat":32.095, "Lon":34.846316, "Alt":52.6, "TS":138}, "Compass":{"Hdg":0.0, "FitErr":""}, "Timestamp":0, "Points":[{"Id":552,"SettingsVersion":0,"Channel":65,"SNR":36.2,"Angles":[{"TSi":139, "TSf":0.259143384615, "RSSI_m":-28.3, "RSSI_s":0.0, "SNR_m":35.9, "SNR_s":0.0, "Master":1, "An":0.00, "Phase":0.00, "Visible":1, "Int_m":-3.54462, "Int_s":0.00000, "Int":-3.54462, "Ant":0, "RxGain":19, "V":1},{"TSi":139, "TSf":0.263758769231, "RSSI_m":-28.3, "RSSI_s":0.0, "SNR_m":37.4, "SNR_s":0.0, "Master":1, "An":0.00, "Phase":0.00, "Visible":1, "Int_m":-3.61846, "Int_s":0.00000, "Int":-3.61846, "Ant":0, "RxGain":19, "V":1},{"TSi":139, "TSf":0.268374153846, "RSSI_m":-28.4, "RSSI_s":0.0, "SNR_m":35.9, "SNR_s":0.0, "Master":1, "An":0.00, "Phase":0.00, "Visible":1, "Int_m":-3.54462, "Int_s":0.00000, "Int":-3.54462, "Ant":0, "RxGain":19, "V":1},{"TSi":139, "TSf":0.342220307692, "RSSI_m":-28.3, "RSSI_s":0.0, "SNR_m":37.4, "SNR_s":0.0, "Master":1, "An":0.00, "Phase":0.00, "Visible":1, "Int_m":-3.61846, "Int_s":0.00000, "Int":-3.61846, "Ant":0, "RxGain":19, "V":1},{"TSi":139, "TSf":0.346835692308, "RSSI_m":-28.3, "RSSI_s":0.0, "SNR_m":40.5, "SNR_s":0.0, "Master":1, "An":0.00, "Phase":0.00, "Visible":1, "Int_m":-3.61846, "Int_s":0.00000, "Int":-3.61846, "Ant":0, "RxGain":19, "V":1},{"TSi":139, "TSf":0.351451076923, "RSSI_m":-28.3, "RSSI_s":0.0, "SNR_m":33.2, "SNR_s":0.0, "Master":1, "An":0.00, "Phase":0.00, "Visible":1, "Int_m":-3.61846, "Int_s":0.00000, "Int":-3.61846, "Ant":0, "RxGain":19, "V":1}],"RSSI":-22.3,"Angle2a":0.0,"Antenna2a":22652576,"Compression":1,"Antenna":22652576,"TSi":139,"TSf":0.351451076923,"Clocks":[0, 0, 1536665497356427, 1536665497398082]}]}'

curl -iv http://localhost:8080/ht/addmany -X POST -d '{"DftrxId":"dftrx0", "Config":{"RxGain":19, "Network":0, "2g_params":{"Arfcn":540, "Ulsc":65, "Dlsc":2}, "3g_params":{"Arfcn":10788, "Ulsc":8758542, "Dlsc":369}, "4g_params":{"Arfcn":1400, "Ulsc":5, "Dlsc":0}, "Alpha":0.017, "Slot":"0", "Celltrack":0, "Watcher":0, "Antenna":0, "GpsSrc":1, "Version":2, "App":1, "GsmMode":1, "Url":"http://localhost:8081/addmany", "Sound":1.000}, "State":{"DlMode":0, "Tune":"ul", "SignalFound":0, "Hsdpa":0, "Standard":"DCS", "Usrp":1, "Capabilities":16678719, "Expiration":1607713200000, "Compression":1, "Error":0, "Remote":0}, "GpsExt":{"Status":0, "Lat":32.0967, "Lon":34.8463, "Alt":52.6, "TS":138}, "Compass":{"Hdg":0.0, "FitErr":""}, "Timestamp":0, "Points":[{"Id":552,"SettingsVersion":0,"Channel":65,"SNR":36.2,"Angles":[{"TSi":139, "TSf":0.259143384615, "RSSI_m":-28.3, "RSSI_s":0.0, "SNR_m":35.9, "SNR_s":0.0, "Master":1, "An":0.00, "Phase":0.00, "Visible":1, "Int_m":-3.54462, "Int_s":0.00000, "Int":-3.54462, "Ant":0, "RxGain":19, "V":1},{"TSi":139, "TSf":0.263758769231, "RSSI_m":-28.3, "RSSI_s":0.0, "SNR_m":37.4, "SNR_s":0.0, "Master":1, "An":0.00, "Phase":0.00, "Visible":1, "Int_m":-3.61846, "Int_s":0.00000, "Int":-3.61846, "Ant":0, "RxGain":19, "V":1},{"TSi":139, "TSf":0.268374153846, "RSSI_m":-28.4, "RSSI_s":0.0, "SNR_m":35.9, "SNR_s":0.0, "Master":1, "An":0.00, "Phase":0.00, "Visible":1, "Int_m":-3.54462, "Int_s":0.00000, "Int":-3.54462, "Ant":0, "RxGain":19, "V":1},{"TSi":139, "TSf":0.342220307692, "RSSI_m":-28.3, "RSSI_s":0.0, "SNR_m":37.4, "SNR_s":0.0, "Master":1, "An":0.00, "Phase":0.00, "Visible":1, "Int_m":-3.61846, "Int_s":0.00000, "Int":-3.61846, "Ant":0, "RxGain":19, "V":1},{"TSi":139, "TSf":0.346835692308, "RSSI_m":-28.3, "RSSI_s":0.0, "SNR_m":40.5, "SNR_s":0.0, "Master":1, "An":0.00, "Phase":0.00, "Visible":1, "Int_m":-3.61846, "Int_s":0.00000, "Int":-3.61846, "Ant":0, "RxGain":19, "V":1},{"TSi":139, "TSf":0.351451076923, "RSSI_m":-28.3, "RSSI_s":0.0, "SNR_m":33.2, "SNR_s":0.0, "Master":1, "An":0.00, "Phase":0.00, "Visible":1, "Int_m":-3.61846, "Int_s":0.00000, "Int":-3.61846, "Ant":0, "RxGain":19, "V":1}],"RSSI":-24.3,"Angle2a":0.0,"Antenna2a":22652576,"Compression":1,"Antenna":22652576,"TSi":139,"TSf":0.351451076923,"Clocks":[0, 0, 1536665497356427, 1536665497398082]},{"Id":553,"SettingsVersion":0,"Channel":65,"SNR":30.2,"Angles":[{"TSi":139, "TSf":0.259143384615, "RSSI_m":-28.3, "RSSI_s":0.0, "SNR_m":35.9, "SNR_s":0.0, "Master":1, "An":0.00, "Phase":0.00, "Visible":1, "Int_m":-3.54462, "Int_s":0.00000, "Int":-3.54462, "Ant":0, "RxGain":19, "V":1},{"TSi":139, "TSf":0.263758769231, "RSSI_m":-28.3, "RSSI_s":0.0, "SNR_m":37.4, "SNR_s":0.0, "Master":1, "An":0.00, "Phase":0.00, "Visible":1, "Int_m":-3.61846, "Int_s":0.00000, "Int":-3.61846, "Ant":0, "RxGain":19, "V":1},{"TSi":139, "TSf":0.268374153846, "RSSI_m":-28.4, "RSSI_s":0.0, "SNR_m":35.9, "SNR_s":0.0, "Master":1, "An":0.00, "Phase":0.00, "Visible":1, "Int_m":-3.54462, "Int_s":0.00000, "Int":-3.54462, "Ant":0, "RxGain":19, "V":1},{"TSi":139, "TSf":0.342220307692, "RSSI_m":-28.3, "RSSI_s":0.0, "SNR_m":37.4, "SNR_s":0.0, "Master":1, "An":0.00, "Phase":0.00, "Visible":1, "Int_m":-3.61846, "Int_s":0.00000, "Int":-3.61846, "Ant":0, "RxGain":19, "V":1},{"TSi":139, "TSf":0.346835692308, "RSSI_m":-28.3, "RSSI_s":0.0, "SNR_m":40.5, "SNR_s":0.0, "Master":1, "An":0.00, "Phase":0.00, "Visible":1, "Int_m":-3.61846, "Int_s":0.00000, "Int":-3.61846, "Ant":0, "RxGain":19, "V":1},{"TSi":139, "TSf":0.351451076923, "RSSI_m":-28.3, "RSSI_s":0.0, "SNR_m":33.2, "SNR_s":0.0, "Master":1, "An":0.00, "Phase":0.00, "Visible":1, "Int_m":-3.61846, "Int_s":0.00000, "Int":-3.61846, "Ant":0, "RxGain":19, "V":1}],"RSSI":-21.3,"Angle2a":0.0,"Antenna2a":22652576,"Compression":1,"Antenna":22652576,"TSi":139,"TSf":0.351451076923,"Clocks":[0, 0, 1536665497356427, 1536665497398082]}]}'

###Automatically

TODO: add auto tests
