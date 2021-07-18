package main

import "log"

func main() {
	m2webRes := ScrapeM2Web()
	chipbinsRes := ScrapeChipBins()
	binViewRes := ScrapeMyBinView()

	InsertIntoDb(m2webRes, chipbinsRes, binViewRes)

	log.Println("Done!")
}
