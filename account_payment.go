package odoo

// AccountPayment represents account.payment model.
type AccountPayment struct {
	AccessToken                           *String     `xmlrpc:"access_token,omitempty"`
	AccessUrl                             *String     `xmlrpc:"access_url,omitempty"`
	AccessWarning                         *String     `xmlrpc:"access_warning,omitempty"`
	ActivityCalendarEventId               *Many2One   `xmlrpc:"activity_calendar_event_id,omitempty"`
	ActivityDateDeadline                  *Time       `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration           *Selection  `xmlrpc:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon                 *String     `xmlrpc:"activity_exception_icon,omitempty"`
	ActivityIds                           *Relation   `xmlrpc:"activity_ids,omitempty"`
	ActivityState                         *Selection  `xmlrpc:"activity_state,omitempty"`
	ActivitySummary                       *String     `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeIcon                      *String     `xmlrpc:"activity_type_icon,omitempty"`
	ActivityTypeId                        *Many2One   `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId                        *Many2One   `xmlrpc:"activity_user_id,omitempty"`
	AlwaysTaxExigible                     *Bool       `xmlrpc:"always_tax_exigible,omitempty"`
	Amount                                *Float      `xmlrpc:"amount,omitempty"`
	AmountAvailableForRefund              *Float      `xmlrpc:"amount_available_for_refund,omitempty"`
	AmountCompanyCurrencySigned           *Float      `xmlrpc:"amount_company_currency_signed,omitempty"`
	AmountPaid                            *Float      `xmlrpc:"amount_paid,omitempty"`
	AmountResidual                        *Float      `xmlrpc:"amount_residual,omitempty"`
	AmountResidualSigned                  *Float      `xmlrpc:"amount_residual_signed,omitempty"`
	AmountSigned                          *Float      `xmlrpc:"amount_signed,omitempty"`
	AmountTax                             *Float      `xmlrpc:"amount_tax,omitempty"`
	AmountTaxSigned                       *Float      `xmlrpc:"amount_tax_signed,omitempty"`
	AmountTotal                           *Float      `xmlrpc:"amount_total,omitempty"`
	AmountTotalInCurrencySigned           *Float      `xmlrpc:"amount_total_in_currency_signed,omitempty"`
	AmountTotalSigned                     *Float      `xmlrpc:"amount_total_signed,omitempty"`
	AmountTotalWords                      *String     `xmlrpc:"amount_total_words,omitempty"`
	AmountUntaxed                         *Float      `xmlrpc:"amount_untaxed,omitempty"`
	AmountUntaxedSigned                   *Float      `xmlrpc:"amount_untaxed_signed,omitempty"`
	AttachmentIds                         *Relation   `xmlrpc:"attachment_ids,omitempty"`
	AuthorizedTransactionIds              *Relation   `xmlrpc:"authorized_transaction_ids,omitempty"`
	AutoPost                              *Selection  `xmlrpc:"auto_post,omitempty"`
	AutoPostOriginId                      *Many2One   `xmlrpc:"auto_post_origin_id,omitempty"`
	AutoPostUntil                         *Time       `xmlrpc:"auto_post_until,omitempty"`
	AvailableJournalIds                   *Relation   `xmlrpc:"available_journal_ids,omitempty"`
	AvailablePartnerBankIds               *Relation   `xmlrpc:"available_partner_bank_ids,omitempty"`
	AvailablePaymentMethodLineIds         *Relation   `xmlrpc:"available_payment_method_line_ids,omitempty"`
	BankPartnerId                         *Many2One   `xmlrpc:"bank_partner_id,omitempty"`
	CampaignId                            *Many2One   `xmlrpc:"campaign_id,omitempty"`
	CommercialPartnerId                   *Many2One   `xmlrpc:"commercial_partner_id,omitempty"`
	CompanyCurrencyId                     *Many2One   `xmlrpc:"company_currency_id,omitempty"`
	CompanyId                             *Many2One   `xmlrpc:"company_id,omitempty"`
	CountryCode                           *String     `xmlrpc:"country_code,omitempty"`
	CreateDate                            *Time       `xmlrpc:"create_date,omitempty"`
	CreateUid                             *Many2One   `xmlrpc:"create_uid,omitempty"`
	CurrencyId                            *Many2One   `xmlrpc:"currency_id,omitempty"`
	Date                                  *Time       `xmlrpc:"date,omitempty"`
	DeliveryDate                          *Time       `xmlrpc:"delivery_date,omitempty"`
	DestinationAccountId                  *Many2One   `xmlrpc:"destination_account_id,omitempty"`
	DestinationJournalId                  *Many2One   `xmlrpc:"destination_journal_id,omitempty"`
	DirectionSign                         *Int        `xmlrpc:"direction_sign,omitempty"`
	DisplayInactiveCurrencyWarning        *Bool       `xmlrpc:"display_inactive_currency_warning,omitempty"`
	DisplayName                           *String     `xmlrpc:"display_name,omitempty"`
	DisplayQrCode                         *Bool       `xmlrpc:"display_qr_code,omitempty"`
	DuplicatedRefIds                      *Relation   `xmlrpc:"duplicated_ref_ids,omitempty"`
	FiscalPositionId                      *Many2One   `xmlrpc:"fiscal_position_id,omitempty"`
	HasMessage                            *Bool       `xmlrpc:"has_message,omitempty"`
	HasReconciledEntries                  *Bool       `xmlrpc:"has_reconciled_entries,omitempty"`
	HidePostButton                        *Bool       `xmlrpc:"hide_post_button,omitempty"`
	HighestName                           *String     `xmlrpc:"highest_name,omitempty"`
	Id                                    *Int        `xmlrpc:"id,omitempty"`
	InalterableHash                       *String     `xmlrpc:"inalterable_hash,omitempty"`
	IncotermLocation                      *String     `xmlrpc:"incoterm_location,omitempty"`
	InvoiceCashRoundingId                 *Many2One   `xmlrpc:"invoice_cash_rounding_id,omitempty"`
	InvoiceDate                           *Time       `xmlrpc:"invoice_date,omitempty"`
	InvoiceDateDue                        *Time       `xmlrpc:"invoice_date_due,omitempty"`
	InvoiceFilterTypeDomain               *String     `xmlrpc:"invoice_filter_type_domain,omitempty"`
	InvoiceHasOutstanding                 *Bool       `xmlrpc:"invoice_has_outstanding,omitempty"`
	InvoiceIncotermId                     *Many2One   `xmlrpc:"invoice_incoterm_id,omitempty"`
	InvoiceLineIds                        *Relation   `xmlrpc:"invoice_line_ids,omitempty"`
	InvoiceOrigin                         *String     `xmlrpc:"invoice_origin,omitempty"`
	InvoiceOutstandingCreditsDebitsWidget *String     `xmlrpc:"invoice_outstanding_credits_debits_widget,omitempty"`
	InvoicePartnerDisplayName             *String     `xmlrpc:"invoice_partner_display_name,omitempty"`
	InvoicePaymentTermId                  *Many2One   `xmlrpc:"invoice_payment_term_id,omitempty"`
	InvoicePaymentsWidget                 *String     `xmlrpc:"invoice_payments_widget,omitempty"`
	InvoicePdfReportFile                  *String     `xmlrpc:"invoice_pdf_report_file,omitempty"`
	InvoicePdfReportId                    *Many2One   `xmlrpc:"invoice_pdf_report_id,omitempty"`
	InvoiceSourceEmail                    *String     `xmlrpc:"invoice_source_email,omitempty"`
	InvoiceUserId                         *Many2One   `xmlrpc:"invoice_user_id,omitempty"`
	InvoiceVendorBillId                   *Many2One   `xmlrpc:"invoice_vendor_bill_id,omitempty"`
	IsBeingSent                           *Bool       `xmlrpc:"is_being_sent,omitempty"`
	IsInternalTransfer                    *Bool       `xmlrpc:"is_internal_transfer,omitempty"`
	IsMatched                             *Bool       `xmlrpc:"is_matched,omitempty"`
	IsMoveSent                            *Bool       `xmlrpc:"is_move_sent,omitempty"`
	IsReconciled                          *Bool       `xmlrpc:"is_reconciled,omitempty"`
	IsStorno                              *Bool       `xmlrpc:"is_storno,omitempty"`
	JournalId                             *Many2One   `xmlrpc:"journal_id,omitempty"`
	LineIds                               *Relation   `xmlrpc:"line_ids,omitempty"`
	MachineInvoice                        *Bool       `xmlrpc:"machine_invoice,omitempty"`
	MachineInvoiceTitle                   *String     `xmlrpc:"machine_invoice_title,omitempty"`
	MachinePurchaseOrder                  *String     `xmlrpc:"machine_purchase_order,omitempty"`
	MadeSequenceHole                      *Bool       `xmlrpc:"made_sequence_hole,omitempty"`
	MediumId                              *Many2One   `xmlrpc:"medium_id,omitempty"`
	MessageAttachmentCount                *Int        `xmlrpc:"message_attachment_count,omitempty"`
	MessageFollowerIds                    *Relation   `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError                       *Bool       `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter                *Int        `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError                    *Bool       `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                            *Relation   `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower                     *Bool       `xmlrpc:"message_is_follower,omitempty"`
	MessageMainAttachmentId               *Many2One   `xmlrpc:"message_main_attachment_id,omitempty"`
	MessageNeedaction                     *Bool       `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter              *Int        `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds                     *Relation   `xmlrpc:"message_partner_ids,omitempty"`
	MoveId                                *Many2One   `xmlrpc:"move_id,omitempty"`
	MoveType                              *Selection  `xmlrpc:"move_type,omitempty"`
	MyActivityDateDeadline                *Time       `xmlrpc:"my_activity_date_deadline,omitempty"`
	Name                                  *String     `xmlrpc:"name,omitempty"`
	Narration                             *String     `xmlrpc:"narration,omitempty"`
	NeedCancelRequest                     *Bool       `xmlrpc:"need_cancel_request,omitempty"`
	NeededTerms                           *String     `xmlrpc:"needed_terms,omitempty"`
	NeededTermsDirty                      *Bool       `xmlrpc:"needed_terms_dirty,omitempty"`
	OutstandingAccountId                  *Many2One   `xmlrpc:"outstanding_account_id,omitempty"`
	PairedInternalTransferPaymentId       *Many2One   `xmlrpc:"paired_internal_transfer_payment_id,omitempty"`
	PartnerBankId                         *Many2One   `xmlrpc:"partner_bank_id,omitempty"`
	PartnerCredit                         *Float      `xmlrpc:"partner_credit,omitempty"`
	PartnerCreditWarning                  *String     `xmlrpc:"partner_credit_warning,omitempty"`
	PartnerId                             *Many2One   `xmlrpc:"partner_id,omitempty"`
	PartnerShippingId                     *Many2One   `xmlrpc:"partner_shipping_id,omitempty"`
	PartnerType                           *Selection  `xmlrpc:"partner_type,omitempty"`
	PaymentId                             *Many2One   `xmlrpc:"payment_id,omitempty"`
	PaymentIds                            *Relation   `xmlrpc:"payment_ids,omitempty"`
	PaymentMethodCode                     *String     `xmlrpc:"payment_method_code,omitempty"`
	PaymentMethodId                       *Many2One   `xmlrpc:"payment_method_id,omitempty"`
	PaymentMethodLineId                   *Many2One   `xmlrpc:"payment_method_line_id,omitempty"`
	PaymentReference                      *String     `xmlrpc:"payment_reference,omitempty"`
	PaymentState                          *Selection  `xmlrpc:"payment_state,omitempty"`
	PaymentTermDetails                    *String     `xmlrpc:"payment_term_details,omitempty"`
	PaymentTokenId                        *Many2One   `xmlrpc:"payment_token_id,omitempty"`
	PaymentTransactionId                  *Many2One   `xmlrpc:"payment_transaction_id,omitempty"`
	PaymentType                           *Selection  `xmlrpc:"payment_type,omitempty"`
	PostedBefore                          *Bool       `xmlrpc:"posted_before,omitempty"`
	PurchaseId                            *Many2One   `xmlrpc:"purchase_id,omitempty"`
	PurchaseOrderCount                    *Int        `xmlrpc:"purchase_order_count,omitempty"`
	PurchaseVendorBillId                  *Many2One   `xmlrpc:"purchase_vendor_bill_id,omitempty"`
	QrCode                                *String     `xmlrpc:"qr_code,omitempty"`
	QrCodeMethod                          *Selection  `xmlrpc:"qr_code_method,omitempty"`
	QuantityTotal                         *Float      `xmlrpc:"quantity_total,omitempty"`
	QuickEditMode                         *Bool       `xmlrpc:"quick_edit_mode,omitempty"`
	QuickEditTotalAmount                  *Float      `xmlrpc:"quick_edit_total_amount,omitempty"`
	QuickEncodingVals                     *String     `xmlrpc:"quick_encoding_vals,omitempty"`
	RatingIds                             *Relation   `xmlrpc:"rating_ids,omitempty"`
	ReconciledBillIds                     *Relation   `xmlrpc:"reconciled_bill_ids,omitempty"`
	ReconciledBillsCount                  *Int        `xmlrpc:"reconciled_bills_count,omitempty"`
	ReconciledInvoiceIds                  *Relation   `xmlrpc:"reconciled_invoice_ids,omitempty"`
	ReconciledInvoicesCount               *Int        `xmlrpc:"reconciled_invoices_count,omitempty"`
	ReconciledInvoicesType                *Selection  `xmlrpc:"reconciled_invoices_type,omitempty"`
	ReconciledStatementLineIds            *Relation   `xmlrpc:"reconciled_statement_line_ids,omitempty"`
	ReconciledStatementLinesCount         *Int        `xmlrpc:"reconciled_statement_lines_count,omitempty"`
	Ref                                   *String     `xmlrpc:"ref,omitempty"`
	RefundsCount                          *Int        `xmlrpc:"refunds_count,omitempty"`
	RequirePartnerBankAccount             *Bool       `xmlrpc:"require_partner_bank_account,omitempty"`
	RestrictModeHashTable                 *Bool       `xmlrpc:"restrict_mode_hash_table,omitempty"`
	ReversalMoveId                        *Relation   `xmlrpc:"reversal_move_id,omitempty"`
	ReversedEntryId                       *Many2One   `xmlrpc:"reversed_entry_id,omitempty"`
	SaleOrderCount                        *Int        `xmlrpc:"sale_order_count,omitempty"`
	SecureSequenceNumber                  *Int        `xmlrpc:"secure_sequence_number,omitempty"`
	SendAndPrintValues                    interface{} `xmlrpc:"send_and_print_values,omitempty"`
	SequenceNumber                        *Int        `xmlrpc:"sequence_number,omitempty"`
	SequencePrefix                        *String     `xmlrpc:"sequence_prefix,omitempty"`
	ShowDeliveryDate                      *Bool       `xmlrpc:"show_delivery_date,omitempty"`
	ShowDiscountDetails                   *Bool       `xmlrpc:"show_discount_details,omitempty"`
	ShowNameWarning                       *Bool       `xmlrpc:"show_name_warning,omitempty"`
	ShowPartnerBankAccount                *Bool       `xmlrpc:"show_partner_bank_account,omitempty"`
	ShowPaymentTermDetails                *Bool       `xmlrpc:"show_payment_term_details,omitempty"`
	ShowResetToDraftButton                *Bool       `xmlrpc:"show_reset_to_draft_button,omitempty"`
	SourceId                              *Many2One   `xmlrpc:"source_id,omitempty"`
	SourcePaymentId                       *Many2One   `xmlrpc:"source_payment_id,omitempty"`
	State                                 *Selection  `xmlrpc:"state,omitempty"`
	StatementId                           *Many2One   `xmlrpc:"statement_id,omitempty"`
	StatementLineId                       *Many2One   `xmlrpc:"statement_line_id,omitempty"`
	StatementLineIds                      *Relation   `xmlrpc:"statement_line_ids,omitempty"`
	StockMoveId                           *Many2One   `xmlrpc:"stock_move_id,omitempty"`
	StockValuationLayerIds                *Relation   `xmlrpc:"stock_valuation_layer_ids,omitempty"`
	StringToHash                          *String     `xmlrpc:"string_to_hash,omitempty"`
	SuitableJournalIds                    *Relation   `xmlrpc:"suitable_journal_ids,omitempty"`
	SuitablePaymentTokenIds               *Relation   `xmlrpc:"suitable_payment_token_ids,omitempty"`
	TaxCalculationRoundingMethod          *Selection  `xmlrpc:"tax_calculation_rounding_method,omitempty"`
	TaxCashBasisCreatedMoveIds            *Relation   `xmlrpc:"tax_cash_basis_created_move_ids,omitempty"`
	TaxCashBasisOriginMoveId              *Many2One   `xmlrpc:"tax_cash_basis_origin_move_id,omitempty"`
	TaxCashBasisRecId                     *Many2One   `xmlrpc:"tax_cash_basis_rec_id,omitempty"`
	TaxCountryCode                        *String     `xmlrpc:"tax_country_code,omitempty"`
	TaxCountryId                          *Many2One   `xmlrpc:"tax_country_id,omitempty"`
	TaxLockDateMessage                    *String     `xmlrpc:"tax_lock_date_message,omitempty"`
	TaxTotals                             *String     `xmlrpc:"tax_totals,omitempty"`
	TeamId                                *Many2One   `xmlrpc:"team_id,omitempty"`
	TimesheetCount                        *Int        `xmlrpc:"timesheet_count,omitempty"`
	TimesheetEncodeUomId                  *Many2One   `xmlrpc:"timesheet_encode_uom_id,omitempty"`
	TimesheetIds                          *Relation   `xmlrpc:"timesheet_ids,omitempty"`
	TimesheetTotalDuration                *Int        `xmlrpc:"timesheet_total_duration,omitempty"`
	ToCheck                               *Bool       `xmlrpc:"to_check,omitempty"`
	TransactionIds                        *Relation   `xmlrpc:"transaction_ids,omitempty"`
	TypeName                              *String     `xmlrpc:"type_name,omitempty"`
	UblCiiXmlFile                         *String     `xmlrpc:"ubl_cii_xml_file,omitempty"`
	UblCiiXmlId                           *Many2One   `xmlrpc:"ubl_cii_xml_id,omitempty"`
	UseElectronicPaymentMethod            *Bool       `xmlrpc:"use_electronic_payment_method,omitempty"`
	UserId                                *Many2One   `xmlrpc:"user_id,omitempty"`
	WebsiteMessageIds                     *Relation   `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                             *Time       `xmlrpc:"write_date,omitempty"`
	WriteUid                              *Many2One   `xmlrpc:"write_uid,omitempty"`
}

