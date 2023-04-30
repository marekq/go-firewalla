package mystructs

import "time"

type JsonToken struct {
	Token string `json:"token"`
	Url   string `json:"url"`
}

type FirewallaDevices []struct {
	Name string `json:"name"`
	IP   string `json:"ip"`
	Mac  string `json:"mac"`
	Gid  string `json:"gid"`
	Intf struct {
		UUID string `json:"uuid"`
		Name string `json:"name"`
	} `json:"intf"`
	MacVendor     string `json:"macVendor"`
	Online        bool   `json:"online"`
	TotalDownload int    `json:"totalDownload"`
	TotalUpload   int    `json:"totalUpload"`
	Reserved      bool   `json:"reserved"`
	LastActive    string `json:"lastActive"`
	FirstFound    string `json:"firstFound"`
	Tags          []struct {
		Name string `json:"name"`
		UID  string `json:"uid"`
	} `json:"tags"`
	BoxName     string `json:"boxName"`
	IsFirewalla bool   `json:"isFirewalla"`
}

type FirewallaAlarms []struct {
	Aid                        string   `json:"aid"`
	Type                       string   `json:"type"`
	Device                     string   `json:"device"`
	AlarmTimestamp             float64  `json:"alarmTimestamp"`
	Timestamp                  float64  `json:"timestamp"`
	PDeviceIP                  string   `json:"p.device.ip"`
	PDestIP                    string   `json:"p.dest.ip"`
	PNoticeType                string   `json:"p.noticeType"`
	PMessage                   string   `json:"p.message"`
	PLocalIsClient             string   `json:"p.local_is_client"`
	Message                    string   `json:"message"`
	PDeviceName                string   `json:"p.device.name"`
	PDeviceID                  string   `json:"p.device.id"`
	PDeviceMac                 string   `json:"p.device.mac"`
	PDeviceMacVendor           string   `json:"p.device.macVendor"`
	PDestName                  string   `json:"p.dest.name"`
	PDestLatitude              string   `json:"p.dest.latitude"`
	PDestLongitude             string   `json:"p.dest.longitude"`
	PDestCountry               string   `json:"p.dest.country"`
	PCloudDecision             string   `json:"p.cloud.decision"`
	PFi                        string   `json:"p.fi"`
	Gid                        string   `json:"gid"`
	BoxName                    string   `json:"boxName"`
	BoxModel                   string   `json:"boxModel"`
	PIntfName                  string   `json:"p.intf.name"`
	PIntfID                    string   `json:"p.intf.id"`
	PTagName                   string   `json:"p.tag.name"`
	PTagIds                    []string `json:"p.tag.ids"`
	TimeToGenerate             string   `json:"time.to.generate"`
	PDestCategory              string   `json:"p.dest.category"`
	PActionBlock               string   `json:"p.action.block"`
	ResultMethod               string   `json:"result_method"`
	PBlockby                   string   `json:"p.blockby"`
	PDestReadableName          string   `json:"p.dest.readableName"`
	Result                     string   `json:"result"`
	ResultPolicy               string   `json:"result_policy"`
	PSeverity                  string   `json:"p.severity"`
	PDestDomain                string   `json:"p.dest.domain"`
	PIntfSubnet                string   `json:"p.intf.subnet"`
	PDestID                    string   `json:"p.dest.id"`
	PDestPort                  string   `json:"p.dest.port"`
	PTransferInboundSize       string   `json:"p.transfer.inbound.size"`
	PTransferOutboundSize      string   `json:"p.transfer.outbound.size"`
	PDevicePort                []int    `json:"p.device.port"`
	PTimestampTimezone         string   `json:"p.timestampTimezone"`
	PTransferOutboundHumansize string   `json:"p.transfer.outbound.humansize"`
	PTransferInboundHumansize  string   `json:"p.transfer.inbound.humansize"`
	PTransferDuration          string   `json:"p.transfer.duration"`
	PProtocol                  string   `json:"p.protocol"`
	PTagNames                  []any    `json:"p.tag.names"`
	PTotalUsage                string   `json:"p.totalUsage"`
	PBeginTs                   string   `json:"p.begin.ts"`
	PFlows                     []struct {
		Count           int    `json:"count"`
		IP              string `json:"ip"`
		Device          string `json:"device"`
		Country         string `json:"country"`
		AggregationHost string `json:"aggregationHost"`
		Host            string `json:"host"`
		Category        string `json:"category"`
		App             string `json:"app"`
	} `json:"p.flows"`
	PDuration            string `json:"p.duration"`
	PPercentage          string `json:"p.percentage"`
	PDestNames           string `json:"p.dest.names"`
	PEndTs               string `json:"p.end.ts"`
	PTotalUsageHumansize string `json:"p.totalUsage.humansize"`
	PDestApp             string `json:"p.dest.app"`
}

