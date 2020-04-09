
package main

import(
	"encoding/json"
	"fmt"
	"time"
)


type HostPlanSyncResult struct {
	Ts                    int64
	ZoneIspHostPlans      map[string]map[int]*HostStatistics //map[zone_isp]map[len(files)]*HostStatistics
	PlanNumHostStatistics map[int]*HostStatistics            //map[len(files)]*HostStatistics
	PeerDistribution      map[string]int                     //zone_isp=>num
}

type HostStatistics struct {
	Count int64
	Hosts []string `json:"-"`
}



func  main()  {
	hostStatistics := HostStatistics{Count:1,Hosts:[]string{"aaa"}}
	peerDistribution := map[string]int{"1_2":1}
	planNumHostStatistics := map[int]*HostStatistics{2:&hostStatistics,0:&hostStatistics}
	zoneIspHostPlans := map[string]map[int]*HostStatistics{"1_2":map[int]*HostStatistics{1:&hostStatistics}}
	ts := time.Now().Unix()

	hostPlanSyncResult := &HostPlanSyncResult{
		Ts                    :ts,
	ZoneIspHostPlans      :zoneIspHostPlans,
	PlanNumHostStatistics :planNumHostStatistics,
	PeerDistribution     :peerDistribution,
	}

	res,_ := json.Marshal(hostPlanSyncResult)
	fmt.Println(string(res))

}