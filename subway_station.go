package main

// 地铁站station,start起点,end终点
// a,b,c,d
type station struct {
	start_station rune
	end_station   []rune
}

var stations []station
var res12 [][]rune

//func main() {
//	station_a := &station{start_station: 'a', end_station: []rune{'b', 'c', 'd'}}
//	station_b := &station{start_station: 'b', end_station: []rune{'c', 'd'}}
//	station_c := &station{start_station: 'c', end_station: []rune{'d'}}
//	stations = make([]station, 0)
//	res12 = make([][]rune, 0)
//	stations = append(stations, *station_a)
//	stations = append(stations, *station_b)
//	stations = append(stations, *station_c)
//	dfs2(stations, []rune{}, 0, 'd')
//	fmt.Println(res12)
//}

func dfs2(stations []station, temp []rune, index int, end_station rune) {
	if len(temp) > 0 && end_station == 'd' {
		arr := make([]rune, 0)
		arr = append(arr, temp...)
		res12 = append(res12, arr)
		return
	}
	for i := 0; i < len(stations[index].end_station); i++ {
		temp = append(temp, stations[index].end_station[i])
		dfs2(stations, temp, i, stations[index].end_station[i])
		temp = temp[:len(temp)-1]
	}
}

// 知道站点去哪，加上信息，去每个站点时间。
// channel go的使用
// 实际问题 需求：几百万用户，排行榜，api->top100积分用户 zset 去重，有序 zrange(0,100)
//桶排序->
// 做任务，领奖励。任务模块->运营，-> api处理模块 ->redis ，button->奖励模块 sdk通用rpc。
// config1<-  config2 vim ->pipeline 动态启动 nacos 服务注册 服务发现
//
