package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

type Transaction struct {
	path         string
	date         string
	success      bool
	errorMessage string
}

type TransactionLog struct {
	transactions []Transaction
	file         *os.File
	writer       *csv.Writer
}

func openTransactionLog(path string) *TransactionLog {
	transactionLog := TransactionLog{
		transactions: make([]Transaction, 0),
	}

	f, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		log.Panic("Failed to open transaction log,", err)
	}
	transactionLog.file = f
	_, err = f.Seek(0, 0)
	if err != nil {
		log.Panic("Failed to seek to start of transaction log,", err)
	}

	r := csv.NewReader(f)

	recs, err := r.ReadAll()
	if err != nil {
		log.Panic("Failed to parse transaction log,", err)
	}

	for i, rec := range recs {
		if len(rec) < 4 {
			log.Panic("Transaction log record", i, "too short")
		}
		path := rec[0]
		success, err := strconv.ParseBool(rec[1])
		if err != nil {
			log.Panic("Failed to parse transaction log record", i, err)
		}
		date := rec[2]
		errorMessage := rec[3]

		trx := Transaction{
			path:         path,
			success:      success,
			date:         date,
			errorMessage: errorMessage,
		}

		transactionLog.transactions = append(transactionLog.transactions, trx)
	}

	transactionLog.file = f
	transactionLog.writer = csv.NewWriter(f)
	return &transactionLog
}

func (t *TransactionLog) addTransaction(path string, success bool, date string, errorMessage string) {

	trx := Transaction{
		path:         path,
		success:      success,
		date:         date,
		errorMessage: errorMessage,
	}
	t.transactions = append(t.transactions, trx)
	t.writer.Write([]string{
		path, strconv.FormatBool(success), date, errorMessage,
	})
	t.writer.Flush()
}

func (t *TransactionLog) find(path string) *Transaction {
	for _, trx := range t.transactions {
		if trx.path == path {
			return &trx
		}
	}
	return nil
}

func (t *TransactionLog) close() {
	t.file.Close()
}
