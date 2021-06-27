package trace

import (
	"time"
)

var FruitPayloads = []*TraceEvent{
	{
		MetaInfo: &DocMetaInfo{
			EventType: HarvestingEventDocType,
			CreatedAt: time.Now(),
		},

		Header: &Header{
			TraceID:  "11111111",
			EventID:  "0e8a4c14-02f2-42b5-939c-a8de8d2c7db6",
			GLN:      "9501101530911",
			GLNType:  HarvestingLocationType,
			GTIN:     "09501101530003",
			GTINType: CrateLevelTradeIDType,
			BatchLot: "AB-123",
		},
		Attributes: &Attributes{
			Quantity:       313,
			FruitType:      DeciduousFruitType,
			FruitName:      Apricots,
			TotalWeight:    212,
			EventWhy:       HarvestingEventDocType + "-traceability",
			EventWhen:      "2021-05-24 18:47:22.151667 +1000 AEST",
			EventWhere:     "Johannesburg, South Africa",
			ExpirationTime: "2021-12-24 18:47:22.1515 +1000 AEST",
		},
	},

	{
		MetaInfo: &DocMetaInfo{
			EventType: PackingRePackingEventDocType,
			CreatedAt: time.Now(),
		},

		Header: &Header{
			TraceID:  "11111111",
			EventID:  "a2743af3-68df-4648-840a-9a9b712b53dc",
			GLN:      "9501101530911",
			GLNType:  ProcessingLocationType,
			GTIN:     "09501101530003",
			GTINType: CrateLevelTradeIDType,
			BatchLot: "AB-123",
		},
		Attributes: &Attributes{
			Quantity:       313,
			FruitType:      DeciduousFruitType,
			FruitName:      Apricots,
			TotalWeight:    212,
			EventWhy:       PackingRePackingEventDocType + "-traceability",
			EventWhen:      "2021-05-26 18:47:22.151667 +1000 AEST",
			EventWhere:     "Durban, South Africa",
			ExpirationTime: "2021-12-24 18:47:22.1515 +1000 AEST",
		},
	},

	{
		MetaInfo: &DocMetaInfo{
			EventType: ShippingEventDocType,
			CreatedAt: time.Now(),
		},

		Header: &Header{
			TraceID:  "11111111",
			EventID:  "a6417988-eae6-432c-84b5-b4c6f9faf0cb",
			GLN:      "9501101530911",
			GLNType:  LoadingLocationType,
			GTIN:     "09501101530003",
			GTINType: PalletTradeIDType,
			BatchLot: "AB-123",
		},
		Attributes: &Attributes{
			Quantity:       313,
			FruitType:      DeciduousFruitType,
			FruitName:      Apricots,
			TotalWeight:    212,
			EventWhy:       ShippingEventDocType + "-traceability",
			EventWhen:      "2021-05-28 18:47:22.151667 +1000 AEST",
			EventWhere:     "Cape Town, South Africa",
			ExpirationTime: "2021-12-24 18:47:22.1515 +1000 AEST",
		},
	},

	{
		MetaInfo: &DocMetaInfo{
			EventType: TransportingDocType,
			CreatedAt: time.Now(),
		},

		Header: &Header{
			TraceID:  "11111111",
			EventID:  "49403a0c-dd5a-4e1f-8afe-add07c136de3",
			GLN:      "9501101530911",
			GLNType:  OutboundLocationType,
			GTIN:     "09501101530003",
			GTINType: PalletTradeIDType,
			BatchLot: "AB-123",
		},
		Attributes: &Attributes{
			Quantity:       313,
			FruitType:      DeciduousFruitType,
			FruitName:      Apricots,
			TotalWeight:    212,
			EventWhy:       TransportingDocType + "-traceability",
			EventWhen:      "2021-05-30 18:47:22.151667 +1000 AEST",
			EventWhere:     "Marseille, France",
			ExpirationTime: "2021-12-24 18:47:22.1515 +1000 AEST",
		},
	},

	{
		MetaInfo: &DocMetaInfo{
			EventType: ReceivingEventDocType,
			CreatedAt: time.Now(),
		},

		Header: &Header{
			TraceID:  "11111111",
			EventID:  "92a7b726-755a-443f-a053-10d2209d89c9",
			GLN:      "9501101530911",
			GLNType:  InboundLocationType,
			GTIN:     "09501101530003",
			GTINType: PalletTradeIDType,
			BatchLot: "AB-123",
		},
		Attributes: &Attributes{
			Quantity:       313,
			FruitType:      DeciduousFruitType,
			FruitName:      Apricots,
			TotalWeight:    212,
			EventWhy:       ReceivingEventDocType + "-traceability",
			EventWhen:      "2021-05-24 18:47:22.151667 +1000 AEST",
			EventWhere:     "Paris , France",
			ExpirationTime: "2021-12-24 18:47:22.1515 +1000 AEST",
		},
	},

	{
		MetaInfo: &DocMetaInfo{
			EventType: SellingConsumptionEventDocType,
			CreatedAt: time.Now(),
		},

		Header: &Header{
			TraceID:  "11111111",
			EventID:  "f201fe93-70f3-4596-97e0-272f58c453cc",
			GLN:      "9501101530911",
			GLNType:  StoreLocationType,
			GTIN:     "09501101530003",
			GTINType: CrateLevelTradeIDType,
			BatchLot: "AB-123",
		},
		Attributes: &Attributes{
			Quantity:       313,
			FruitType:      DeciduousFruitType,
			FruitName:      Apricots,
			TotalWeight:    212,
			EventWhy:       SellingConsumptionEventDocType + "-traceability",
			EventWhen:      "2021-06-24 18:47:22.151667 +1000 AEST",
			EventWhere:     "Madrid , Spain",
			ExpirationTime: "2021-12-24 18:47:22.1515 +1000 AEST",
		},
	},

	{
		MetaInfo: &DocMetaInfo{
			EventType: HarvestingEventDocType,
			CreatedAt: time.Now(),
		},

		Header: &Header{
			TraceID:  "22222222",
			EventID:  "4eff9fe6-ceb1-4ff5-a26e-3943e64f169e",
			GLN:      "9501101530911",
			GLNType:  HarvestingLocationType,
			GTIN:     "09501101530003",
			GTINType: CrateLevelTradeIDType,
			BatchLot: "AB-123",
		},
		Attributes: &Attributes{
			Quantity:       500,
			FruitType:      CitrusFruitType,
			FruitName:      Oranges,
			TotalWeight:    300,
			EventWhy:       HarvestingEventDocType + "-traceability",
			EventWhen:      "2021-03-16 18:47:22.151667 +1000 AEST",
			EventWhere:     "Johannesburg, South Africa",
			ExpirationTime: "2021-10-01 18:47:22.1515 +1000 AEST",
		},
	},

	{
		MetaInfo: &DocMetaInfo{
			EventType: PackingRePackingEventDocType,
			CreatedAt: time.Now(),
		},

		Header: &Header{
			TraceID:  "22222222",
			EventID:  "9e7a82d7-9659-4ae2-8996-a4a4da28692d",
			GLN:      "9501101530911",
			GLNType:  ProcessingLocationType,
			GTIN:     "09501101530003",
			GTINType: CrateLevelTradeIDType,
			BatchLot: "AB-123",
		},
		Attributes: &Attributes{
			Quantity:       500,
			FruitType:      CitrusFruitType,
			FruitName:      Oranges,
			TotalWeight:    300,
			EventWhy:       PackingRePackingEventDocType + "-traceability",
			EventWhen:      "2021-04-16 18:47:22.151667 +1000 AEST",
			EventWhere:     "Durban, South Africa",
			ExpirationTime: "2021-10-01 18:47:22.1515 +1000 AEST",
		},
	},

	{
		MetaInfo: &DocMetaInfo{
			EventType: ShippingEventDocType,
			CreatedAt: time.Now(),
		},

		Header: &Header{
			TraceID:  "22222222",
			EventID:  "aa3cff89-986c-4f8c-8c6f-337996758f86",
			GLN:      "9501101530911",
			GLNType:  LoadingLocationType,
			GTIN:     "09501101530003",
			GTINType: PalletTradeIDType,
			BatchLot: "AB-123",
		},
		Attributes: &Attributes{
			Quantity:       500,
			FruitType:      CitrusFruitType,
			FruitName:      Oranges,
			TotalWeight:    300,
			EventWhy:       ShippingEventDocType + "-traceability",
			EventWhen:      "2021-06-06 18:47:22.151667 +1000 AEST",
			EventWhere:     "Cape Town, South Africa",
			ExpirationTime: "2021-10-01 18:47:22.1515 +1000 AEST",
		},
	},

	{
		MetaInfo: &DocMetaInfo{
			EventType: TransportingDocType,
			CreatedAt: time.Now(),
		},

		Header: &Header{
			TraceID:  "22222222",
			EventID:  "45b4586f-a747-42aa-8145-ebe7ae4d5c05",
			GLN:      "9501101530911",
			GLNType:  OutboundLocationType,
			GTIN:     "09501101530003",
			GTINType: PalletTradeIDType,
			BatchLot: "AB-123",
		},
		Attributes: &Attributes{
			Quantity:       500,
			FruitType:      CitrusFruitType,
			FruitName:      Oranges,
			TotalWeight:    300,
			EventWhy:       TransportingDocType + "-traceability",
			EventWhen:      "2021-06-8 18:47:22.151667 +1000 AEST",
			EventWhere:     "Marseille, France",
			ExpirationTime: "2021-10-01 18:47:22.1515 +1000 AEST",
		},
	},

	{
		MetaInfo: &DocMetaInfo{
			EventType: ReceivingEventDocType,
			CreatedAt: time.Now(),
		},

		Header: &Header{
			TraceID:  "22222222",
			EventID:  "14604350-a9dd-4efa-b0e1-3996bf31e10f",
			GLN:      "9501101530911",
			GLNType:  InboundLocationType,
			GTIN:     "09501101530003",
			GTINType: PalletTradeIDType,
			BatchLot: "AB-123",
		},
		Attributes: &Attributes{
			Quantity:       500,
			FruitType:      CitrusFruitType,
			FruitName:      Oranges,
			TotalWeight:    300,
			EventWhy:       ReceivingEventDocType + "-traceability",
			EventWhen:      "2021-06-10 18:47:22.151667 +1000 AEST",
			EventWhere:     "Paris , France",
			ExpirationTime: "2021-10-01 18:47:22.1515 +1000 AEST",
		},
	},

	{
		MetaInfo: &DocMetaInfo{
			EventType: SellingConsumptionEventDocType,
			CreatedAt: time.Now(),
		},

		Header: &Header{
			TraceID:  "22222222",
			EventID:  "ef90e218-f917-459f-848f-d1f7dad3008f",
			GLN:      "9501101530911",
			GLNType:  StoreLocationType,
			GTIN:     "09501101530003",
			GTINType: CrateLevelTradeIDType,
			BatchLot: "AB-123",
		},
		Attributes: &Attributes{
			Quantity:       500,
			FruitType:      CitrusFruitType,
			FruitName:      Oranges,
			TotalWeight:    300,
			EventWhy:       SellingConsumptionEventDocType + "-traceability",
			EventWhen:      "2021-06-12 18:47:22.151667 +1000 AEST",
			EventWhere:     "Rome , Italy",
			ExpirationTime: "2021-10-01 18:47:22.1515 +1000 AEST",
		},
	},
	
	{
		MetaInfo: &DocMetaInfo{
			EventType: HarvestingEventDocType,
			CreatedAt: time.Now(),
		},
		
		Header: &Header{
			TraceID:  "33333333",
			EventID:  "dbe9235b-2c77-45f2-8590-58783edb9052",
			GLN:      "9501101530911",
			GLNType:  HarvestingLocationType,
			GTIN:     "09501101530003",
			GTINType: CrateLevelTradeIDType,
			BatchLot: "AB-123",
		},
		Attributes: &Attributes{
			Quantity:       600,
			FruitType:      SubtropicalFruitType,
			FruitName:      Mangoes,
			TotalWeight:    400,
			EventWhy:       HarvestingEventDocType + "-traceability",
			EventWhen:      "2021-02-06 8:47:22.151667 +1000 AEST",
			EventWhere:     "Johannesburg, South Africa",
			ExpirationTime: "2021-09-01 8:47:22.1515 +1000 AEST",
		},
	},
	
	{
		MetaInfo: &DocMetaInfo{
			EventType: PackingRePackingEventDocType,
			CreatedAt: time.Now(),
		},
		
		Header: &Header{
			TraceID:  "33333333",
			EventID:  "9e7a82d7-9659-4ae2-8996-a4a4da28692d",
			GLN:      "9501101530911",
			GLNType:  ProcessingLocationType,
			GTIN:     "09501101530003",
			GTINType: CrateLevelTradeIDType,
			BatchLot: "AB-123",
		},
		Attributes: &Attributes{
			Quantity:       600,
			FruitType:      SubtropicalFruitType,
			FruitName:      Mangoes,
			TotalWeight:    400,
			EventWhy:       PackingRePackingEventDocType + "-traceability",
			EventWhen:      "2021-02-08 8:47:22.151667 +1000 AEST",
			EventWhere:     "Durban, South Africa",
			ExpirationTime: "2021-09-01 8:47:22.1515 +1000 AEST",
		},
	},
	
	{
		MetaInfo: &DocMetaInfo{
			EventType: ShippingEventDocType,
			CreatedAt: time.Now(),
		},
		
		Header: &Header{
			TraceID:  "33333333",
			EventID:  "4bde384b-1fa6-4fd4-b68b-41dcb788ddfb",
			GLN:      "9501101530911",
			GLNType:  LoadingLocationType,
			GTIN:     "09501101530003",
			GTINType: PalletTradeIDType,
			BatchLot: "AB-123",
		},
		Attributes: &Attributes{
			Quantity:       600,
			FruitType:      SubtropicalFruitType,
			FruitName:      Mangoes,
			TotalWeight:    400,
			EventWhy:       ShippingEventDocType + "-traceability",
			EventWhen:      "2021-02-10 8:47:22.151667 +1000 AEST",
			EventWhere:     "Cape Town, South Africa",
			ExpirationTime: "2021-09-01 8:47:22.1515 +1000 AEST",
		},
	},
	
	{
		MetaInfo: &DocMetaInfo{
			EventType: TransportingDocType,
			CreatedAt: time.Now(),
		},
		
		Header: &Header{
			TraceID:  "33333333",
			EventID:  "acfb9a2a-ad0d-4350-b036-5e8958255c4d",
			GLN:      "9501101530911",
			GLNType:  OutboundLocationType,
			GTIN:     "09501101530003",
			GTINType: PalletTradeIDType,
			BatchLot: "AB-123",
		},
		Attributes: &Attributes{
			Quantity:       600,
			FruitType:      SubtropicalFruitType,
			FruitName:      Mangoes,
			TotalWeight:    400,
			EventWhy:       TransportingDocType + "-traceability",
			EventWhen:      "2021-02-12 8:47:22.151667 +1000 AEST",
			EventWhere:     "Marseille, France",
			ExpirationTime: "2021-09-01 8:47:22.1515 +1000 AEST",
		},
	},
	
	{
		MetaInfo: &DocMetaInfo{
			EventType: ReceivingEventDocType,
			CreatedAt: time.Now(),
		},
		
		Header: &Header{
			TraceID:  "33333333",
			EventID:  "00394017-11fd-4e9d-8797-1dfefd7d6571",
			GLN:      "9501101530911",
			GLNType:  InboundLocationType,
			GTIN:     "09501101530003",
			GTINType: PalletTradeIDType,
			BatchLot: "AB-123",
		},
		Attributes: &Attributes{
			Quantity:       600,
			FruitType:      SubtropicalFruitType,
			FruitName:      Mangoes,
			TotalWeight:    400,
			EventWhy:       ReceivingEventDocType + "-traceability",
			EventWhen:      "2021-02-14 8:47:22.151667 +1000 AEST",
			EventWhere:     "Paris , France",
			ExpirationTime: "2021-09-01 8:47:22.1515 +1000 AEST",
		},
	},
	
	{
		MetaInfo: &DocMetaInfo{
			EventType: SellingConsumptionEventDocType,
			CreatedAt: time.Now(),
		},
		
		Header: &Header{
			TraceID:  "33333333",
			EventID:  "9490fbea-26cb-464e-bf42-cccd4e3bbf47",
			GLN:      "9501101530911",
			GLNType:  StoreLocationType,
			GTIN:     "09501101530003",
			GTINType: CrateLevelTradeIDType,
			BatchLot: "AB-123",
		},
		Attributes: &Attributes{
			Quantity:       600,
			FruitType:      SubtropicalFruitType,
			FruitName:      Mangoes,
			TotalWeight:    400,
			EventWhy:       SellingConsumptionEventDocType + "-traceability",
			EventWhen:      "2021-02-18 8:47:22.151667 +1000 AEST",
			EventWhere:     "Zurich , Switzerland",
			ExpirationTime: "2021-09-01 8:47:22.1515 +1000 AEST",
		},
	},
}