type FirewallaAlarmDetail struct {
	PDeviceIP                  string   `json:"p.device.ip"`
	PDeviceID                  string   `json:"p.device.id"`
	PDestDomain                string   `json:"p.dest.domain"`
	PDestLatitude              string   `json:"p.dest.latitude"`
	Timestamp                  float64  `json:"timestamp"`
	PDeviceName                string   `json:"p.device.name"`
	PTagIds                    []string `json:"p.tag.ids"`
	PIntfSubnet                string   `json:"p.intf.subnet"`
	PDestID                    string   `json:"p.dest.id"`
	Message                    string   `json:"message"`
	PDestPort                  string   `json:"p.dest.port"`
	PDestName                  string   `json:"p.dest.name"`
	PIntfID                    string   `json:"p.intf.id"`
	PTransferInboundSize       string   `json:"p.transfer.inbound.size"`
	PDeviceMac                 string   `json:"p.device.mac"`
	PTransferOutboundSize      string   `json:"p.transfer.outbound.size"`
	PDestCountry               string   `json:"p.dest.country"`
	PLocalIsClient             string   `json:"p.local_is_client"`
	Device                     string   `json:"device"`
	PDestIP                    string   `json:"p.dest.ip"`
	AlarmTimestamp             float64  `json:"alarmTimestamp"`
	PDevicePort                []int    `json:"p.device.port"`
	PTimestampTimezone         string   `json:"p.timestampTimezone"`
	PDestLongitude             string   `json:"p.dest.longitude"`
	PTransferOutboundHumansize string   `json:"p.transfer.outbound.humansize"`
	PTransferInboundHumansize  string   `json:"p.transfer.inbound.humansize"`
	PIntfName                  string   `json:"p.intf.name"`
	PDeviceMacVendor           string   `json:"p.device.macVendor"`
	PCloudDecision             string   `json:"p.cloud.decision"`
	Aid                        string   `json:"aid"`
	PTransferDuration          string   `json:"p.transfer.duration"`
	PProtocol                  string   `json:"p.protocol"`
	Type                       string   `json:"type"`
	PFi                        string   `json:"p.fi"`
	PTagNames                  []struct {
		UID  string `json:"uid"`
		Name string `json:"name"`
	} `json:"p.tag.names"`
	Gid                        string    `json:"gid"`
	BoxName                    string    `json:"boxName"`
	BoxModel                   string    `json:"boxModel"`
	PTagName                   string    `json:"p.tag.name"`
	TimeToGenerate             string    `json:"time.to.generate"`
	EDestIPRange               string    `json:"e.dest.ip.range"`
	EDestIPCidr                string    `json:"e.dest.ip.cidr"`
	EDestIPOrg                 string    `json:"e.dest.ip.org"`
	EDestIPCountry             string    `json:"e.dest.ip.country"`
	EDestIPCity                string    `json:"e.dest.ip.city"`
	EDestDomain                string    `json:"e.dest.domain"`
	EDestDomainCreatedDate     time.Time `json:"e.dest.domain.createdDate"`
	EDestDomainLastUpdatedDate time.Time `json:"e.dest.domain.lastUpdatedDate"`
	EDestDomainRegister        string    `json:"e.dest.domain.register"`
	EDestSslSubject            string    `json:"e.dest.ssl.subject"`
	EDestSslServerName         string    `json:"e.dest.ssl.server_name"`
	EDestSslCN                 string    `json:"e.dest.ssl.CN"`
	EDestSslO                  string    `json:"e.dest.ssl.O"`
	ETransfer                  string    `json:"e.transfer"`
	PDeviceOnline              bool      `json:"p.device.online"`
	Intf                       struct {
		UUID string `json:"uuid"`
		Name string `json:"name"`
	} `json:"intf"`
	Tags []struct {
		UID    string `json:"uid"`
		Name   string `json:"name"`
		Policy struct {
		} `json:"policy"`
		Gid     string `json:"gid"`
		Devices []struct {
			IP    string   `json:"ip"`
			Mac   string   `json:"mac"`
			Names []string `json:"names"`
			Name  string   `json:"name"`
		} `json:"devices"`
	} `json:"tags"`
	PDevicePortInfo []any `json:"p.device.port.info"`
	PDestPortInfo   []struct {
		Description string `json:"description"`
		Name        string `json:"name"`
		Protocol    string `json:"protocol"`
		Port        string `json:"port"`
	} `json:"p.dest.port.info"`
}

