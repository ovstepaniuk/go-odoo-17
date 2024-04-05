package odoo

// IapAccountInfo represents iap.account.info model.
type IapAccountInfo struct {
	AccountId         *Many2One `xmlrpc:"account_id,omitempty"`
	AccountToken      *String   `xmlrpc:"account_token,omitempty"`
	AccountUuidHashed *String   `xmlrpc:"account_uuid_hashed,omitempty"`
	Balance           *Float    `xmlrpc:"balance,omitempty"`
	CreateDate        *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid         *Many2One `xmlrpc:"create_uid,omitempty"`
	Description       *String   `xmlrpc:"description,omitempty"`
	DisplayName       *String   `xmlrpc:"display_name,omitempty"`
	Id                *Int      `xmlrpc:"id,omitempty"`
	ServiceName       *String   `xmlrpc:"service_name,omitempty"`
	UnitName          *String   `xmlrpc:"unit_name,omitempty"`
	WarnMe            *Bool     `xmlrpc:"warn_me,omitempty"`
	WarningEmail      *String   `xmlrpc:"warning_email,omitempty"`
	WarningThreshold  *Float    `xmlrpc:"warning_threshold,omitempty"`
	WriteDate         *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid          *Many2One `xmlrpc:"write_uid,omitempty"`
}

// IapAccountInfos represents array of iap.account.info model.
type IapAccountInfos []IapAccountInfo

// IapAccountInfoModel is the odoo model name.
const IapAccountInfoModel = "iap.account.info"

// Many2One convert IapAccountInfo to *Many2One.
func (iai *IapAccountInfo) Many2One() *Many2One {
	return NewMany2One(iai.Id.Get(), "")
}

// CreateIapAccountInfo creates a new iap.account.info model and returns its id.
func (c *Client) CreateIapAccountInfo(iai *IapAccountInfo) (int64, error) {
	ids, err := c.CreateIapAccountInfos([]*IapAccountInfo{iai})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateIapAccountInfo creates a new iap.account.info model and returns its id.
func (c *Client) CreateIapAccountInfos(iais []*IapAccountInfo) ([]int64, error) {
	var vv []interface{}
	for _, v := range iais {
		vv = append(vv, v)
	}
	return c.Create(IapAccountInfoModel, vv, nil)
}

// UpdateIapAccountInfo updates an existing iap.account.info record.
func (c *Client) UpdateIapAccountInfo(iai *IapAccountInfo) error {
	return c.UpdateIapAccountInfos([]int64{iai.Id.Get()}, iai)
}

// UpdateIapAccountInfos updates existing iap.account.info records.
// All records (represented by ids) will be updated by iai values.
func (c *Client) UpdateIapAccountInfos(ids []int64, iai *IapAccountInfo) error {
	return c.Update(IapAccountInfoModel, ids, iai, nil)
}

// DeleteIapAccountInfo deletes an existing iap.account.info record.
func (c *Client) DeleteIapAccountInfo(id int64) error {
	return c.DeleteIapAccountInfos([]int64{id})
}

// DeleteIapAccountInfos deletes existing iap.account.info records.
func (c *Client) DeleteIapAccountInfos(ids []int64) error {
	return c.Delete(IapAccountInfoModel, ids)
}

// GetIapAccountInfo gets iap.account.info existing record.
func (c *Client) GetIapAccountInfo(id int64) (*IapAccountInfo, error) {
	iais, err := c.GetIapAccountInfos([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*iais)[0]), nil
}

// GetIapAccountInfos gets iap.account.info existing records.
func (c *Client) GetIapAccountInfos(ids []int64) (*IapAccountInfos, error) {
	iais := &IapAccountInfos{}
	if err := c.Read(IapAccountInfoModel, ids, nil, iais); err != nil {
		return nil, err
	}
	return iais, nil
}

// FindIapAccountInfo finds iap.account.info record by querying it with criteria.
func (c *Client) FindIapAccountInfo(criteria *Criteria) (*IapAccountInfo, error) {
	iais := &IapAccountInfos{}
	if err := c.SearchRead(IapAccountInfoModel, criteria, NewOptions().Limit(1), iais); err != nil {
		return nil, err
	}
	return &((*iais)[0]), nil
}

// FindIapAccountInfos finds iap.account.info records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIapAccountInfos(criteria *Criteria, options *Options) (*IapAccountInfos, error) {
	iais := &IapAccountInfos{}
	if err := c.SearchRead(IapAccountInfoModel, criteria, options, iais); err != nil {
		return nil, err
	}
	return iais, nil
}

// FindIapAccountInfoIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIapAccountInfoIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(IapAccountInfoModel, criteria, options)
}

// FindIapAccountInfoId finds record id by querying it with criteria.
func (c *Client) FindIapAccountInfoId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IapAccountInfoModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
