package odoo

// AccountMoveSend represents account.move.send model.
type AccountMoveSend struct {
	CheckboxDownload         *Bool       `xmlrpc:"checkbox_download,omitempty"`
	CheckboxSendByPost       *Bool       `xmlrpc:"checkbox_send_by_post,omitempty"`
	CheckboxSendMail         *Bool       `xmlrpc:"checkbox_send_mail,omitempty"`
	CheckboxUblCiiLabel      *String     `xmlrpc:"checkbox_ubl_cii_label,omitempty"`
	CheckboxUblCiiXml        *Bool       `xmlrpc:"checkbox_ubl_cii_xml,omitempty"`
	CompanyId                *Many2One   `xmlrpc:"company_id,omitempty"`
	CreateDate               *Time       `xmlrpc:"create_date,omitempty"`
	CreateUid                *Many2One   `xmlrpc:"create_uid,omitempty"`
	DisplayMailComposer      *Bool       `xmlrpc:"display_mail_composer,omitempty"`
	DisplayName              *String     `xmlrpc:"display_name,omitempty"`
	EnableDownload           *Bool       `xmlrpc:"enable_download,omitempty"`
	EnableSendByPost         *Bool       `xmlrpc:"enable_send_by_post,omitempty"`
	EnableSendMail           *Bool       `xmlrpc:"enable_send_mail,omitempty"`
	EnableUblCiiXml          *Bool       `xmlrpc:"enable_ubl_cii_xml,omitempty"`
	Id                       *Int        `xmlrpc:"id,omitempty"`
	MailAttachmentsWidget    interface{} `xmlrpc:"mail_attachments_widget,omitempty"`
	MailBody                 *String     `xmlrpc:"mail_body,omitempty"`
	MailLang                 *String     `xmlrpc:"mail_lang,omitempty"`
	MailPartnerIds           *Relation   `xmlrpc:"mail_partner_ids,omitempty"`
	MailSubject              *String     `xmlrpc:"mail_subject,omitempty"`
	MailTemplateId           *Many2One   `xmlrpc:"mail_template_id,omitempty"`
	Mode                     *Selection  `xmlrpc:"mode,omitempty"`
	MoveIds                  *Relation   `xmlrpc:"move_ids,omitempty"`
	SendByPostCost           *Int        `xmlrpc:"send_by_post_cost,omitempty"`
	SendByPostReadonly       *Bool       `xmlrpc:"send_by_post_readonly,omitempty"`
	SendByPostWarningMessage *String     `xmlrpc:"send_by_post_warning_message,omitempty"`
	SendMailReadonly         *Bool       `xmlrpc:"send_mail_readonly,omitempty"`
	SendMailWarningMessage   interface{} `xmlrpc:"send_mail_warning_message,omitempty"`
	ShowUblCompanyWarning    *Bool       `xmlrpc:"show_ubl_company_warning,omitempty"`
	UblPartnerWarning        *String     `xmlrpc:"ubl_partner_warning,omitempty"`
	WriteDate                *Time       `xmlrpc:"write_date,omitempty"`
	WriteUid                 *Many2One   `xmlrpc:"write_uid,omitempty"`
}

// AccountMoveSends represents array of account.move.send model.
type AccountMoveSends []AccountMoveSend

// AccountMoveSendModel is the odoo model name.
const AccountMoveSendModel = "account.move.send"

// Many2One convert AccountMoveSend to *Many2One.
func (ams *AccountMoveSend) Many2One() *Many2One {
	return NewMany2One(ams.Id.Get(), "")
}

// CreateAccountMoveSend creates a new account.move.send model and returns its id.
func (c *Client) CreateAccountMoveSend(ams *AccountMoveSend) (int64, error) {
	ids, err := c.CreateAccountMoveSends([]*AccountMoveSend{ams})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountMoveSend creates a new account.move.send model and returns its id.
func (c *Client) CreateAccountMoveSends(amss []*AccountMoveSend) ([]int64, error) {
	var vv []interface{}
	for _, v := range amss {
		vv = append(vv, v)
	}
	return c.Create(AccountMoveSendModel, vv, nil)
}

// UpdateAccountMoveSend updates an existing account.move.send record.
func (c *Client) UpdateAccountMoveSend(ams *AccountMoveSend) error {
	return c.UpdateAccountMoveSends([]int64{ams.Id.Get()}, ams)
}

// UpdateAccountMoveSends updates existing account.move.send records.
// All records (represented by ids) will be updated by ams values.
func (c *Client) UpdateAccountMoveSends(ids []int64, ams *AccountMoveSend) error {
	return c.Update(AccountMoveSendModel, ids, ams, nil)
}

// DeleteAccountMoveSend deletes an existing account.move.send record.
func (c *Client) DeleteAccountMoveSend(id int64) error {
	return c.DeleteAccountMoveSends([]int64{id})
}

// DeleteAccountMoveSends deletes existing account.move.send records.
func (c *Client) DeleteAccountMoveSends(ids []int64) error {
	return c.Delete(AccountMoveSendModel, ids)
}

// GetAccountMoveSend gets account.move.send existing record.
func (c *Client) GetAccountMoveSend(id int64) (*AccountMoveSend, error) {
	amss, err := c.GetAccountMoveSends([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*amss)[0]), nil
}

// GetAccountMoveSends gets account.move.send existing records.
func (c *Client) GetAccountMoveSends(ids []int64) (*AccountMoveSends, error) {
	amss := &AccountMoveSends{}
	if err := c.Read(AccountMoveSendModel, ids, nil, amss); err != nil {
		return nil, err
	}
	return amss, nil
}

// FindAccountMoveSend finds account.move.send record by querying it with criteria.
func (c *Client) FindAccountMoveSend(criteria *Criteria) (*AccountMoveSend, error) {
	amss := &AccountMoveSends{}
	if err := c.SearchRead(AccountMoveSendModel, criteria, NewOptions().Limit(1), amss); err != nil {
		return nil, err
	}
	return &((*amss)[0]), nil
}

// FindAccountMoveSends finds account.move.send records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountMoveSends(criteria *Criteria, options *Options) (*AccountMoveSends, error) {
	amss := &AccountMoveSends{}
	if err := c.SearchRead(AccountMoveSendModel, criteria, options, amss); err != nil {
		return nil, err
	}
	return amss, nil
}

// FindAccountMoveSendIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountMoveSendIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountMoveSendModel, criteria, options)
}

// FindAccountMoveSendId finds record id by querying it with criteria.
func (c *Client) FindAccountMoveSendId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountMoveSendModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
