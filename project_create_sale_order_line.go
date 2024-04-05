package odoo

// ProjectCreateSaleOrderLine represents project.create.sale.order.line model.
type ProjectCreateSaleOrderLine struct {
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	CurrencyId  *Many2One `xmlrpc:"currency_id,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	EmployeeId  *Many2One `xmlrpc:"employee_id,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	PriceUnit   *Float    `xmlrpc:"price_unit,omitempty"`
	ProductId   *Many2One `xmlrpc:"product_id,omitempty"`
	WizardId    *Many2One `xmlrpc:"wizard_id,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// ProjectCreateSaleOrderLines represents array of project.create.sale.order.line model.
type ProjectCreateSaleOrderLines []ProjectCreateSaleOrderLine

// ProjectCreateSaleOrderLineModel is the odoo model name.
const ProjectCreateSaleOrderLineModel = "project.create.sale.order.line"

// Many2One convert ProjectCreateSaleOrderLine to *Many2One.
func (pcsol *ProjectCreateSaleOrderLine) Many2One() *Many2One {
	return NewMany2One(pcsol.Id.Get(), "")
}

// CreateProjectCreateSaleOrderLine creates a new project.create.sale.order.line model and returns its id.
func (c *Client) CreateProjectCreateSaleOrderLine(pcsol *ProjectCreateSaleOrderLine) (int64, error) {
	ids, err := c.CreateProjectCreateSaleOrderLines([]*ProjectCreateSaleOrderLine{pcsol})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProjectCreateSaleOrderLine creates a new project.create.sale.order.line model and returns its id.
func (c *Client) CreateProjectCreateSaleOrderLines(pcsols []*ProjectCreateSaleOrderLine) ([]int64, error) {
	var vv []interface{}
	for _, v := range pcsols {
		vv = append(vv, v)
	}
	return c.Create(ProjectCreateSaleOrderLineModel, vv, nil)
}

// UpdateProjectCreateSaleOrderLine updates an existing project.create.sale.order.line record.
func (c *Client) UpdateProjectCreateSaleOrderLine(pcsol *ProjectCreateSaleOrderLine) error {
	return c.UpdateProjectCreateSaleOrderLines([]int64{pcsol.Id.Get()}, pcsol)
}

// UpdateProjectCreateSaleOrderLines updates existing project.create.sale.order.line records.
// All records (represented by ids) will be updated by pcsol values.
func (c *Client) UpdateProjectCreateSaleOrderLines(ids []int64, pcsol *ProjectCreateSaleOrderLine) error {
	return c.Update(ProjectCreateSaleOrderLineModel, ids, pcsol, nil)
}

// DeleteProjectCreateSaleOrderLine deletes an existing project.create.sale.order.line record.
func (c *Client) DeleteProjectCreateSaleOrderLine(id int64) error {
	return c.DeleteProjectCreateSaleOrderLines([]int64{id})
}

// DeleteProjectCreateSaleOrderLines deletes existing project.create.sale.order.line records.
func (c *Client) DeleteProjectCreateSaleOrderLines(ids []int64) error {
	return c.Delete(ProjectCreateSaleOrderLineModel, ids)
}

// GetProjectCreateSaleOrderLine gets project.create.sale.order.line existing record.
func (c *Client) GetProjectCreateSaleOrderLine(id int64) (*ProjectCreateSaleOrderLine, error) {
	pcsols, err := c.GetProjectCreateSaleOrderLines([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pcsols)[0]), nil
}

// GetProjectCreateSaleOrderLines gets project.create.sale.order.line existing records.
func (c *Client) GetProjectCreateSaleOrderLines(ids []int64) (*ProjectCreateSaleOrderLines, error) {
	pcsols := &ProjectCreateSaleOrderLines{}
	if err := c.Read(ProjectCreateSaleOrderLineModel, ids, nil, pcsols); err != nil {
		return nil, err
	}
	return pcsols, nil
}

// FindProjectCreateSaleOrderLine finds project.create.sale.order.line record by querying it with criteria.
func (c *Client) FindProjectCreateSaleOrderLine(criteria *Criteria) (*ProjectCreateSaleOrderLine, error) {
	pcsols := &ProjectCreateSaleOrderLines{}
	if err := c.SearchRead(ProjectCreateSaleOrderLineModel, criteria, NewOptions().Limit(1), pcsols); err != nil {
		return nil, err
	}
	return &((*pcsols)[0]), nil
}

// FindProjectCreateSaleOrderLines finds project.create.sale.order.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectCreateSaleOrderLines(criteria *Criteria, options *Options) (*ProjectCreateSaleOrderLines, error) {
	pcsols := &ProjectCreateSaleOrderLines{}
	if err := c.SearchRead(ProjectCreateSaleOrderLineModel, criteria, options, pcsols); err != nil {
		return nil, err
	}
	return pcsols, nil
}

// FindProjectCreateSaleOrderLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectCreateSaleOrderLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ProjectCreateSaleOrderLineModel, criteria, options)
}

// FindProjectCreateSaleOrderLineId finds record id by querying it with criteria.
func (c *Client) FindProjectCreateSaleOrderLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProjectCreateSaleOrderLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
