package models

type BTC_blacklist struct {
    Address string `json:"address"`
    FromSources int64 `json:"from_sources"`
    ProcedureType string `json:"procedure_type"`
    CrimeTags int64 `json:"crime_tags"`
	ActorCategory string `json:"actor_category"`
	SiteLink string `json:"site_link"`
	Price int `json:"price"`
	Email string `json:"email"`
	RawData string `json:"source_raw_data"`
	Confidence int `json:"confidence"`
	LastUpdatedTime int `json:"last_updated_time"`
	FirstSeenTime int `json:"first_seen_time"`
	Status int `json:"status"`
	Description string `json:"description"`

}

func GetAddress(pageNum int, pageSize int, maps interface {}) (addresses []BTC_blacklist) {
    db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&addresses)

    return
}

func GetAddressTotal(maps interface {}) (count int){
    db.Model(&BTC_blacklist{}).Where(maps).Count(&count)

    return
}
