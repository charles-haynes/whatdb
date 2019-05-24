//Package whatapi is a wrapper for the What.CD JSON API (https://github.com/WhatCD/Gazelle/wiki/JSON-API-Documentation).
package whatapi

import (
	"net/url"

	"github.com/charles-haynes/whatapi"
	"github.com/jmoiron/sqlx"
)

//NewWhatAPI creates a new client for the What.CD API using the provided URL.
func NewWhatDB(url, agent string) (WhatDB, error) {
	api, err := whatapi.NewWhatAPI(url, agent)
	if err != nil {
		return nil, err
	}
	db, err := sqlx.Connect("sqlite3", "./.whatapi.cache.db")
	if err != nil {
		return nil, err
	}
	whatDB := WhatDBStruct{
		WhatAPI: api,
		DB:      db,
	}
	return &whatDB, nil
}

//WhatDB represents a caching client for the What.CD API.
type WhatDB interface {
	GetJSON(requestURL string, responseObj interface{}) error
	Do(action string, params url.Values, result interface{}) error
	CreateDownloadURL(id int) (string, error)
	Login(username, password string) error
	Logout() error
	GetAccount() (whatapi.Account, error)
	GetMailbox(params url.Values) (whatapi.Mailbox, error)
	GetConversation(id int) (whatapi.Conversation, error)
	GetNotifications(params url.Values) (whatapi.Notifications, error)
	GetAnnouncements() (whatapi.Announcements, error)
	GetSubscriptions(params url.Values) (whatapi.Subscriptions, error)
	GetCategories() (whatapi.Categories, error)
	GetForum(id int, params url.Values) (whatapi.Forum, error)
	GetThread(id int, params url.Values) (whatapi.Thread, error)
	GetArtistBookmarks() (whatapi.ArtistBookmarks, error)
	GetTorrentBookmarks() (whatapi.TorrentBookmarks, error)
	GetArtist(id int, params url.Values) (whatapi.Artist, error)
	GetRequest(id int, params url.Values) (whatapi.Request, error)
	GetTorrent(id int, params url.Values) (whatapi.GetTorrentStruct, error)
	GetTorrentGroup(id int, params url.Values) (whatapi.TorrentGroup, error)
	SearchTorrents(searchStr string, params url.Values) (
		whatapi.TorrentSearch, error)
	SearchRequests(searchStr string, params url.Values) (
		whatapi.RequestsSearch, error)
	SearchUsers(searchStr string, params url.Values) (
		whatapi.UserSearch, error)
	GetTopTenTorrents(params url.Values) (whatapi.TopTenTorrents, error)
	GetTopTenTags(params url.Values) (whatapi.TopTenTags, error)
	GetTopTenUsers(params url.Values) (whatapi.TopTenUsers, error)
	GetSimilarArtists(id, limit int) (whatapi.SimilarArtists, error)
	ParseHTML(s string) (string, error)
}

//WhatDBStruct represents a client for the What.CD API plus db cache.
type WhatDBStruct struct {
	whatapi.WhatAPI
	DB *sqlx.DB
}

//GetJSON sends a HTTP GET request to the API and decodes the JSON response into responseObj.
func (w *WhatDBStruct) GetJSON(requestURL string, responseObj interface{}) error {
	return w.WhatAPI.GetJSON(requestURL, responseObj)
}

func (w *WhatDBStruct) Do(action string, params url.Values, result interface{}) error {
	return w.WhatAPI.Do(action, params, result)
}

//CreateDownloadURL constructs a download URL using the provided torrent id.
func (w *WhatDBStruct) CreateDownloadURL(id int) (string, error) {
	return w.WhatAPI.CreateDownloadURL(id)
}

//Login logs in to the API using the provided credentials.
func (w *WhatDBStruct) Login(username, password string) error {
	return w.WhatAPI.Login(username, password)
}

//Logout logs out of the API, ending the current session.
func (w *WhatDBStruct) Logout() error {
	return w.WhatAPI.Logout()
}

//GetAccount retrieves account information for the current user.
func (w *WhatDBStruct) GetAccount() (whatapi.Account, error) {
	return w.WhatAPI.GetAccount()
}

//GetMailbox retrieves mailbox information for the current user using
//the provided parameters.
func (w *WhatDBStruct) GetMailbox(params url.Values) (whatapi.Mailbox, error) {
	return w.WhatAPI.GetMailbox(params)
}

//GetConversation retrieves conversation information for the current
//user using the provided conversation id and parameters.
func (w *WhatDBStruct) GetConversation(id int) (whatapi.Conversation, error) {
	return w.WhatAPI.GetConversation(id)
}

//GetNotifications retrieves notification information using the
//specifed parameters.
func (w *WhatDBStruct) GetNotifications(params url.Values) (
	whatapi.Notifications, error) {
	return w.WhatAPI.GetNotifications(params)
}

