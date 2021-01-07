package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"text/tabwriter"
	"time"
)

var (
	help = flag.Bool("h", false, "show this usage information")

	f = flag.String("f", "Mon Jan 2 15:04 MST", "time format layout")
)

func init() {
	log.SetFlags(0)
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS] TIMEZONE [TIMEZONE]\n\n", os.Args[0])
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()
	if *help {
		flag.Usage()
		return
	}
	now := time.Now()
	tz1 := ""
	tz2 := ""
	switch flag.NArg() {
	case 1:
		tz1 = now.Location().String()
		tz2 = flag.Arg(0)
	case 2:
		tz1 = flag.Arg(0)
		tz2 = flag.Arg(1)
	default:
		flag.Usage()
		os.Exit(2)
		return
	}
	err := compare(now, tz1, tz2)
	if err != nil {
		flag.Usage()
		os.Exit(1)
	}
}

func compare(now time.Time, tz1, tz2 string) error {
	now = now.Round(time.Hour)
	loc1, err := time.LoadLocation(tz1)
	if err != nil {
		return err
	}
	loc2, err := time.LoadLocation(tz2)
	if err != nil {
		return err
	}
	t1 := now.In(loc1)
	t2 := now.In(loc2)
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	fmt.Fprintf(w, "%s\t%s\t\n", tz1, tz2)
	for i := 0; i < 24; i++ {
		t1 = t1.Add(time.Hour)
		t2 = t2.Add(time.Hour)
		fmt.Fprintf(w, "%s\t%s\t\n", t1.Format(*f), t2.Format(*f))
	}
	return w.Flush()
}
