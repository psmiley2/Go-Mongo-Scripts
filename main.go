/*
Decription: Removes all documents from a mongodb collection

Usage:
Flag: -url=URLOFDB (will default to localhost if not flag is specified)
First Arg: name of table
Second Arg: name of collection
*/

package main

import (
	"flag"
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
)

func main() {
	urlPtr := flag.String("url", "localhost", "URL of the mongo DB.")
	flag.Parse()

	session, err := mgo.Dial(*urlPtr)
	if err != nil {
		log.Fatal(err)
	}

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	table := flag.Arg(0)
	collection := flag.Arg(1)
	c := session.DB(table).C(collection)

	c.RemoveAll(nil)

	m := fmt.Sprintf("Deleted collection %s from table %s", collection, table)
	fmt.Println(m)

}
