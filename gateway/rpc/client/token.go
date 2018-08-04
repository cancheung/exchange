package client

import (
	"context"
	cf "digicon/gateway/conf"
	proto "digicon/proto/rpc"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	log "github.com/sirupsen/logrus"
)

type TokenRPCCli struct {
	conn proto.TokenRPCService
}

func (s *TokenRPCCli) CallAdmin(name string) (rsp *proto.AdminResponse, err error) {
	rsp, err = s.conn.AdminCmd(context.TODO(), &proto.AdminRequest{})
	if err != nil {
		log.Errorln(err.Error())
		return
	}
	return
}

func (s *TokenRPCCli) CallEntrustOrder(p *proto.EntrustOrderRequest) (rsp *proto.CommonErrResponse, err error) {
	rsp, err = s.conn.EntrustOrder(context.TODO(), p)
	if err != nil {
		log.Errorln(err.Error())
		return
	}
	return
}

func (s *TokenRPCCli) CallSelfSymbols(p *proto.SelfSymbolsRequest) (rsp *proto.SelfSymbolsResponse, err error) {
	/*
		rsp, err = s.conn.SelfSymbols(context.TODO(), p)
		if err != nil {
			log.Errorln(err.Error())
			return
		}
	*/
	return
}

func (s *TokenRPCCli) CallEntrustQuene(p *proto.EntrustQueneRequest) (rsp *proto.EntrustQueneResponse, err error) {
	rsp, err = s.conn.EntrustQuene(context.TODO(), p)
	if err != nil {
		log.Errorln(err.Error())
		return
	}
	return
}

func (s *TokenRPCCli) CallEntrustList(p *proto.EntrustHistoryRequest) (rsp *proto.EntrustListResponse, err error) {
	rsp, err = s.conn.EntrustList(context.TODO(), p)
	if err != nil {
		log.Errorln(err.Error())
		return
	}
	return
}

func (s *TokenRPCCli) CallEntrustHistory(p *proto.EntrustHistoryRequest) (rsp *proto.EntrustHistoryResponse, err error) {
	rsp, err = s.conn.EntrustHistory(context.TODO(), p)
	if err != nil {
		log.Errorln(err.Error())
		return
	}
	return
}

func (s *TokenRPCCli) CallBibiHistory(p *proto.BibiHistoryRequest) (rsp *proto.BibiHistoryResponse, err error) {
	rsp, err = s.conn.BibiHistory(context.TODO(), p)
	if err != nil {
		log.Errorln(err.Error())
		return
	}
	return
}

func (s *TokenRPCCli) CallDelEntrust(p *proto.DelEntrustRequest) (rsp *proto.DelEntrustResponse, err error) {
	rsp, err = s.conn.DelEntrust(context.TODO(), p)
	if err != nil {
		log.Errorln(err.Error())
		return
	}
	return
}

func (s *TokenRPCCli) CallTrade(p *proto.TradeRequest) (rsp *proto.TradeRespone, err error) {
	rsp, err = s.conn.Trade(context.TODO(), p)
	if err != nil {
		log.Errorln(err.Error())
		return
	}
	return
}

func (s *TokenRPCCli) CallTokenBalance(p *proto.TokenBalanceRequest) (rsp *proto.TokenBalanceResponse, err error) {
	rsp, err = s.conn.TokenBalance(context.TODO(), p)
	if err != nil {
		log.Errorln(err.Error())
		return
	}
	return
}

func (s *TokenRPCCli) CallTokenBalanceList(p *proto.TokenBalanceListRequest) (rsp *proto.TokenBalanceListResponse, err error) {
	rsp, err = s.conn.TokenBalanceList(context.TODO(), p)
	if err != nil {
		log.Errorln(err.Error())
		return
	}
	return
}

func (s *TokenRPCCli) CallTokenTradeList(p *proto.TokenTradeListRequest) (rsp *proto.TokenTradeListResponse, err error) {
	rsp, err = s.conn.TokenTradeList(context.TODO(), p)
	if err != nil {
		log.Errorln(err.Error())
		return
	}
	return
}

