//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"time"
)

type Invoice struct {
	InvoiceID         int32 `sql:"primary_key"`
	CustomerID        int32
	InvoiceDate       time.Time
	BillingAddress    *string
	BillingCity       *string
	BillingState      *string
	BillingCountry    *string
	BillingPostalCode *string
	Total             float64
}
