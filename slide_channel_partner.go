package odoo

import "fmt"

type SlideChannelPartner struct {
	Id                   *Int      `xmlrpc:"id,omptempty"`
	ChannelId            *Many2One `xmlrpc:"channel_id,omptempty"`
	Completion           *Int      `xmlrpc:"completion,omptempty"`
	CompletedSlidesCount *Int      `xmlrpc:"completed_slides_count,omptempty"`
	PartnerId            *Many2One `xmlrpc:"partner_id,omptempty"`
	CreateUid            *Many2One `xmlrpc:"create_uid,omptempty"`
	WriteUid             *Many2One `xmlrpc:"write_uid,omptempty"`
	Completed            *Bool     `xmlrpc:"completed,omptempty"`
	CreateDate           *Time     `xmlrpc:"create_date,omptempty"`
	WriteDate            *Time     `xmlrpc:"write_date,omptempty"`
}

type SlideChannelPartners []SlideChannelPartner

const SlideChannelPartnerModel = "slide.channel.partner"

func (scp *SlideChannelPartner) Many2One() *Many2One {
	return NewMany2One(scp.Id.Get(), "")
}

func (c *Client) CreateSlideChannelPartner(scp *SlideChannelPartner) (int64, error) {
	ids, err := c.CreateSlideChannelPartners([]*SlideChannelPartner{scp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

func (c *Client) CreateSlideChannelPartners(scp []*SlideChannelPartner) ([]int64, error) {
	var vv []interface{}
	for _, v := range scp {
		vv = append(vv, v)
	}
	return c.Create(SlideChannelPartnerModel, vv)
}

func (c *Client) UpdateSlideChannelPartner(scp *SlideChannelPartner) error {
	return c.UpdateSlideChannelPartners([]int64{scp.Id.Get()}, scp)
}

func (c *Client) UpdateSlideChannelPartners(ids []int64, scp *SlideChannelPartner) error {
	return c.Update(SlideChannelPartnerModel, ids, scp)
}

func (c *Client) DeleteSlideChannelPartner(id int64) error {
	return c.DeleteSlideChannelPartners([]int64{id})
}

func (c *Client) DeleteSlideChannelPartners(ids []int64) error {
	return c.Delete(SlideChannelPartnerModel, ids)
}

func (c *Client) GetSlideChannelPartner(id int64) (*SlideChannelPartner, error) {
	scps, err := c.GetSlideChannelPartners([]int64{id})
	if err != nil {
		return nil, err
	}
	if scps != nil && len(*scps) > 0 {
		return &((*scps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of slide.channel.partner not found", id)
}

func (c *Client) GetSlideChannelPartners(ids []int64) (*SlideChannelPartners, error) {
	scps := &SlideChannelPartners{}
	if err := c.Read(SlideChannelPartnerModel, ids, nil, scps); err != nil {
		return nil, err
	}
	return scps, nil
}

func (c *Client) FindSlideChannelPartner(criteria *Criteria) (*SlideChannelPartner, error) {
	scps := &SlideChannelPartners{}
	if err := c.SearchRead(SlideChannelPartnerModel, criteria, NewOptions().Limit(1), scps); err != nil {
		return nil, err
	}
	if scps != nil && len(*scps) > 0 {
		return &((*scps)[0]), nil
	}
	return nil, fmt.Errorf("slide.channel.partner was not found with criteria %v", criteria)
}

func (c *Client) FindSlideChannelPartners(criteria *Criteria, options *Options) (*SlideChannelPartners, error) {
	scps := &SlideChannelPartners{}
	if err := c.SearchRead(SlideChannelPartnerModel, criteria, options, scps); err != nil {
		return nil, err
	}
	return scps, nil
}

func (c *Client) FindSlideChannelPartnerIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SlideChannelPartnerModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

func (c *Client) FindSlideChannelPartnerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SlideChannelPartnerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("slide.channel.partner was not found with criteria %v and options %v", criteria, options)
}
