package odoo

import "fmt"

type SlideChannelInvite struct {
	Id         *Int      `xmlrpc:"id,omptempty"`
	TemplateId *Many2One `xmlrpc:"template_id,omptempty"`
	ChannelId  *Many2One `xmlrpc:"channel_id,omptempty"`
	CreateUid  *Many2One `xmlrpc:"create_uid,omptempty"`
	WriteUid   *Many2One `xmlrpc:"write_uid,omptempty"`
	Lang       *String   `xmlrpc:"lang,omptempty"`
	Subject    *String   `xmlrpc:"subject,omptempty"`
	Body       *String   `xmlrpc:"body,omptempty"`
	CreateDate *Time     `xmlrpc:"create_date,omptempty"`
	WriteDate  *Time     `xmlrpc:"create_date,omptempty"`
}

type SlideChannelInvites []SlideChannelInvite

const SlideChannelInviteModel = "slide.channel.invite"

func (sci *SlideChannelInvite) Many2One() *Many2One {
	return NewMany2One(sci.Id.Get(), "")
}

func (c *Client) CreateSlideChannelInvite(sci *SlideChannelInvite) (int64, error) {
	ids, err := c.CreateSlideChannelInvites([]*SlideChannelInvite{sci})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

func (c *Client) CreateSlideChannelInvites(sci []*SlideChannelInvite) ([]int64, error) {
	var vv []interface{}
	for _, v := range sci {
		vv = append(vv, v)
	}
	return c.Create(SlideChannelInviteModel, vv)
}

func (c *Client) UpdateSlideChannelInvite(sci *SlideChannelInvite) error {
	return c.UpdateSlideChannelInvites([]int64{sci.Id.Get()}, sci)
}

func (c *Client) UpdateSlideChannelInvites(ids []int64, sci *SlideChannelInvite) error {
	return c.Update(SlideChannelInviteModel, ids, sci)
}

func (c *Client) DeleteSlideChannelInvite(id int64) error {
	return c.DeleteSlideChannelInvites([]int64{id})
}

func (c *Client) DeleteSlideChannelInvites(ids []int64) error {
	return c.Delete(SlideChannelInviteModel, ids)
}

func (c *Client) GetSlideChannelInvite(id int64) (*SlideChannelInvite, error) {
	scis, err := c.GetSlideChannelInvites([]int64{id})
	if err != nil {
		return nil, err
	}
	if scis != nil && len(*scis) > 0 {
		return &((*scis)[0]), nil
	}
	return nil, fmt.Errorf("id %v of slide.channel.invite not found", id)
}

func (c *Client) GetSlideChannelInvites(ids []int64) (*SlideChannelInvites, error) {
	scis := &SlideChannelInvites{}
	if err := c.Read(SlideChannelInviteModel, ids, nil, scis); err != nil {
		return nil, err
	}
	return scis, nil
}

func (c *Client) FindSlideChannelInvite(criteria *Criteria) (*SlideChannelInvite, error) {
	scis := &SlideChannelInvites{}
	if err := c.SearchRead(SlideChannelInviteModel, criteria, NewOptions().Limit(1), scis); err != nil {
		return nil, err
	}
	if scis != nil && len(*scis) > 0 {
		return &((*scis)[0]), nil
	}
	return nil, fmt.Errorf("slide.channel.invite was not found with criteria %v", criteria)
}

func (c *Client) FindSlideChannelInvites(criteria *Criteria, options *Options) (*SlideChannelInvites, error) {
	scis := &SlideChannelInvites{}
	if err := c.SearchRead(SlideChannelInviteModel, criteria, options, scis); err != nil {
		return nil, err
	}
	return scis, nil
}

func (c *Client) FindSlideChannelInviteIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SlideChannelInviteModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

func (c *Client) FindSlideChannelInviteId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SlideChannelInviteModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("slide.channel.invite was not found with criteria %v and options %v", criteria, options)
}
