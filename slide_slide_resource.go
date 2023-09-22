package odoo

import "fmt"

type SlideSlideResource struct {
	Id           *Int      `xmlrpc:"id,omptempty"`
	SlideId      *Many2One `xmlrpc:"slide_id,omptempty"`
	CreateUid    *Many2One `xmlrpc:"create_uid,omptempty"`
	ResourceType *String   `xmlrpc:"resource_type,omptempty"`
	Name         *String   `xmlrpc:"name,omptempty"`
	FileName     *String   `xmlrpc:"file_name,omptempty"`
	Link         *String   `xmlrpc:"link,omptempty"`
	WriteUid     *Many2One `xmlrpc:"write_uid,omptempty"`
	CreateDate   *Time     `xmlrpc:"create_date,omptempty"`
	WriteDate    *Time     `xmlrpc:"write_date,omptempty"`
}

type SlideSlideResources []SlideSlideResource

const SlideSlideResourceModel = "slide.slide.resource"

func (ssr *SlideSlideResource) Many2One() *Many2One {
	return NewMany2One(ssr.Id.Get(), "")
}

func (c *Client) CreateSlideSlideResource(ssr *SlideSlideResource) (int64, error) {
	ids, err := c.CreateSlideSlideResources([]*SlideSlideResource{ssr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

func (c *Client) CreateSlideSlideResources(ssr []*SlideSlideResource) ([]int64, error) {
	var vv []interface{}
	for _, v := range ssr {
		vv = append(vv, v)
	}
	return c.Create(SlideSlideResourceModel, vv)
}

func (c *Client) UpdateSlideSlideResource(ssr *SlideSlideResource) error {
	return c.UpdateSlideSlideResources([]int64{ssr.Id.Get()}, ssr)
}

func (c *Client) UpdateSlideSlideResources(ids []int64, ssr *SlideSlideResource) error {
	return c.Update(SlideSlideResourceModel, ids, ssr)
}

func (c *Client) DeleteSlideSlideResource(id int64) error {
	return c.DeleteSlideSlideResources([]int64{id})
}

func (c *Client) DeleteSlideSlideResources(ids []int64) error {
	return c.Delete(SlideSlideResourceModel, ids)
}

func (c *Client) GetSlideSlideResource(id int64) (*SlideSlideResource, error) {
	ssr, err := c.GetSlideSlideResources([]int64{id})
	if err != nil {
		return nil, err
	}
	if ssr != nil && len(*ssr) > 0 {
		return &((*ssr)[0]), nil
	}
	return nil, fmt.Errorf("id %v of slide.slide.resource not found", id)
}

func (c *Client) GetSlideSlideResources(ids []int64) (*SlideSlideResources, error) {
	ssrs := &SlideSlideResources{}
	if err := c.Read(SlideSlideResourceModel, ids, nil, ssrs); err != nil {
		return nil, err
	}
	return ssrs, nil
}

func (c *Client) FindSlideSlideResource(criteria *Criteria) (*SlideSlideResource, error) {
	ssrs := &SlideSlideResources{}
	if err := c.SearchRead(SlideSlideResourceModel, criteria, NewOptions().Limit(1), ssrs); err != nil {
		return nil, err
	}
	if ssrs != nil && len(*ssrs) > 0 {
		return &((*ssrs)[0]), nil
	}
	return nil, fmt.Errorf("slide.slide.resource was not found with criteria %v", criteria)
}

func (c *Client) FindSlideSlideResources(criteria *Criteria, options *Options) (*SlideSlideResources, error) {
	ssrs := &SlideSlideResources{}
	if err := c.SearchRead(SlideSlideResourceModel, criteria, options, ssrs); err != nil {
		return nil, err
	}
	return ssrs, nil
}

func (c *Client) FindSlideSlideResourceIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SlideSlideResourceModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

func (c *Client) FindSlideSlideResourceId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SlideSlideResourceModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("slide.slide.resource was not found with criteria %v and options %v", criteria, options)
}
