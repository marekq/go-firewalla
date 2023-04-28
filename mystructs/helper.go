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
	TotalDownload int    `json:"totalDownload,omitempty"`
	TotalUpload   int    `json:"totalUpload,omitempty"`
	Reserved      bool   `json:"reserved,omitempty"`
	LastActive    string `json:"lastActive"`
	FirstFound    string `json:"firstFound"`
	Tags          []struct {
		Name string `json:"name"`
		UID  string `json:"uid"`
	} `json:"tags,omitempty"`
	BoxName     string `json:"boxName"`
	IsFirewalla bool   `json:"isFirewalla,omitempty"`
}

type FirewallaAlarms []struct {
	Aid                        string   `json:"aid,omitempty"`
	Type                       string   `json:"type,omitempty"`
	Device                     string   `json:"device,omitempty"`
	AlarmTimestamp             float64  `json:"alarmTimestamp,omitempty"`
	Timestamp                  float64  `json:"timestamp,omitempty"`
	PDeviceIP                  string   `json:"p.device.ip,omitempty"`
	PDestIP                    string   `json:"p.dest.ip,omitempty"`
	PNoticeType                string   `json:"p.noticeType,omitempty"`
	PMessage                   string   `json:"p.message,omitempty"`
	PLocalIsClient             string   `json:"p.local_is_client,omitempty"`
	Message                    string   `json:"message,omitempty"`
	PDeviceName                string   `json:"p.device.name,omitempty"`
	PDeviceID                  string   `json:"p.device.id,omitempty"`
	PDeviceMac                 string   `json:"p.device.mac,omitempty"`
	PDeviceMacVendor           string   `json:"p.device.macVendor,omitempty"`
	PDestName                  string   `json:"p.dest.name,omitempty"`
	PDestLatitude              string   `json:"p.dest.latitude,omitempty"`
	PDestLongitude             string   `json:"p.dest.longitude,omitempty"`
	PDestCountry               string   `json:"p.dest.country,omitempty"`
	PCloudDecision             string   `json:"p.cloud.decision,omitempty"`
	PFi                        string   `json:"p.fi,omitempty"`
	Gid                        string   `json:"gid,omitempty"`
	BoxName                    string   `json:"boxName,omitempty"`
	BoxModel                   string   `json:"boxModel,omitempty"`
	PIntfName                  string   `json:"p.intf.name,omitempty"`
	PIntfID                    string   `json:"p.intf.id,omitempty"`
	PTagName                   string   `json:"p.tag.name,omitempty"`
	PTagIds                    []string `json:"p.tag.ids,omitempty"`
	TimeToGenerate             string   `json:"time.to.generate,omitempty"`
	PDestCategory              string   `json:"p.dest.category,omitempty"`
	PActionBlock               string   `json:"p.action.block,omitempty"`
	ResultMethod               string   `json:"result_method,omitempty"`
	PBlockby                   string   `json:"p.blockby,omitempty"`
	PDestReadableName          string   `json:"p.dest.readableName,omitempty"`
	Result                     string   `json:"result,omitempty"`
	ResultPolicy               string   `json:"result_policy,omitempty"`
	PSeverity                  string   `json:"p.severity,omitempty"`
	PDestDomain                string   `json:"p.dest.domain,omitempty"`
	PIntfSubnet                string   `json:"p.intf.subnet,omitempty"`
	PDestID                    string   `json:"p.dest.id,omitempty"`
	PDestPort                  string   `json:"p.dest.port,omitempty"`
	PTransferInboundSize       string   `json:"p.transfer.inbound.size,omitempty"`
	PTransferOutboundSize      string   `json:"p.transfer.outbound.size,omitempty"`
	PDevicePort                []int    `json:"p.device.port,omitempty"`
	PTimestampTimezone         string   `json:"p.timestampTimezone,omitempty"`
	PTransferOutboundHumansize string   `json:"p.transfer.outbound.humansize,omitempty"`
	PTransferInboundHumansize  string   `json:"p.transfer.inbound.humansize,omitempty"`
	PTransferDuration          string   `json:"p.transfer.duration,omitempty"`
	PProtocol                  string   `json:"p.protocol,omitempty"`
	PTagNames                  []any    `json:"p.tag.names,omitempty"`
	PTotalUsage                string   `json:"p.totalUsage,omitempty"`
	PBeginTs                   string   `json:"p.begin.ts,omitempty"`
	PFlows                     []struct {
		Count           int    `json:"count,omitempty"`
		IP              string `json:"ip,omitempty"`
		Device          string `json:"device,omitempty"`
		Country         string `json:"country,omitempty"`
		AggregationHost string `json:"aggregationHost,omitempty"`
		Host            string `json:"host,omitempty"`
		Category        string `json:"category,omitempty"`
		App             string `json:"app,omitempty"`
	} `json:"p.flows,omitempty"`
	PDuration            string `json:"p.duration,omitempty"`
	PPercentage          string `json:"p.percentage,omitempty"`
	PDestNames           string `json:"p.dest.names,omitempty"`
	PEndTs               string `json:"p.end.ts,omitempty"`
	PTotalUsageHumansize string `json:"p.totalUsage.humansize,omitempty"`
	PDestApp             string `json:"p.dest.app,omitempty"`
}

