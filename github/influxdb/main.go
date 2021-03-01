package main

import (
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"time"
)

func main(){
	const token = "m3N1ReeLvQzc8vrW7SM5TsgoZlQh17SI3dFhCZHbRtagTCpj9NglW_-m6B3DAA3CENDKYYY-m0LOS_ZPOKoIVA=="
	const bucket = "bucket01"
	const org = "0990"


	client := influxdb2.NewClient("http://10.225.136.212:8086", token)
	// always close client at the end
	defer client.Close()

	wapi:=client.WriteAPI(org,bucket)
	//wapi.WriteRecord(fmt.Sprintf("stat,unit=temperature avg=%f,max=%f",23.5,45.0))
	//wapi.Flush()

	p:=influxdb2.NewPoint("stat", map[string]string{
		"unit":"temperature",
	}, map[string]interface{}{
		"avg":24.3,
		"max":46.7,
	},time.Now())

	wapi.WritePoint(p)
	wapi.Flush()
	time.Sleep(time.Second*3)
}
