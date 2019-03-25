// Autogenerated by Frugal Compiler (3.1.1)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package golang

import (
	"fmt"

	"github.com/Workiva/frugal/lib/gopherjs/frugal"
	"github.com/Workiva/frugal/lib/gopherjs/thrift"
)

const ConstI32FromBase = 582

func init() {
}

type BaseHealthCondition int64

const (
	BaseHealthCondition_PASS    BaseHealthCondition = 1
	BaseHealthCondition_WARN    BaseHealthCondition = 2
	BaseHealthCondition_FAIL    BaseHealthCondition = 3
	BaseHealthCondition_UNKNOWN BaseHealthCondition = 4
)

func (p BaseHealthCondition) String() string {
	switch p {
	case BaseHealthCondition_PASS:
		return "PASS"
	case BaseHealthCondition_WARN:
		return "WARN"
	case BaseHealthCondition_FAIL:
		return "FAIL"
	case BaseHealthCondition_UNKNOWN:
		return "UNKNOWN"
	}
	return "<UNSET>"
}

func BaseHealthConditionFromString(s string) (BaseHealthCondition, error) {
	switch s {
	case "PASS":
		return BaseHealthCondition_PASS, nil
	case "WARN":
		return BaseHealthCondition_WARN, nil
	case "FAIL":
		return BaseHealthCondition_FAIL, nil
	case "UNKNOWN":
		return BaseHealthCondition_UNKNOWN, nil
	}
	return BaseHealthCondition(0), fmt.Errorf("not a valid BaseHealthCondition string")
}

type Thing struct {
	AnID    int32
	AString string
}

func NewThing() *Thing {
	return &Thing{}
}

func (p *Thing) GetAnID() int32 {
	return p.AnID
}

func (p *Thing) GetAString() string {
	return p.AString
}

func (p *Thing) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if v, err := iprot.ReadI32(); err != nil {
				return thrift.PrependError("error reading field 1: ", err)
			} else {
				p.AnID = v
			}
		case 2:
			if v, err := iprot.ReadString(); err != nil {
				return thrift.PrependError("error reading field 2: ", err)
			} else {
				p.AString = v
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *Thing) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("thing"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := frugal.WriteI32(oprot, p.AnID, "an_id", 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T::an_id:1 ", p), err)
	}
	if err := frugal.WriteString(oprot, p.AString, "a_string", 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T::a_string:2 ", p), err)
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *Thing) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Thing(%+v)", *p)
}

type NestedThing struct {
	Things []*Thing
}

func NewNestedThing() *NestedThing {
	return &NestedThing{}
}

func (p *NestedThing) GetThings() []*Thing {
	return p.Things
}

func (p *NestedThing) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			_, size, err := iprot.ReadListBegin()
			if err != nil {
				return thrift.PrependError("error reading list begin: ", err)
			}
			p.Things = make([]*Thing, 0, size)
			for i := 0; i < size; i++ {
				elem24 := NewThing()
				if err := elem24.Read(iprot); err != nil {
					return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", elem24), err)
				}
				p.Things = append(p.Things, elem24)
			}
			if err := iprot.ReadListEnd(); err != nil {
				return thrift.PrependError("error reading list end: ", err)
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *NestedThing) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("nested_thing"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *NestedThing) writeField1(oprot thrift.TProtocol) error {
	if err := oprot.WriteFieldBegin("things", thrift.LIST, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:things: ", p), err)
	}
	if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Things)); err != nil {
		return thrift.PrependError("error writing list begin: ", err)
	}
	for _, v := range p.Things {
		if err := v.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return thrift.PrependError("error writing list end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:things: ", p), err)
	}
	return nil
}

func (p *NestedThing) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("NestedThing(%+v)", *p)
}

type APIException struct {
}

func NewAPIException() *APIException {
	return &APIException{}
}

func (p *APIException) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if err := iprot.Skip(fieldTypeId); err != nil {
			return err
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *APIException) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("api_exception"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *APIException) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("APIException(%+v)", *p)
}

func (p *APIException) Error() string {
	return p.String()
}