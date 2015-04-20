// DO NOT EDIT!
// Code generated by ffjson <https://github.com/pquerna/ffjson>
// source: request.go
// DO NOT EDIT!

package openrtb

import (
	"bytes"
	"encoding/json"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

func (mj *Request) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *Request) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	var err error
	var obj []byte
	_ = obj
	_ = err
	if mj.Id != nil {
		buf.WriteString(`{ "id":`)
		fflib.WriteJsonString(buf, string(*mj.Id))
	} else {
		buf.WriteString(`{ "id":null`)
	}
	buf.WriteByte(',')
	if len(mj.Imp) != 0 {
		buf.WriteString(`"imp":`)
		if mj.Imp != nil {
			buf.WriteString(`[`)
			for i, v := range mj.Imp {
				if i != 0 {
					buf.WriteString(`,`)
				}

				{
					err = v.MarshalJSONBuf(buf)
					if err != nil {
						return err
					}
				}

			}
			buf.WriteString(`]`)
		} else {
			buf.WriteString(`null`)
		}
		buf.WriteByte(',')
	}
	if mj.Site != nil {
		if true {
			buf.WriteString(`"site":`)

			{
				err = mj.Site.MarshalJSONBuf(buf)
				if err != nil {
					return err
				}
			}

			buf.WriteByte(',')
		}
	}
	if mj.App != nil {
		if true {
			buf.WriteString(`"app":`)

			{
				err = mj.App.MarshalJSONBuf(buf)
				if err != nil {
					return err
				}
			}

			buf.WriteByte(',')
		}
	}
	if mj.Device != nil {
		if true {
			buf.WriteString(`"device":`)

			{
				err = mj.Device.MarshalJSONBuf(buf)
				if err != nil {
					return err
				}
			}

			buf.WriteByte(',')
		}
	}
	if mj.User != nil {
		if true {
			buf.WriteString(`"user":`)

			{
				err = mj.User.MarshalJSONBuf(buf)
				if err != nil {
					return err
				}
			}

			buf.WriteByte(',')
		}
	}
	if mj.At != nil {
		buf.WriteString(`"at":`)
		fflib.FormatBits2(buf, uint64(*mj.At), 10, *mj.At < 0)
	} else {
		buf.WriteString(`"at":null`)
	}
	buf.WriteByte(',')
	if mj.Tmax != nil {
		if true {
			buf.WriteString(`"tmax":`)
			fflib.FormatBits2(buf, uint64(*mj.Tmax), 10, *mj.Tmax < 0)
			buf.WriteByte(',')
		}
	}
	if len(mj.Wseat) != 0 {
		buf.WriteString(`"wseat":`)
		if mj.Wseat != nil {
			buf.WriteString(`[`)
			for i, v := range mj.Wseat {
				if i != 0 {
					buf.WriteString(`,`)
				}
				fflib.WriteJsonString(buf, string(v))
			}
			buf.WriteString(`]`)
		} else {
			buf.WriteString(`null`)
		}
		buf.WriteByte(',')
	}
	if mj.Allimps != nil {
		if true {
			buf.WriteString(`"allimps":`)
			fflib.FormatBits2(buf, uint64(*mj.Allimps), 10, *mj.Allimps < 0)
			buf.WriteByte(',')
		}
	}
	if len(mj.Cur) != 0 {
		buf.WriteString(`"cur":`)
		if mj.Cur != nil {
			buf.WriteString(`[`)
			for i, v := range mj.Cur {
				if i != 0 {
					buf.WriteString(`,`)
				}
				fflib.WriteJsonString(buf, string(v))
			}
			buf.WriteString(`]`)
		} else {
			buf.WriteString(`null`)
		}
		buf.WriteByte(',')
	}
	if len(mj.Bcat) != 0 {
		buf.WriteString(`"bcat":`)
		if mj.Bcat != nil {
			buf.WriteString(`[`)
			for i, v := range mj.Bcat {
				if i != 0 {
					buf.WriteString(`,`)
				}
				fflib.WriteJsonString(buf, string(v))
			}
			buf.WriteString(`]`)
		} else {
			buf.WriteString(`null`)
		}
		buf.WriteByte(',')
	}
	if len(mj.Badv) != 0 {
		buf.WriteString(`"badv":`)
		if mj.Badv != nil {
			buf.WriteString(`[`)
			for i, v := range mj.Badv {
				if i != 0 {
					buf.WriteString(`,`)
				}
				fflib.WriteJsonString(buf, string(v))
			}
			buf.WriteString(`]`)
		} else {
			buf.WriteString(`null`)
		}
		buf.WriteByte(',')
	}
	if mj.Regs != nil {
		if true {
			buf.WriteString(`"regs":`)

			{
				err = mj.Regs.MarshalJSONBuf(buf)
				if err != nil {
					return err
				}
			}

			buf.WriteByte(',')
		}
	}
	if len(mj.Ext) != 0 {
		buf.WriteString(`"ext":`)
		/* Falling back. type=openrtb.Extensions kind=map */
		err = buf.Encode(mj.Ext)
		if err != nil {
			return err
		}
		buf.WriteByte(',')
	}
	if mj.Pmp != nil {
		if true {
			buf.WriteString(`"pmp":`)

			{
				err = mj.Pmp.MarshalJSONBuf(buf)
				if err != nil {
					return err
				}
			}

			buf.WriteByte(',')
		}
	}
	buf.Rewind(1)
	buf.WriteByte('}')
	return nil
}

