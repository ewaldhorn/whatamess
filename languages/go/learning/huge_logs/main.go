package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/actforgood/bigcsvreader"
)

const noOfColumns = 29

var rowCount = 0

type ALBLog struct {
	ClientIP             string
	TargetProcessingTime string
	TargetStatusCode     string
	ReceivedBytes        string
	RequestUrl           string
	RequestCreationTime  string
	SSLProtocol          string
	TraceID              string
	RawRow               []string
}

type count32 int32

func (c *count32) inc() int32 {
	return atomic.AddInt32((*int32)(c), 1)
}
func (c *count32) get() int32 {
	return atomic.LoadInt32((*int32)(c))
}
func (c *count32) reset() {
	atomic.StoreInt32((*int32)(c), 0)
}

var counter count32

func main() {

	homedir, _ := os.UserHomeDir()
	csvLocation := homedir + "/logs/"
	entries, err := os.ReadDir(csvLocation)
	if err != nil {
		log.Fatal(err)
	}

	var TotalProcessRows int32
	var FilteredProcessedRows int

	f, err := os.Create("filtered.csv")
	w := csv.NewWriter(f)
	w.Write([]string{"request_url", "target_processing_time"})
	w.Flush()
	f.Close()

	for _, e := range entries {
		if strings.HasSuffix(e.Name(), ".log") {
			totalRows, filteredRows := processFile(csvLocation, e.Name())
			TotalProcessRows = TotalProcessRows + totalRows
			FilteredProcessedRows = FilteredProcessedRows + filteredRows
		}
	}

	fmt.Fprintf(os.Stdout, "Total Rows: %d Filtered Rows: %d", TotalProcessRows, FilteredProcessedRows)
}

func processFile(csvLocation string, fileName string) (int32, int) {

	counter.reset()
	// initialize the big csv reader
	bigCSV := bigcsvreader.New()
	bigCSV.SetFilePath(csvLocation + fileName)
	bigCSV.ColumnsCount = noOfColumns
	bigCSV.MaxGoroutinesNo = 16
	bigCSV.BufferSize = 81920
	bigCSV.FileHasHeader = true
	bigCSV.ColumnsDelimiter = ' '

	ctx, cancelCtx := context.WithCancel(context.Background())
	defer cancelCtx()
	var wg sync.WaitGroup
	var filteredRows int
	// start multi-thread reading
	rowsChans, errsChan := bigCSV.Read(ctx)
	resultChan := make(chan ALBLog)

	for i := 0; i < len(rowsChans); i++ {
		wg.Add(1)
		go rowWorker(rowsChans[i], &wg, resultChan)
	}

	wg.Add(1)
	go errWorker(errsChan, &wg)

	var writeCSVWaitGroup sync.WaitGroup
	go func() {
		writeCSVWaitGroup.Add(1)
		var result []ALBLog
		for filteredLog := range resultChan {
			result = append(result, filteredLog)
		}
		fmt.Fprintf(os.Stdout, "%s Filtered rows %d\n", fileName, len(result))
		filteredRows = len(result)

		f, err := os.OpenFile("filtered.csv", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		defer f.Close()

		if err != nil {
			log.Fatalln("failed to open file", err)
		}

		w := csv.NewWriter(f)
		defer w.Flush()

		if err != nil {
			log.Fatal(err)
		}

		for _, row := range result {
			w.Write(row.RawRow)
		}

		writeCSVWaitGroup.Done()

	}()

	wg.Wait()
	close(resultChan)

	writeCSVWaitGroup.Wait()
	fmt.Fprintf(os.Stdout, " %s Total rows: %d\n", fileName, counter.get())

	return counter.get(), filteredRows
}

func rowWorker(rowsChan bigcsvreader.RowsChan, waitGr *sync.WaitGroup, resultChan chan<- ALBLog) {
	for row := range rowsChan {
		processRow(row, resultChan)
	}
	waitGr.Done()
}

func errWorker(errsChan bigcsvreader.ErrsChan, waitGr *sync.WaitGroup) {
	for err := range errsChan {
		handleError(err)
	}
	waitGr.Done()
}

// processRow can be used to implement business logic
// like validation / converting to a struct / persisting row into a storage.
func processRow(row []string, resultChan chan<- ALBLog) {
	clientIP := strings.Split(row[3], ":")[0]
	requestUrl := strings.Split(row[12], " ")[1]
	urlObject, err := url.Parse(requestUrl)

	if err != nil {
		handleError(err)
	}

	counter.inc()

	RawRow := []string{urlObject.Path, row[6]}
	alblog := ALBLog{
		ClientIP:             clientIP,
		TargetProcessingTime: row[6],
		TargetStatusCode:     row[9],
		ReceivedBytes:        row[10],
		RequestUrl:           row[12],
		RequestCreationTime:  row[21],
		RawRow:               RawRow,
	}

	if doSomeFiltering() {
		resultChan <- alblog
	}
}

// handleError handles the error.
// errors can be fatal like file does not exist, or row related like a given row could not be parsed, etc...
func handleError(err error) {
	fmt.Println("error", err)
}
