package storagessl

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/caddyserver/certmagic"
)

/**
 */
func (rd *StorageParam) Store(ctx context.Context, key string, value []byte) error {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("sslcert", "sslcert")
	part.Write(value)
	Key, _ := writer.CreateFormField("key")
	Key.Write([]byte(key))
	Db, _ := writer.CreateFormField("database")
	Db.Write([]byte(rd.DataBaseName))
	writer.Close()
	Server := fmt.Sprintf("http://%s:%s/cm/sslstorage", rd.Host, rd.Port)
	r, _ := http.NewRequest("POST", Server, body)
	r = r.WithContext(ctx)

	r.Header.Add("Content-Type", writer.FormDataContentType())
	r.Header.Add("access_key", rd.AccessKey)
	r.Header.Add("command", "update")
	client := &http.Client{}
	Response, err := client.Do(r)

	if err != nil {
		return err
	} else if Response.StatusCode == 404 {
		return errors.New("error on store certs")
	}
	return nil
}

func (rd *StorageParam) Load(ctx context.Context, key string) ([]byte, error) {

	form := url.Values{}
	form.Add("key", key)
	form.Add("database", rd.DataBaseName)

	Server := fmt.Sprintf("http://%s:%s/cm/sslstorage", rd.Host, rd.Port)
	r, _ := http.NewRequest("POST", Server, strings.NewReader(form.Encode()))
	r = r.WithContext(ctx)

	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("access_key", rd.AccessKey)
	r.Header.Add("command", "load")
	client := &http.Client{}
	Response, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	if Response.StatusCode == 404 {
		return nil, fs.ErrNotExist
	} else if Response.StatusCode == 501 {
		return nil, errors.New("error on load certs")
	}
	buffer, err := io.ReadAll(Response.Body)

	return buffer, err
}

// Delete deletes key.
func (rd *StorageParam) Delete(ctx context.Context, key string) error {
	form := url.Values{}
	form.Add("key", key)
	form.Add("database", rd.DataBaseName)

	Server := fmt.Sprintf("http://%s:%s/cm/sslstorage", rd.Host, rd.Port)
	r, _ := http.NewRequest("POST", Server, strings.NewReader(form.Encode()))
	r = r.WithContext(ctx)

	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("access_key", rd.AccessKey)
	r.Header.Add("command", "delete")
	client := &http.Client{}
	Response, err := client.Do(r)
	if err != nil {
		return err
	} else if Response.StatusCode == 404 {
		return fs.ErrNotExist
	}
	return nil
}

// Exists returns true if the key exists
func (rd *StorageParam) Exists(ctx context.Context, key string) bool {
	form := url.Values{}
	form.Add("key", key)
	form.Add("database", rd.DataBaseName)

	Server := fmt.Sprintf("http://%s:%s/cm/sslstorage", rd.Host, rd.Port)
	r, _ := http.NewRequest("POST", Server, strings.NewReader(form.Encode()))
	r = r.WithContext(ctx)

	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("access_key", rd.AccessKey)
	r.Header.Add("command", "exists")
	client := &http.Client{}
	Response, err := client.Do(r)
	if err != nil {
		return false
	}
	if Response.StatusCode == 200 {
		return true
	}
	return false
}

// List returns all keys that match prefix.
func (rd *StorageParam) List(ctx context.Context, prefix string, recursive bool) ([]string, error) {
	form := url.Values{}
	form.Add("prefix", prefix)
	form.Add("recursive", strconv.FormatBool(recursive))
	form.Add("database", rd.DataBaseName)

	Server := fmt.Sprintf("http://%s:%s/cm/sslstorage", rd.Host, rd.Port)
	r, _ := http.NewRequest("POST", Server, strings.NewReader(form.Encode()))
	r = r.WithContext(ctx)

	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("access_key", rd.AccessKey)
	r.Header.Add("command", "exists")
	client := &http.Client{}
	_, err := client.Do(r)

	if err != nil {
		return nil, err
	}

	return nil, nil
}

// Stat returns information about key.
func (rd *StorageParam) Stat(ctx context.Context, key string) (certmagic.KeyInfo, error) {

	form := url.Values{}
	form.Add("key", key)
	form.Add("database", rd.DataBaseName)

	var result certmagic.KeyInfo
	Server := fmt.Sprintf("http://%s:%s/cm/sslstorage", rd.Host, rd.Port)
	r, _ := http.NewRequest("POST", Server, strings.NewReader(form.Encode()))
	r = r.WithContext(ctx)

	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("access_key", rd.AccessKey)
	r.Header.Add("command", "stat")
	client := &http.Client{}
	Response, err := client.Do(r)
	if err != nil {
		return result, err
	}

	body, err := ioutil.ReadAll(Response.Body)
	if err != nil {
		return result, err
	}
	json.Unmarshal(body, &result)
	return result, nil
}

func (rd *StorageParam) Lock(ctx context.Context, key string) error {
	form := url.Values{}
	form.Add("key", key)
	form.Add("database", rd.DataBaseName)

	Server := fmt.Sprintf("http://%s:%s/cm/sslstorage", rd.Host, rd.Port)
	r, _ := http.NewRequest("POST", Server, strings.NewReader(form.Encode()))
	r = r.WithContext(ctx)

	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("access_key", rd.AccessKey)
	r.Header.Add("command", "lock")
	client := &http.Client{}
	Response, err := client.Do(r)
	if err != nil {
		return err
	} else if Response.StatusCode == 200 {
		return nil
	}
	return errors.New("not lock: " + key)
}

func (rd *StorageParam) Unlock(ctx context.Context, key string) error {
	form := url.Values{}
	form.Add("key", key)
	form.Add("database", rd.DataBaseName)

	Server := fmt.Sprintf("http://%s:%s/cm/sslstorage", rd.Host, rd.Port)
	r, _ := http.NewRequest("POST", Server, strings.NewReader(form.Encode()))
	r = r.WithContext(ctx)

	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("access_key", rd.AccessKey)
	r.Header.Add("command", "unlock")
	client := &http.Client{}
	Response, err := client.Do(r)
	if err != nil {
		return err
	} else if Response.StatusCode == 200 {
		return nil
	}
	return errors.New("not unlock: " + key)
}
