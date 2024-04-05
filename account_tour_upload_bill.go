package odoo

// AccountTourUploadBill represents account.tour.upload.bill model.
type AccountTourUploadBill struct {
	AttachmentIds  *Relation  `xmlrpc:"attachment_ids,omitempty"`
	CreateDate     *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid      *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName    *String    `xmlrpc:"display_name,omitempty"`
	Id             *Int       `xmlrpc:"id,omitempty"`
	PreviewInvoice *String    `xmlrpc:"preview_invoice,omitempty"`
	Selection      *Selection `xmlrpc:"selection,omitempty"`
	WriteDate      *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid       *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AccountTourUploadBills represents array of account.tour.upload.bill model.
type AccountTourUploadBills []AccountTourUploadBill

// AccountTourUploadBillModel is the odoo model name.
const AccountTourUploadBillModel = "account.tour.upload.bill"

// Many2One convert AccountTourUploadBill to *Many2One.
func (atub *AccountTourUploadBill) Many2One() *Many2One {
	return NewMany2One(atub.Id.Get(), "")
}

// CreateAccountTourUploadBill creates a new account.tour.upload.bill model and returns its id.
func (c *Client) CreateAccountTourUploadBill(atub *AccountTourUploadBill) (int64, error) {
	ids, err := c.CreateAccountTourUploadBills([]*AccountTourUploadBill{atub})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountTourUploadBill creates a new account.tour.upload.bill model and returns its id.
func (c *Client) CreateAccountTourUploadBills(atubs []*AccountTourUploadBill) ([]int64, error) {
	var vv []interface{}
	for _, v := range atubs {
		vv = append(vv, v)
	}
	return c.Create(AccountTourUploadBillModel, vv, nil)
}

// UpdateAccountTourUploadBill updates an existing account.tour.upload.bill record.
func (c *Client) UpdateAccountTourUploadBill(atub *AccountTourUploadBill) error {
	return c.UpdateAccountTourUploadBills([]int64{atub.Id.Get()}, atub)
}

// UpdateAccountTourUploadBills updates existing account.tour.upload.bill records.
// All records (represented by ids) will be updated by atub values.
func (c *Client) UpdateAccountTourUploadBills(ids []int64, atub *AccountTourUploadBill) error {
	return c.Update(AccountTourUploadBillModel, ids, atub, nil)
}

// DeleteAccountTourUploadBill deletes an existing account.tour.upload.bill record.
func (c *Client) DeleteAccountTourUploadBill(id int64) error {
	return c.DeleteAccountTourUploadBills([]int64{id})
}

// DeleteAccountTourUploadBills deletes existing account.tour.upload.bill records.
func (c *Client) DeleteAccountTourUploadBills(ids []int64) error {
	return c.Delete(AccountTourUploadBillModel, ids)
}

// GetAccountTourUploadBill gets account.tour.upload.bill existing record.
func (c *Client) GetAccountTourUploadBill(id int64) (*AccountTourUploadBill, error) {
	atubs, err := c.GetAccountTourUploadBills([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*atubs)[0]), nil
}

// GetAccountTourUploadBills gets account.tour.upload.bill existing records.
func (c *Client) GetAccountTourUploadBills(ids []int64) (*AccountTourUploadBills, error) {
	atubs := &AccountTourUploadBills{}
	if err := c.Read(AccountTourUploadBillModel, ids, nil, atubs); err != nil {
		return nil, err
	}
	return atubs, nil
}

// FindAccountTourUploadBill finds account.tour.upload.bill record by querying it with criteria.
func (c *Client) FindAccountTourUploadBill(criteria *Criteria) (*AccountTourUploadBill, error) {
	atubs := &AccountTourUploadBills{}
	if err := c.SearchRead(AccountTourUploadBillModel, criteria, NewOptions().Limit(1), atubs); err != nil {
		return nil, err
	}
	return &((*atubs)[0]), nil
}

// FindAccountTourUploadBills finds account.tour.upload.bill records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountTourUploadBills(criteria *Criteria, options *Options) (*AccountTourUploadBills, error) {
	atubs := &AccountTourUploadBills{}
	if err := c.SearchRead(AccountTourUploadBillModel, criteria, options, atubs); err != nil {
		return nil, err
	}
	return atubs, nil
}

// FindAccountTourUploadBillIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountTourUploadBillIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountTourUploadBillModel, criteria, options)
}

// FindAccountTourUploadBillId finds record id by querying it with criteria.
func (c *Client) FindAccountTourUploadBillId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountTourUploadBillModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
