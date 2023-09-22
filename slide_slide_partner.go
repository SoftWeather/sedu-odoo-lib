package odoo

import "fmt"

type SlideSlidePartner struct {
	Id                *Int      `xmlrpc:"id,omptempty"`
	SlideId           *Many2One `xmlrpc:"slide_id,omptempty"`
	ChannelId         *Many2One `xmlrpc:"channel_id,omptempty"`
	PartnerId         *Many2One `xmlrpc:"partner_id,omptempty"`
	Vote              *Int      `xmlrpc:"vote,omptempty"`
	QuizAttemptsCount *Int      `xmlrpc:"quiz_attempts_count,omptempty"`
	CreateUid         *Many2One `xmlrpc:"create_uid,omptempty"`
	WriteUid          *Many2One `xmlrpc:"write_uid,omptempty"`
	Completed         *Bool     `xmlrpc:"completed,omptempty"`
	CreateDate        *Time     `xmlrpc:"create_date,omptempty"`
	WriteDate         *Time     `xmlrpc:"write_date,omptempty"`
}

type SlideSlidePartners []SlideSlidePartner

const SlideSlidePartnerModel = "slide.slide.partner"

func (ssp *SlideSlidePartner) Many2One() *Many2One {
	return NewMany2One(ssp.Id.Get(), "")
}

func (c *Client) CreateSlideSlidePartner(ssp *SlideSlidePartner) (int64, error) {
	ids, err := c.CreateSlideSlidePartners([]*SlideSlidePartner{ssp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

func (c *Client) CreateSlideSlidePartners(ssp []*SlideSlidePartner) ([]int64, error) {
	var vv []interface{}
	for _, v := range ssp {
		vv = append(vv, v)
	}
	return c.Create(SlideSlidePartnerModel, vv)
}

func (c *Client) UpdateSlideSlidePartner(ssp *SlideSlidePartner) error {
	return c.UpdateSlideSlidePartners([]int64{ssp.Id.Get()}, ssp)
}

func (c *Client) UpdateSlideSlidePartners(ids []int64, ssp *SlideSlidePartner) error {
	return c.Update(SlideSlidePartnerModel, ids, ssp)
}

func (c *Client) DeleteSlideSlidePartner(id int64) error {
	return c.DeleteSlideSlidePartners([]int64{id})
}

func (c *Client) DeleteSlideSlidePartners(ids []int64) error {
	return c.Delete(SlideSlidePartnerModel, ids)
}

func (c *Client) GetSlideSlidePartner(id int64) (*SlideSlidePartner, error) {
	ssp, err := c.GetSlideSlidePartners([]int64{id})
	if err != nil {
		return nil, err
	}
	if ssp != nil && len(*ssp) > 0 {
		return &((*ssp)[0]), nil
	}
	return nil, fmt.Errorf("id %v of slide.slide.partner not found", id)
}

func (c *Client) GetSlideSlidePartners(ids []int64) (*SlideSlidePartners, error) {
	ssps := &SlideSlidePartners{}
	if err := c.Read(SlideSlidePartnerModel, ids, nil, ssps); err != nil {
		return nil, err
	}
	return ssps, nil
}

func (c *Client) FindSlideSlidePartner(criteria *Criteria) (*SlideSlidePartner, error) {
	ssps := &SlideSlidePartners{}
	if err := c.SearchRead(SlideSlidePartnerModel, criteria, NewOptions().Limit(1), ssps); err != nil {
		return nil, err
	}
	if ssps != nil && len(*ssps) > 0 {
		return &((*ssps)[0]), nil
	}
	return nil, fmt.Errorf("slide.slide.partner was not found with criteria %v", criteria)
}

func (c *Client) FindSlideSlidePartners(criteria *Criteria, options *Options) (*SlideSlidePartners, error) {
	ssps := &SlideSlidePartners{}
	if err := c.SearchRead(SlideSlidePartnerModel, criteria, options, ssps); err != nil {
		return nil, err
	}
	return ssps, nil
}

func (c *Client) FindSlideSlidePartnerIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SlideSlidePartnerModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

func (c *Client) FindSlideSlidePartnerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SlideSlidePartnerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("slide.slide.partner was not found with criteria %v and options %v", criteria, options)
}
