package main

import "net/url"

type VersioningMethod interface {
	RedirectURL(Product, Requester) url.URL
}

type VersionPrefixedURL struct{}

//func (*VersionPrefixedURL) RedirectURL(p Product, r Requester) url.URL {
//return &url.URL
//}

type Product struct {
	Name          string
	Host          string
	VersionMethod VersioningMethod
}

type Requester struct {
	Name string
	TargetedVersion
}
