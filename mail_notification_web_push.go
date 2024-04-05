package odoo

// MailNotificationWebPush represents mail.notification.web.push model.
type MailNotificationWebPush struct {
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Payload     *String   `xmlrpc:"payload,omitempty"`
	UserDevice  *Many2One `xmlrpc:"user_device,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// MailNotificationWebPushs represents array of mail.notification.web.push model.
type MailNotificationWebPushs []MailNotificationWebPush

// MailNotificationWebPushModel is the odoo model name.
const MailNotificationWebPushModel = "mail.notification.web.push"

// Many2One convert MailNotificationWebPush to *Many2One.
func (mnwp *MailNotificationWebPush) Many2One() *Many2One {
	return NewMany2One(mnwp.Id.Get(), "")
}

// CreateMailNotificationWebPush creates a new mail.notification.web.push model and returns its id.
func (c *Client) CreateMailNotificationWebPush(mnwp *MailNotificationWebPush) (int64, error) {
	ids, err := c.CreateMailNotificationWebPushs([]*MailNotificationWebPush{mnwp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailNotificationWebPush creates a new mail.notification.web.push model and returns its id.
func (c *Client) CreateMailNotificationWebPushs(mnwps []*MailNotificationWebPush) ([]int64, error) {
	var vv []interface{}
	for _, v := range mnwps {
		vv = append(vv, v)
	}
	return c.Create(MailNotificationWebPushModel, vv, nil)
}

// UpdateMailNotificationWebPush updates an existing mail.notification.web.push record.
func (c *Client) UpdateMailNotificationWebPush(mnwp *MailNotificationWebPush) error {
	return c.UpdateMailNotificationWebPushs([]int64{mnwp.Id.Get()}, mnwp)
}

// UpdateMailNotificationWebPushs updates existing mail.notification.web.push records.
// All records (represented by ids) will be updated by mnwp values.
func (c *Client) UpdateMailNotificationWebPushs(ids []int64, mnwp *MailNotificationWebPush) error {
	return c.Update(MailNotificationWebPushModel, ids, mnwp, nil)
}

// DeleteMailNotificationWebPush deletes an existing mail.notification.web.push record.
func (c *Client) DeleteMailNotificationWebPush(id int64) error {
	return c.DeleteMailNotificationWebPushs([]int64{id})
}

// DeleteMailNotificationWebPushs deletes existing mail.notification.web.push records.
func (c *Client) DeleteMailNotificationWebPushs(ids []int64) error {
	return c.Delete(MailNotificationWebPushModel, ids)
}

// GetMailNotificationWebPush gets mail.notification.web.push existing record.
func (c *Client) GetMailNotificationWebPush(id int64) (*MailNotificationWebPush, error) {
	mnwps, err := c.GetMailNotificationWebPushs([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mnwps)[0]), nil
}

// GetMailNotificationWebPushs gets mail.notification.web.push existing records.
func (c *Client) GetMailNotificationWebPushs(ids []int64) (*MailNotificationWebPushs, error) {
	mnwps := &MailNotificationWebPushs{}
	if err := c.Read(MailNotificationWebPushModel, ids, nil, mnwps); err != nil {
		return nil, err
	}
	return mnwps, nil
}

// FindMailNotificationWebPush finds mail.notification.web.push record by querying it with criteria.
func (c *Client) FindMailNotificationWebPush(criteria *Criteria) (*MailNotificationWebPush, error) {
	mnwps := &MailNotificationWebPushs{}
	if err := c.SearchRead(MailNotificationWebPushModel, criteria, NewOptions().Limit(1), mnwps); err != nil {
		return nil, err
	}
	return &((*mnwps)[0]), nil
}

// FindMailNotificationWebPushs finds mail.notification.web.push records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailNotificationWebPushs(criteria *Criteria, options *Options) (*MailNotificationWebPushs, error) {
	mnwps := &MailNotificationWebPushs{}
	if err := c.SearchRead(MailNotificationWebPushModel, criteria, options, mnwps); err != nil {
		return nil, err
	}
	return mnwps, nil
}

// FindMailNotificationWebPushIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailNotificationWebPushIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MailNotificationWebPushModel, criteria, options)
}

// FindMailNotificationWebPushId finds record id by querying it with criteria.
func (c *Client) FindMailNotificationWebPushId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailNotificationWebPushModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
