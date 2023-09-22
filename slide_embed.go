package odoo

import "fmt"

type SlideEmbed struct {
	Id         *Int      `xmlrpc:"id,omptempty"`
	SlideId    *Many2One `xmlrpc:"slide_id,omptempty"`
	CreateUid  *Many2One `xmlrpc:"create_uid,omptempty"`
	WriteUid   *Many2One `xmlrpc:"write_uid,omptempty"`
	Url        *String   `xmlrpc:"url,omptempty"`
	CreateDate *Time     `xmlrpc:"create_date,omptempty"`
	WriteDate  *Time     `xmlrpc:"write_date,omptempty"`
}

type SlideEmbeds []SlideEmbed

const SlideEmbedModel = "slide.embed"

func (se *SlideEmbed) Many2One() *Many2One {
	return NewMany2One(se.Id.Get(), "")
}

func (c *Client) CreateSlideEmbed(se *SlideEmbed) (int64, error) {
	ids, err := c.CreateSlideEmbeds([]*SlideEmbed{se})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

func (c *Client) CreateSlideEmbeds(se []*SlideEmbed) ([]int64, error) {
	var vv []interface{}
	for _, v := range se {
		vv = append(vv, v)
	}
	return c.Create(SlideEmbedModel, vv)
}

func (c *Client) UpdateSlideEmbed(se *SlideEmbed) error {
	return c.UpdateSlideEmbeds([]int64{se.Id.Get()}, se)
}

func (c *Client) UpdateSlideEmbeds(ids []int64, se *SlideEmbed) error {
	return c.Update(SlideEmbedModel, ids, se)
}

func (c *Client) DeleteSlideEmbed(id int64) error {
	return c.DeleteSlideEmbeds([]int64{id})
}

func (c *Client) DeleteSlideEmbeds(ids []int64) error {
	return c.Delete(SlideEmbedModel, ids)
}

func (c *Client) GetSlideEmbed(id int64) (*SlideEmbed, error) {
	ses, err := c.GetSlideEmbeds([]int64{id})
	if err != nil {
		return nil, err
	}
	if ses != nil && len(*ses) > 0 {
		return &((*ses)[0]), nil
	}
	return nil, fmt.Errorf("id %v of slide.embed not found", id)
}

func (c *Client) GetSlideEmbeds(ids []int64) (*SlideEmbeds, error) {
	ses := &SlideEmbeds{}
	if err := c.Read(SlideEmbedModel, ids, nil, ses); err != nil {
		return nil, err
	}
	return ses, nil
}

func (c *Client) FindSlideEmbed(criteria *Criteria) (*SlideEmbed, error) {
	ses := &SlideEmbeds{}
	if err := c.SearchRead(SlideEmbedModel, criteria, NewOptions().Limit(1), ses); err != nil {
		return nil, err
	}
	if ses != nil && len(*ses) > 0 {
		return &((*ses)[0]), nil
	}
	return nil, fmt.Errorf("slide.embed was not found with criteria %v", criteria)
}

func (c *Client) FindSlideEmbeds(criteria *Criteria, options *Options) (*SlideEmbeds, error) {
	ses := &SlideEmbeds{}
	if err := c.SearchRead(SlideEmbedModel, criteria, options, ses); err != nil {
		return nil, err
	}
	return ses, nil
}

func (c *Client) FindSlideEmbedIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SlideEmbedModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

func (c *Client) FindSlideEmbedId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SlideEmbedModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("slide.embed was not found with criteria %v and options %v", criteria, options)
}