// AccountPayments represents array of account.payment model.
type AccountPayments []AccountPayment

// AccountPaymentModel is the odoo model name.
const AccountPaymentModel = "account.payment"

// Many2One convert AccountPayment to *Many2One.
func (ap *AccountPayment) Many2One() *Many2One {
	return NewMany2One(ap.Id.Get(), "")
}

// CreateAccountPayment creates a new account.payment model and returns its id.
func (c *Client) CreateAccountPayment(ap *AccountPayment) (int64, error) {
	ids, err := c.CreateAccountPayments([]*AccountPayment{ap})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountPayment creates a new account.payment model and returns its id.
func (c *Client) CreateAccountPayments(aps []*AccountPayment) ([]int64, error) {
	var vv []interface{}
	for _, v := range aps {
		vv = append(vv, v)
	}
	return c.Create(AccountPaymentModel, vv, nil)
}

// UpdateAccountPayment updates an existing account.payment record.
func (c *Client) UpdateAccountPayment(ap *AccountPayment) error {
	return c.UpdateAccountPayments([]int64{ap.Id.Get()}, ap)
}

// UpdateAccountPayments updates existing account.payment records.
// All records (represented by ids) will be updated by ap values.
func (c *Client) UpdateAccountPayments(ids []int64, ap *AccountPayment) error {
	return c.Update(AccountPaymentModel, ids, ap, nil)
}

// DeleteAccountPayment deletes an existing account.payment record.
func (c *Client) DeleteAccountPayment(id int64) error {
	return c.DeleteAccountPayments([]int64{id})
}

// DeleteAccountPayments deletes existing account.payment records.
func (c *Client) DeleteAccountPayments(ids []int64) error {
	return c.Delete(AccountPaymentModel, ids)
}

// GetAccountPayment gets account.payment existing record.
func (c *Client) GetAccountPayment(id int64) (*AccountPayment, error) {
	aps, err := c.GetAccountPayments([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aps)[0]), nil
}

// GetAccountPayments gets account.payment existing records.
func (c *Client) GetAccountPayments(ids []int64) (*AccountPayments, error) {
	aps := &AccountPayments{}
	if err := c.Read(AccountPaymentModel, ids, nil, aps); err != nil {
		return nil, err
	}
	return aps, nil
}

// FindAccountPayment finds account.payment record by querying it with criteria.
func (c *Client) FindAccountPayment(criteria *Criteria) (*AccountPayment, error) {
	aps := &AccountPayments{}
	if err := c.SearchRead(AccountPaymentModel, criteria, NewOptions().Limit(1), aps); err != nil {
		return nil, err
	}
	return &((*aps)[0]), nil
}

// FindAccountPayments finds account.payment records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountPayments(criteria *Criteria, options *Options) (*AccountPayments, error) {
	aps := &AccountPayments{}
	if err := c.SearchRead(AccountPaymentModel, criteria, options, aps); err != nil {
		return nil, err
	}
	return aps, nil
}

// FindAccountPaymentIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountPaymentIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountPaymentModel, criteria, options)
}

// FindAccountPaymentId finds record id by querying it with criteria.
func (c *Client) FindAccountPaymentId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountPaymentModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
