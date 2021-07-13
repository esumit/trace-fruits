# Trace Fruits 

###### About 

![Traceability](/images/Traceability_2068763.png) ![Fruits](/images/Fruits_1896737.png) 

Trace fruits is a conceptual demo to visualise fruits supply chain events stored in an immutable ledger.  It uses splunk to analyse 
generated data from hyperledger-fabric's blockchain ledgers.


This demo concept picked from the GS1 global traceability standard, that provides guidelines on fresh fruit and vegetable 
traceability. It includes following major supply chain events for traceability : 

-  Harvesting - The producer harvests the crop and packs the products into in cases. 
-  Packing/Repacking - The packer/repacker transforms ungraded commodities into products. 
-  Shipping - The packer/repacker palletizes the cases of product.
-  Transporting - The carrier arrives and loads the pallets onto the truck. 
-  Receiving - The pallets arrive in the retailer or foodservice operator’s distribution centre.
-  Selling or Consumption - The products have arrived at the retail store and have been placed on the shelves.

_harvesting-traceability event as stored in ledger_
 
````
{
  "_id": "0e8a4c14-02f2-42b5-939c-a8de8d2c7db6",
  "_rev": "1-c1b69236a1b423ec7806837549c32c8c",
  "attributes": {
    "event-when": "2021-05-24 18:47:22.151667 +1000 AEST",
    "event-where": "Johannesburg, South Africa",
    "event-why": "harvesting-traceability",
    "expiration-datetime": "2021-12-24 18:47:22.1515 +1000 AEST",
    "fruit-name": "Apricots",
    "fruit-type": "deciduous",
    "quantity": 313,
    "weight": 212
  },
  "header": {
    "batch-lot": "AB-123",
    "event-id": "0e8a4c14-02f2-42b5-939c-a8de8d2c7db6",
    "gln": "9501101530911",
    "gln-location-type": "planting",
    "gtin": "09501101530003",
    "gtin-trade-id-type": "crate-trade-id",
    "trace-id": "11111111"
  },
  "metaInfo": {
    "createdAt": "2021-06-26T18:45:04.140564684Z",
    "docType": "harvesting"
  },
  "~version": "\u0000CgMBBAA="
}
````
A trace-id is unique for set of product's supply chain events. 

#### Demo Recording

https://tinyurl.com/trace-fruits-demo 


#### How conceptually it works ?

- In every logical step of fruits supply chain, a fresh fruit supply chain event with unique trace-id got inserted into the blockchain ledger e.g.
harvesting event or shipping event. This event contains most important logical information to identify and prove this 
supply chain event.

- The inserted unique trace-ids relates to a batch of fruits products, and that trace-id inserted on each and every event
from start to finish of that batch of fruits products e.g. from harvesting to retail. 

- A complete audit trail of events from start to finish can be queried from ledger via trace-id 

- Individuals events can be queried from ledger 

- events inserted along the way of supply chain 


#### How To 

##### Step-1 : Clone this repository to linux flavour machine e.g.  Ubuntu 20.04.2 LTS

##### Step-2 : Install Prerequisites of hyperledger fabric 

##### Step-3 : Install Samples, Binaries, and Docker Images 

Execute the command to pull down the binaries and images

```
curl -sSL http://bit.ly/2ysbOFE | bash -s -- 1.4.6 1.4.6 0.4.18
```

- hyperledger/fabric-baseos 0.4.18
- hyperledger/fabric-ca 1.4.6
- hyperledger/fabric-orderer 1.4.6
- hyperledger/fabric-peer 1.4.6


##### Step-5 : Setup network, necessary dockers  , orgs, peer , channels , Splunk 

Inside this repo, call start shell script. It will setup necessary dockers images e.g. couch, 
Splunk, fabric network, orgs, peers, channels

```
./start.sh

```

![After a successful run of start.sh](/images/Start-sh-successful-output.png)


##### Step-6 : Run init ledgers from command line 

````
./init_ledgers.sh
````  

It will insert 18 supply events into the via trace_fruits_cc chaincode to the ledgers.

![After a successful run of init-ledgers.sh](/images/Init-ledgers-sh-successful.png)

##### Step-7 : Open couchdb Fauxton , a native web-based interface built into CouchDB :

````
http://<ip address>:5984/_utils/
````   

 Click on databases 
 Click on fruitsreceiver_trace_fruits_cc
 
 ![After a successful init of events in world state](/images/couch-db-database-enteries.png)
 ![After a successful init of events in world state](/images/After_init_successful_enteries_in_couchdb.png)
 ![After a successful init of events in world state](/images/A-recordin-couch-db.png)


##### Step-9 : Run Send Txns to continuously send events to Trace Fruits ledgers

````
./send_txns.sh
````

![After a successful run of send txns](/images/Send-Txns-Successfully-Sending-Events.png)

##### Step-10 : Open splunk enterprise :
 
 ```
 http://<ip-address>:8000/en-US/account/login 
 
 username : admin 
 Password : changeme 
  
 if local then ip address is 127.0.0.1 
 ```  
 
 Open Splunk App for Hyperledger Fabric 
 Click on Trace Fruits On the Side Bar
 
 ![Click on Trace Fruits On the Side Bar](/images/On-Splunk-Enterprise-Refer-To-Trace-Fruits.png)
 ![Navigate on the About Page of Trace Fruits](/images/Dashboard-Page-Of-Trace-Fruits.png)
 ![Sample Data Analysis of Ledger generated data](/images/Trace-Fruit-Ledgers.png)
 ![Sample Data Analysis Of Trace Fruits Events Payloads.png](/images/Trace-Fruits-Events-Payloads.png)
  


###### To shutdown the environment 

`./stop.sh`.

 ![On successful run of stop sh](/images/Stop-sh-successful.png)
 
 ##### Note
 
 - .checkpoints is an empty file -> ./.checkpoints:/usr/src/app/.checkpoints in docker-compose-splunk.yaml
 - splunk-apps is an empty dir   -> - ./splunk-apps:/opt/splunk/etc/apps in docker-compose-splunk.yaml
 - trace-fruits-apps contains trace-fruits splunk app available tar.gz, and a regular folder
 - trace-fruits chaincode available in trace-fruits-cc directory
 - trace-fruits-splunk-app.tar.gz also available in s3 https://trace-fruits.s3.ap-southeast-2.amazonaws.com/trace-fruits-splunk-app.tar.gz
 - Audit trail query is not implemented yet
 - Tar file of Trace-Fruits-1 doesn't work if created from mac,  
 
 #### References 
 
 - https://hyperledger-fabric.readthedocs.io/en/release-2.2/
 - https://github.com/splunk/fabric-logger
 - https://github.com/splunk/fabric-logger/tree/master/examples/vaccine-demo
 - https://splunkbase.splunk.com/app/4605/
 - https://www.splunkdlt.com/
 