type FirewallaAlarmDetail struct {
	PDeviceIP                  string   `json:"p.device.ip,omitempty"`
	PDeviceID                  string   `json:"p.device.id,omitempty"`
	PDestDomain                string   `json:"p.dest.domain,omitempty"`
	PDestLatitude              string   `json:"p.dest.latitude,omitempty"`
	Timestamp                  float64  `json:"timestamp,omitempty"`
	PDeviceName                string   `json:"p.device.name,omitempty"`
	PTagIds                    []string `json:"p.tag.ids,omitempty"`
	PIntfSubnet                string   `json:"p.intf.subnet,omitempty"`
	PDestID                    string   `json:"p.dest.id,omitempty"`
	Message                    string   `json:"message,omitempty"`
	PDestPort                  string   `json:"p.dest.port,omitempty"`
	PDestName                  string   `json:"p.dest.name,omitempty"`
	PIntfID                    string   `json:"p.intf.id,omitempty"`
	PTransferInboundSize       string   `json:"p.transfer.inbound.size,omitempty"`
	PDeviceMac                 string   `json:"p.device.mac,omitempty"`
	PTransferOutboundSize      string   `json:"p.transfer.outbound.size,omitempty"`
	PDestCountry               string   `json:"p.dest.country,omitempty"`
	PLocalIsClient             string   `json:"p.local_is_client,omitempty"`
	Device                     string   `json:"device,omitempty"`
	PDestIP                    string   `json:"p.dest.ip,omitempty"`
	AlarmTimestamp             float64  `json:"alarmTimestamp,omitempty"`
	PDevicePort                []int    `json:"p.device.port,omitempty"`
	PTimestampTimezone         string   `json:"p.timestampTimezone,omitempty"`
	PDestLongitude             string   `json:"p.dest.longitude,omitempty"`
	PTransferOutboundHumansize string   `json:"p.transfer.outbound.humansize,omitempty"`
	PTransferInboundHumansize  string   `json:"p.transfer.inbound.humansize,omitempty"`
	PIntfName                  string   `json:"p.intf.name,omitempty"`
	PDeviceMacVendor           string   `json:"p.device.macVendor,omitempty"`
	PCloudDecision             string   `json:"p.cloud.decision,omitempty"`
	Aid                        string   `json:"aid,omitempty"`
	PTransferDuration          string   `json:"p.transfer.duration,omitempty"`
	PProtocol                  string   `json:"p.protocol,omitempty"`
	Type                       string   `json:"type,omitempty"`
	PFi                        string   `json:"p.fi,omitempty"`
	PTagNames                  []struct {
		UID  string `json:"uid,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"p.tag.names,omitempty"`
	Gid                        string    `json:"gid,omitempty"`
	BoxName                    string    `json:"boxName,omitempty"`
	BoxModel                   string    `json:"boxModel,omitempty"`
	PTagName                   string    `json:"p.tag.name,omitempty"`
	TimeToGenerate             string    `json:"time.to.generate,omitempty"`
	EDestIPRange               string    `json:"e.dest.ip.range,omitempty"`
	EDestIPCidr                string    `json:"e.dest.ip.cidr,omitempty"`
	EDestIPOrg                 string    `json:"e.dest.ip.org,omitempty"`
	EDestIPCountry             string    `json:"e.dest.ip.country,omitempty"`
	EDestIPCity                string    `json:"e.dest.ip.city,omitempty"`
	EDestDomain                string    `json:"e.dest.domain,omitempty"`
	EDestDomainCreatedDate     time.Time `json:"e.dest.domain.createdDate,omitempty"`
	EDestDomainLastUpdatedDate time.Time `json:"e.dest.domain.lastUpdatedDate,omitempty"`
	EDestDomainRegister        string    `json:"e.dest.domain.register,omitempty"`
	EDestSslSubject            string    `json:"e.dest.ssl.subject,omitempty"`
	EDestSslServerName         string    `json:"e.dest.ssl.server_name,omitempty"`
	EDestSslCN                 string    `json:"e.dest.ssl.CN,omitempty"`
	EDestSslO                  string    `json:"e.dest.ssl.O,omitempty"`
	ETransfer                  string    `json:"e.transfer,omitempty"`
	PDeviceOnline              bool      `json:"p.device.online,omitempty"`
	Intf                       struct {
		UUID string `json:"uuid,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"intf,omitempty"`
	Tags []struct {
		UID    string `json:"uid,omitempty"`
		Name   string `json:"name,omitempty"`
		Policy struct {
		} `json:"policy,omitempty"`
		Gid     string `json:"gid,omitempty"`
		Devices []struct {
			IP    string   `json:"ip,omitempty"`
			Mac   string   `json:"mac,omitempty"`
			Names []string `json:"names,omitempty"`
			Name  string   `json:"name,omitempty"`
		} `json:"devices,omitempty"`
	} `json:"tags,omitempty"`
	PDevicePortInfo []any `json:"p.device.port.info,omitempty"`
	PDestPortInfo   []struct {
		Description string `json:"description,omitempty"`
		Name        string `json:"name,omitempty"`
		Protocol    string `json:"protocol,omitempty"`
		Port        string `json:"port,omitempty"`
	} `json:"p.dest.port.info,omitempty"`
}
