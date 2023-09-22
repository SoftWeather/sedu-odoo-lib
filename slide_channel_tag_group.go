package odoo

import "fmt"

type SlideChannelTagGroup struct {
	Id          *Int       `xmlrpc:"id,omptempty"`
	Sequence    *Int       `xmlrpc:"sequence,omptempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omptempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omptempty"`
	Name        *Selection `xmlrpc:"name,omptempty"`
	IsPublished *Bool      `xmlrpc:"is_published,omptempty"`
	CreateDate  *Time      `xmlrpc:"create_date,omptempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omptempty"`
}

type SlideChannelTagGroups []SlideChannelTagGroup

const SlideChannelTagGroupModel = "slide.channel.tag.group"

func (sctg *SlideChannelTagGroup) Many2One() *Many2One {
	return NewMany2One(sctg.Id.Get(), "")
}

func (c *Client) CreateSlideChannelTagGroup(sctg *SlideChannelTagGroup) (int64, error) {
	ids, err := c.CreateSlideChannelTagGroups([]*SlideChannelTagGroup{sctg})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

func (c *Client) CreateSlideChannelTagGroups(sctg []*SlideChannelTagGroup) ([]int64, error) {
	var vv []interface{}
	for _, v := range sctg {
		vv = append(vv, v)
	}
	return c.Create(SlideChannelTagGroupModel, vv)
}

func (c *Client) UpdateSlideChannelTagGroup(sctg *SlideChannelTagGroup) error {
	return c.UpdateSlideChannelTagGroups([]int64{sctg.Id.Get()}, sctg)
}

func (c *Client) UpdateSlideChannelTagGroups(ids []int64, sctg *SlideChannelTagGroup) error {
	return c.Update(SlideChannelTagGroupModel, ids, sctg)
}

func (c *Client) DeleteSlideChannelTagGroup(id int64) error {
	return c.DeleteSlideChannelTagGroups([]int64{id})
}

func (c *Client) DeleteSlideChannelTagGroups(ids []int64) error {
	return c.Delete(SlideChannelTagGroupModel, ids)
}

func (c *Client) GetSlideChannelTagGroup(id int64) (*SlideChannelTagGroup, error) {
	sctgs, err := c.GetSlideChannelTagGroups([]int64{id})
	if err != nil {
		return nil, err
	}
	if sctgs != nil && len(*sctgs) > 0 {
		return &((*sctgs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of slide.channel.tag.group not found", id)
}

func (c *Client) GetSlideChannelTagGroups(ids []int64) (*SlideChannelTagGroups, error) {
	sctgs := &SlideChannelTagGroups{}
	if err := c.Read(SlideChannelTagGroupModel, ids, nil, sctgs); err != nil {
		return nil, err
	}
	return sctgs, nil
}

func (c *Client) FindSlideChannelTagGroup(criteria *Criteria) (*SlideChannelTagGroup, error) {
	sctgs := &SlideChannelTagGroups{}
	if err := c.SearchRead(SlideChannelTagGroupModel, criteria, NewOptions().Limit(1), sctgs); err != nil {
		return nil, err
	}
	if sctgs != nil && len(*sctgs) > 0 {
		return &((*sctgs)[0]), nil
	}
	return nil, fmt.Errorf("slide.channel.tag.group was not found with criteria %v", criteria)
}

func (c *Client) FindSlideChannelTagGroups(criteria *Criteria, options *Options) (*SlideChannelTagGroups, error) {
	sctgs := &SlideChannelTagGroups{}
	if err := c.SearchRead(SlideChannelTagGroupModel, criteria, options, sctgs); err != nil {
		return nil, err
	}
	return sctgs, nil
}

func (c *Client) FindSlideChannelTagGroupIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SlideChannelTagGroupModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

func (c *Client) FindSlideChannelTagGroupId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SlideChannelTagGroupModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("slide.channel.tag.group was not found with criteria %v and options %v", criteria, options)
}
