package odoo

import "fmt"

type SlideAnswer struct {
	Id         *Int       `xmlrpc:"id,omptempty"`
	Sequence   *Int       `xmlrpc:"sequence,omptempty"`
	QuestionId *Many2One  `xmlrpc:"question_id,omptempty"`
	CreateUid  *Many2One  `xmlrpc:"create_uid,omptempty"`
	WriteUid   *Many2One  `xmlrpc:"write_uid,omptempty"`
	TextValue  *Selection `xmlrpc:"text_value,omptempty"`
	Comment    *Selection `xmlrpc:"comment,omptempty"`
	IsCorrect  *Bool      `xmlrpc:"is_correct,omptempty"`
	CreateDate *Time      `xmlrpc:"create_date,omptempty"`
	WriteDate  *Time      `xmlrpc:"write_date,omptempty"`
}

type SlideAnswers []SlideAnswer

const SlideAnswerModel = "slide.answer"

func (sa *SlideAnswer) Many2One() *Many2One {
	return NewMany2One(sa.Id.Get(), "")
}

func (c *Client) CreateSlideAnswer(sa *SlideAnswer) (int64, error) {
	ids, err := c.CreateSlideAnswers([]*SlideAnswer{sa})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

func (c *Client) CreateSlideAnswers(sa []*SlideAnswer) ([]int64, error) {
	var vv []interface{}
	for _, v := range sa {
		vv = append(vv, v)
	}
	return c.Create(SlideAnswerModel, vv)
}

func (c *Client) UpdateSlideAnswer(sa *SlideAnswer) error {
	return c.UpdateSlideAnswers([]int64{sa.Id.Get()}, sa)
}

func (c *Client) UpdateSlideAnswers(ids []int64, sa *SlideAnswer) error {
	return c.Update(SlideAnswerModel, ids, sa)
}

func (c *Client) DeleteSlideAnswer(id int64) error {
	return c.DeleteSlideAnswers([]int64{id})
}

func (c *Client) DeleteSlideAnswers(ids []int64) error {
	return c.Delete(SlideAnswerModel, ids)
}

func (c *Client) GetSlideAnswer(id int64) (*SlideAnswer, error) {
	sas, err := c.GetSlideAnswers([]int64{id})
	if err != nil {
		return nil, err
	}
	if sas != nil && len(*sas) > 0 {
		return &((*sas)[0]), nil
	}
	return nil, fmt.Errorf("id %v of slide.answer not found", id)
}

func (c *Client) GetSlideAnswers(ids []int64) (*SlideAnswers, error) {
	sas := &SlideAnswers{}
	if err := c.Read(SlideAnswerModel, ids, nil, sas); err != nil {
		return nil, err
	}
	return sas, nil
}

func (c *Client) FindSlideAnswer(criteria *Criteria) (*SlideAnswer, error) {
	sas := &SlideAnswers{}
	if err := c.SearchRead(SlideAnswerModel, criteria, NewOptions().Limit(1), sas); err != nil {
		return nil, err
	}
	if sas != nil && len(*sas) > 0 {
		return &((*sas)[0]), nil
	}
	return nil, fmt.Errorf("slide.answer was not found with criteria %v", criteria)
}

func (c *Client) FindSlideAnswers(criteria *Criteria, options *Options) (*SlideAnswers, error) {
	sas := &SlideAnswers{}
	if err := c.SearchRead(SlideAnswerModel, criteria, options, sas); err != nil {
		return nil, err
	}
	return sas, nil
}

func (c *Client) FindSlideAnswerIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SlideAnswerModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

func (c *Client) FindSlideAnswerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SlideAnswerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("slide.answer was not found with criteria %v and options %v", criteria, options)
}
