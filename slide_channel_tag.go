package odoo

import "fmt"

type SlideChannelTag struct {
	Id            *Int       `xmlrpc:"id,omptempty"`
	Sequence      *Int       `xmlrpc:"sequence,omptempty"`
	GroupId       *Many2One  `xmlrpc:"group_id,omptempty"`
	GroupSequence *Many2One  `xmlrpc:"group_sequence,omptempty"`
	Color         *Int       `xmlrpc:"color,omptempty"`
	CreateUid     *Many2One  `xmlrpc:"create_uid,omptempty"`
	WriteUid      *Many2One  `xmlrpc:"write_uid,omptempty"`
	Name          *Selection `xmlrpc:"name,omptempty"`
	CreateDate    *Time      `xmlrpc:"create_date,omptempty"`
	WriteDate     *Time      `xmlrpc:"write_date,omptempty"`
}

type SlideChannelTags []SlideChannelTag

const SlideChannelTagModel = "slide.channel.tag"

func (sct *SlideChannelTag) Many2One() *Many2One {
	return NewMany2One(sct.Id.Get(), "")
}

func (c *Client) CreateSlideChannelTag(sct *SlideChannelTag) (int64, error) {
	ids, err := c.CreateSlideChannelTags([]*SlideChannelTag{sct})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

func (c *Client) CreateSlideChannelTags(sct []*SlideChannelTag) ([]int64, error) {
	var vv []interface{}
	for _, v := range sct {
		vv = append(vv, v)
	}
	return c.Create(SlideChannelTagModel, vv)
}

func (c *Client) UpdateSlideChannelTag(sct *SlideChannelTag) error {
	return c.UpdateSlideChannelTags([]int64{sct.Id.Get()}, sct)
}

func (c *Client) UpdateSlideChannelTags(ids []int64, sct *SlideChannelTag) error {
	return c.Update(SlideChannelTagModel, ids, sct)
}

func (c *Client) DeleteSlideChannelTag(id int64) error {
	return c.DeleteSlideChannelTags([]int64{id})
}

func (c *Client) DeleteSlideChannelTags(ids []int64) error {
	return c.Delete(SlideChannelTagModel, ids)
}

func (c *Client) GetSlideChannelTag(id int64) (*SlideChannelTag, error) {
	scts, err := c.GetSlideChannelTags([]int64{id})
	if err != nil {
		return nil, err
	}
	if scts != nil && len(*scts) > 0 {
		return &((*scts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of slide.channel.tag not found", id)
}

func (c *Client) GetSlideChannelTags(ids []int64) (*SlideChannelTags, error) {
	scts := &SlideChannelTags{}
	if err := c.Read(SlideChannelTagModel, ids, nil, scts); err != nil {
		return nil, err
	}
	return scts, nil
}

func (c *Client) FindSlideChannelTag(criteria *Criteria) (*SlideChannelTag, error) {
	scts := &SlideChannelTags{}
	if err := c.SearchRead(SlideChannelTagModel, criteria, NewOptions().Limit(1), scts); err != nil {
		return nil, err
	}
	if scts != nil && len(*scts) > 0 {
		return &((*scts)[0]), nil
	}
	return nil, fmt.Errorf("slide.channel.tag was not found with criteria %v", criteria)
}

func (c *Client) FindSlideChannelTags(criteria *Criteria, options *Options) (*SlideChannelTags, error) {
	scts := &SlideChannelTags{}
	if err := c.SearchRead(SlideChannelTagModel, criteria, options, scts); err != nil {
		return nil, err
	}
	return scts, nil
}

func (c *Client) FindSlideChannelTagIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SlideChannelTagModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

func (c *Client) FindSlideChannelTagId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SlideChannelTagModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("slide.channel.tag was not found with criteria %v and options %v", criteria, options)
}
