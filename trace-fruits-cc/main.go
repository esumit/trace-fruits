package main

import (
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/s7techlab/cckit/router"
	"github.com/s7techlab/cckit/router/param"
	"github.com/trace-fruits-cc/trace"
)

type TraceFruitsCC struct {
	router *router.Group
}

func (t *TraceFruitsCC) registerEventHarvesting(c router.Context) (interface{}, error) {
	rq := c.Param(`rq`).(trace.TraceEvent)
	fmt.Println("registerEventHarvesting: ", rq)

	rs := &trace.TraceEventRs{
		EventID: rq.Header.EventID,
		TraceId: rq.Header.TraceID,
	}

	txTimestamp, _ := c.Time()

	the := &trace.TraceEvent{
		Header:     rq.Header,
		Attributes: rq.Attributes,
		MetaInfo: &trace.DocMetaInfo{
			EventType: trace.HarvestingEventDocType,
			CreatedAt: txTimestamp,
		},
	}

	return rs,
		c.State().Insert(
			the.Header.EventID,
			the)
}

func (t *TraceFruitsCC) registerEventPacking(c router.Context) (interface{}, error) {
	rq := c.Param(`rq`).(trace.TraceEvent)
	fmt.Println("registerEventPacking: ", rq)

	rs := &trace.TraceEventRs{
		EventID: rq.Header.EventID,
		TraceId: rq.Header.TraceID,
	}

	txTimestamp, _ := c.Time()

	tpe := &trace.TraceEvent{
		Header:     rq.Header,
		Attributes: rq.Attributes,
		MetaInfo: &trace.DocMetaInfo{
			EventType: trace.PackingRePackingEventDocType,
			CreatedAt: txTimestamp,
		},
	}

	return rs,
		c.State().Insert(
			tpe.Header.EventID,
			tpe)
}

func (t *TraceFruitsCC) registerEventShipping(c router.Context) (interface{}, error) {
	rq := c.Param(`rq`).(trace.TraceEvent)
	fmt.Println("registerEventPacking: ", rq)

	rs := &trace.TraceEventRs{
		EventID: rq.Header.EventID,
		TraceId: rq.Header.TraceID,
	}

	txTimestamp, _ := c.Time()

	tse := &trace.TraceEvent{
		Header:     rq.Header,
		Attributes: rq.Attributes,
		MetaInfo: &trace.DocMetaInfo{
			EventType: trace.ShippingEventDocType,
			CreatedAt: txTimestamp,
		},
	}

	return rs,
		c.State().Insert(
			tse.Header.EventID,
			tse)
}

func (t *TraceFruitsCC) registerEventTransporting(c router.Context) (interface{}, error) {
	rq := c.Param(`rq`).(trace.TraceEvent)
	fmt.Println("registerEventPacking: ", rq)

	rs := &trace.TraceEventRs{
		EventID: rq.Header.EventID,
		TraceId: rq.Header.TraceID,
	}

	txTimestamp, _ := c.Time()

	tte := &trace.TraceEvent{
		Header:     rq.Header,
		Attributes: rq.Attributes,
		MetaInfo: &trace.DocMetaInfo{
			EventType: trace.TransportingDocType,
			CreatedAt: txTimestamp,
		},
	}

	return rs,
		c.State().Insert(
			tte.Header.EventID,
			tte)
}

func (t *TraceFruitsCC) registerEventReceiving(c router.Context) (interface{}, error) {
	rq := c.Param(`rq`).(trace.TraceEvent)
	fmt.Println("registerEventPacking: ", rq)

	rs := &trace.TraceEventRs{
		EventID: rq.Header.EventID,
		TraceId: rq.Header.TraceID,
	}

	txTimestamp, _ := c.Time()

	tre := &trace.TraceEvent{
		Header:     rq.Header,
		Attributes: rq.Attributes,
		MetaInfo: &trace.DocMetaInfo{
			EventType: trace.ReceivingEventDocType,
			CreatedAt: txTimestamp,
		},
	}

	return rs,
		c.State().Insert(
			tre.Header.EventID,
			tre)
}