func (s *TokenRPCCli) CallTransferToCurrency(p *proto.TransferToCurrencyRequest) (rsp *proto.TransferToCurrencyResponse, err error) {
	rsp, err = s.conn.TransferToCurrency(context.TODO(), p)
	if err != nil {
		log.Errorln(err.Error())
		return
	}
	return
}

type KLineData struct {
	ID     int64   `json:"id"`     // K线ID
	Open   float64 `json:"open"`   // 开盘价
	Close  float64 `json:"close"`  // 收盘价, 当K线为最晚的一根时, 时最新成交价
	Low    float64 `json:"low"`    // 最低价
	High   float64 `json:"high"`   // 最高价
	Amount float64 `json:"amount"` // 成交量
	Vol    float64 `json:"vol"`    // 成交额, 即SUM(每一笔成交价 * 该笔的成交数量)
	Count  int64   `json:"count"`  // 成交笔数
}

func (s *TokenRPCCli) CallHistoryKline(symbol, period string, size int32) (rsp []KLineData, err error) {

	d := []KLineData{
		{1529640000, 0.127713, 0.127659, 0.127659, 0.127752, 3.4064, 0.4350326320, 41}, {1529639700, 0.127646, 0.127673, 0.127541, 0.127763, 4.4231, 0.5647776425, 48}, {1529639400, 0.127546, 0.127652, 0.127464, 0.127713, 4.5327, 0.5784328741, 57}, {1529639100, 0.127519, 0.127665, 0.127423, 0.127666, 8.0874, 1.0309396295, 65}, {1529638800, 0.127726, 0.127533, 0.127477, 0.127776, 25.1192, 3.2072313492, 67}, {1529638500, 0.127780, 0.127751, 0.127700, 0.127780, 4.1302, 0.5275327868, 54}, {1529638200, 0.127840, 0.127780, 0.127700, 0.127856, 3.3718, 0.4308459752, 52}, {1529637900, 0.127739, 0.127871, 0.127700, 0.127871, 5.3257, 0.6805732671, 55}, {1529637600, 0.127905, 0.127766, 0.127689, 0.127911, 13.0647, 1.6705417017, 60}, {1529637300, 0.127650, 0.127889, 0.127622, 0.127936, 6.4575, 0.8249809507, 63}, {1529637000, 0.127731, 0.127709, 0.127622, 0.127876, 6.4776, 0.8275950108, 60}, {1529636700, 0.127787, 0.127746, 0.127538, 0.127802, 17.2538, 2.2029310889, 61}, {1529636400, 0.127601, 0.127639, 0.127518, 0.127714, 8.3045, 1.0596382244, 55}, {1529636100, 0.127600, 0.127520, 0.127142, 0.127629, 15.0450, 1.9158201144, 77}, {1529635800, 0.127390, 0.127554, 0.127250, 0.127875, 34.3561, 4.3858343183, 82}, {1529635500, 0.127391, 0.127403, 0.127233, 0.127510, 3.1475, 0.4009152823, 56}, {1529635200, 0.127149, 0.127387, 0.127102, 0.127512, 25.9077, 3.2949373915, 94}, {1529634900, 0.127098, 0.127142, 0.126985, 0.127186, 9.7191, 1.2356509727, 72}, {1529634600, 0.127063, 0.127077, 0.126939, 0.127301, 33.9769, 4.3188899814, 156}, {1529634300, 0.127214, 0.127083, 0.127022, 0.127244, 12.3607, 1.5707209293, 108}, {1529634000, 0.127019, 0.127170, 0.127007, 0.127275, 6.9972, 0.8896400217, 56}, {1529633700, 0.127095, 0.127152, 0.126936, 0.127215, 4.7928, 0.6091053818, 58}, {1529633400, 0.127185, 0.127120, 0.126979, 0.127327, 23.2096, 2.9515769354, 66}, {1529633100, 0.127309, 0.127150, 0.127082, 0.127508, 19.0711, 2.4268084420, 77}, {1529632800, 0.127172, 0.127407, 0.127095, 0.127426, 10.9179, 1.3893219766, 62}, {1529632500, 0.127301, 0.127211, 0.127100, 0.127402, 20.1473, 2.5645633613, 92}, {1529632200, 0.128087, 0.127301, 0.127300, 0.128196, 71.9906, 9.1966737600, 185}, {1529631900, 0.128268, 0.128128, 0.128000, 0.128404, 42.2440, 5.4133713249, 125}, {1529631600, 0.129325, 0.128276, 0.128130, 0.129358, 94.0219, 12.1087642367, 197}, {1529631300, 0.129437, 0.129357, 0.129228, 0.129437, 19.7776, 2.5578151439, 77}, {1529631000, 0.129563, 0.129461, 0.129305, 0.129570, 6.3767, 0.8256741001, 55}, {1529630700, 0.129974, 0.129566, 0.129555, 0.129981, 52.3924, 6.7995900575, 134}, {1529630400, 0.130117, 0.129961, 0.129955, 0.130202, 3.5868, 0.4663971867, 55}, {1529630100, 0.129948, 0.130118, 0.129948, 0.130213, 4.6164, 0.6001863358, 55}, {1529629800, 0.129937, 0.129948, 0.129870, 0.130008, 17.3729, 2.2573415080, 54}, {1529629500, 0.129898, 0.129938, 0.129869, 0.129968, 6.0017, 0.7795571585, 55}, {1529629200, 0.130002, 0.129900, 0.129873, 0.130032, 3.9288, 0.5103913238, 52}, {1529628900, 0.130081, 0.130030, 0.129923, 0.130085, 3.0381, 0.3950688313, 51}, {1529628600, 0.129926, 0.130020, 0.129888, 0.130103, 3.4539, 0.4488837601, 47}, {1529628300, 0.129982, 0.129903, 0.129872, 0.130080, 2.7533, 0.3578225310, 41}, {1529628000, 0.129892, 0.130050, 0.129822, 0.130108, 8.0117, 1.0417620052, 77}, {1529627700, 0.129841, 0.129892, 0.129821, 0.129962, 11.8519, 1.5395410832, 44}, {1529627400, 0.129855, 0.129868, 0.129821, 0.129890, 10.2626, 1.3328551545, 44}, {1529627100, 0.129834, 0.129871, 0.129822, 0.129890, 6.9479, 0.9023286181, 60}, {1529626800, 0.129845, 0.129860, 0.129821, 0.129869, 3.6525, 0.4742660833, 54}, {1529626500, 0.129836, 0.129845, 0.129821, 0.129869, 3.2651, 0.4239700280, 46}, {1529626200, 0.129828, 0.129828, 0.129822, 0.129836, 3.2111, 0.4168926905, 48}, {1529625900, 0.129791, 0.129824, 0.129791, 0.129836, 3.2268, 0.4189150509, 52}, {1529625600, 0.129714, 0.129791, 0.129653, 0.129836, 4.4287, 0.5745855123, 51}, {1529625300, 0.129683, 0.129713, 0.129652, 0.129774, 2.8157, 0.3651926296, 46}, {1529625000, 0.129709, 0.129660, 0.129652, 0.129739, 2.7366, 0.3549139498, 46}, {1529624700, 0.129612, 0.129661, 0.129564, 0.129758, 6.4228, 0.8327722056, 60}, {1529624400, 0.129754, 0.129641, 0.129607, 0.129818, 28.9741, 3.7588754026, 73}, {1529624100, 0.129750, 0.129754, 0.129624, 0.129884, 2.9657, 0.3846855893, 44}, {1529623800, 0.129863, 0.129796, 0.129628, 0.129863, 11.7513, 1.5239769934, 64}, {1529623500, 0.129739, 0.129795, 0.129694, 0.129889, 6.3054, 0.8183735216, 70}, {1529623200, 0.129798, 0.129819, 0.129754, 0.129838, 5.2047, 0.6755262471, 60}, {1529622900, 0.129896, 0.129813, 0.129789, 0.129915, 3.7507, 0.4869658166, 56}, {1529622600, 0.129891, 0.129915, 0.129822, 0.129935, 2.4625, 0.3198200740, 43}, {1529622300, 0.129894, 0.129863, 0.129798, 0.129972, 1.9284, 0.2504863916, 33}, {1529622000, 0.129871, 0.129906, 0.129789, 0.129971, 3.0170, 0.3918800692, 53}, {1529621700, 0.129801, 0.129896, 0.129761, 0.129972, 9.8791, 1.2829423315, 56}, {1529621400, 0.129954, 0.129775, 0.129775, 0.129954, 7.5781, 0.9840235459, 49}, {1529621100, 0.129862, 0.129904, 0.129837, 0.129971, 4.8677, 0.6321924788, 62}, {1529620800, 0.129709, 0.129864, 0.129599, 0.129872, 15.6991, 2.0380897972, 84}, {1529620500, 0.129482, 0.129822, 0.129460, 0.129829, 26.0300, 3.3730042377, 99}, {1529620200, 0.129608, 0.129482, 0.129475, 0.129608, 8.3632, 1.0829780218, 57}, {1529619900, 0.129563, 0.129537, 0.129475, 0.129608, 5.0384, 0.6526279939, 56}, {1529619600, 0.129583, 0.129569, 0.129503, 0.129619, 5.6184, 0.7278108506, 62}, {1529619300, 0.129584, 0.129534, 0.129534, 0.129634, 3.7650, 0.4878615973, 57}, {1529619000, 0.129569, 0.129592, 0.129532, 0.129634, 4.2855, 0.5552670591, 53}, {1529618700, 0.129582, 0.129581, 0.129532, 0.129598, 8.0141, 1.0382118710, 62}, {1529618400, 0.129573, 0.129616, 0.129536, 0.129634, 16.3226, 2.1153460928, 59}, {1529618100, 0.129555, 0.129573, 0.129537, 0.129573, 7.4726, 0.9681839247, 62}, {1529617800, 0.129540, 0.129555, 0.129532, 0.129573, 9.2539, 1.1989397667, 73}, {1529617500, 0.129552, 0.129557, 0.129532, 0.129573, 4.0459, 0.5241455655, 50}, {1529617200, 0.129615, 0.129536, 0.129527, 0.129615, 5.2619, 0.6817282209, 53}, {1529616900, 0.129569, 0.129539, 0.129486, 0.129613, 2.5637, 0.3321297446, 43}, {1529616600, 0.129603, 0.129548, 0.129486, 0.129678, 9.3160, 1.2073404679, 58}, {1529616300, 0.129612, 0.129609, 0.129587, 0.129679, 4.0265, 0.5219301624, 62}, {1529616000, 0.129555, 0.129580, 0.129549, 0.129678, 10.9861, 1.4235691682, 55}, {1529615700, 0.129577, 0.129563, 0.129553, 0.129577, 3.9800, 0.5156951820, 43}, {1529615400, 0.129557, 0.129570, 0.129511, 0.129641, 3.2964, 0.4271809407, 51}, {1529615100, 0.129502, 0.129545, 0.129482, 0.129633, 9.1564, 1.1861794469, 50}, {1529614800, 0.129484, 0.129460, 0.129460, 0.129621, 3.0136, 0.3902729773, 45}, {1529614500, 0.129584, 0.129484, 0.129460, 0.129618, 6.3948, 0.8282736018, 58}, {1529614200, 0.129638, 0.129605, 0.129540, 0.129718, 5.6425, 0.7314290018, 79}, {1529613900, 0.129608, 0.129703, 0.129542, 0.129828, 6.4210, 0.8328812199, 61}, {1529613600, 0.129788, 0.129614, 0.129531, 0.129809, 11.8709, 1.5387129549, 75}, {1529613300, 0.129737, 0.129757, 0.129699, 0.129860, 2.6012, 0.3376078174, 42}, {1529613000, 0.129674, 0.129784, 0.129652, 0.129827, 3.7690, 0.4889112935, 49}, {1529612700, 0.129735, 0.129767, 0.129656, 0.129806, 4.5723, 0.5932420105, 61}, {1529612400, 0.129568, 0.129735, 0.129531, 0.129762, 10.9048, 1.4133124146, 49}, {1529612100, 0.129636, 0.129568, 0.129531, 0.129681, 3.8387, 0.4974273437, 53}, {1529611800, 0.129464, 0.129531, 0.129464, 0.129681, 5.0939, 0.6600637148, 54}, {1529611500, 0.129469, 0.129463, 0.129460, 0.129469, 11.0091, 1.4253346661, 21}, {1529611200, 0.129538, 0.129469, 0.129469, 0.129592, 2.0959, 0.2714364884, 26}, {1529610900, 0.129553, 0.129531, 0.129467, 0.129575, 3.3214, 0.4301833025, 53}, {1529610600, 0.129526, 0.129533, 0.129478, 0.129578, 3.1428, 0.4070720269, 43}, {1529610300, 0.129541, 0.129500, 0.129500, 0.129570, 4.2855, 0.5550882369, 48}, {1529610000, 0.129546, 0.129541, 0.129500, 0.129574, 3.1829, 0.4123053689, 49}, {1529609700, 0.129520, 0.129561, 0.129515, 0.129592, 3.0465, 0.3946746127, 52}, {1529609400, 0.129528, 0.129540, 0.129479, 0.129591, 2.8817, 0.3732611398, 48}, {1529609100, 0.129535, 0.129496, 0.129479, 0.129592, 3.7580, 0.4867469242, 51}, {1529608800, 0.129553, 0.129535, 0.129481, 0.129592, 3.4476, 0.4466388944, 49}, {1529608500, 0.129527, 0.129540, 0.129497, 0.129625, 25.4076, 3.2916349966, 53}, {1529608200, 0.129490, 0.129526, 0.129397, 0.129585, 2.9918, 0.3874979825, 45}, {1529607900, 0.129558, 0.129443, 0.129385, 0.129585, 17.9374, 2.3232804255, 60}, {1529607600, 0.129414, 0.129559, 0.129414, 0.129586, 4.8989, 0.6344789914, 47}, {1529607300, 0.129446, 0.129545, 0.129414, 0.129590, 3.0401, 0.3936651182, 47}, {1529607000, 0.129499, 0.129500, 0.129417, 0.129592, 2.5649, 0.3321583638, 47}, {1529606700, 0.129502, 0.129434, 0.129434, 0.129561, 2.7023, 0.3499437833, 46}, {1529606400, 0.129474, 0.129503, 0.129414, 0.129597, 4.3822, 0.5674869947, 55}, {1529606100, 0.129461, 0.129474, 0.129447, 0.129597, 4.4689, 0.5787124609, 56}, {1529605800, 0.129529, 0.129461, 0.129453, 0.129597, 5.0507, 0.6540372345, 61}, {1529605500, 0.129543, 0.129507, 0.129445, 0.129586, 5.0127, 0.6492798591, 53}, {1529605200, 0.129501, 0.129503, 0.129428, 0.129586, 4.6528, 0.6026121267, 64}, {1529604900, 0.129501, 0.129498, 0.129373, 0.129594, 2.6557, 0.3438511856, 46}, {1529604600, 0.129454, 0.129501, 0.129431, 0.129546, 2.9229, 0.3784740617, 47}, {1529604300, 0.129532, 0.129468, 0.129460, 0.129588, 4.3973, 0.5695275522, 55}, {1529604000, 0.129539, 0.129532, 0.129466, 0.129597, 4.0246, 0.5212352127, 57}, {1529603700, 0.129515, 0.129538, 0.129492, 0.129595, 2.8037, 0.3631774154, 47}, {1529603400, 0.129506, 0.129532, 0.129467, 0.129600, 2.7852, 0.3607649984, 47}, {1529603100, 0.129508, 0.129553, 0.129441, 0.129589, 4.5101, 0.5840103301, 53}, {1529602800, 0.129453, 0.129509, 0.129453, 0.129577, 2.9325, 0.3797914849, 48}, {1529602500, 0.129538, 0.129516, 0.129453, 0.129685, 2.5622, 0.3319536624, 42}, {1529602200, 0.129540, 0.129490, 0.129423, 0.129678, 2.4768, 0.3208684265, 42}, {1529601900, 0.129357, 0.129476, 0.129322, 0.129709, 11.8884, 1.5394702674, 57}, {1529601600, 0.129405, 0.129388, 0.129294, 0.129477, 7.7253, 0.9998509385, 68}, {1529601300, 0.129719, 0.129336, 0.129214, 0.129818, 13.5641, 1.7561498766, 87}, {1529601000, 0.129732, 0.129669, 0.129620, 0.129861, 5.4314, 0.7045549608, 59}, {1529600700, 0.129753, 0.129755, 0.129621, 0.129885, 2.5138, 0.3261328133, 46}, {1529600400, 0.129779, 0.129730, 0.129621, 0.129885, 3.0565, 0.3966799279, 50}, {1529600100, 0.129702, 0.129885, 0.129656, 0.129886, 2.5069, 0.3253740123, 41}, {1529599800, 0.129831, 0.129687, 0.129660, 0.129900, 7.6803, 0.9963683252, 58}, {1529599500, 0.129928, 0.129828, 0.129815, 0.130024, 3.7383, 0.4855703999, 50}, {1529599200, 0.129957, 0.129941, 0.129818, 0.130025, 5.4721, 0.7110866971, 55}, {1529598900, 0.129827, 0.129956, 0.129757, 0.129965, 3.8530, 0.5003899164, 54}, {1529598600, 0.129965, 0.129836, 0.129836, 0.129965, 3.9314, 0.5107915737, 48}, {1529598300, 0.129893, 0.129912, 0.129849, 0.129965, 12.8487, 1.6696412771, 62}, {1529598000, 0.129929, 0.129939, 0.129826, 0.130085, 41.9216, 5.4459941423, 136}, {1529597700, 0.129984, 0.129919, 0.129869, 0.130437, 28.7751, 3.7405423488, 204}, {1529597400, 0.130114, 0.129965, 0.129931, 0.130600, 50.6588, 6.6061934340, 79}, {1529597100, 0.130273, 0.130304, 0.130188, 0.130304, 13.2736, 1.7292028268, 18}, {1529596800, 0.130103, 0.130186, 0.130103, 0.130274, 2.0475, 0.2665305782, 35}, {1529596500, 0.130132, 0.130178, 0.130084, 0.130255, 10.6230, 1.3824351426, 80}, {1529596200, 0.130112, 0.130085, 0.130081, 0.130144, 3.1998, 0.4163318303, 54}, {1529595900, 0.130020, 0.130112, 0.129960, 0.130143, 10.5785, 1.3760527209, 59}, {1529595600, 0.129994, 0.130020, 0.129917, 0.130158, 8.1437, 1.0587300682, 93}, {1529595300, 0.129882, 0.129995, 0.129882, 0.130102, 3.9110, 0.5084449337, 54},
	}

	return d, nil
}

func NewTokenRPCCli() (u *TokenRPCCli) {
	consul_addr := cf.Cfg.MustValue("consul", "addr")
	r := consul.NewRegistry(registry.Addrs(consul_addr))
	service := micro.NewService(
		micro.Name("token.client"),
		micro.Registry(r),
	)
	service.Init()

	service_name := cf.Cfg.MustValue("base", "service_client_token")
	greeter := proto.NewTokenRPCService(service_name, service.Client())
	u = &TokenRPCCli{
		conn: greeter,
	}
	return
}