type FirewallaDeviceDetail struct {
	Name string `json:"name"`
	IP   string `json:"ip"`
	Mac  string `json:"mac"`
	Gid  string `json:"gid"`
	Intf struct {
		UUID string `json:"uuid"`
		Name string `json:"name"`
	} `json:"intf"`
	MacVendor     string `json:"macVendor"`
	Online        bool   `json:"online"`
	TotalDownload int    `json:"totalDownload"`
	TotalUpload   int    `json:"totalUpload"`
	Reserved      bool   `json:"reserved"`
	LastActive    string `json:"lastActive"`
	FirstFound    string `json:"firstFound"`
	Tags          []struct {
		Name string `json:"name"`
		UID  string `json:"uid"`
	} `json:"tags"`
	BoxName           string `json:"boxName"`
	LocalDomain       string `json:"localDomain"`
	LocalDomainSuffix string `json:"localDomainSuffix"`
	Last12Months      struct {
		Upload        [][]int `json:"upload"`
		TotalUpload   int     `json:"totalUpload"`
		Download      [][]int `json:"download"`
		TotalDownload int     `json:"totalDownload"`
	} `json:"last12Months"`
	Last30 struct {
		Upload        [][]int `json:"upload"`
		TotalUpload   int     `json:"totalUpload"`
		Download      [][]int `json:"download"`
		TotalDownload int     `json:"totalDownload"`
	} `json:"last30"`
	Last60 struct {
		Upload        [][]int `json:"upload"`
		TotalUpload   int     `json:"totalUpload"`
		Download      [][]int `json:"download"`
		TotalDownload int     `json:"totalDownload"`
	} `json:"last60"`
	Last24 struct {
		Upload        [][]int `json:"upload"`
		TotalUpload   int     `json:"totalUpload"`
		Download      [][]int `json:"download"`
		TotalDownload int     `json:"totalDownload"`
	} `json:"last24"`
	Policy struct {
		Monitor bool `json:"monitor"`
	} `json:"policy"`
	Flows struct {
		CategoryDetails struct {
		} `json:"categoryDetails"`
		Download []struct {
			Device     string   `json:"device"`
			Port       []string `json:"port"`
			Fd         string   `json:"fd"`
			Count      int      `json:"count"`
			IP         string   `json:"ip"`
			Begin      int      `json:"begin"`
			End        int      `json:"end"`
			Country    string   `json:"country"`
			Host       string   `json:"host"`
			DeviceName string   `json:"deviceName"`
			MacVendor  string   `json:"macVendor"`
			Tags       []string `json:"tags"`
			TagIds     []string `json:"tagIds"`
			PortInfo   struct {
				Description string `json:"description"`
				Name        string `json:"name"`
				Protocol    string `json:"protocol"`
				Port        string `json:"port"`
			} `json:"portInfo"`
			Type string `json:"type"`
		} `json:"download"`
		Recent []struct {
			Ltype       string   `json:"ltype"`
			Ts          float64  `json:"ts"`
			Fd          string   `json:"fd"`
			Count       int      `json:"count"`
			Duration    float64  `json:"duration"`
			Intf        string   `json:"intf"`
			Tags        []string `json:"tags"`
			Protocol    string   `json:"protocol"`
			Port        int      `json:"port"`
			DevicePort  int      `json:"devicePort"`
			IP          string   `json:"ip"`
			DeviceIP    string   `json:"deviceIP"`
			Upload      int      `json:"upload"`
			Download    int      `json:"download"`
			Device      string   `json:"device"`
			Country     string   `json:"country"`
			Host        string   `json:"host"`
			DeviceName  string   `json:"deviceName"`
			MacVendor   string   `json:"macVendor"`
			TagIds      []string `json:"tagIds"`
			NetworkName string   `json:"networkName"`
			OnWan       bool     `json:"onWan"`
			IntfInfo    struct {
				Name string `json:"name"`
				Type string `json:"type"`
				UUID string `json:"uuid"`
			} `json:"intfInfo"`
			PortInfo struct {
				Description string `json:"description"`
				Name        string `json:"name"`
				Protocol    string `json:"protocol"`
				Port        string `json:"port"`
			} `json:"portInfo"`
			DevicePortInfo struct {
				Protocol string `json:"protocol"`
			} `json:"devicePortInfo"`
			Type string `json:"type"`
		} `json:"recent"`
		Upload []struct {
			Device     string   `json:"device"`
			Port       []string `json:"port"`
			Fd         string   `json:"fd"`
			Count      int      `json:"count"`
			IP         string   `json:"ip"`
			Begin      int      `json:"begin"`
			End        int      `json:"end"`
			Country    string   `json:"country"`
			Host       string   `json:"host"`
			DeviceName string   `json:"deviceName"`
			MacVendor  string   `json:"macVendor"`
			Tags       []string `json:"tags"`
			TagIds     []string `json:"tagIds"`
			PortInfo   struct {
				Description string `json:"description"`
				Name        string `json:"name"`
				Protocol    string `json:"protocol"`
				Port        string `json:"port"`
			} `json:"portInfo"`
			Type string `json:"type"`
		} `json:"upload"`
	} `json:"flows"`
	NewLast24 struct {
		Upload        [][]int `json:"upload"`
		TotalUpload   int     `json:"totalUpload"`
		Download      [][]int `json:"download"`
		TotalDownload int     `json:"totalDownload"`
		Conn          [][]int `json:"conn"`
		TotalConn     int     `json:"totalConn"`
		IPB           [][]int `json:"ipB"`
		TotalIPB      int     `json:"totalIpB"`
		DNS           [][]int `json:"dns"`
		TotalDNS      int     `json:"totalDns"`
		DNSB          [][]int `json:"dnsB"`
		TotalDNSB     int     `json:"totalDnsB"`
	} `json:"newLast24"`
	Ipv6 []any `json:"ipv6"`
}

