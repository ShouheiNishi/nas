// Code generated by generate.sh, DO NOT EDIT.

package nasMessage

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/free5gc/nas/nasType"
)

type PDUSessionModificationReject struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.PDUSessionID
	nasType.PTI
	nasType.PDUSESSIONMODIFICATIONREJECTMessageIdentity
	nasType.Cause5GSM
	*nasType.BackoffTimerValue
	*nasType.ExtendedProtocolConfigurationOptions
}

func NewPDUSessionModificationReject(iei uint8) (pDUSessionModificationReject *PDUSessionModificationReject) {
	pDUSessionModificationReject = &PDUSessionModificationReject{}
	return pDUSessionModificationReject
}

const (
	PDUSessionModificationRejectBackoffTimerValueType                    uint8 = 0x37
	PDUSessionModificationRejectExtendedProtocolConfigurationOptionsType uint8 = 0x7B
)

func (a *PDUSessionModificationReject) EncodePDUSessionModificationReject(buffer *bytes.Buffer) error {
	if err := binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return fmt.Errorf("NAS encode error (PDUSessionModificationReject/ExtendedProtocolDiscriminator): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.PDUSessionID.Octet); err != nil {
		return fmt.Errorf("NAS encode error (PDUSessionModificationReject/PDUSessionID): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.PTI.Octet); err != nil {
		return fmt.Errorf("NAS encode error (PDUSessionModificationReject/PTI): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.PDUSESSIONMODIFICATIONREJECTMessageIdentity.Octet); err != nil {
		return fmt.Errorf("NAS encode error (PDUSessionModificationReject/PDUSESSIONMODIFICATIONREJECTMessageIdentity): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.Cause5GSM.Octet); err != nil {
		return fmt.Errorf("NAS encode error (PDUSessionModificationReject/Cause5GSM): %w", err)
	}
	if a.BackoffTimerValue != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.BackoffTimerValue.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionModificationReject/BackoffTimerValue): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.BackoffTimerValue.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionModificationReject/BackoffTimerValue): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.BackoffTimerValue.Octet); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionModificationReject/BackoffTimerValue): %w", err)
		}
	}
	if a.ExtendedProtocolConfigurationOptions != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.ExtendedProtocolConfigurationOptions.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionModificationReject/ExtendedProtocolConfigurationOptions): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.ExtendedProtocolConfigurationOptions.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionModificationReject/ExtendedProtocolConfigurationOptions): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolConfigurationOptions.Buffer); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionModificationReject/ExtendedProtocolConfigurationOptions): %w", err)
		}
	}
	return nil
}

func (a *PDUSessionModificationReject) DecodePDUSessionModificationReject(byteArray *[]byte) error {
	buffer := bytes.NewBuffer(*byteArray)
	if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return fmt.Errorf("NAS decode error (PDUSessionModificationReject/ExtendedProtocolDiscriminator): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.PDUSessionID.Octet); err != nil {
		return fmt.Errorf("NAS decode error (PDUSessionModificationReject/PDUSessionID): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.PTI.Octet); err != nil {
		return fmt.Errorf("NAS decode error (PDUSessionModificationReject/PTI): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.PDUSESSIONMODIFICATIONREJECTMessageIdentity.Octet); err != nil {
		return fmt.Errorf("NAS decode error (PDUSessionModificationReject/PDUSESSIONMODIFICATIONREJECTMessageIdentity): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.Cause5GSM.Octet); err != nil {
		return fmt.Errorf("NAS decode error (PDUSessionModificationReject/Cause5GSM): %w", err)
	}
	for buffer.Len() > 0 {
		var ieiN uint8
		var tmpIeiN uint8
		if err := binary.Read(buffer, binary.BigEndian, &ieiN); err != nil {
			return fmt.Errorf("NAS decode error (PDUSessionModificationReject/iei): %w", err)
		}
		// fmt.Println(ieiN)
		if ieiN >= 0x80 {
			tmpIeiN = (ieiN & 0xf0) >> 4
		} else {
			tmpIeiN = ieiN
		}
		// fmt.Println("type", tmpIeiN)
		switch tmpIeiN {
		case PDUSessionModificationRejectBackoffTimerValueType:
			a.BackoffTimerValue = nasType.NewBackoffTimerValue(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.BackoffTimerValue.Len); err != nil {
				return fmt.Errorf("NAS decode error (PDUSessionModificationReject/BackoffTimerValue): %w", err)
			}
			if a.BackoffTimerValue.Len != 1 {
				return fmt.Errorf("invalid ie length (PDUSessionModificationReject/BackoffTimerValue): %d", a.BackoffTimerValue.Len)
			}
			a.BackoffTimerValue.SetLen(a.BackoffTimerValue.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, &a.BackoffTimerValue.Octet); err != nil {
				return fmt.Errorf("NAS decode error (PDUSessionModificationReject/BackoffTimerValue): %w", err)
			}
		case PDUSessionModificationRejectExtendedProtocolConfigurationOptionsType:
			a.ExtendedProtocolConfigurationOptions = nasType.NewExtendedProtocolConfigurationOptions(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolConfigurationOptions.Len); err != nil {
				return fmt.Errorf("NAS decode error (PDUSessionModificationReject/ExtendedProtocolConfigurationOptions): %w", err)
			}
			if a.ExtendedProtocolConfigurationOptions.Len < 1 {
				return fmt.Errorf("invalid ie length (PDUSessionModificationReject/ExtendedProtocolConfigurationOptions): %d", a.ExtendedProtocolConfigurationOptions.Len)
			}
			a.ExtendedProtocolConfigurationOptions.SetLen(a.ExtendedProtocolConfigurationOptions.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.ExtendedProtocolConfigurationOptions.Buffer[:a.ExtendedProtocolConfigurationOptions.GetLen()]); err != nil {
				return fmt.Errorf("NAS decode error (PDUSessionModificationReject/ExtendedProtocolConfigurationOptions): %w", err)
			}
		default:
		}
	}
	return nil
}
