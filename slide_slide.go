package odoo

import "fmt"

type SlideSlide struct {
	Id                        *Int       `xmlrpc:"id,omptempty"`
	MessageMainAttachmentId   *Many2One  `xmlrpc:"message_main_attachment_id,omptempty"`
	Sequence                  *Int       `xmlrpc:"sequence,omptempty"`
	UserId                    *Many2One  `xmlrpc:"user_id,omptempty"`
	ChannelId                 *Many2One  `xmlrpc:"channel_id,omptempty"`
	CategoryId                *Many2One  `xmlrpc:"category_id,omptempty"`
	QuizDirstAttemptReward    *Int       `xmlrpc:"quiz_first_attempt_reward,omptempty"`
	QuizSecondAttemptReward   *Int       `xmlrpc:"quiz_second_attempt_reward,omptempty"`
	QuizThirdAttemptReward    *Int       `xmlrpc:"quiz_third_attempt_reward,omptempty"`
	QuizFourthAttemptReward   *Int       `xmlrpc:"quiz_fourth_attempt_reward,omptempty"`
	Likes                     *Int       `xmlrpc:"likes,omptempty"`
	Dislikes                  *Int       `xmlrpc:"dislikes,omptempty"`
	SlideViews                *Int       `xmlrpc:"slide_views,omptempty"`
	PublicViews               *Int       `xmlrpc:"public_views,omptempty"`
	TotalViews                *Int       `xmlrpc:"total_views,omptempty"`
	NbrDocument               *Int       `xmlrpc:"nbr_document,omptempty"`
	NbrVideo                  *Int       `xmlrpc:"nbr_video,omptempty"`
	NbrInfographic            *Int       `xmlrpc:"nbr_infographic,omptempty"`
	NbrArticle                *Int       `xmlrpc:"nbr_article,omptempty"`
	NbrQuiz                   *Int       `xmlrpc:"nbr_quiz,omptempty"`
	TotalSlides               *Int       `xmlrpc:"total_slides,omptempty"`
	CreateUid                 *Many2One  `xmlrpc:"create_uid,omptempty"`
	WriteUid                  *Many2One  `xmlrpc:"write_uid,omptempty"`
	WebsiteMetaOgImg          *String    `xmlrpc:"website_meta_og_img,omptempty"`
	SlideCategory             *String    `xmlrpc:"slide_category,omptempty"`
	SourceType                *String    `xmlrpc:"source_type,omptempty"`
	Url                       *String    `xmlrpc:"url,omptempty"`
	SlideType                 *String    `xmlrpc:"slide_type,omptempty"`
	WebsiteMetaTitle          *Selection `xmlrpc:"website_meta_title,omptempty"`
	WebsiteMetaDescription    *Selection `xmlrpc:"website_meta_description,omptempty"`
	WebsiteMetaKeywords       *Selection `xmlrpc:"website_meta_keywords,omptempty"`
	SeoName                   *Selection `xmlrpc:"seo_name,omptempty"`
	Name                      *Selection `xmlrpc:"name,omptempty"`
	Description               *Selection `xmlrpc:"description,omptempty"`
	HtmlContent               *Selection `xmlrpc:"html_content,omptempty"`
	CompletionTime            *Float     `xmlrpc:"completion_time,omptempty"`
	IsPublished               *Bool      `xmlrpc:"is_published,omptempty"`
	Active                    *Bool      `xmlrpc:"active,omptempty"`
	IsPreview                 *Bool      `xmlrpc:"is_preview,omptempty"`
	IsCategory                *Bool      `xmlrpc:"is_category,omptempty"`
	SlideResourceDownloadable *Bool      `xmlrpc:"slide_resource_downloadable,omptempty"`
	DatePublished             *Time      `xmlrpc:"date_published,omptempty"`
	CreateDate                *Time      `xmlrpc:"create_date,omptempty"`
	WriteDate                 *Time      `xmlrpc:"write_date,omptempty"`
}

type SlideSlides []SlideSlide

const SlideSlideModel = "slide.slide"

func (ss *SlideSlide) Many2One() *Many2One {
	return NewMany2One(ss.Id.Get(), "")
}

func (c *Client) CreateSlideSlide(ss *SlideSlide) (int64, error) {
	ids, err := c.CreateSlideSlides([]*SlideSlide{ss})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

func (c *Client) CreateSlideSlides(ss []*SlideSlide) ([]int64, error) {
	var vv []interface{}
	for _, v := range ss {
		vv = append(vv, v)
	}
	return c.Create(SlideSlideModel, vv)
}

func (c *Client) UpdateSlideSlide(ss *SlideSlide) error {
	return c.UpdateSlideSlides([]int64{ss.Id.Get()}, ss)
}

func (c *Client) UpdateSlideSlides(ids []int64, ss *SlideSlide) error {
	return c.Update(SlideSlideModel, ids, ss)
}

func (c *Client) DeleteSlideSlide(id int64) error {
	return c.DeleteSlideSlides([]int64{id})
}

func (c *Client) DeleteSlideSlides(ids []int64) error {
	return c.Delete(SlideSlideModel, ids)
}

func (c *Client) GetSlideSlide(id int64) (*SlideSlide, error) {
	sss, err := c.GetSlideSlides([]int64{id})
	if err != nil {
		return nil, err
	}
	if sss != nil && len(*sss) > 0 {
		return &((*sss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of slide.slide not found", id)
}

func (c *Client) GetSlideSlides(ids []int64) (*SlideSlides, error) {
	sss := &SlideSlides{}
	if err := c.Read(SlideSlideModel, ids, nil, sss); err != nil {
		return nil, err
	}
	return sss, nil
}

func (c *Client) FindSlideSlide(criteria *Criteria) (*SlideSlide, error) {
	sss := &SlideSlides{}
	if err := c.SearchRead(SlideSlideModel, criteria, NewOptions().Limit(1), sss); err != nil {
		return nil, err
	}
	if sss != nil && len(*sss) > 0 {
		return &((*sss)[0]), nil
	}
	return nil, fmt.Errorf("slide.slide was not found with criteria %v", criteria)
}

func (c *Client) FindSlideSlides(criteria *Criteria, options *Options) (*SlideSlides, error) {
	sss := &SlideSlides{}
	if err := c.SearchRead(SlideSlideModel, criteria, options, sss); err != nil {
		return nil, err
	}
	return sss, nil
}

func (c *Client) FindSlideSlideIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SlideSlideModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

func (c *Client) FindSlideSlideId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SlideSlideModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("slide.slide was not found with criteria %v and options %v", criteria, options)
}
