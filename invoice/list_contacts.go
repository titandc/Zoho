package invoice

import (
	"fmt"
	"github.com/schmorrison/Zoho"
)

//https://www.zoho.com/invoice/api/v3/#Contacts_List_Contacts
//func (c *ZohoInvoiceAPI) ListContacts(request interface{}, organizationId string, params map[string]zoho.Parameter) (data ListContactsResponse, err error) {
func (c *ZohoInvoiceAPI) ListContacts() (data ListContactsResponse, err error) {

	// Renew token if necessary
	if c.Zoho.Token.CheckExpiry() {
		err := c.Zoho.RefreshTokenRequest()
		if err != nil {
			return ListContactsResponse{}, err
		}
	}

	endpoint := zoho.Endpoint{
		Name:         ContactsModule,
		URL:          fmt.Sprintf(InvoiceAPIEndPoint+"%s", ContactsModule),
		Method:       zoho.HTTPGet,
		ResponseData: &ListContactsResponse{},
		URLParameters: map[string]zoho.Parameter{
			"filter_by": "",
		},
	}

	/*for k, v := range params {
		endpoint.URLParameters[k] = v
	}
	*/

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return ListContactsResponse{}, fmt.Errorf("Failed to retrieve expense reports: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*ListContactsResponse); ok {
		// Check if the request succeeded
		if v.Code != 0 {
			return *v, fmt.Errorf("Failed to list contacts: %s", v.Message)
		}
		return *v, nil
	}
	return ListContactsResponse{}, fmt.Errorf("Data retrieved was not 'ListContactsResponse'")
}

// ListContactsResponse is the data returned by GetExpenseReports
type ListContactsResponse struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	Contacts []struct {
		ContactID                     string  `json:"contact_id"`
		ContactName                   string  `json:"contact_name"`
		CompanyName                   string  `json:"company_name"`
		ContactType                   string  `json:"contact_type"`
		Status                        string  `json:"status"`
		PaymentTerms                  int64   `json:"payment_terms"`
		PaymentTermsLabel             string  `json:"payment_terms_label"`
		CurrencyID                    string  `json:"currency_id"`
		CurrencyCode                  string  `json:"currency_code"`
		OutstandingReceivableAmount   float64 `json:"outstanding_receivable_amount"`
		UnusedCreditsReceivableAmount float64 `json:"unused_credits_receivable_amount"`
		FirstName                     string  `json:"first_name"`
		LastName                      string  `json:"last_name"`
		Email                         string  `json:"email"`
		Phone                         string  `json:"phone"`
		Mobile                        string  `json:"mobile"`
		CreatedTime                   string  `json:"created_time"`
		LastModifiedTime              string  `json:"last_modified_time"`
		/*CustomFields  []struct {
			CustomfieldID string `json:"customfield_id"`
			Label         string `json:"label"`
			Value         string `json:"value"`
		} `json:"custom_fields"`*/
	} `json:"contacts"`
}
