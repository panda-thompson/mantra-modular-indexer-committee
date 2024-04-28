package cases

import (
	"fmt"
	"github.com/RiemaLabs/modular-indexer-committee"
	"log"
	"testing"
	"time"

	"github.com/RiemaLabs/modular-indexer-committee/ord"
	"github.com/RiemaLabs/modular-indexer-committee/ord/stateless"
)

func Test_OPI(t *testing.T) {
	var latestHeight uint = stateless.BRC20StartHeight + ord.BitcoinConfirmations
	records, err := stateless.LoadOPIRecords("./data/785000-ordi.csv")
	if err != nil {
		log.Fatalf(fmt.Sprintf("error happened: %v", err))
	}
	ordGetterTest, arguments := loadMain(782000)
	queue, err := main.CatchupStage(ordGetterTest, &arguments, stateless.BRC20StartHeight-1, latestHeight)
	if err != nil {
		log.Fatalf(fmt.Sprintf("error happened: %v", err))
	}
	ordGetterTest.LatestBlockHeight = latestHeight
	go main.ServiceStage(ordGetterTest, &arguments, queue, 10*time.Millisecond)
	for {
		if ordGetterTest.LatestBlockHeight == queue.LatestHeight() {
			queue.Header.VerifyState(&records)
			log.Printf("Block: %d is verified!\n", ordGetterTest.LatestBlockHeight)
			ordGetterTest.LatestBlockHeight++
		}
		if ordGetterTest.LatestBlockHeight >= 780000 {
			break
		}
		time.Sleep(10 * time.Millisecond)
	}
}