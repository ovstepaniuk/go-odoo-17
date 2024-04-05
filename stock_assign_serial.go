package odoo

// StockAssignSerial represents stock.assign.serial model.
type StockAssignSerial struct {
	CreateDate       *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid        *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName      *String   `xmlrpc:"display_name,omitempty"`
	Id               *Int      `xmlrpc:"id,omitempty"`
	MoveId           *Many2One `xmlrpc:"move_id,omitempty"`
	NextSerialCount  *Int      `xmlrpc:"next_serial_count,omitempty"`
	NextSerialNumber *String   `xmlrpc:"next_serial_number,omitempty"`
	ProductId        *Many2One `xmlrpc:"product_id,omitempty"`
	WriteDate        *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid         *Many2One `xmlrpc:"write_uid,omitempty"`
}

// StockAssignSerials represents array of stock.assign.serial model.
type StockAssignSerials []StockAssignSerial

// StockAssignSerialModel is the odoo model name.
const StockAssignSerialModel = "stock.assign.serial"

// Many2One convert StockAssignSerial to *Many2One.
func (sas *StockAssignSerial) Many2One() *Many2One {
	return NewMany2One(sas.Id.Get(), "")
}

// CreateStockAssignSerial creates a new stock.assign.serial model and returns its id.
func (c *Client) CreateStockAssignSerial(sas *StockAssignSerial) (int64, error) {
	ids, err := c.CreateStockAssignSerials([]*StockAssignSerial{sas})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockAssignSerial creates a new stock.assign.serial model and returns its id.
func (c *Client) CreateStockAssignSerials(sass []*StockAssignSerial) ([]int64, error) {
	var vv []interface{}
	for _, v := range sass {
		vv = append(vv, v)
	}
	return c.Create(StockAssignSerialModel, vv, nil)
}

// UpdateStockAssignSerial updates an existing stock.assign.serial record.
func (c *Client) UpdateStockAssignSerial(sas *StockAssignSerial) error {
	return c.UpdateStockAssignSerials([]int64{sas.Id.Get()}, sas)
}

// UpdateStockAssignSerials updates existing stock.assign.serial records.
// All records (represented by ids) will be updated by sas values.
func (c *Client) UpdateStockAssignSerials(ids []int64, sas *StockAssignSerial) error {
	return c.Update(StockAssignSerialModel, ids, sas, nil)
}

// DeleteStockAssignSerial deletes an existing stock.assign.serial record.
func (c *Client) DeleteStockAssignSerial(id int64) error {
	return c.DeleteStockAssignSerials([]int64{id})
}

// DeleteStockAssignSerials deletes existing stock.assign.serial records.
func (c *Client) DeleteStockAssignSerials(ids []int64) error {
	return c.Delete(StockAssignSerialModel, ids)
}

// GetStockAssignSerial gets stock.assign.serial existing record.
func (c *Client) GetStockAssignSerial(id int64) (*StockAssignSerial, error) {
	sass, err := c.GetStockAssignSerials([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sass)[0]), nil
}

// GetStockAssignSerials gets stock.assign.serial existing records.
func (c *Client) GetStockAssignSerials(ids []int64) (*StockAssignSerials, error) {
	sass := &StockAssignSerials{}
	if err := c.Read(StockAssignSerialModel, ids, nil, sass); err != nil {
		return nil, err
	}
	return sass, nil
}

// FindStockAssignSerial finds stock.assign.serial record by querying it with criteria.
func (c *Client) FindStockAssignSerial(criteria *Criteria) (*StockAssignSerial, error) {
	sass := &StockAssignSerials{}
	if err := c.SearchRead(StockAssignSerialModel, criteria, NewOptions().Limit(1), sass); err != nil {
		return nil, err
	}
	return &((*sass)[0]), nil
}

// FindStockAssignSerials finds stock.assign.serial records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockAssignSerials(criteria *Criteria, options *Options) (*StockAssignSerials, error) {
	sass := &StockAssignSerials{}
	if err := c.SearchRead(StockAssignSerialModel, criteria, options, sass); err != nil {
		return nil, err
	}
	return sass, nil
}

// FindStockAssignSerialIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockAssignSerialIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StockAssignSerialModel, criteria, options)
}

// FindStockAssignSerialId finds record id by querying it with criteria.
func (c *Client) FindStockAssignSerialId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockAssignSerialModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
