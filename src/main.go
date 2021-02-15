package main

/*
 * The contents of this file are subject to the terms of the
 * Common Development and Distribution License, Version 1.0 only
 * (the "License").  You may not use this file except in compliance
 * with the License.
 *
 * See the file LICENSE in this distribution for details.
 * A copy of the CDDL is also available via the Internet at
 * http://www.opensource.org/licenses/cddl1.txt
 *
 * When distributing Covered Code, include this CDDL HEADER in each
 * file and include the contents of the LICENSE file from this
 * distribution.
 */

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/araddon/dateparse"
	"github.com/rodaine/numwords"
)

func main() {
	// dateadder expects exactly one argument.
	if len(os.Args) != 2 {
		fmt.Println("dateadder expects exactly one parameter. read the readme.")
		os.Exit(1)
	}

	// Sanitize before using:
	formatstring := strings.Replace(os.Args[1], "  ", " ", -1)

	// We expect the formatstring to be in this format:
	//   <date> <plus/in/+> <num> <units>.
	// As both <date> and <num> could technically contain spaces,
	// dateadder needs to find the other two tokens first.
	split := strings.Split(formatstring, " ")

	plustoken := -1
	unitstoken := -1
	units := ""

	for i, token := range split {
		if token == "plus" || token == "in" || token == "+" {
			plustoken = i
			continue
		}

		if token == "day" || token == "days" || token == "week" || token == "weeks" || token == "month" || token == "months" || token == "year" || token == "years" {
			unitstoken = i
			units = token
			continue
		}
	}

	if plustoken == -1 || unitstoken == -1 {
		// Could not find the tokens.
		fmt.Println("Failed to find a plus token and/or the units to add where they were expected. Aborting.")
		os.Exit(1)
	}

	// So everything < <plustoken> is the date, everything > <plustoken> and < <unitstoken> is the number.
	lastdatetoken := plustoken - 1
	firstnumtoken := plustoken + 1
	lastnumtoken := unitstoken - 1

	if lastnumtoken == plustoken {
		// "plus days" or something.
		fmt.Printf("Failed to find the number of %s. Aborting.\n", units)
		os.Exit(1)
	}

	datestring := ""
	if lastdatetoken > 0 {
		datestring = strings.Join(split[:lastdatetoken], " ")
	} else if lastdatetoken == 0 {
		// Looks like there is no date.
		// Assume "today" by default.
		datestring = "today"
	} else {
		datestring = split[0]
	}
	
	numstring := ""
	if firstnumtoken == lastnumtoken {
		numstring = split[firstnumtoken]
	} else {
		numstring = strings.Join(split[firstnumtoken:lastnumtoken], " ")
	}

	date_in := time.Now()
	date_out := time.Now()

	if datestring == "yesterday" {
		date_in = date_in.AddDate(0, 0, -1)
	} else if datestring == "tomorrow" {
		date_in = date_in.AddDate(0, 0, 1)
	} else if datestring != "today" {
		// Try to parse the input date.
		d, err := dateparse.ParseLocal(datestring)
		if err != nil {
			fmt.Printf("dateadder has determined '%s' to be your input date. That does not seem valid. Aborting.\n", datestring)
			os.Exit(1)
		}

		date_in = d
	}

	num_in := 0
	if IsNumeric(numstring) {
		int, _ := strconv.Atoi(numstring)
		num_in = int
	} else {
		// Try to parse the input number.
		i, err := numwords.ParseInt(numstring)
		if err != nil {
			fmt.Printf("dateadder has determined '%s' to be your input number. That does not seem valid. Aborting.\n", numstring)
			os.Exit(1)
		}

		num_in = i
	}

	if units == "day" || units == "days" {
		date_out = date_out.AddDate(0, 0, num_in)
	} else if units == "week" || units == "weeks" {
		date_out = date_out.AddDate(0, 0, num_in*7)
	} else if units == "month" || units == "months" {
		date_out = date_out.AddDate(0, num_in, 0)
	} else if units == "year" || units == "years" {
		date_out = date_out.AddDate(num_in, 0, 0)
	}

	fmt.Printf("%s\n", date_out.Format("Monday, January _2 2006"))
}

func IsNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