const (
	ffj_t_Requestbase = iota
	ffj_t_Requestno_such_key

	ffj_t_Request_Id

	ffj_t_Request_Imp

	ffj_t_Request_Site

	ffj_t_Request_App

	ffj_t_Request_Device

	ffj_t_Request_User

	ffj_t_Request_At

	ffj_t_Request_Tmax

	ffj_t_Request_Wseat

	ffj_t_Request_Allimps

	ffj_t_Request_Cur

	ffj_t_Request_Bcat

	ffj_t_Request_Badv

	ffj_t_Request_Regs

	ffj_t_Request_Ext

	ffj_t_Request_Pmp
)

var ffj_key_Request_Id = []byte("id")

var ffj_key_Request_Imp = []byte("imp")

var ffj_key_Request_Site = []byte("site")

var ffj_key_Request_App = []byte("app")

var ffj_key_Request_Device = []byte("device")

var ffj_key_Request_User = []byte("user")

var ffj_key_Request_At = []byte("at")

var ffj_key_Request_Tmax = []byte("tmax")

var ffj_key_Request_Wseat = []byte("wseat")

var ffj_key_Request_Allimps = []byte("allimps")

var ffj_key_Request_Cur = []byte("cur")

var ffj_key_Request_Bcat = []byte("bcat")

var ffj_key_Request_Badv = []byte("badv")

var ffj_key_Request_Regs = []byte("regs")

var ffj_key_Request_Ext = []byte("ext")

var ffj_key_Request_Pmp = []byte("pmp")

func (uj *Request) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return uj.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

