// Code generated by generate.sh, DO NOT EDIT.

package nasMessage

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/free5gc/nas/nasType"
)

type PDUSessionReleaseComplete struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.PDUSessionID
	nasType.PTI
	nasType.PDUSESSIONRELEASECOMPLETEMessageIdentity
	*nasType.Cause5GSM
	*nasType.ExtendedProtocolConfigurationOptions
}

func NewPDUSessionReleaseComplete(iei uint8) (pDUSessionReleaseComplete *PDUSessionReleaseComplete) {
	pDUSessionReleaseComplete = &PDUSessionReleaseComplete{}
	return pDUSessionReleaseComplete
}

const (
	PDUSessionReleaseCompleteCause5GSMType                            uint8 = 0x59
	PDUSessionReleaseCompleteExtendedProtocolConfigurationOptionsType uint8 = 0x7B
)

func (a *PDUSessionReleaseComplete) EncodePDUSessionReleaseComplete(buffer *bytes.Buffer) error {
	if err := binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return fmt.Errorf("NAS encode error (PDUSessionReleaseComplete/ExtendedProtocolDiscriminator): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.PDUSessionID.Octet); err != nil {
		return fmt.Errorf("NAS encode error (PDUSessionReleaseComplete/PDUSessionID): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.PTI.Octet); err != nil {
		return fmt.Errorf("NAS encode error (PDUSessionReleaseComplete/PTI): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.PDUSESSIONRELEASECOMPLETEMessageIdentity.Octet); err != nil {
		return fmt.Errorf("NAS encode error (PDUSessionReleaseComplete/PDUSESSIONRELEASECOMPLETEMessageIdentity): %w", err)
	}
	if a.Cause5GSM != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.Cause5GSM.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionReleaseComplete/Cause5GSM): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.Cause5GSM.Octet); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionReleaseComplete/Cause5GSM): %w", err)
		}
	}
	if a.ExtendedProtocolConfigurationOptions != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.ExtendedProtocolConfigurationOptions.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionReleaseComplete/ExtendedProtocolConfigurationOptions): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.ExtendedProtocolConfigurationOptions.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionReleaseComplete/ExtendedProtocolConfigurationOptions): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolConfigurationOptions.Buffer); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionReleaseComplete/ExtendedProtocolConfigurationOptions): %w", err)
		}
	}
	return nil
}

func (a *PDUSessionReleaseComplete) DecodePDUSessionReleaseComplete(byteArray *[]byte) error {
	buffer := bytes.NewBuffer(*byteArray)
	if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return fmt.Errorf("NAS decode error (PDUSessionReleaseComplete/ExtendedProtocolDiscriminator): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.PDUSessionID.Octet); err != nil {
		return fmt.Errorf("NAS decode error (PDUSessionReleaseComplete/PDUSessionID): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.PTI.Octet); err != nil {
		return fmt.Errorf("NAS decode error (PDUSessionReleaseComplete/PTI): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.PDUSESSIONRELEASECOMPLETEMessageIdentity.Octet); err != nil {
		return fmt.Errorf("NAS decode error (PDUSessionReleaseComplete/PDUSESSIONRELEASECOMPLETEMessageIdentity): %w", err)
	}
	for buffer.Len() > 0 {
		var ieiN uint8
		var tmpIeiN uint8
		if err := binary.Read(buffer, binary.BigEndian, &ieiN); err != nil {
			return fmt.Errorf("NAS decode error (PDUSessionReleaseComplete/iei): %w", err)
		}
		// fmt.Println(ieiN)
		if ieiN >= 0x80 {
			tmpIeiN = (ieiN & 0xf0) >> 4
		} else {
			tmpIeiN = ieiN
		}
		// fmt.Println("type", tmpIeiN)
		switch tmpIeiN {
		case PDUSessionReleaseCompleteCause5GSMType:
			a.Cause5GSM = nasType.NewCause5GSM(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.Cause5GSM.Octet); err != nil {
				return fmt.Errorf("NAS decode error (PDUSessionReleaseComplete/Cause5GSM): %w", err)
			}
		case PDUSessionReleaseCompleteExtendedProtocolConfigurationOptionsType:
			a.ExtendedProtocolConfigurationOptions = nasType.NewExtendedProtocolConfigurationOptions(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolConfigurationOptions.Len); err != nil {
				return fmt.Errorf("NAS decode error (PDUSessionReleaseComplete/ExtendedProtocolConfigurationOptions): %w", err)
			}
			if a.ExtendedProtocolConfigurationOptions.Len < 1 {
				return fmt.Errorf("invalid ie length (PDUSessionReleaseComplete/ExtendedProtocolConfigurationOptions): %d", a.ExtendedProtocolConfigurationOptions.Len)
			}
			a.ExtendedProtocolConfigurationOptions.SetLen(a.ExtendedProtocolConfigurationOptions.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.ExtendedProtocolConfigurationOptions.Buffer[:a.ExtendedProtocolConfigurationOptions.GetLen()]); err != nil {
				return fmt.Errorf("NAS decode error (PDUSessionReleaseComplete/ExtendedProtocolConfigurationOptions): %w", err)
			}
		default:
		}
	}
	return nil
}
