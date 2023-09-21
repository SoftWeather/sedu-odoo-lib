package odoo

import "fmt"

type SlideChannel struct {
	Id                      *Int       `xmlrpc:"id,omptempty"`
	WebsiteId               *Int       `xmlrpc:"website_id,omptempty"`
	MessageMainAttachmentId *Int       `xmlrpc:"message_main_attachment_id,omptempty"`
	Sequence                *Int       `xmlrpc:"sequence,omptempty"`
	UserId                  *Many2One  `xmlrpc:"user_id,omptempty"`
	Color                   *Int       `xmlrpc:"color,omptempty"`
	PromotedSlideId         *Many2One  `xmlrpc:"promoted_slide_id,omptempty"`
	NBRDocument             *Int       `xmlrpc:"nbr_document,omptempty"`
	NBRVideo                *Int       `xmlrpc:"nbr_video,omptempty"`
	NBRInfographic          *Int       `xmlrpc:"nbr_infographic,omptempty"`
	NBRQuiz                 *Int       `xmlrpc:"nbr_quiz,omptempty"`
	TotalSlides             *Int       `xmlrpc:"total_slides,omptempty"`
	TotalViews              *Int       `xmlrpc:"total_views,omptempty"`
	TotalVotes              *Int       `xmlrpc:"total_votes,omptempty"`
	PublishTemplateId       *Many2One  `xmlrpc:"publish_template_id,omptempty"`
	ShareChannelTemplateId  *Many2One  `xmlrpc:"share_channel_template_id,omptempty"`
	ShareSlideTemplateId    *Many2One  `xmlrpc:"share_slide_template_id,omptempty"`
	CompletedTemplateId     *Many2One  `xmlrpc:"completed_template_id,omptempty"`
	KarmaGenSlideVote       *Int       `xmlrpc:"karma_gen_slide_vote,omptempty"`
	KarmaGenChannelRank     *Int       `xmlrpc:"karma_gen_channel_rank,omptempty"`
	KarmaGenChannelFinish   *Int       `xmlrpc:"karma_gen_channel_finish,omptempty"`
	KarmaReview             *Int       `xmlrpc:"karma_review,omptempty"`
	KarmaSlideComment       *Int       `xmlrpc:"karma_slide_comment,omptempty"`
	KarmaSlideVote          *Int       `xmlrpc:"karma_slide_vote,omptempty"`
	CreateUid               *Many2One  `xmlrpc:"create_uid,omptempty"`
	WriteUid                *Many2One  `xmlrpc:"write_uid,omptempty"`
	WebsiteMetaOgImg        *String    `xmlrpc:"website_meta_og_img,omptempty"`
	ChannelType             *String    `xmlrpc:"channel_type,omptempty"`
	PromoteStrategy         *String    `xmlrpc:"promote_strategy,omptempty"`
	AccessToken             *String    `xmlrpc:"access_token,omptempty"`
	Enroll                  *String    `xmlrpc:"enroll,omptempty"`
	Visibility              *String    `xmlrpc:"visibility,omptempty"`
	SlideLastUpdate         *Time      `xmlrpc:"slide_last_update,omptempty"`
	WebsiteMetaTitle        *Selection `xmlrpc:"website_meta_title,omptempty"`
	WebsiteMetaDescription  *Selection `xmlrpc:"website_meta_description,omptempty"`
	WebsiteMetaKeywords     *Selection `xmlrpc:"website_meta_keywords,omptempty"`
	SeoName                 *String    `xmlrpc:"seo_name,omptempty"`
	Name                    *String    `xmlrpc:"name,omptempty"`
	Description             *String    `xmlrpc:"description,omptempty"`
	DescriptionShort        *String    `xmlrpc:"description_short,omptempty"`
	DescriptionHtml         *String    `xmlrpc:"description_html,omptempty"`
	EnrollMsg               *Selection `xmlrpc:"enroll_msg,omptempty"`
	CoverProperties         *String    `xmlrpc:"cover_properties,omptempty"`
	TotalTime               *Float     `xmlrpc:"total_time,omptempty"`
	IsPublished             *Bool      `xmlrpc:"is_published,omptempty"`
	Active                  *Bool      `xmlrpc:"active,omptempty"`
	AllowComment            *Bool      `xmlrpc:"allow_comment,omptempty"`
	CreateDate              *Time      `xmlrpc:"create_date,omptempty"`
	WriteDate               *Time      `xmlrpc:"write_date,omptempty"`
	RatingLastValue         *Float     `xmlrpc:"rating_last_value,omptempty"`
}

type SlideChannels []SlideChannel

const SlideChannelModel = "slide.channel"

func (sc *SlideChannel) Many2One() *Many2One {
	return NewMany2One(sc.Id.Get(), "")
}

func (c *Client) CreateSlideChannel(sc *SlideChannel) (int64, error) {
	ids, err := c.CreateSlideChannels([]*SlideChannel{sc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

func (c *Client) CreateSlideChannels(sc []*SlideChannel) ([]int64, error) {
	var vv []interface{}
	for _, v := range sc {
		vv = append(vv, v)
	}
	return c.Create(SlideChannelModel, vv)
}

func (c *Client) UpdateSlideChannel(sc *SlideChannel) error {
	return c.UpdateSlideChannels([]int64{sc.Id.Get()}, sc)
}

func (c *Client) UpdateSlideChannels(ids []int64, sc *SlideChannel) error {
	return c.Update(SlideChannelModel, ids, sc)
}

func (c *Client) DeleteSlideChannel(id int64) error {
	return c.DeleteSlideChannels([]int64{id})
}

func (c *Client) DeleteSlideChannels(ids []int64) error {
	return c.Delete(SlideChannelModel, ids)
}

func (c *Client) GetSlideChannel(id int64) (*SlideChannel, error) {
	scs, err := c.GetSlideChannels([]int64{id})
	if err != nil {
		return nil, err
	}
	if scs != nil && len(*scs) > 0 {
		return &((*scs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of slide.channel not found", id)
}

func (c *Client) GetSlideChannels(ids []int64) (*SlideChannels, error) {
	scs := &SlideChannels{}
	if err := c.Read(SlideChannelModel, ids, nil, scs); err != nil {
		return nil, err
	}
	return scs, nil
}

func (c *Client) FindSlideChannel(criteria *Criteria) (*SlideChannel, error) {
	scs := &SlideChannels{}
	if err := c.SearchRead(SlideChannelModel, criteria, NewOptions().Limit(1), scs); err != nil {
		return nil, err
	}
	if scs != nil && len(*scs) > 0 {
		return &((*scs)[0]), nil
	}
	return nil, fmt.Errorf("slide.channel was not found with criteria %v", criteria)
}

func (c *Client) FindSlideChannels(criteria *Criteria, options *Options) (*SlideChannels, error) {
	scs := &SlideChannels{}
	if err := c.SearchRead(SlideChannelModel, criteria, options, scs); err != nil {
		return nil, err
	}
	return scs, nil
}

func (c *Client) FindSlideChannelIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SlideChannelModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

func (c *Client) FindSlideChannelId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SlideChannelModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("slide.channel was not found with criteria %v and options %v", criteria, options)
}
