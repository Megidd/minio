package cmd

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/minio/minio/pkg/madmin"
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *TierConfigMgr) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Tiers":
			var zb0002 uint32
			zb0002, err = dc.ReadMapHeader()
			if err != nil {
				err = msgp.WrapError(err, "Tiers")
				return
			}
			if z.Tiers == nil {
				z.Tiers = make(map[string]madmin.TierConfig, zb0002)
			} else if len(z.Tiers) > 0 {
				for key := range z.Tiers {
					delete(z.Tiers, key)
				}
			}
			for zb0002 > 0 {
				zb0002--
				var za0001 string
				var za0002 madmin.TierConfig
				za0001, err = dc.ReadString()
				if err != nil {
					err = msgp.WrapError(err, "Tiers")
					return
				}
				err = za0002.DecodeMsg(dc)
				if err != nil {
					err = msgp.WrapError(err, "Tiers", za0001)
					return
				}
				z.Tiers[za0001] = za0002
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *TierConfigMgr) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "Tiers"
	err = en.Append(0x81, 0xa5, 0x54, 0x69, 0x65, 0x72, 0x73)
	if err != nil {
		return
	}
	err = en.WriteMapHeader(uint32(len(z.Tiers)))
	if err != nil {
		err = msgp.WrapError(err, "Tiers")
		return
	}
	for za0001, za0002 := range z.Tiers {
		err = en.WriteString(za0001)
		if err != nil {
			err = msgp.WrapError(err, "Tiers")
			return
		}
		err = za0002.EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "Tiers", za0001)
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *TierConfigMgr) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "Tiers"
	o = append(o, 0x81, 0xa5, 0x54, 0x69, 0x65, 0x72, 0x73)
	o = msgp.AppendMapHeader(o, uint32(len(z.Tiers)))
	for za0001, za0002 := range z.Tiers {
		o = msgp.AppendString(o, za0001)
		o, err = za0002.MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "Tiers", za0001)
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *TierConfigMgr) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Tiers":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Tiers")
				return
			}
			if z.Tiers == nil {
				z.Tiers = make(map[string]madmin.TierConfig, zb0002)
			} else if len(z.Tiers) > 0 {
				for key := range z.Tiers {
					delete(z.Tiers, key)
				}
			}
			for zb0002 > 0 {
				var za0001 string
				var za0002 madmin.TierConfig
				zb0002--
				za0001, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Tiers")
					return
				}
				bts, err = za0002.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Tiers", za0001)
					return
				}
				z.Tiers[za0001] = za0002
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *TierConfigMgr) Msgsize() (s int) {
	s = 1 + 6 + msgp.MapHeaderSize
	if z.Tiers != nil {
		for za0001, za0002 := range z.Tiers {
			_ = za0002
			s += msgp.StringPrefixSize + len(za0001) + za0002.Msgsize()
		}
	}
	return
}
