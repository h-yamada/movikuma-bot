package config

import (
	"errors"

	validator "github.com/asaskevich/govalidator"
)

var Conf Config

type Config struct {
	Facebook facebook
	Movikuma movikuma
}

type facebook struct {
	Token string
}

type movikuma struct {
	ElasticSearchUri string `toml:"elasticsearch"`
	DetailPageUri    string `toml:"detail_uri"`
	ImageUri         string `toml:"image_uri"`
}

func (c Config) Validate() error {
	if validator.IsNull(c.Facebook.Token) {
		return errors.New("Facebook Token")
	}

	if !validator.IsURL(c.Movikuma.ElasticSearchUri) {
		return errors.New("elasticsearch")
	}
	if !validator.IsURL(c.Movikuma.DetailPageUri) {
		return errors.New("detail_uri")
	}
	if !validator.IsURL(c.Movikuma.ImageUri) {
		return errors.New("image_uri")
	}
	return nil
}