func (t *TraceFruitsCC) registerEventSellingConsumption(c router.Context) (interface{}, error) {
	rq := c.Param(`rq`).(trace.TraceEvent)
	fmt.Println("registerEventPacking: ", rq)

	rs := &trace.TraceEventRs{
		EventID: rq.Header.EventID,
		TraceId: rq.Header.TraceID,
	}

	txTimestamp, _ := c.Time()

	tsc := &trace.TraceEvent{
		Header:     rq.Header,
		Attributes: rq.Attributes,
		MetaInfo: &trace.DocMetaInfo{
			EventType: trace.SellingConsumptionEventDocType,
			CreatedAt: txTimestamp,
		},
	}

	return rs,
		c.State().Insert(
			tsc.Header.EventID,
			tsc)
}

func (t *TraceFruitsCC) fetchEventChain(c router.Context) (interface{}, error) {
	rq := c.Param(`rq`).(trace.TraceEvent)
	fmt.Println("registerEventHarvesting: ", rq)

	return nil, nil
}

func (t *TraceFruitsCC) fetchEventSupplyById(c router.Context) (interface{}, error) {
	fmt.Println("fetchEventSupply: ")

	data, err := c.State().Get(
		c.ParamString(`trace_id`), &trace.TraceEvent{})

	if err != nil {
		return nil, err
	}

	te, ok := data.(trace.TraceEvent)

	if !ok {
		return nil, err
	}

	return &te, nil
}

// InitLedger adds a base set of assets to the ledger
func (t *TraceFruitsCC) traceEventInitLedger(c router.Context) (interface{}, error) {

	eventIds := []string{}
	for _, fruit := range trace.FruitPayloads {

		fruit.MetaInfo.CreatedAt, _ = c.Time()
		err := c.State().Insert(
			fruit.Header.EventID,
			fruit)

		if err != nil {
			return nil, fmt.Errorf("failed to put to world state. %v", err)
		}

		eventIds = append(eventIds, fruit.Header.EventID)
	}

	return eventIds, nil
}


func NewTraceFruitsCC() *router.Chaincode {
	r := router.New(`TraceFruits`)
	tf := &TraceFruitsCC{r}

	r.Init(func(c router.Context) (i interface{}, e error) {
		return nil, nil
	})

	r.Group(`EventHarvesting`).
		Invoke(`Register`, tf.registerEventHarvesting, param.Struct(`rq`, &trace.TraceEvent{}))

	r.Group(`EventPacking`).
		Invoke(`Register`, tf.registerEventPacking, param.Struct(`rq`, &trace.TraceEvent{}))

	r.Group(`EventShipping`).
		Invoke(`Register`, tf.registerEventShipping, param.Struct(`rq`, &trace.TraceEvent{}))

	r.Group(`EventTransporting`).
		Invoke(`Register`, tf.registerEventTransporting, param.Struct(`rq`, &trace.TraceEvent{}))

	r.Group(`EventReceiving`).
		Invoke(`Register`, tf.registerEventReceiving, param.Struct(`rq`, &trace.TraceEvent{}))

	r.Group(`EventSellingConsumption`).
		Invoke(`Register`, tf.registerEventSellingConsumption, param.Struct(`rq`, &trace.TraceEvent{}))

	r.Group(`EventChain`).
		Query(`ByTraceId`, tf.fetchEventChain, param.String(`trace_id`))

	r.Group(`EventSupplyFetch`).
		Query(`ByEventId`, tf.fetchEventSupplyById, param.String(`trace_id`))

	r.Group(`TraceEvent`).
		Invoke(`InitLedger`, tf.traceEventInitLedger)

	return router.NewChaincode(tf.router)
}

func main() {
	if err := shim.Start(NewTraceFruitsCC()); err != nil {
		fmt.Printf("Error starting TraceFruits chaincode: %s", err)
	}
}
