// Code generated by generate.sh, DO NOT EDIT.

//go:build ignore
// +build ignore

package main

import (
	"reflect"

	"github.com/free5gc/nas/nasType"
)

var nasTypeTable map[string]reflect.Type = map[string]reflect.Type{
	"ABBA":                                               reflect.TypeOf(nasType.ABBA{}),
	"Additional5GSecurityInformation":                    reflect.TypeOf(nasType.Additional5GSecurityInformation{}),
	"AdditionalGUTI":                                     reflect.TypeOf(nasType.AdditionalGUTI{}),
	"AdditionalInformation":                              reflect.TypeOf(nasType.AdditionalInformation{}),
	"AllowedNSSAI":                                       reflect.TypeOf(nasType.AllowedNSSAI{}),
	"AllowedPDUSessionStatus":                            reflect.TypeOf(nasType.AllowedPDUSessionStatus{}),
	"AllowedSSCMode":                                     reflect.TypeOf(nasType.AllowedSSCMode{}),
	"AlwaysonPDUSessionIndication":                       reflect.TypeOf(nasType.AlwaysonPDUSessionIndication{}),
	"AlwaysonPDUSessionRequested":                        reflect.TypeOf(nasType.AlwaysonPDUSessionRequested{}),
	"AuthenticationFailureMessageIdentity":               reflect.TypeOf(nasType.AuthenticationFailureMessageIdentity{}),
	"AuthenticationFailureParameter":                     reflect.TypeOf(nasType.AuthenticationFailureParameter{}),
	"AuthenticationParameterAUTN":                        reflect.TypeOf(nasType.AuthenticationParameterAUTN{}),
	"AuthenticationParameterRAND":                        reflect.TypeOf(nasType.AuthenticationParameterRAND{}),
	"AuthenticationRejectMessageIdentity":                reflect.TypeOf(nasType.AuthenticationRejectMessageIdentity{}),
	"AuthenticationRequestMessageIdentity":               reflect.TypeOf(nasType.AuthenticationRequestMessageIdentity{}),
	"AuthenticationResponseMessageIdentity":              reflect.TypeOf(nasType.AuthenticationResponseMessageIdentity{}),
	"AuthenticationResponseParameter":                    reflect.TypeOf(nasType.AuthenticationResponseParameter{}),
	"AuthenticationResultMessageIdentity":                reflect.TypeOf(nasType.AuthenticationResultMessageIdentity{}),
	"AuthorizedQosFlowDescriptions":                      reflect.TypeOf(nasType.AuthorizedQosFlowDescriptions{}),
	"AuthorizedQosRules":                                 reflect.TypeOf(nasType.AuthorizedQosRules{}),
	"BackoffTimerValue":                                  reflect.TypeOf(nasType.BackoffTimerValue{}),
	"Capability5GMM":                                     reflect.TypeOf(nasType.Capability5GMM{}),
	"Capability5GSM":                                     reflect.TypeOf(nasType.Capability5GSM{}),
	"Cause5GMM":                                          reflect.TypeOf(nasType.Cause5GMM{}),
	"Cause5GSM":                                          reflect.TypeOf(nasType.Cause5GSM{}),
	"ConfigurationUpdateCommandMessageIdentity":          reflect.TypeOf(nasType.ConfigurationUpdateCommandMessageIdentity{}),
	"ConfigurationUpdateCompleteMessageIdentity":         reflect.TypeOf(nasType.ConfigurationUpdateCompleteMessageIdentity{}),
	"ConfigurationUpdateIndication":                      reflect.TypeOf(nasType.ConfigurationUpdateIndication{}),
	"ConfiguredNSSAI":                                    reflect.TypeOf(nasType.ConfiguredNSSAI{}),
	"DLNASTRANSPORTMessageIdentity":                      reflect.TypeOf(nasType.DLNASTRANSPORTMessageIdentity{}),
	"DNN":                                                reflect.TypeOf(nasType.DNN{}),
	"DeregistrationAcceptMessageIdentity":                reflect.TypeOf(nasType.DeregistrationAcceptMessageIdentity{}),
	"DeregistrationRequestMessageIdentity":               reflect.TypeOf(nasType.DeregistrationRequestMessageIdentity{}),
	"EAPMessage":                                         reflect.TypeOf(nasType.EAPMessage{}),
	"EPSNASMessageContainer":                             reflect.TypeOf(nasType.EPSNASMessageContainer{}),
	"EmergencyNumberList":                                reflect.TypeOf(nasType.EmergencyNumberList{}),
	"EquivalentPlmns":                                    reflect.TypeOf(nasType.EquivalentPlmns{}),
	"ExtendedEmergencyNumberList":                        reflect.TypeOf(nasType.ExtendedEmergencyNumberList{}),
	"ExtendedProtocolConfigurationOptions":               reflect.TypeOf(nasType.ExtendedProtocolConfigurationOptions{}),
	"ExtendedProtocolDiscriminator":                      reflect.TypeOf(nasType.ExtendedProtocolDiscriminator{}),
	"FullNameForNetwork":                                 reflect.TypeOf(nasType.FullNameForNetwork{}),
	"GUTI5G":                                             reflect.TypeOf(nasType.GUTI5G{}),
	"IMEISV":                                             reflect.TypeOf(nasType.IMEISV{}),
	"IMEISVRequest":                                      reflect.TypeOf(nasType.IMEISVRequest{}),
	"IdentityRequestMessageIdentity":                     reflect.TypeOf(nasType.IdentityRequestMessageIdentity{}),
	"IdentityResponseMessageIdentity":                    reflect.TypeOf(nasType.IdentityResponseMessageIdentity{}),
	"IntegrityProtectionMaximumDataRate":                 reflect.TypeOf(nasType.IntegrityProtectionMaximumDataRate{}),
	"LADNIndication":                                     reflect.TypeOf(nasType.LADNIndication{}),
	"LADNInformation":                                    reflect.TypeOf(nasType.LADNInformation{}),
	"LastVisitedRegisteredTAI":                           reflect.TypeOf(nasType.LastVisitedRegisteredTAI{}),
	"LocalTimeZone":                                      reflect.TypeOf(nasType.LocalTimeZone{}),
	"MICOIndication":                                     reflect.TypeOf(nasType.MICOIndication{}),
	"MappedEPSBearerContexts":                            reflect.TypeOf(nasType.MappedEPSBearerContexts{}),
	"MaximumNumberOfSupportedPacketFilters":              reflect.TypeOf(nasType.MaximumNumberOfSupportedPacketFilters{}),
	"MessageAuthenticationCode":                          reflect.TypeOf(nasType.MessageAuthenticationCode{}),
	"MobileIdentity":                                     reflect.TypeOf(nasType.MobileIdentity{}),
	"MobileIdentity5GS":                                  reflect.TypeOf(nasType.MobileIdentity5GS{}),
	"NASMessageContainer":                                reflect.TypeOf(nasType.NASMessageContainer{}),
	"NSSAIInclusionMode":                                 reflect.TypeOf(nasType.NSSAIInclusionMode{}),
	"NegotiatedDRXParameters":                            reflect.TypeOf(nasType.NegotiatedDRXParameters{}),
	"NetworkDaylightSavingTime":                          reflect.TypeOf(nasType.NetworkDaylightSavingTime{}),
	"NetworkFeatureSupport5GS":                           reflect.TypeOf(nasType.NetworkFeatureSupport5GS{}),
	"NetworkSlicingIndication":                           reflect.TypeOf(nasType.NetworkSlicingIndication{}),
	"NgksiAndDeregistrationType":                         reflect.TypeOf(nasType.NgksiAndDeregistrationType{}),
	"NgksiAndRegistrationType5GS":                        reflect.TypeOf(nasType.NgksiAndRegistrationType5GS{}),
	"Non3GppDeregistrationTimerValue":                    reflect.TypeOf(nasType.Non3GppDeregistrationTimerValue{}),
	"NoncurrentNativeNASKeySetIdentifier":                reflect.TypeOf(nasType.NoncurrentNativeNASKeySetIdentifier{}),
	"NotificationMessageIdentity":                        reflect.TypeOf(nasType.NotificationMessageIdentity{}),
	"NotificationResponseMessageIdentity":                reflect.TypeOf(nasType.NotificationResponseMessageIdentity{}),
	"OldPDUSessionID":                                    reflect.TypeOf(nasType.OldPDUSessionID{}),
	"OperatordefinedAccessCategoryDefinitions":           reflect.TypeOf(nasType.OperatordefinedAccessCategoryDefinitions{}),
	"PDUAddress":                                         reflect.TypeOf(nasType.PDUAddress{}),
	"PDUSESSIONAUTHENTICATIONCOMMANDMessageIdentity":     reflect.TypeOf(nasType.PDUSESSIONAUTHENTICATIONCOMMANDMessageIdentity{}),
	"PDUSESSIONAUTHENTICATIONCOMPLETEMessageIdentity":    reflect.TypeOf(nasType.PDUSESSIONAUTHENTICATIONCOMPLETEMessageIdentity{}),
	"PDUSESSIONAUTHENTICATIONRESULTMessageIdentity":      reflect.TypeOf(nasType.PDUSESSIONAUTHENTICATIONRESULTMessageIdentity{}),
	"PDUSESSIONESTABLISHMENTACCEPTMessageIdentity":       reflect.TypeOf(nasType.PDUSESSIONESTABLISHMENTACCEPTMessageIdentity{}),
	"PDUSESSIONESTABLISHMENTREJECTMessageIdentity":       reflect.TypeOf(nasType.PDUSESSIONESTABLISHMENTREJECTMessageIdentity{}),
	"PDUSESSIONESTABLISHMENTREQUESTMessageIdentity":      reflect.TypeOf(nasType.PDUSESSIONESTABLISHMENTREQUESTMessageIdentity{}),
	"PDUSESSIONMODIFICATIONCOMMANDMessageIdentity":       reflect.TypeOf(nasType.PDUSESSIONMODIFICATIONCOMMANDMessageIdentity{}),
	"PDUSESSIONMODIFICATIONCOMMANDREJECTMessageIdentity": reflect.TypeOf(nasType.PDUSESSIONMODIFICATIONCOMMANDREJECTMessageIdentity{}),
	"PDUSESSIONMODIFICATIONCOMPLETEMessageIdentity":      reflect.TypeOf(nasType.PDUSESSIONMODIFICATIONCOMPLETEMessageIdentity{}),
	"PDUSESSIONMODIFICATIONREJECTMessageIdentity":        reflect.TypeOf(nasType.PDUSESSIONMODIFICATIONREJECTMessageIdentity{}),
	"PDUSESSIONMODIFICATIONREQUESTMessageIdentity":       reflect.TypeOf(nasType.PDUSESSIONMODIFICATIONREQUESTMessageIdentity{}),
	"PDUSESSIONRELEASECOMMANDMessageIdentity":            reflect.TypeOf(nasType.PDUSESSIONRELEASECOMMANDMessageIdentity{}),
	"PDUSESSIONRELEASECOMPLETEMessageIdentity":           reflect.TypeOf(nasType.PDUSESSIONRELEASECOMPLETEMessageIdentity{}),
	"PDUSESSIONRELEASEREJECTMessageIdentity":             reflect.TypeOf(nasType.PDUSESSIONRELEASEREJECTMessageIdentity{}),
	"PDUSESSIONRELEASEREQUESTMessageIdentity":            reflect.TypeOf(nasType.PDUSESSIONRELEASEREQUESTMessageIdentity{}),
	"PDUSessionID":                                       reflect.TypeOf(nasType.PDUSessionID{}),
	"PDUSessionReactivationResult":                       reflect.TypeOf(nasType.PDUSessionReactivationResult{}),
	"PDUSessionReactivationResultErrorCause":             reflect.TypeOf(nasType.PDUSessionReactivationResultErrorCause{}),
	"PDUSessionStatus":                                   reflect.TypeOf(nasType.PDUSessionStatus{}),
	"PDUSessionType":                                     reflect.TypeOf(nasType.PDUSessionType{}),
	"PTI":                                                reflect.TypeOf(nasType.PTI{}),
	"PayloadContainer":                                   reflect.TypeOf(nasType.PayloadContainer{}),
	"PduSessionID2Value":                                 reflect.TypeOf(nasType.PduSessionID2Value{}),
	"Plain5GSNASMessage":                                 reflect.TypeOf(nasType.Plain5GSNASMessage{}),
	"RQTimerValue":                                       reflect.TypeOf(nasType.RQTimerValue{}),
	"RegistrationAcceptMessageIdentity":                  reflect.TypeOf(nasType.RegistrationAcceptMessageIdentity{}),
	"RegistrationCompleteMessageIdentity":                reflect.TypeOf(nasType.RegistrationCompleteMessageIdentity{}),
	"RegistrationRejectMessageIdentity":                  reflect.TypeOf(nasType.RegistrationRejectMessageIdentity{}),
	"RegistrationRequestMessageIdentity":                 reflect.TypeOf(nasType.RegistrationRequestMessageIdentity{}),
	"RegistrationResult5GS":                              reflect.TypeOf(nasType.RegistrationResult5GS{}),
	"RejectedNSSAI":                                      reflect.TypeOf(nasType.RejectedNSSAI{}),
	"ReplayedS1UESecurityCapabilities":                   reflect.TypeOf(nasType.ReplayedS1UESecurityCapabilities{}),
	"ReplayedUESecurityCapabilities":                     reflect.TypeOf(nasType.ReplayedUESecurityCapabilities{}),
	"RequestType":                                        reflect.TypeOf(nasType.RequestType{}),
	"RequestedDRXParameters":                             reflect.TypeOf(nasType.RequestedDRXParameters{}),
	"RequestedNSSAI":                                     reflect.TypeOf(nasType.RequestedNSSAI{}),
	"RequestedQosFlowDescriptions":                       reflect.TypeOf(nasType.RequestedQosFlowDescriptions{}),
	"RequestedQosRules":                                  reflect.TypeOf(nasType.RequestedQosRules{}),
	"S1UENetworkCapability":                              reflect.TypeOf(nasType.S1UENetworkCapability{}),
	"SMPDUDNRequestContainer":                            reflect.TypeOf(nasType.SMPDUDNRequestContainer{}),
	"SMSIndication":                                      reflect.TypeOf(nasType.SMSIndication{}),
	"SNSSAI":                                             reflect.TypeOf(nasType.SNSSAI{}),
	"SORTransparentContainer":                            reflect.TypeOf(nasType.SORTransparentContainer{}),
	"SSCMode":                                            reflect.TypeOf(nasType.SSCMode{}),
	"STATUSMessageIdentity5GMM":                          reflect.TypeOf(nasType.STATUSMessageIdentity5GMM{}),
	"STATUSMessageIdentity5GSM":                          reflect.TypeOf(nasType.STATUSMessageIdentity5GSM{}),
	"SecurityModeCommandMessageIdentity":                 reflect.TypeOf(nasType.SecurityModeCommandMessageIdentity{}),
	"SecurityModeCompleteMessageIdentity":                reflect.TypeOf(nasType.SecurityModeCompleteMessageIdentity{}),
	"SecurityModeRejectMessageIdentity":                  reflect.TypeOf(nasType.SecurityModeRejectMessageIdentity{}),
	"SelectedEPSNASSecurityAlgorithms":                   reflect.TypeOf(nasType.SelectedEPSNASSecurityAlgorithms{}),
	"SelectedNASSecurityAlgorithms":                      reflect.TypeOf(nasType.SelectedNASSecurityAlgorithms{}),
	"SelectedSSCModeAndSelectedPDUSessionType":           reflect.TypeOf(nasType.SelectedSSCModeAndSelectedPDUSessionType{}),
	"SequenceNumber":                                     reflect.TypeOf(nasType.SequenceNumber{}),
	"ServiceAcceptMessageIdentity":                       reflect.TypeOf(nasType.ServiceAcceptMessageIdentity{}),
	"ServiceAreaList":                                    reflect.TypeOf(nasType.ServiceAreaList{}),
	"ServiceRejectMessageIdentity":                       reflect.TypeOf(nasType.ServiceRejectMessageIdentity{}),
	"ServiceRequestMessageIdentity":                      reflect.TypeOf(nasType.ServiceRequestMessageIdentity{}),
	"ServiceTypeAndNgksi":                                reflect.TypeOf(nasType.ServiceTypeAndNgksi{}),
	"SessionAMBR":                                        reflect.TypeOf(nasType.SessionAMBR{}),
	"ShortNameForNetwork":                                reflect.TypeOf(nasType.ShortNameForNetwork{}),
	"SpareHalfOctetAndAccessType":                        reflect.TypeOf(nasType.SpareHalfOctetAndAccessType{}),
	"SpareHalfOctetAndDeregistrationType":                reflect.TypeOf(nasType.SpareHalfOctetAndDeregistrationType{}),
	"SpareHalfOctetAndIdentityType":                      reflect.TypeOf(nasType.SpareHalfOctetAndIdentityType{}),
	"SpareHalfOctetAndNgksi":                             reflect.TypeOf(nasType.SpareHalfOctetAndNgksi{}),
	"SpareHalfOctetAndPayloadContainerType":              reflect.TypeOf(nasType.SpareHalfOctetAndPayloadContainerType{}),
	"SpareHalfOctetAndSecurityHeaderType":                reflect.TypeOf(nasType.SpareHalfOctetAndSecurityHeaderType{}),
	"T3346Value":                                         reflect.TypeOf(nasType.T3346Value{}),
	"T3502Value":                                         reflect.TypeOf(nasType.T3502Value{}),
	"T3512Value":                                         reflect.TypeOf(nasType.T3512Value{}),
	"TAIList":                                            reflect.TypeOf(nasType.TAIList{}),
	"TMSI5GS":                                            reflect.TypeOf(nasType.TMSI5GS{}),
	"UESecurityCapability":                               reflect.TypeOf(nasType.UESecurityCapability{}),
	"UEStatus":                                           reflect.TypeOf(nasType.UEStatus{}),
	"ULNASTRANSPORTMessageIdentity":                      reflect.TypeOf(nasType.ULNASTRANSPORTMessageIdentity{}),
	"UesUsageSetting":                                    reflect.TypeOf(nasType.UesUsageSetting{}),
	"UniversalTimeAndLocalTimeZone":                      reflect.TypeOf(nasType.UniversalTimeAndLocalTimeZone{}),
	"UpdateType5GS":                                      reflect.TypeOf(nasType.UpdateType5GS{}),
	"UplinkDataStatus":                                   reflect.TypeOf(nasType.UplinkDataStatus{}),
}
