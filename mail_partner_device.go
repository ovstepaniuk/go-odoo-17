package odoo

// MailPartnerDevice represents mail.partner.device model.
type MailPartnerDevice struct {
	CreateDate     *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid      *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName    *String   `xmlrpc:"display_name,omitempty"`
	Endpoint       *String   `xmlrpc:"endpoint,omitempty"`
	ExpirationTime *Time     `xmlrpc:"expiration_time,omitempty"`
	Id             *Int      `xmlrpc:"id,omitempty"`
	Keys           *String   `xmlrpc:"keys,omitempty"`
	PartnerId      *Many2One `xmlrpc:"partner_id,omitempty"`
	WriteDate      *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid       *Many2One `xmlrpc:"write_uid,omitempty"`
}

// MailPartnerDevices represents array of mail.partner.device model.
type MailPartnerDevices []MailPartnerDevice

// MailPartnerDeviceModel is the odoo model name.
const MailPartnerDeviceModel = "mail.partner.device"

// Many2One convert MailPartnerDevice to *Many2One.
func (mpd *MailPartnerDevice) Many2One() *Many2One {
	return NewMany2One(mpd.Id.Get(), "")
}

// CreateMailPartnerDevice creates a new mail.partner.device model and returns its id.
func (c *Client) CreateMailPartnerDevice(mpd *MailPartnerDevice) (int64, error) {
	ids, err := c.CreateMailPartnerDevices([]*MailPartnerDevice{mpd})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailPartnerDevice creates a new mail.partner.device model and returns its id.
func (c *Client) CreateMailPartnerDevices(mpds []*MailPartnerDevice) ([]int64, error) {
	var vv []interface{}
	for _, v := range mpds {
		vv = append(vv, v)
	}
	return c.Create(MailPartnerDeviceModel, vv, nil)
}

// UpdateMailPartnerDevice updates an existing mail.partner.device record.
func (c *Client) UpdateMailPartnerDevice(mpd *MailPartnerDevice) error {
	return c.UpdateMailPartnerDevices([]int64{mpd.Id.Get()}, mpd)
}

// UpdateMailPartnerDevices updates existing mail.partner.device records.
// All records (represented by ids) will be updated by mpd values.
func (c *Client) UpdateMailPartnerDevices(ids []int64, mpd *MailPartnerDevice) error {
	return c.Update(MailPartnerDeviceModel, ids, mpd, nil)
}

// DeleteMailPartnerDevice deletes an existing mail.partner.device record.
func (c *Client) DeleteMailPartnerDevice(id int64) error {
	return c.DeleteMailPartnerDevices([]int64{id})
}

// DeleteMailPartnerDevices deletes existing mail.partner.device records.
func (c *Client) DeleteMailPartnerDevices(ids []int64) error {
	return c.Delete(MailPartnerDeviceModel, ids)
}

// GetMailPartnerDevice gets mail.partner.device existing record.
func (c *Client) GetMailPartnerDevice(id int64) (*MailPartnerDevice, error) {
	mpds, err := c.GetMailPartnerDevices([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mpds)[0]), nil
}

// GetMailPartnerDevices gets mail.partner.device existing records.
func (c *Client) GetMailPartnerDevices(ids []int64) (*MailPartnerDevices, error) {
	mpds := &MailPartnerDevices{}
	if err := c.Read(MailPartnerDeviceModel, ids, nil, mpds); err != nil {
		return nil, err
	}
	return mpds, nil
}

// FindMailPartnerDevice finds mail.partner.device record by querying it with criteria.
func (c *Client) FindMailPartnerDevice(criteria *Criteria) (*MailPartnerDevice, error) {
	mpds := &MailPartnerDevices{}
	if err := c.SearchRead(MailPartnerDeviceModel, criteria, NewOptions().Limit(1), mpds); err != nil {
		return nil, err
	}
	return &((*mpds)[0]), nil
}

// FindMailPartnerDevices finds mail.partner.device records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailPartnerDevices(criteria *Criteria, options *Options) (*MailPartnerDevices, error) {
	mpds := &MailPartnerDevices{}
	if err := c.SearchRead(MailPartnerDeviceModel, criteria, options, mpds); err != nil {
		return nil, err
	}
	return mpds, nil
}

// FindMailPartnerDeviceIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailPartnerDeviceIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MailPartnerDeviceModel, criteria, options)
}

// FindMailPartnerDeviceId finds record id by querying it with criteria.
func (c *Client) FindMailPartnerDeviceId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailPartnerDeviceModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
