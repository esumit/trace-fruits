package trace

import "time"

type DocType = string

type DocMetaInfo struct {
	EventType DocType   `json:"docType,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

const (
	HarvestingEventDocType         DocType = "harvesting"
	PackingRePackingEventDocType   DocType = "packing-repacking"
	ShippingEventDocType           DocType = "shipping"
	TransportingDocType            DocType = "transporting"
	ReceivingEventDocType          DocType = "receiving"
	SellingConsumptionEventDocType DocType = "selling-consumption"
)

type LocationType = string

const (
	HarvestingLocationType      LocationType = "planting"
	ProcessingLocationType      LocationType = "packing"
	LoadingLocationType         LocationType = "loading"
	OutboundLocationType        LocationType = "outbound"
	InboundLocationType         LocationType = "inbound"
	StoreLocationType           LocationType = "store"
)

type TradeIDType = string

const (
	CrateLevelTradeIDType TradeIDType = "crate-trade-id"
	PalletTradeIDType     TradeIDType = "pallet-trade-id"
)

type FruitType = string

const (
	DeciduousFruitType   FruitType = "deciduous"
	CitrusFruitType      FruitType = "citrus"
	SubtropicalFruitType FruitType = "subtropical"
)

type Deciduous = string

const (
	Apricots   Deciduous = "Apricots"
	Peaches    Deciduous = "peaches"
	Nectarines Deciduous = "nectarines"
	Plums      Deciduous = "plums"
	Grapes     Deciduous = "grapes"
	Pears      Deciduous = "pears"
	Apples     Deciduous = "apples"
)

type Citrus = string

const (
	Oranges     Citrus = "oranges"
	Grapefruit  Citrus = "grapefruit"
	Lemons      Citrus = "lemons"
	EasyPeelers Citrus = "easy-peelers"
)

type SubTropical = string

const (
	Mangoes    SubTropical = "mangoes"
	Litchis    SubTropical = "litchis"
	Melons     SubTropical = "melons"
	Avocados   SubTropical = "avocados"
	PineApples SubTropical = "pine-apples"
)

type Header struct {
	TraceID  string `json:"trace-id,omitempty"`
	EventID  string `json:"event-id,omitempty"`
	GLN      string `json:"gln,omitempty"`
	GLNType  string `json:"gln-location-type,omitempty"`
	GTIN     string `json:"gtin,omitempty"`
	GTINType string `json:"gtin-trade-id-type,omitempty"`
	BatchLot string `json:"batch-lot,omitempty"`
}

type Attributes struct {
	Quantity       int     `json:"quantity,omitempty"`
	FruitType      string  `json:"fruit-type,omitempty"`
	FruitName      string  `json:"fruit-name,omitempty"`
	TotalWeight    float32 `json:"weight,omitempty"`
	EventWhy       string  `json:"event-why,omitempty"`
	EventWhen      string  `json:"event-when,omitempty"`
	EventWhere     string  `json:"event-where,omitempty"`
	ExpirationTime string  `json:"expiration-datetime,omitempty"`
}

type TraceEvent struct {
	MetaInfo   *DocMetaInfo `json:"metaInfo,omitempty"`
	Header     *Header      `json:"header,omitempty"`
	Attributes *Attributes  `json:"attributes,omitempty"`
}

type TraceEventRs struct {
	EventID string `json:"event-id,omitempty"`
	TraceId string `json:"trace-id,omitempty"`
}
