package main

import (
	"log"

	"github.com/kytart/goav/avcodec"
	"github.com/kytart/goav/avdevice"
	"github.com/kytart/goav/avfilter"
	"github.com/kytart/goav/avformat"
	"github.com/kytart/goav/avutil"
	"github.com/kytart/goav/swresample"
	"github.com/kytart/goav/swscale"
)

func main() {

	// Register all formats and codecs
	avformat.AvRegisterAll()
	avcodec.AvcodecRegisterAll()

	log.Printf("AvFilter Version:\t%v", avfilter.AvfilterVersion())
	log.Printf("AvDevice Version:\t%v", avdevice.AvdeviceVersion())
	log.Printf("SWScale Version:\t%v", swscale.SwscaleVersion())
	log.Printf("AvUtil Version:\t%v", avutil.AvutilVersion())
	log.Printf("AvCodec Version:\t%v", avcodec.AvcodecVersion())
	log.Printf("Resample Version:\t%v", swresample.SwresampleLicense())

}
