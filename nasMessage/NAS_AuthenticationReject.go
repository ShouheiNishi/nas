// Code generated by generate.sh, DO NOT EDIT.

package nasMessage

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/free5gc/nas/nasType"
)

type AuthenticationReject struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.AuthenticationRejectMessageIdentity
	*nasType.EAPMessage
}

func NewAuthenticationReject(iei uint8) (authenticationReject *AuthenticationReject) {
	authenticationReject = &AuthenticationReject{}
	return authenticationReject
}

const (
	AuthenticationRejectEAPMessageType uint8 = 0x78
)

func (a *AuthenticationReject) EncodeAuthenticationReject(buffer *bytes.Buffer) error {
	if err := binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return fmt.Errorf("NAS encode error (AuthenticationReject/ExtendedProtocolDiscriminator): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return fmt.Errorf("NAS encode error (AuthenticationReject/SpareHalfOctetAndSecurityHeaderType): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.AuthenticationRejectMessageIdentity.Octet); err != nil {
		return fmt.Errorf("NAS encode error (AuthenticationReject/AuthenticationRejectMessageIdentity): %w", err)
	}
	if a.EAPMessage != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.EAPMessage.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (AuthenticationReject/EAPMessage): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.EAPMessage.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (AuthenticationReject/EAPMessage): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.EAPMessage.Buffer); err != nil {
			return fmt.Errorf("NAS encode error (AuthenticationReject/EAPMessage): %w", err)
		}
	}
	return nil
}

func (a *AuthenticationReject) DecodeAuthenticationReject(byteArray *[]byte) error {
	buffer := bytes.NewBuffer(*byteArray)
	if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return fmt.Errorf("NAS decode error (AuthenticationReject/ExtendedProtocolDiscriminator): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return fmt.Errorf("NAS decode error (AuthenticationReject/SpareHalfOctetAndSecurityHeaderType): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.AuthenticationRejectMessageIdentity.Octet); err != nil {
		return fmt.Errorf("NAS decode error (AuthenticationReject/AuthenticationRejectMessageIdentity): %w", err)
	}
	for buffer.Len() > 0 {
		var ieiN uint8
		var tmpIeiN uint8
		if err := binary.Read(buffer, binary.BigEndian, &ieiN); err != nil {
			return fmt.Errorf("NAS decode error (AuthenticationReject/iei): %w", err)
		}
		// fmt.Println(ieiN)
		if ieiN >= 0x80 {
			tmpIeiN = (ieiN & 0xf0) >> 4
		} else {
			tmpIeiN = ieiN
		}
		// fmt.Println("type", tmpIeiN)
		switch tmpIeiN {
		case AuthenticationRejectEAPMessageType:
			a.EAPMessage = nasType.NewEAPMessage(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.EAPMessage.Len); err != nil {
				return fmt.Errorf("NAS decode error (AuthenticationReject/EAPMessage): %w", err)
			}
			if a.EAPMessage.Len < 4 || a.EAPMessage.Len > 1500 {
				return fmt.Errorf("invalid ie length (AuthenticationReject/EAPMessage): %d", a.EAPMessage.Len)
			}
			a.EAPMessage.SetLen(a.EAPMessage.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.EAPMessage.Buffer[:a.EAPMessage.GetLen()]); err != nil {
				return fmt.Errorf("NAS decode error (AuthenticationReject/EAPMessage): %w", err)
			}
		default:
		}
	}
	return nil
}
