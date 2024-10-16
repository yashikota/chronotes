package model

type Language struct {
    Name         string  `json:"name"`
    TotalSeconds float64 `json:"total_seconds"`
    Digital      string  `json:"digital"`
    Decimal      string  `json:"decimal"`
    Text         string  `json:"text"`
    Hours        int     `json:"hours"`
    Minutes      int     `json:"minutes"`
    Seconds      int     `json:"seconds"`
    Percent      float64 `json:"percent"`
}

type GrandTotal struct {
    Hours        int     `json:"hours"`
    Minutes      int     `json:"minutes"`
    TotalSeconds float64 `json:"total_seconds"`
    Digital      string  `json:"digital"`
    Decimal      string  `json:"decimal"`
    Text         string  `json:"text"`
}

type Editor struct {
    Name         string  `json:"name"`
    TotalSeconds float64 `json:"total_seconds"`
    Digital      string  `json:"digital"`
    Decimal      string  `json:"decimal"`
    Text         string  `json:"text"`
    Hours        int     `json:"hours"`
    Minutes      int     `json:"minutes"`
    Seconds      int     `json:"seconds"`
    Percent      float64 `json:"percent"`
}

type OperatingSystem struct {
    Name         string  `json:"name"`
    TotalSeconds float64 `json:"total_seconds"`
    Digital      string  `json:"digital"`
    Decimal      string  `json:"decimal"`
    Text         string  `json:"text"`
    Hours        int     `json:"hours"`
    Minutes      int     `json:"minutes"`
    Seconds      int     `json:"seconds"`
    Percent      float64 `json:"percent"`
}

type Category struct {
    Name         string  `json:"name"`
    TotalSeconds float64 `json:"total_seconds"`
    Digital      string  `json:"digital"`
    Decimal      string  `json:"decimal"`
    Text         string  `json:"text"`
    Hours        int     `json:"hours"`
    Minutes      int     `json:"minutes"`
    Seconds      int     `json:"seconds"`
    Percent      float64 `json:"percent"`
}

type Dependency struct {
    Name         string  `json:"name"`
    TotalSeconds float64 `json:"total_seconds"`
    Digital      string  `json:"digital"`
    Decimal      string  `json:"decimal"`
    Text         string  `json:"text"`
    Hours        int     `json:"hours"`
    Minutes      int     `json:"minutes"`
    Seconds      int     `json:"seconds"`
    Percent      float64 `json:"percent"`
}

type Data struct {
    Languages       []Language       `json:"languages"`
    GrandTotal      GrandTotal       `json:"grand_total"`
    Editors         []Editor         `json:"editors"`
    OperatingSystems []OperatingSystem `json:"operating_systems"`
    Categories      []Category       `json:"categories"`
    Dependencies    []Dependency     `json:"dependencies"`
}

type WakatimeResponse struct {
    Data []Data `json:"data"`
}

type LanguageSummary struct {
    Name    string  `json:"name"`
    Time   float64 `json:"time"`
}