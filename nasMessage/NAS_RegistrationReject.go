// Code generated by generate.sh, DO NOT EDIT.

package nasMessage

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/free5gc/nas/nasType"
)

type RegistrationReject struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.RegistrationRejectMessageIdentity
	nasType.Cause5GMM
	*nasType.T3346Value
	*nasType.T3502Value
	*nasType.EAPMessage
}

func NewRegistrationReject(iei uint8) (registrationReject *RegistrationReject) {
	registrationReject = &RegistrationReject{}
	return registrationReject
}

const (
	RegistrationRejectT3346ValueType uint8 = 0x5F
	RegistrationRejectT3502ValueType uint8 = 0x16
	RegistrationRejectEAPMessageType uint8 = 0x78
)

func (a *RegistrationReject) EncodeRegistrationReject(buffer *bytes.Buffer) error {
	if err := binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return fmt.Errorf("NAS encode error (RegistrationReject/ExtendedProtocolDiscriminator): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return fmt.Errorf("NAS encode error (RegistrationReject/SpareHalfOctetAndSecurityHeaderType): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.RegistrationRejectMessageIdentity.Octet); err != nil {
		return fmt.Errorf("NAS encode error (RegistrationReject/RegistrationRejectMessageIdentity): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.Cause5GMM.Octet); err != nil {
		return fmt.Errorf("NAS encode error (RegistrationReject/Cause5GMM): %w", err)
	}
	if a.T3346Value != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.T3346Value.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationReject/T3346Value): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.T3346Value.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationReject/T3346Value): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.T3346Value.Octet); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationReject/T3346Value): %w", err)
		}
	}
	if a.T3502Value != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.T3502Value.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationReject/T3502Value): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.T3502Value.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationReject/T3502Value): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.T3502Value.Octet); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationReject/T3502Value): %w", err)
		}
	}
	if a.EAPMessage != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.EAPMessage.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationReject/EAPMessage): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.EAPMessage.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationReject/EAPMessage): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.EAPMessage.Buffer); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationReject/EAPMessage): %w", err)
		}
	}
	return nil
}

func (a *RegistrationReject) DecodeRegistrationReject(byteArray *[]byte) error {
	buffer := bytes.NewBuffer(*byteArray)
	if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return fmt.Errorf("NAS decode error (RegistrationReject/ExtendedProtocolDiscriminator): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return fmt.Errorf("NAS decode error (RegistrationReject/SpareHalfOctetAndSecurityHeaderType): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.RegistrationRejectMessageIdentity.Octet); err != nil {
		return fmt.Errorf("NAS decode error (RegistrationReject/RegistrationRejectMessageIdentity): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.Cause5GMM.Octet); err != nil {
		return fmt.Errorf("NAS decode error (RegistrationReject/Cause5GMM): %w", err)
	}
	for buffer.Len() > 0 {
		var ieiN uint8
		var tmpIeiN uint8
		if err := binary.Read(buffer, binary.BigEndian, &ieiN); err != nil {
			return fmt.Errorf("NAS decode error (RegistrationReject/iei): %w", err)
		}
		// fmt.Println(ieiN)
		if ieiN >= 0x80 {
			tmpIeiN = (ieiN & 0xf0) >> 4
		} else {
			tmpIeiN = ieiN
		}
		// fmt.Println("type", tmpIeiN)
		switch tmpIeiN {
		case RegistrationRejectT3346ValueType:
			a.T3346Value = nasType.NewT3346Value(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.T3346Value.Len); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationReject/T3346Value): %w", err)
			}
			a.T3346Value.SetLen(a.T3346Value.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, &a.T3346Value.Octet); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationReject/T3346Value): %w", err)
			}
		case RegistrationRejectT3502ValueType:
			a.T3502Value = nasType.NewT3502Value(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.T3502Value.Len); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationReject/T3502Value): %w", err)
			}
			a.T3502Value.SetLen(a.T3502Value.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, &a.T3502Value.Octet); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationReject/T3502Value): %w", err)
			}
		case RegistrationRejectEAPMessageType:
			a.EAPMessage = nasType.NewEAPMessage(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.EAPMessage.Len); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationReject/EAPMessage): %w", err)
			}
			a.EAPMessage.SetLen(a.EAPMessage.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.EAPMessage.Buffer[:a.EAPMessage.GetLen()]); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationReject/EAPMessage): %w", err)
			}
		default:
		}
	}
	return nil
}