type FirewallaFlowlogDetail []struct {
	Count        int      `json:"count"`
	Device       string   `json:"device"`
	Download     int      `json:"download,omitempty"`
	Duration     float64  `json:"duration,omitempty"`
	Host         string   `json:"host"`
	IP           string   `json:"ip"`
	Ts           float64  `json:"ts"`
	Country      string   `json:"country"`
	DevicePort   string   `json:"devicePort"`
	Intf         string   `json:"intf"`
	Total        int      `json:"total,omitempty"`
	Type         string   `json:"type"`
	DeviceIP     string   `json:"deviceIP"`
	Fd           string   `json:"fd"`
	Gid          string   `json:"gid"`
	GidExtracted string   `json:"gid_extracted"`
	Ltype        string   `json:"ltype"`
	Port         string   `json:"port"`
	Protocol     string   `json:"protocol"`
	Upload       int      `json:"upload,omitempty"`
	DeviceName   string   `json:"deviceName"`
	MacVendor    string   `json:"macVendor"`
	Tags         []string `json:"tags"`
	TagIds       []string `json:"tagIds"`
	NetworkName  string   `json:"networkName"`
	OnWan        bool     `json:"onWan"`
	IntfInfo     struct {
		Name string `json:"name"`
		Type string `json:"type"`
		UUID string `json:"uuid"`
	} `json:"intfInfo"`
	PortInfo struct {
		Description string `json:"description"`
		Name        string `json:"name"`
		Protocol    string `json:"protocol"`
		Port        string `json:"port"`
	} `json:"portInfo"`
	DevicePortInfo struct {
		Protocol string `json:"protocol"`
	} `json:"devicePortInfo"`
	Pid       string `json:"pid,omitempty"`
	WanIntf   string `json:"wanIntf,omitempty"`
	Blocked   bool   `json:"blocked,omitempty"`
	BlockPid  string `json:"blockPid,omitempty"`
	BlockType string `json:"blockType,omitempty"`
	Category  string `json:"category,omitempty"`
	App       string `json:"app,omitempty"`
}

type FirewallaFlowlogBody struct {
	Startts int64 `json:"startts"`
	Endts   int64 `json:"endts"`
}