func (uj *Request) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error = nil
	currentKey := ffj_t_Requestbase
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffj_t_Requestno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'a':

					if bytes.Equal(ffj_key_Request_App, kn) {
						currentKey = ffj_t_Request_App
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Request_At, kn) {
						currentKey = ffj_t_Request_At
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Request_Allimps, kn) {
						currentKey = ffj_t_Request_Allimps
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'b':

					if bytes.Equal(ffj_key_Request_Bcat, kn) {
						currentKey = ffj_t_Request_Bcat
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Request_Badv, kn) {
						currentKey = ffj_t_Request_Badv
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'c':

					if bytes.Equal(ffj_key_Request_Cur, kn) {
						currentKey = ffj_t_Request_Cur
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'd':

					if bytes.Equal(ffj_key_Request_Device, kn) {
						currentKey = ffj_t_Request_Device
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'e':

					if bytes.Equal(ffj_key_Request_Ext, kn) {
						currentKey = ffj_t_Request_Ext
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'i':

					if bytes.Equal(ffj_key_Request_Id, kn) {
						currentKey = ffj_t_Request_Id
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Request_Imp, kn) {
						currentKey = ffj_t_Request_Imp
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'p':

					if bytes.Equal(ffj_key_Request_Pmp, kn) {
						currentKey = ffj_t_Request_Pmp
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'r':

					if bytes.Equal(ffj_key_Request_Regs, kn) {
						currentKey = ffj_t_Request_Regs
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 's':

					if bytes.Equal(ffj_key_Request_Site, kn) {
						currentKey = ffj_t_Request_Site
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 't':

					if bytes.Equal(ffj_key_Request_Tmax, kn) {
						currentKey = ffj_t_Request_Tmax
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'u':

					if bytes.Equal(ffj_key_Request_User, kn) {
						currentKey = ffj_t_Request_User
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'w':

					if bytes.Equal(ffj_key_Request_Wseat, kn) {
						currentKey = ffj_t_Request_Wseat
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.SimpleLetterEqualFold(ffj_key_Request_Pmp, kn) {
					currentKey = ffj_t_Request_Pmp
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Request_Ext, kn) {
					currentKey = ffj_t_Request_Ext
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Request_Regs, kn) {
					currentKey = ffj_t_Request_Regs
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Request_Badv, kn) {
					currentKey = ffj_t_Request_Badv
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Request_Bcat, kn) {
					currentKey = ffj_t_Request_Bcat
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Request_Cur, kn) {
					currentKey = ffj_t_Request_Cur
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Request_Allimps, kn) {
					currentKey = ffj_t_Request_Allimps
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Request_Wseat, kn) {
					currentKey = ffj_t_Request_Wseat
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Request_Tmax, kn) {
					currentKey = ffj_t_Request_Tmax
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Request_At, kn) {
					currentKey = ffj_t_Request_At
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Request_User, kn) {
					currentKey = ffj_t_Request_User
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Request_Device, kn) {
					currentKey = ffj_t_Request_Device
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Request_App, kn) {
					currentKey = ffj_t_Request_App
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Request_Site, kn) {
					currentKey = ffj_t_Request_Site
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Request_Imp, kn) {
					currentKey = ffj_t_Request_Imp
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Request_Id, kn) {
					currentKey = ffj_t_Request_Id
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffj_t_Requestno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffj_t_Request_Id:
					goto handle_Id

				case ffj_t_Request_Imp:
					goto handle_Imp

				case ffj_t_Request_Site:
					goto handle_Site

				case ffj_t_Request_App:
					goto handle_App

				case ffj_t_Request_Device:
					goto handle_Device

				case ffj_t_Request_User:
					goto handle_User

				case ffj_t_Request_At:
					goto handle_At

				case ffj_t_Request_Tmax:
					goto handle_Tmax

				case ffj_t_Request_Wseat:
					goto handle_Wseat

				case ffj_t_Request_Allimps:
					goto handle_Allimps

				case ffj_t_Request_Cur:
					goto handle_Cur

				case ffj_t_Request_Bcat:
					goto handle_Bcat

				case ffj_t_Request_Badv:
					goto handle_Badv

				case ffj_t_Request_Regs:
					goto handle_Regs

				case ffj_t_Request_Ext:
					goto handle_Ext

				case ffj_t_Request_Pmp:
					goto handle_Pmp

				case ffj_t_Requestno_such_key:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_Id:

	/* handler: uj.Id type=string kind=string */

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

			uj.Id = nil

		} else {

			var tval string
			tval = string(fs.Output.String())
			uj.Id = &tval

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Imp:

	/* handler: uj.Imp type=[]openrtb.Impression kind=slice */

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			uj.Imp = nil
		} else {

			uj.Imp = make([]Impression, 0)

			wantVal := true

			for {

				var v Impression

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: v type=openrtb.Impression kind=struct */

				{
					if tok == fflib.FFTok_null {

						state = fflib.FFParse_after_value
						goto mainparse
					}

					err = v.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
					if err != nil {
						return err
					}
					state = fflib.FFParse_after_value
				}

				uj.Imp = append(uj.Imp, v)
				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Site:

	/* handler: uj.Site type=openrtb.Site kind=struct */

	{
		if tok == fflib.FFTok_null {

			uj.Site = nil

			state = fflib.FFParse_after_value
			goto mainparse
		}

		if uj.Site == nil {
			uj.Site = new(Site)
		}

		err = uj.Site.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
		if err != nil {
			return err
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_App:

	/* handler: uj.App type=openrtb.App kind=struct */

	{
		if tok == fflib.FFTok_null {

			uj.App = nil

			state = fflib.FFParse_after_value
			goto mainparse
		}

		if uj.App == nil {
			uj.App = new(App)
		}

		err = uj.App.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
		if err != nil {
			return err
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Device:

	/* handler: uj.Device type=openrtb.Device kind=struct */

	{
		if tok == fflib.FFTok_null {

			uj.Device = nil

			state = fflib.FFParse_after_value
			goto mainparse
		}

		if uj.Device == nil {
			uj.Device = new(Device)
		}

		err = uj.Device.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
		if err != nil {
			return err
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_User:

	/* handler: uj.User type=openrtb.User kind=struct */

	{
		if tok == fflib.FFTok_null {

			uj.User = nil

			state = fflib.FFParse_after_value
			goto mainparse
		}

		if uj.User == nil {
			uj.User = new(User)
		}

		err = uj.User.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
		if err != nil {
			return err
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_At:

	/* handler: uj.At type=int kind=int */

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

			uj.At = nil

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			ttypval := int(tval)
			uj.At = &ttypval

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Tmax:

	/* handler: uj.Tmax type=int kind=int */

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

			uj.Tmax = nil

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			ttypval := int(tval)
			uj.Tmax = &ttypval

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Wseat:

	/* handler: uj.Wseat type=[]string kind=slice */

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			uj.Wseat = nil
		} else {

			uj.Wseat = make([]string, 0)

			wantVal := true

			for {

				var v string

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: v type=string kind=string */

				{

					{
						if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
							return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
						}
					}

					if tok == fflib.FFTok_null {

					} else {

						v = string(fs.Output.String())

					}
				}

				uj.Wseat = append(uj.Wseat, v)
				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Allimps:

	/* handler: uj.Allimps type=int kind=int */

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

			uj.Allimps = nil

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			ttypval := int(tval)
			uj.Allimps = &ttypval

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Cur:

	/* handler: uj.Cur type=[]string kind=slice */

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			uj.Cur = nil
		} else {

			uj.Cur = make([]string, 0)

			wantVal := true

			for {

				var v string

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: v type=string kind=string */

				{

					{
						if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
							return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
						}
					}

					if tok == fflib.FFTok_null {

					} else {

						v = string(fs.Output.String())

					}
				}

				uj.Cur = append(uj.Cur, v)
				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Bcat:

	/* handler: uj.Bcat type=[]string kind=slice */

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			uj.Bcat = nil
		} else {

			uj.Bcat = make([]string, 0)

			wantVal := true

			for {

				var v string

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: v type=string kind=string */

				{

					{
						if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
							return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
						}
					}

					if tok == fflib.FFTok_null {

					} else {

						v = string(fs.Output.String())

					}
				}

				uj.Bcat = append(uj.Bcat, v)
				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Badv:

	/* handler: uj.Badv type=[]string kind=slice */

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			uj.Badv = nil
		} else {

			uj.Badv = make([]string, 0)

			wantVal := true

			for {

				var v string

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: v type=string kind=string */

				{

					{
						if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
							return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
						}
					}

					if tok == fflib.FFTok_null {

					} else {

						v = string(fs.Output.String())

					}
				}

				uj.Badv = append(uj.Badv, v)
				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Regs:

	/* handler: uj.Regs type=openrtb.Regulations kind=struct */

	{
		if tok == fflib.FFTok_null {

			uj.Regs = nil

			state = fflib.FFParse_after_value
			goto mainparse
		}

		if uj.Regs == nil {
			uj.Regs = new(Regulations)
		}

		err = uj.Regs.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
		if err != nil {
			return err
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Ext:

	/* handler: uj.Ext type=openrtb.Extensions kind=map */

	{
		/* Falling back. type=openrtb.Extensions kind=map */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &uj.Ext)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Pmp:

	/* handler: uj.Pmp type=openrtb.Pmp kind=struct */

	{
		if tok == fflib.FFTok_null {

			uj.Pmp = nil

			state = fflib.FFParse_after_value
			goto mainparse
		}

		if uj.Pmp == nil {
			uj.Pmp = new(Pmp)
		}

		err = uj.Pmp.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
		if err != nil {
			return err
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:
	return nil
}
