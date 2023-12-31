package client

type Credentials struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
	User         string `json:"user"`
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
}

type GetOperationsRequest struct {
	UserID      string
	OperationID string
}

type GetOperationsResponse struct {
	Operation Operation
}

type (
	OperationList []Operation

	Operation struct {
		Success                  string              `json:"success"`
		Hash                     string              `json:"hash"`
		UserName                 string              `json:"userName"`
		CustomerName             string              `json:"customerName"`
		Email                    string              `json:"email"`
		LocalTimezone            string              `json:"localTimezone"`
		Pan                      string              `json:"pan"`
		IssuerName               string              `json:"issuerName"`
		CurrencyCode             string              `json:"currencyCode"`
		PaymentMethod            string              `json:"paymentMethod"`
		DateTime                 string              `json:"dateTime"`
		PaymentMethodName        string              `json:"paymentMethodName"`
		PaymentMethodType        string              `json:"paymentMethodType"`
		Operation                string              `json:"operation"`
		AuthenticationType       string              `json:"authenticationType"`
		Province                 string              `json:"province"`
		Type                     string              `json:"type"`
		ShowName                 string              `json:"showName"`
		AvailableDate            string              `json:"availableDate"`
		Commission               string              `json:"commission"`
		CommissionTax            string              `json:"commissionTax"`
		IssuerImageURL           string              `json:"issuerImageUrl"`
		CardHolderName           string              `json:"cardHolderName"`
		TicketNumber             string              `json:"ticketNumber"`
		AuthorizationCode        string              `json:"authorizationCode"`
		CardHolderDocumentType   string              `json:"cardHolderDocumentType"`
		CardHolderDocument       string              `json:"cardHolderDocument"`
		ReferenceNumber          string              `json:"referenceNumber"`
		TaxText                  string              `json:"taxText"`
		PromotionName            string              `json:"promotionName"`
		PaymentButton            string              `json:"paymentButton"`
		TransactionDate          string              `json:"transactionDate"`
		InputMode                string              `json:"inputMode"`
		ProcessorReferenceNumber string              `json:"processorReferenceNumber"`
		Buyer                    Buyer               `json:"buyer"`
		Terminal                 TrxTerminal         `json:"terminal"`
		ID                       int64               `json:"id"`
		UserID                   int64               `json:"userId"`
		DateTimeTimestamp        int64               `json:"dateTimeTimestamp"`
		DateTimeTimestampUTC     int64               `json:"dateTimeTimestampUTC"`
		CreatedAt                int64               `json:"createdAt"`
		PaymentMethodID          int64               `json:"paymentMethodId"`
		BachClosureNumber        int64               `json:"bachClosureNumber"`
		TipAmount                float64             `json:"tipAmount"`
		TaxAmount                float64             `json:"taxAmount"`
		Total                    float64             `json:"total"`
		TotalGross               float64             `json:"totalGross"`
		TotalCash                float64             `json:"totalCash"`
		DiscountAmount           float64             `json:"discountAmount"`
		CommissionTaxAmount      float64             `json:"commissionTaxAmount"`
		CommissionAmount         float64             `json:"commissionAmount"`
		Installments             float64             `json:"installments"`
		InstallmentsAmount       float64             `json:"installmentsAmount"`
		InvoiceTotalAmount       float64             `json:"invoiceTotalAmount"`
		FinancialCostTax         float64             `json:"financialCostTax"`
		FinancialCostTaxAmount   float64             `json:"financialCostTaxAmount"`
		PaidWithDCC              float64             `json:"paidWithDcc"`
		Tax                      float64             `json:"tax"`
		SubtotalAmount           float64             `json:"subtotalAmount"`
		IsInverse                bool                `json:"isInverse"`
		CanAnnulate              bool                `json:"canAnnulate"`
		Canceled                 bool                `json:"canceled"`
		HasTip                   bool                `json:"hasTip"`
		HasRefund                bool                `json:"hasRefund"`
		AdjustTipAvailable       bool                `json:"adjustTipAvailable"`
		CanRefundWithoutCard     bool                `json:"canRefundWithoutCard"`
		CanRefund                bool                `json:"canRefund"`
		IsClose                  bool                `json:"isClose"`
		TaxReturnApply           bool                `json:"taxReturnApply"`
		CostPerInstallments      CostPerInstallments `json:"costPerInstallments"`
		SaleProducts             []SaleProduct       `json:"saleProducts"`
		WithholdingTaxes         []WithholdingTax    `json:"withholdingTaxes"`
		OperationCode            interface{}         `json:"operationCode"`
		CustomerID               interface{}         `json:"customerId"`
		Latitude                 interface{}         `json:"latitude"`
		Longitude                interface{}         `json:"longitude"`
		InvoiceNumber            interface{}         `json:"invoiceNumber"`
		AditionalNumber          interface{}         `json:"aditionalNumber"`
		TrxID                    interface{}         `json:"trxId"`
		Taxes                    []TaxGeopagos       `json:"taxes"`
		Refunds                  []interface{}       `json:"refunds"`
	}

	CostPerInstallments struct {
		Name          string  `json:"name"`
		Value         float64 `json:"value"`
		WithInterests bool    `json:"withInterests"`
	}

	Buyer struct {
		Email    string `json:"email"`
		Province string `json:"province"`
	}

	SaleProduct struct {
		ProductImageURL        string        `json:"productImageUrl"`
		ProductName            string        `json:"productName"`
		ProductBackgroundColor string        `json:"productBackgroundColor"`
		ID                     int64         `json:"id"`
		Quantity               float64       `json:"quantity"`
		DiscountAmount         float64       `json:"discountAmount"`
		Total                  float64       `json:"total"`
		UnitPrice              float64       `json:"unitPrice"`
		IsCustomAmount         bool          `json:"isCustomAmount"`
		Taxes                  []interface{} `json:"taxes"`
	}

	TrxTerminal struct {
		Type string `json:"type"`
		Icon string `json:"icon"`
	}

	WithholdingTax struct {
		Name          string  `json:"name"`
		Description   string  `json:"description"`
		TaxableIncome float64 `json:"taxableIncome"`
		Rate          float64 `json:"rate"`
		Withheld      float64 `json:"withheld"`
	}

	TaxGeopagos struct {
		Name   string  `json:"name"`
		Amount float64 `json:"amount"`
	}
)

func (p OperationList) Len() int {
	return len(p)
}

func (p OperationList) Less(i, j int) bool {
	return p[i].DateTimeTimestampUTC > p[j].DateTimeTimestampUTC
}

func (p OperationList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
