package mystructs

type JsonToken struct {
	Token string `json:"token"`
	Url   string `json:"url"`
}

type FirewallaDevices []struct {
	Date          string `json:"date"`
	Name          string `json:"name"`
	IP            string `json:"ip"`
	Mac           string `json:"mac"`
	Id            string `json:"id"`
	Gid           string `json:"gid"`
	MacVendor     string `json:"macVendor"`
	Online        bool   `json:"online"`
	LastSeen      string `json:"lastSeen"`
	TotalDownload int    `json:"totalDownload"`
	TotalUpload   int    `json:"totalUpload"`
	Network       struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"network"`
	Group struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"group"`
	ipReserved bool `json:"ipReserved"`
}

type FirewallaAlarms struct {
	Count   int `json:"count"`
	Results []struct {
		Date       string  `json:"date"`
		Aid        int     `json:"aid"`
		TypeId     int     `json:"type"`
		TypeName   string  `json:"_type"`
		Message    string  `json:"message"`
		Ts         float64 `json:"ts"`
		Gid        string  `json:"gid"`
		deviceType string  `json:"deviceType"`
		Device     struct {
			Name    string `json:"name"`
			MAC     string `json:"mac"`
			IP      string `json:"ip"`
			Network struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"network"`
			Group struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"group"`
		} `json:"device"`
		Remote struct {
			Domain     string `json:"domain"`
			Ip         string `json:"ip"`
			Port       string `json:"port"`
			Region     string `json:"region"`
			RootDomain string `json:"rootDomain"`
		} `json:"remote"`
		Transfer struct {
			Download int     `json:"download"`
			Upload   int     `json:"upload"`
			Total    int     `json:"total"`
			Duration float64 `json:"duration"`
		} `json:"transfer"`
	} `json:"results"`
	NextCursor *string `json:"next_cursor"`
}

type FirewallaAlarmDetail struct {
	Date       string  `json:"date"`
	Aid        int     `json:"aid"`
	TypeId     int     `json:"type"`
	TypeName   string  `json:"_type"`
	Message    string  `json:"message"`
	Ts         float64 `json:"ts"`
	Gid        string  `json:"gid"`
	deviceType string  `json:"deviceType"`
	Device     struct {
		Name    string `json:"name"`
		MAC     string `json:"mac"`
		IP      string `json:"ip"`
		Network struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"network"`
		Group struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"group"`
	} `json:"device"`
	Remote struct {
		Domain     string `json:"domain"`
		Ip         string `json:"ip"`
		Port       string `json:"port"`
		Region     string `json:"region"`
		RootDomain string `json:"rootDomain"`
	} `json:"remote"`
	Transfer struct {
		Download int     `json:"download"`
		Upload   int     `json:"upload"`
		Total    int     `json:"total"`
		Duration float64 `json:"duration"`
	} `json:"transfer"`
}

type FirewallaFlowlog struct {
	Results []struct {
		Date      string  `json:"date"`
		Ts        float64 `json:"ts"`
		Gid       string  `json:"gid"`
		Protocol  string  `json:"protocol"`
		Direction string  `json:"direction"`
		Block     bool    `json:"block"`
		BlockType *string `json:"blockType,omitempty"`
		Download  *int64  `json:"download,omitempty"`
		Upload    *int64  `json:"upload,omitempty"`
		Duration  *int64  `json:"duration,omitempty"`
		Count     int     `json:"count"`
		Device    struct {
			ID   string `json:"id"`
			IP   string `json:"ip"`
			Name string `json:"name"`
		} `json:"device"`
		Source struct {
			ID   string `json:"id"`
			Name string `json:"name"`
			IP   string `json:"ip"`
		} `json:"source,omitempty"`
		Destination struct {
			ID   string `json:"id"`
			Name string `json:"name"`
			IP   string `json:"ip"`
		} `json:"destination,omitempty"`
		Region   *string `json:"region,omitempty"`
		Category *string `json:"category,omitempty"`
		Network  struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"network"`
	} `json:"results"`
	Count      int    `json:"count"`
	NextCursor string `json:"next_cursor"`
}
