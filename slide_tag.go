package odoo

import "fmt"

type SlideTag struct {
	Id         *Int       `xmlrpc:"id,omptempty"`
	CreateUid  *Many2One  `xmlrpc:"create_uid,omptempty"`
	WriteUid   *Many2One  `xmlrpc:"write_uid,omptempty"`
	Name       *Selection `xmlrpc:"name,omptempty"`
	CreateDate *Time      `xmlrpc:"create_date,omptempty"`
	WriteDate  *Time      `xmlrpc:"write_date,omptempty"`
}

type SlideTags []SlideTag

const SlideTagModel = "slide.tag"

func (st *SlideTag) Many2One() *Many2One {
	return NewMany2One(st.Id.Get(), "")
}

func (c *Client) CreateSlideTag(st *SlideTag) (int64, error) {
	ids, err := c.CreateSlideTags([]*SlideTag{st})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

func (c *Client) CreateSlideTags(st []*SlideTag) ([]int64, error) {
	var vv []interface{}
	for _, v := range st {
		vv = append(vv, v)
	}
	return c.Create(SlideTagModel, vv)
}

func (c *Client) UpdateSlideTag(st *SlideTag) error {
	return c.UpdateSlideTags([]int64{st.Id.Get()}, st)
}

func (c *Client) UpdateSlideTags(ids []int64, st *SlideTag) error {
	return c.Update(SlideTagModel, ids, st)
}

func (c *Client) DeleteSlideTag(id int64) error {
	return c.DeleteSlideTags([]int64{id})
}

func (c *Client) DeleteSlideTags(ids []int64) error {
	return c.Delete(SlideTagModel, ids)
}

func (c *Client) GetSlideTag(id int64) (*SlideTag, error) {
	st, err := c.GetSlideTags([]int64{id})
	if err != nil {
		return nil, err
	}
	if st != nil && len(*st) > 0 {
		return &((*st)[0]), nil
	}
	return nil, fmt.Errorf("id %v of slide.tag not found", id)
}

func (c *Client) GetSlideTags(ids []int64) (*SlideTags, error) {
	sts := &SlideTags{}
	if err := c.Read(SlideTagModel, ids, nil, sts); err != nil {
		return nil, err
	}
	return sts, nil
}

func (c *Client) FindSlideTag(criteria *Criteria) (*SlideTag, error) {
	sts := &SlideTags{}
	if err := c.SearchRead(SlideTagModel, criteria, NewOptions().Limit(1), sts); err != nil {
		return nil, err
	}
	if sts != nil && len(*sts) > 0 {
		return &((*sts)[0]), nil
	}
	return nil, fmt.Errorf("slide.tag was not found with criteria %v", criteria)
}

func (c *Client) FindSlideTags(criteria *Criteria, options *Options) (*SlideTags, error) {
	sts := &SlideTags{}
	if err := c.SearchRead(SlideTagModel, criteria, options, sts); err != nil {
		return nil, err
	}
	return sts, nil
}

func (c *Client) FindSlideTagIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SlideTagModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

func (c *Client) FindSlideTagId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SlideTagModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("slide.tag was not found with criteria %v and options %v", criteria, options)
}
