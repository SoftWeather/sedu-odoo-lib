package odoo

import "fmt"

type SlideQuestion struct {
	Id         *Int       `xmlrpc:"id,omptempty"`
	Sequence   *Int       `xmlrpc:"sequence,omptempty"`
	SlideId    *Many2One  `xmlrpc:"slide_id,omptempty"`
	CreateUid  *Many2One  `xmlrpc:"create_uid,omptempty"`
	WriteUid   *Many2One  `xmlrpc:"write_uid,omptempty"`
	Question   *Selection `xmlrpc:"question,omptempty"`
	CreateDate *Time      `xmlrpc:"create_date,omptempty"`
	WriteDate  *Time      `xmlrpc:"write_date,omptempty"`
}

type SlideQuestions []SlideQuestion

const SlideQuestionModel = "slide.question"

func (sq *SlideQuestion) Many2One() *Many2One {
	return NewMany2One(sq.Id.Get(), "")
}

func (c *Client) CreateSlideQuestion(sq *SlideQuestion) (int64, error) {
	ids, err := c.CreateSlideQuestions([]*SlideQuestion{sq})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

func (c *Client) CreateSlideQuestions(sq []*SlideQuestion) ([]int64, error) {
	var vv []interface{}
	for _, v := range sq {
		vv = append(vv, v)
	}
	return c.Create(SlideQuestionModel, vv)
}

func (c *Client) UpdateSlideQuestion(sq *SlideQuestion) error {
	return c.UpdateSlideQuestions([]int64{sq.Id.Get()}, sq)
}

func (c *Client) UpdateSlideQuestions(ids []int64, sq *SlideQuestion) error {
	return c.Update(SlideQuestionModel, ids, sq)
}

func (c *Client) DeleteSlideQuestion(id int64) error {
	return c.DeleteSlideQuestions([]int64{id})
}

func (c *Client) DeleteSlideQuestions(ids []int64) error {
	return c.Delete(SlideQuestionModel, ids)
}

func (c *Client) GetSlideQuestion(id int64) (*SlideQuestion, error) {
	sqs, err := c.GetSlideQuestions([]int64{id})
	if err != nil {
		return nil, err
	}
	if sqs != nil && len(*sqs) > 0 {
		return &((*sqs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of slide.question not found", id)
}

func (c *Client) GetSlideQuestions(ids []int64) (*SlideQuestions, error) {
	sqs := &SlideQuestions{}
	if err := c.Read(SlideQuestionModel, ids, nil, sqs); err != nil {
		return nil, err
	}
	return sqs, nil
}

func (c *Client) FindSlideQuestion(criteria *Criteria) (*SlideQuestion, error) {
	sqs := &SlideQuestions{}
	if err := c.SearchRead(SlideQuestionModel, criteria, NewOptions().Limit(1), sqs); err != nil {
		return nil, err
	}
	if sqs != nil && len(*sqs) > 0 {
		return &((*sqs)[0]), nil
	}
	return nil, fmt.Errorf("slide.question was not found with criteria %v", criteria)
}

func (c *Client) FindSlideQuestions(criteria *Criteria, options *Options) (*SlideQuestions, error) {
	sqs := &SlideQuestions{}
	if err := c.SearchRead(SlideQuestionModel, criteria, options, sqs); err != nil {
		return nil, err
	}
	return sqs, nil
}

func (c *Client) FindSlideQuestionIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SlideQuestionModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

func (c *Client) FindSlideQuestionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SlideQuestionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("slide.question was not found with criteria %v and options %v", criteria, options)
}
