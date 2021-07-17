package main

import (
	_"fmt"
	"m2web"
	"chipbins"
	"binview"
)


func main() {
	m2web.ScrapeM2Web()
	chipbins.ScrapeChipBins()
	binview.ScrapeMyBinView()
}