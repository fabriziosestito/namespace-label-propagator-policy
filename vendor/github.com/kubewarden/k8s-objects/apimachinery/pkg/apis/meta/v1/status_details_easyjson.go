// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonFdf086a1DecodeGithubComKubewardenK8sObjectsApimachineryPkgApisMetaV1(in *jlexer.Lexer, out *StatusDetails) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "causes":
			if in.IsNull() {
				in.Skip()
				out.Causes = nil
			} else {
				in.Delim('[')
				if out.Causes == nil {
					if !in.IsDelim(']') {
						out.Causes = make([]*StatusCause, 0, 8)
					} else {
						out.Causes = []*StatusCause{}
					}
				} else {
					out.Causes = (out.Causes)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *StatusCause
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(StatusCause)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.Causes = append(out.Causes, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "group":
			out.Group = string(in.String())
		case "kind":
			out.Kind = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "retryAfterSeconds":
			out.RetryAfterSeconds = int32(in.Int32())
		case "uid":
			out.UID = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonFdf086a1EncodeGithubComKubewardenK8sObjectsApimachineryPkgApisMetaV1(out *jwriter.Writer, in StatusDetails) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Causes) != 0 {
		const prefix string = ",\"causes\":"
		first = false
		out.RawString(prefix[1:])
		{
			out.RawByte('[')
			for v2, v3 := range in.Causes {
				if v2 > 0 {
					out.RawByte(',')
				}
				if v3 == nil {
					out.RawString("null")
				} else {
					(*v3).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	if in.Group != "" {
		const prefix string = ",\"group\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Group))
	}
	if in.Kind != "" {
		const prefix string = ",\"kind\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Kind))
	}
	if in.Name != "" {
		const prefix string = ",\"name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	if in.RetryAfterSeconds != 0 {
		const prefix string = ",\"retryAfterSeconds\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.RetryAfterSeconds))
	}
	if in.UID != "" {
		const prefix string = ",\"uid\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.UID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v StatusDetails) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonFdf086a1EncodeGithubComKubewardenK8sObjectsApimachineryPkgApisMetaV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v StatusDetails) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonFdf086a1EncodeGithubComKubewardenK8sObjectsApimachineryPkgApisMetaV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *StatusDetails) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonFdf086a1DecodeGithubComKubewardenK8sObjectsApimachineryPkgApisMetaV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *StatusDetails) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonFdf086a1DecodeGithubComKubewardenK8sObjectsApimachineryPkgApisMetaV1(l, v)
}