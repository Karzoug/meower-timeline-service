package entity

import (
	"errors"

	"github.com/rs/xid"
)

const encodedLen = 41 // string encoded len

type Post struct {
	AuthorID xid.ID
	PostID   xid.ID
	IsRepost bool
}

func (p Post) String() string {
	enc, _ := p.MarshalBinary()
	return string(enc)
}

func FromString(post string) (Post, error) {
	p := &Post{}
	err := p.UnmarshalText([]byte(post))
	return *p, err
}

func (p *Post) UnmarshalText(text []byte) error {
	if len(text) != encodedLen {
		return errors.New("invalid post format")
	}

	if err := p.AuthorID.UnmarshalText(text[0:20]); err != nil {
		return err
	}
	if err := p.PostID.UnmarshalText(text[20:40]); err != nil {
		return err
	}
	p.IsRepost = text[40] == '1'

	return nil
}

func (p Post) MarshalBinary() (data []byte, err error) {
	b := make([]byte, 41)

	p.AuthorID.Encode(b[0:20])
	p.PostID.Encode(b[20:40])
	if p.IsRepost {
		b[40] = '1'
	}

	return b, nil
}

func (p *Post) UnmarshalBinary(data []byte) error {
	if err := p.AuthorID.UnmarshalText(data[0:20]); err != nil {
		return err
	}
	if err := p.PostID.UnmarshalText(data[20:40]); err != nil {
		return err
	}
	p.IsRepost = data[40] == '1'

	return nil
}