//GetAnnouncements retrieves announcement information.
func (w *WhatDBStruct) GetAnnouncements() (whatapi.Announcements, error) {
	return w.WhatAPI.GetAnnouncements()
}

//GetSubscriptions retrieves forum subscription information for the
//current user using the provided parameters.
func (w *WhatDBStruct) GetSubscriptions(params url.Values) (
	whatapi.Subscriptions, error) {
	return w.WhatAPI.GetSubscriptions(params)
}

//GetCategories retrieves forum category information.
func (w *WhatDBStruct) GetCategories() (whatapi.Categories, error) {
	return w.WhatAPI.GetCategories()
}

//GetForum retrieves forum information using the provided forum id and
//parameters.
func (w *WhatDBStruct) GetForum(id int, params url.Values) (
	whatapi.Forum, error) {
	return w.WhatAPI.GetForum(id, params)
}

//GetThread retrieves forum thread information using the provided
//thread id and parameters.
func (w *WhatDBStruct) GetThread(id int, params url.Values) (
	whatapi.Thread, error) {
	return w.WhatAPI.GetThread(id, params)
}

//GetArtistBookmarks retrieves artist bookmark information for the current user.
func (w *WhatDBStruct) GetArtistBookmarks() (whatapi.ArtistBookmarks, error) {
	return w.WhatAPI.GetArtistBookmarks()
}

//GetTorrentBookmarks retrieves torrent bookmark information for the
//current user.
func (w *WhatDBStruct) GetTorrentBookmarks() (whatapi.TorrentBookmarks, error) {
	return w.WhatAPI.GetTorrentBookmarks()
}

//GetArtist retrieves artist information using the provided artist id
//and parameters.
func (w *WhatDBStruct) GetArtist(id int, params url.Values) (
	whatapi.Artist, error) {
	return w.WhatAPI.GetArtist(id, params)
}

//GetRequest retrieves request information using the provided request
//id and parameters.
func (w *WhatDBStruct) GetRequest(id int, params url.Values) (
	whatapi.Request, error) {
	return w.WhatAPI.GetRequest(id, params)
}

//GetTorrent retrieves torrent information using the provided torrent
//id and parameters.
func (w *WhatDBStruct) GetTorrent(id int, params url.Values) (
	whatapi.GetTorrentStruct, error) {
	return w.WhatAPI.GetTorrent(id, params)
}

//GetTorrentGroup retrieves torrent group information using the
//provided torrent group id and parameters.
func (w *WhatDBStruct) GetTorrentGroup(id int, params url.Values) (
	whatapi.TorrentGroup, error) {
	return w.WhatAPI.GetTorrentGroup(id, params)
}

//SearchTorrents retrieves torrent search results using the provided
//search string and parameters.
func (w *WhatDBStruct) SearchTorrents(searchStr string, params url.Values) (
	whatapi.TorrentSearch, error) {
	return w.WhatAPI.SearchTorrents(searchStr, params)
}

//SearchRequests retrieves request search results using the provided
//search string and parameters.
func (w *WhatDBStruct) SearchRequests(searchStr string, params url.Values) (
	whatapi.RequestsSearch, error) {
	return w.WhatAPI.SearchRequests(searchStr, params)
}

//SearchUsers retrieves user search results using the provided search
//string and parameters.
func (w *WhatDBStruct) SearchUsers(searchStr string, params url.Values) (
	whatapi.UserSearch, error) {
	return w.WhatAPI.SearchUsers(searchStr, params)
}

//GetTopTenTorrents retrieves "top ten torrents" information using the
//provided parameters.
func (w *WhatDBStruct) GetTopTenTorrents(params url.Values) (
	whatapi.TopTenTorrents, error) {
	return w.WhatAPI.GetTopTenTorrents(params)
}

//GetTopTenTags retrieves "top ten tags" information using the
//provided parameters.
func (w *WhatDBStruct) GetTopTenTags(params url.Values) (
	whatapi.TopTenTags, error) {
	return w.WhatAPI.GetTopTenTags(params)
}

//GetTopTenUsers retrieves "top tem users" information using the
//provided parameters.
func (w *WhatDBStruct) GetTopTenUsers(params url.Values) (
	whatapi.TopTenUsers, error) {
	return w.WhatAPI.GetTopTenUsers(params)
}

//GetSimilarArtists retrieves similar artist information using the
//provided artist id and limit.
func (w *WhatDBStruct) GetSimilarArtists(id, limit int) (
	whatapi.SimilarArtists, error) {
	return w.WhatAPI.GetSimilarArtists(id, limit)
}

// ParseHTML takes an HTML formatted string and passes it to the server
// to be converted into BBCode (only available on some Gazelle servers)
func (w *WhatDBStruct) ParseHTML(s string) (string, error) {
	return w.WhatAPI.ParseHTML(s)
}
