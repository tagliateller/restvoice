package main

import (
	"github.com/rwirdemann/restvoice/kapitel06/database"
	"github.com/rwirdemann/restvoice/kapitel06/rest"
	"github.com/rwirdemann/restvoice/kapitel06/usecase"
)

func main() {
	repository := database.NewFakeRepository()
	adapter := rest.NewAdapter()

	createInvoice := usecase.NewCreateInvoice(repository)
	createInvoiceHandler := adapter.MakeCreateInvoiceHandler(createInvoice)
	adapter.HandleFunc("/customers/{customerId:[0-9]+}/invoices", createInvoiceHandler).Methods("POST")

	createBooking := usecase.NewCreateBooking(repository)
	createBookingHandler := adapter.MakeCreateBookingHandler(createBooking)
	adapter.HandleFunc("/customers/{customerId:[0-9]+}/invoices/{invoiceId:[0-9]+}/bookings",
		createBookingHandler).Methods("POST")

	updateInvoice := usecase.NewUpdateInvoice(repository)
	adapter.HandleFunc("/customers/{customerId:[0-9]+}/invoices/{invoiceId:[0-9]+}",
		adapter.MakeUpdateInvoiceHandler(updateInvoice)).Methods("PUT")

	getInvoice := usecase.NewGetInvoice(repository)
	adapter.HandleFunc("/customers/{customerId:[0-9]+}/invoices/{invoiceId:[0-9]+}",
		adapter.MakeGetInvoiceHandler(getInvoice)).Methods("GET")

	adapter.ListenAndServe()
}
