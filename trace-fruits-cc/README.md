# traceability-apis
A Sample Demo Traceability APIS for Fresh Fruit Supply Chain on Blockchain. 

*Note: Its just to explain the concept of Fresh Fruit  traceability from harvesing to the selling for consumtion 
on blockchain as immutable database. It doesn't have any relationship with actual fruit-chain*  

###### Harvesting: 

The producer harvests the crop and packs the products into in cases. Each of the cases gets a label with GTIN + batch/lot ID, and the related data are recorded.

######  Packing/Repacking: 

The packer/repacker transforms ungraded commodities into products. After that, the packer/repacker packs the products into cases.
To maintain traceability the inputs and outputs of the process are recorded on batch/lot level.

######  Shipping: 

The packer/repacker palletizes the cases of product. To maintain traceability the warehouse records the links between product IDs (GTIN + batch/lot ID) and pallet IDs (SSCC).

Subsequently, the pallets are moved to the outbound staging area to be collected by the carrier.

###### Transporting: 

The carrier arrives and loads the pallets onto the truck. The driver uses his mobile device to identify each of the pallets. The link between the pallets and the truck is recorded. Now, by tracking the truck also the pallets and goods can be tracked.

######  Receiving: 

The pallets arrive in the retailer or foodservice operator’s distribution centre.
The incoming goods department inspects the received goods by scanning the SSCCs on the pallet label and comparing the data against the pre-registered information in the system.
When all checks are ok, the goods will be marked as available in the inventory management system.

######  Selling or Consumption: 

The products have arrived at the retail store and have been placed on the shelves.
A consumer has decided to buy two products. At the checkout, the clerk scans the barcode on the products. The system automatically checks the expiry date.
The sales are recorded, in addition to the GTIN also the batch/lot ID is registered.
In foodservice, orders are placed by and shipped to Foodservice operators, the incoming goods department inspects the received goods by scanning the SSCCs on the pallet label or the GTINs on the individual cases and comparing the data against the waybill/invoice.
When all checks are ok, the goods will be marked as available in the inventory management system. Then it can be used (i.e. depleted).


Reference : 

 GS1 Global Traceability Standard 
