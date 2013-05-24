// Copyright 2013 Google. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file or at
// https://developers.google.com/open-source/licenses/bsd

package github

import (
	"fmt"
	"net/url"
	"strconv"
)

// UsersService handles communication with the user related
// methods of the GitHub API.
//
// GitHub API docs: http://developer.github.com/v3/users/
type UsersService struct {
	client *Client
}

type User struct {
	Login       string `json:"login,omitempty"`
	ID          int    `json:"id,omitempty"`
	URL         string `json:"url,omitempty"`
	AvatarURL   string `json:"avatar_url,omitempty"`
	GravatarID  string `json:"gravatar_id,omitempty"`
	Name        string `json:"name,omitempty"`
	Company     string `json:"company,omitempty"`
	Blog        string `json:"blog,omitempty"`
	Location    string `json:"location,omitempty"`
	Email       string `json:"email,omitempty"`
	Hireable    bool   `json:"hireable,omitempty"`
	PublicRepos int    `json:"public_repos,omitempty"`
	Followers   int    `json:"followers,omitempty"`
	Following   int    `json:"following,omitempty"`
}

// Get fetches a user.  Passing the empty string will fetch the authenticated
// user.
func (s *UsersService) Get(user string) (*User, error) {
	var url_ string
	if user != "" {
		url_ = fmt.Sprintf("users/%v", user)
	} else {
		url_ = "user"
	}
	req, err := s.client.NewRequest("GET", url_, nil)
	if err != nil {
		return nil, err
	}

	u := new(User)
	_, err = s.client.Do(req, u)
	return u, err
}

// Edit the authenticated user.
func (s *UsersService) Edit(user *User) (*User, error) {
	url_ := "user"
	req, err := s.client.NewRequest("PATCH", url_, user)
	if err != nil {
		return nil, err
	}

	u := new(User)
	_, err = s.client.Do(req, u)
	return u, err
}

// UserListOptions specifies optional paramters to the UsersService.List
// method.
type UserListOptions struct {
	// ID of the last user seen
	Since int
}

// List all users.
func (s *UsersService) List(opt *UserListOptions) ([]User, error) {
	url_ := "users"
	if opt != nil {
		params := url.Values{
			"since": []string{strconv.Itoa(opt.Since)},
		}
		url_ += "?" + params.Encode()
	}

	req, err := s.client.NewRequest("GET", url_, nil)
	if err != nil {
		return nil, err
	}

	users := new([]User)
	_, err = s.client.Do(req, users)
	return *users, err
}
