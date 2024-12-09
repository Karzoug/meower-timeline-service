package entity

import "github.com/rs/xid"

type Post struct {
	AuthorID xid.ID
	PostID   xid.ID
	IsRepost bool
}

func (p Post) MarshalBinary() (data []byte, err error) {
	b := make([]byte, 41)

	p.AuthorID.Encode(b[0:20])
	p.PostID.Encode(b[20:40])
	if p.IsRepost {
		b[40] = 1
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
	p.IsRepost = data[40] == 1

	return nil
}
