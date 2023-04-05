// Code generated by generate.sh, DO NOT EDIT.

package nasMessage

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/free5gc/nas/nasType"
)

type DeregistrationAcceptUEOriginatingDeregistration struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.DeregistrationAcceptMessageIdentity
}

func NewDeregistrationAcceptUEOriginatingDeregistration(iei uint8) (deregistrationAcceptUEOriginatingDeregistration *DeregistrationAcceptUEOriginatingDeregistration) {
	deregistrationAcceptUEOriginatingDeregistration = &DeregistrationAcceptUEOriginatingDeregistration{}
	return deregistrationAcceptUEOriginatingDeregistration
}

func (a *DeregistrationAcceptUEOriginatingDeregistration) EncodeDeregistrationAcceptUEOriginatingDeregistration(buffer *bytes.Buffer) error {
	if err := binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return fmt.Errorf("NAS encode error (DeregistrationAcceptUEOriginatingDeregistration/ExtendedProtocolDiscriminator): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return fmt.Errorf("NAS encode error (DeregistrationAcceptUEOriginatingDeregistration/SpareHalfOctetAndSecurityHeaderType): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.DeregistrationAcceptMessageIdentity.Octet); err != nil {
		return fmt.Errorf("NAS encode error (DeregistrationAcceptUEOriginatingDeregistration/DeregistrationAcceptMessageIdentity): %w", err)
	}
	return nil
}

func (a *DeregistrationAcceptUEOriginatingDeregistration) DecodeDeregistrationAcceptUEOriginatingDeregistration(byteArray *[]byte) error {
	buffer := bytes.NewBuffer(*byteArray)
	if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return fmt.Errorf("NAS decode error (DeregistrationAcceptUEOriginatingDeregistration/ExtendedProtocolDiscriminator): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return fmt.Errorf("NAS decode error (DeregistrationAcceptUEOriginatingDeregistration/SpareHalfOctetAndSecurityHeaderType): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.DeregistrationAcceptMessageIdentity.Octet); err != nil {
		return fmt.Errorf("NAS decode error (DeregistrationAcceptUEOriginatingDeregistration/DeregistrationAcceptMessageIdentity): %w", err)
	}
	for buffer.Len() > 0 {
		var ieiN uint8
		var tmpIeiN uint8
		if err := binary.Read(buffer, binary.BigEndian, &ieiN); err != nil {
			return fmt.Errorf("NAS decode error (DeregistrationAcceptUEOriginatingDeregistration/iei): %w", err)
		}
		// fmt.Println(ieiN)
		if ieiN >= 0x80 {
			tmpIeiN = (ieiN & 0xf0) >> 4
		} else {
			tmpIeiN = ieiN
		}
		// fmt.Println("type", tmpIeiN)
		switch tmpIeiN {
		default:
		}
	}
	return nil
}
