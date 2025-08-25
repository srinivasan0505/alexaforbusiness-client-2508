package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// SkillSummary represents the SkillSummary schema from the OpenAPI specification
type SkillSummary struct {
	Enablementtype interface{} `json:"EnablementType,omitempty"`
	Skillid interface{} `json:"SkillId,omitempty"`
	Skillname interface{} `json:"SkillName,omitempty"`
	Skilltype interface{} `json:"SkillType,omitempty"`
	Supportslinking interface{} `json:"SupportsLinking,omitempty"`
}

// Room represents the Room schema from the OpenAPI specification
type Room struct {
	Description interface{} `json:"Description,omitempty"`
	Profilearn interface{} `json:"ProfileArn,omitempty"`
	Providercalendarid interface{} `json:"ProviderCalendarId,omitempty"`
	Roomarn interface{} `json:"RoomArn,omitempty"`
	Roomname interface{} `json:"RoomName,omitempty"`
}

// GetProfileResponse represents the GetProfileResponse schema from the OpenAPI specification
type GetProfileResponse struct {
	Profile GetProfileResponseProfile `json:"Profile,omitempty"`
}

// Device represents the Device schema from the OpenAPI specification
type Device struct {
	Networkprofileinfo DeviceNetworkProfileInfo `json:"NetworkProfileInfo,omitempty"`
	Macaddress interface{} `json:"MacAddress,omitempty"`
	Softwareversion interface{} `json:"SoftwareVersion,omitempty"`
	Devicearn interface{} `json:"DeviceArn,omitempty"`
	Deviceserialnumber interface{} `json:"DeviceSerialNumber,omitempty"`
	Devicestatusinfo DeviceDeviceStatusInfo `json:"DeviceStatusInfo,omitempty"`
	Roomarn interface{} `json:"RoomArn,omitempty"`
	Devicename interface{} `json:"DeviceName,omitempty"`
	Devicestatus interface{} `json:"DeviceStatus,omitempty"`
	Devicetype interface{} `json:"DeviceType,omitempty"`
}

// IPDialIn represents the IPDialIn schema from the OpenAPI specification
type IPDialIn struct {
	Commsprotocol interface{} `json:"CommsProtocol"`
	Endpoint interface{} `json:"Endpoint"`
}

// DeleteBusinessReportScheduleResponse represents the DeleteBusinessReportScheduleResponse schema from the OpenAPI specification
type DeleteBusinessReportScheduleResponse struct {
}

// CreateProactiveJoin represents the CreateProactiveJoin schema from the OpenAPI specification
type CreateProactiveJoin struct {
	Enabledbymotion bool `json:"EnabledByMotion"`
}

// DeleteGatewayGroupResponse represents the DeleteGatewayGroupResponse schema from the OpenAPI specification
type DeleteGatewayGroupResponse struct {
}

// UpdateContactResponse represents the UpdateContactResponse schema from the OpenAPI specification
type UpdateContactResponse struct {
}

// UpdateGatewayRequest represents the UpdateGatewayRequest schema from the OpenAPI specification
type UpdateGatewayRequest struct {
	Description interface{} `json:"Description,omitempty"`
	Gatewayarn interface{} `json:"GatewayArn"`
	Name interface{} `json:"Name,omitempty"`
	Softwareversion interface{} `json:"SoftwareVersion,omitempty"`
}

// SearchSkillGroupsResponse represents the SearchSkillGroupsResponse schema from the OpenAPI specification
type SearchSkillGroupsResponse struct {
	Skillgroups interface{} `json:"SkillGroups,omitempty"`
	Totalcount interface{} `json:"TotalCount,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// GatewayGroup represents the GatewayGroup schema from the OpenAPI specification
type GatewayGroup struct {
	Name interface{} `json:"Name,omitempty"`
	Arn interface{} `json:"Arn,omitempty"`
	Description interface{} `json:"Description,omitempty"`
}

// Contact represents the Contact schema from the OpenAPI specification
type Contact struct {
	Contactarn interface{} `json:"ContactArn,omitempty"`
	Displayname interface{} `json:"DisplayName,omitempty"`
	Firstname interface{} `json:"FirstName,omitempty"`
	Lastname interface{} `json:"LastName,omitempty"`
	Phonenumber interface{} `json:"PhoneNumber,omitempty"`
	Phonenumbers interface{} `json:"PhoneNumbers,omitempty"`
	Sipaddresses interface{} `json:"SipAddresses,omitempty"`
}

// DeleteSkillGroupResponse represents the DeleteSkillGroupResponse schema from the OpenAPI specification
type DeleteSkillGroupResponse struct {
}

// CreateConferenceProviderRequestPSTNDialIn represents the CreateConferenceProviderRequestPSTNDialIn schema from the OpenAPI specification
type CreateConferenceProviderRequestPSTNDialIn struct {
	Phonenumber interface{} `json:"PhoneNumber"`
	Countrycode interface{} `json:"CountryCode"`
	Oneclickiddelay interface{} `json:"OneClickIdDelay"`
	Oneclickpindelay interface{} `json:"OneClickPinDelay"`
}

// DeviceEvent represents the DeviceEvent schema from the OpenAPI specification
type DeviceEvent struct {
	Timestamp interface{} `json:"Timestamp,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
	Value interface{} `json:"Value,omitempty"`
}

// PutInvitationConfigurationResponse represents the PutInvitationConfigurationResponse schema from the OpenAPI specification
type PutInvitationConfigurationResponse struct {
}

// UpdateSkillGroupResponse represents the UpdateSkillGroupResponse schema from the OpenAPI specification
type UpdateSkillGroupResponse struct {
}

// PutConferencePreferenceResponse represents the PutConferencePreferenceResponse schema from the OpenAPI specification
type PutConferencePreferenceResponse struct {
}

// RegisterAVSDeviceRequest represents the RegisterAVSDeviceRequest schema from the OpenAPI specification
type RegisterAVSDeviceRequest struct {
	Usercode interface{} `json:"UserCode"`
	Amazonid interface{} `json:"AmazonId"`
	Clientid interface{} `json:"ClientId"`
	Deviceserialnumber interface{} `json:"DeviceSerialNumber,omitempty"`
	Productid interface{} `json:"ProductId"`
	Roomarn interface{} `json:"RoomArn,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
}

// ProfileData represents the ProfileData schema from the OpenAPI specification
type ProfileData struct {
	Locale interface{} `json:"Locale,omitempty"`
	Wakeword interface{} `json:"WakeWord,omitempty"`
	Profilename interface{} `json:"ProfileName,omitempty"`
	Temperatureunit interface{} `json:"TemperatureUnit,omitempty"`
	Profilearn interface{} `json:"ProfileArn,omitempty"`
	Timezone interface{} `json:"Timezone,omitempty"`
	Address interface{} `json:"Address,omitempty"`
	Distanceunit interface{} `json:"DistanceUnit,omitempty"`
	Isdefault interface{} `json:"IsDefault,omitempty"`
}

// SearchDevicesRequest represents the SearchDevicesRequest schema from the OpenAPI specification
type SearchDevicesRequest struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Sortcriteria interface{} `json:"SortCriteria,omitempty"`
	Filters interface{} `json:"Filters,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
}

// CreateAddressBookRequest represents the CreateAddressBookRequest schema from the OpenAPI specification
type CreateAddressBookRequest struct {
	Name interface{} `json:"Name"`
	Tags interface{} `json:"Tags,omitempty"`
	Clientrequesttoken interface{} `json:"ClientRequestToken,omitempty"`
	Description interface{} `json:"Description,omitempty"`
}

// DeviceNetworkProfileInfo represents the DeviceNetworkProfileInfo schema from the OpenAPI specification
type DeviceNetworkProfileInfo struct {
	Certificatearn interface{} `json:"CertificateArn,omitempty"`
	Certificateexpirationtime interface{} `json:"CertificateExpirationTime,omitempty"`
	Networkprofilearn interface{} `json:"NetworkProfileArn,omitempty"`
}

// DeleteConferenceProviderRequest represents the DeleteConferenceProviderRequest schema from the OpenAPI specification
type DeleteConferenceProviderRequest struct {
	Conferenceproviderarn interface{} `json:"ConferenceProviderArn"`
}

// SmartHomeAppliance represents the SmartHomeAppliance schema from the OpenAPI specification
type SmartHomeAppliance struct {
	Description interface{} `json:"Description,omitempty"`
	Friendlyname interface{} `json:"FriendlyName,omitempty"`
	Manufacturername interface{} `json:"ManufacturerName,omitempty"`
}

// PutInvitationConfigurationRequest represents the PutInvitationConfigurationRequest schema from the OpenAPI specification
type PutInvitationConfigurationRequest struct {
	Organizationname interface{} `json:"OrganizationName"`
	Privateskillids interface{} `json:"PrivateSkillIds,omitempty"`
	Contactemail interface{} `json:"ContactEmail,omitempty"`
}

// DisassociateDeviceFromRoomResponse represents the DisassociateDeviceFromRoomResponse schema from the OpenAPI specification
type DisassociateDeviceFromRoomResponse struct {
}

// DisassociateSkillFromSkillGroupResponse represents the DisassociateSkillFromSkillGroupResponse schema from the OpenAPI specification
type DisassociateSkillFromSkillGroupResponse struct {
}

// DeleteProfileResponse represents the DeleteProfileResponse schema from the OpenAPI specification
type DeleteProfileResponse struct {
}

// CreateBusinessReportScheduleResponse represents the CreateBusinessReportScheduleResponse schema from the OpenAPI specification
type CreateBusinessReportScheduleResponse struct {
	Schedulearn interface{} `json:"ScheduleArn,omitempty"`
}

// Category represents the Category schema from the OpenAPI specification
type Category struct {
	Categoryid interface{} `json:"CategoryId,omitempty"`
	Categoryname interface{} `json:"CategoryName,omitempty"`
}

// CreateConferenceProviderRequestIPDialIn represents the CreateConferenceProviderRequestIPDialIn schema from the OpenAPI specification
type CreateConferenceProviderRequestIPDialIn struct {
	Endpoint interface{} `json:"Endpoint"`
	Commsprotocol interface{} `json:"CommsProtocol"`
}

// RejectSkillRequest represents the RejectSkillRequest schema from the OpenAPI specification
type RejectSkillRequest struct {
	Skillid interface{} `json:"SkillId"`
}

// NetworkProfileData represents the NetworkProfileData schema from the OpenAPI specification
type NetworkProfileData struct {
	Securitytype interface{} `json:"SecurityType,omitempty"`
	Ssid interface{} `json:"Ssid,omitempty"`
	Certificateauthorityarn interface{} `json:"CertificateAuthorityArn,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Eapmethod interface{} `json:"EapMethod,omitempty"`
	Networkprofilearn interface{} `json:"NetworkProfileArn,omitempty"`
	Networkprofilename interface{} `json:"NetworkProfileName,omitempty"`
}

// DeleteUserResponse represents the DeleteUserResponse schema from the OpenAPI specification
type DeleteUserResponse struct {
}

// SkillGroup represents the SkillGroup schema from the OpenAPI specification
type SkillGroup struct {
	Skillgrouparn interface{} `json:"SkillGroupArn,omitempty"`
	Skillgroupname interface{} `json:"SkillGroupName,omitempty"`
	Description interface{} `json:"Description,omitempty"`
}

// UntagResourceRequest represents the UntagResourceRequest schema from the OpenAPI specification
type UntagResourceRequest struct {
	Tagkeys interface{} `json:"TagKeys"`
	Arn interface{} `json:"Arn"`
}

// DeleteSkillAuthorizationRequest represents the DeleteSkillAuthorizationRequest schema from the OpenAPI specification
type DeleteSkillAuthorizationRequest struct {
	Skillid interface{} `json:"SkillId"`
	Roomarn interface{} `json:"RoomArn,omitempty"`
}

// AssociateSkillWithSkillGroupRequest represents the AssociateSkillWithSkillGroupRequest schema from the OpenAPI specification
type AssociateSkillWithSkillGroupRequest struct {
	Skillid interface{} `json:"SkillId"`
	Skillgrouparn interface{} `json:"SkillGroupArn,omitempty"`
}

// CreateAddressBookResponse represents the CreateAddressBookResponse schema from the OpenAPI specification
type CreateAddressBookResponse struct {
	Addressbookarn interface{} `json:"AddressBookArn,omitempty"`
}

// PutSkillAuthorizationRequest represents the PutSkillAuthorizationRequest schema from the OpenAPI specification
type PutSkillAuthorizationRequest struct {
	Roomarn interface{} `json:"RoomArn,omitempty"`
	Skillid interface{} `json:"SkillId"`
	Authorizationresult interface{} `json:"AuthorizationResult"`
}

// PSTNDialIn represents the PSTNDialIn schema from the OpenAPI specification
type PSTNDialIn struct {
	Countrycode interface{} `json:"CountryCode"`
	Oneclickiddelay interface{} `json:"OneClickIdDelay"`
	Oneclickpindelay interface{} `json:"OneClickPinDelay"`
	Phonenumber interface{} `json:"PhoneNumber"`
}

// PutConferencePreferenceRequest represents the PutConferencePreferenceRequest schema from the OpenAPI specification
type PutConferencePreferenceRequest struct {
	Conferencepreference PutConferencePreferenceRequestConferencePreference `json:"ConferencePreference"`
}

// BusinessReportScheduleLastBusinessReport represents the BusinessReportScheduleLastBusinessReport schema from the OpenAPI specification
type BusinessReportScheduleLastBusinessReport struct {
	Downloadurl interface{} `json:"DownloadUrl,omitempty"`
	Failurecode interface{} `json:"FailureCode,omitempty"`
	S3location BusinessReportS3Location `json:"S3Location,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Deliverytime interface{} `json:"DeliveryTime,omitempty"`
}

// RoomSkillParameter represents the RoomSkillParameter schema from the OpenAPI specification
type RoomSkillParameter struct {
	Parameterkey interface{} `json:"ParameterKey"`
	Parametervalue interface{} `json:"ParameterValue"`
}

// AssociateSkillWithUsersRequest represents the AssociateSkillWithUsersRequest schema from the OpenAPI specification
type AssociateSkillWithUsersRequest struct {
	Skillid interface{} `json:"SkillId"`
}

// DeleteContactResponse represents the DeleteContactResponse schema from the OpenAPI specification
type DeleteContactResponse struct {
}

// ProactiveJoin represents the ProactiveJoin schema from the OpenAPI specification
type ProactiveJoin struct {
	Enabledbymotion bool `json:"EnabledByMotion,omitempty"`
}

// DeleteAddressBookRequest represents the DeleteAddressBookRequest schema from the OpenAPI specification
type DeleteAddressBookRequest struct {
	Addressbookarn interface{} `json:"AddressBookArn"`
}

// MeetingRoomConfiguration represents the MeetingRoomConfiguration schema from the OpenAPI specification
type MeetingRoomConfiguration struct {
	Instantbooking MeetingRoomConfigurationInstantBooking `json:"InstantBooking,omitempty"`
	Proactivejoin ProactiveJoin `json:"ProactiveJoin,omitempty"`
	Requirecheckin MeetingRoomConfigurationRequireCheckIn `json:"RequireCheckIn,omitempty"`
	Roomutilizationmetricsenabled interface{} `json:"RoomUtilizationMetricsEnabled,omitempty"`
	Endofmeetingreminder MeetingRoomConfigurationEndOfMeetingReminder `json:"EndOfMeetingReminder,omitempty"`
}

// UpdateProfileResponse represents the UpdateProfileResponse schema from the OpenAPI specification
type UpdateProfileResponse struct {
}

// ListBusinessReportSchedulesResponse represents the ListBusinessReportSchedulesResponse schema from the OpenAPI specification
type ListBusinessReportSchedulesResponse struct {
	Businessreportschedules interface{} `json:"BusinessReportSchedules,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// DeleteNetworkProfileResponse represents the DeleteNetworkProfileResponse schema from the OpenAPI specification
type DeleteNetworkProfileResponse struct {
}

// GetGatewayGroupRequest represents the GetGatewayGroupRequest schema from the OpenAPI specification
type GetGatewayGroupRequest struct {
	Gatewaygrouparn interface{} `json:"GatewayGroupArn"`
}

// UpdateAddressBookResponse represents the UpdateAddressBookResponse schema from the OpenAPI specification
type UpdateAddressBookResponse struct {
}

// CreateSkillGroupRequest represents the CreateSkillGroupRequest schema from the OpenAPI specification
type CreateSkillGroupRequest struct {
	Clientrequesttoken interface{} `json:"ClientRequestToken,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Skillgroupname interface{} `json:"SkillGroupName"`
	Tags interface{} `json:"Tags,omitempty"`
}

// SendAnnouncementRequestContent represents the SendAnnouncementRequestContent schema from the OpenAPI specification
type SendAnnouncementRequestContent struct {
	Textlist interface{} `json:"TextList,omitempty"`
	Audiolist interface{} `json:"AudioList,omitempty"`
	Ssmllist interface{} `json:"SsmlList,omitempty"`
}

// SendAnnouncementResponse represents the SendAnnouncementResponse schema from the OpenAPI specification
type SendAnnouncementResponse struct {
	Announcementarn interface{} `json:"AnnouncementArn,omitempty"`
}

// StartSmartHomeApplianceDiscoveryRequest represents the StartSmartHomeApplianceDiscoveryRequest schema from the OpenAPI specification
type StartSmartHomeApplianceDiscoveryRequest struct {
	Roomarn interface{} `json:"RoomArn"`
}

// ListSmartHomeAppliancesResponse represents the ListSmartHomeAppliancesResponse schema from the OpenAPI specification
type ListSmartHomeAppliancesResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Smarthomeappliances interface{} `json:"SmartHomeAppliances,omitempty"`
}

// CreateInstantBooking represents the CreateInstantBooking schema from the OpenAPI specification
type CreateInstantBooking struct {
	Durationinminutes interface{} `json:"DurationInMinutes"`
	Enabled interface{} `json:"Enabled"`
}

// Reviews represents the Reviews schema from the OpenAPI specification
type Reviews struct {
}

// SearchContactsRequest represents the SearchContactsRequest schema from the OpenAPI specification
type SearchContactsRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Sortcriteria interface{} `json:"SortCriteria,omitempty"`
	Filters interface{} `json:"Filters,omitempty"`
}

// AssociateSkillWithUsersResponse represents the AssociateSkillWithUsersResponse schema from the OpenAPI specification
type AssociateSkillWithUsersResponse struct {
}

// SearchUsersResponse represents the SearchUsersResponse schema from the OpenAPI specification
type SearchUsersResponse struct {
	Totalcount interface{} `json:"TotalCount,omitempty"`
	Users interface{} `json:"Users,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// Profile represents the Profile schema from the OpenAPI specification
type Profile struct {
	Meetingroomconfiguration ProfileMeetingRoomConfiguration `json:"MeetingRoomConfiguration,omitempty"`
	Maxvolumelimit interface{} `json:"MaxVolumeLimit,omitempty"`
	Profilearn interface{} `json:"ProfileArn,omitempty"`
	Addressbookarn interface{} `json:"AddressBookArn,omitempty"`
	Locale interface{} `json:"Locale,omitempty"`
	Temperatureunit interface{} `json:"TemperatureUnit,omitempty"`
	Setupmodedisabled interface{} `json:"SetupModeDisabled,omitempty"`
	Wakeword interface{} `json:"WakeWord,omitempty"`
	Dataretentionoptin interface{} `json:"DataRetentionOptIn,omitempty"`
	Profilename interface{} `json:"ProfileName,omitempty"`
	Address interface{} `json:"Address,omitempty"`
	Isdefault interface{} `json:"IsDefault,omitempty"`
	Distanceunit interface{} `json:"DistanceUnit,omitempty"`
	Pstnenabled interface{} `json:"PSTNEnabled,omitempty"`
	Timezone interface{} `json:"Timezone,omitempty"`
}

// UpdateMeetingRoomConfigurationInstantBooking represents the UpdateMeetingRoomConfigurationInstantBooking schema from the OpenAPI specification
type UpdateMeetingRoomConfigurationInstantBooking struct {
	Enabled interface{} `json:"Enabled,omitempty"`
	Durationinminutes interface{} `json:"DurationInMinutes,omitempty"`
}

// SendInvitationRequest represents the SendInvitationRequest schema from the OpenAPI specification
type SendInvitationRequest struct {
	Userarn interface{} `json:"UserArn,omitempty"`
}

// CreateRoomRequest represents the CreateRoomRequest schema from the OpenAPI specification
type CreateRoomRequest struct {
	Roomname interface{} `json:"RoomName"`
	Tags interface{} `json:"Tags,omitempty"`
	Clientrequesttoken interface{} `json:"ClientRequestToken,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Profilearn interface{} `json:"ProfileArn,omitempty"`
	Providercalendarid interface{} `json:"ProviderCalendarId,omitempty"`
}

// DeviceData represents the DeviceData schema from the OpenAPI specification
type DeviceData struct {
	Macaddress interface{} `json:"MacAddress,omitempty"`
	Devicearn interface{} `json:"DeviceArn,omitempty"`
	Devicetype interface{} `json:"DeviceType,omitempty"`
	Roomarn interface{} `json:"RoomArn,omitempty"`
	Roomname interface{} `json:"RoomName,omitempty"`
	Createdtime interface{} `json:"CreatedTime,omitempty"`
	Devicename interface{} `json:"DeviceName,omitempty"`
	Deviceserialnumber interface{} `json:"DeviceSerialNumber,omitempty"`
	Devicestatusinfo DeviceDeviceStatusInfo `json:"DeviceStatusInfo,omitempty"`
	Networkprofilearn interface{} `json:"NetworkProfileArn,omitempty"`
	Softwareversion interface{} `json:"SoftwareVersion,omitempty"`
	Devicestatus interface{} `json:"DeviceStatus,omitempty"`
	Networkprofilename interface{} `json:"NetworkProfileName,omitempty"`
}

// DisassociateSkillFromUsersResponse represents the DisassociateSkillFromUsersResponse schema from the OpenAPI specification
type DisassociateSkillFromUsersResponse struct {
}

// UpdateGatewayGroupRequest represents the UpdateGatewayGroupRequest schema from the OpenAPI specification
type UpdateGatewayGroupRequest struct {
	Description interface{} `json:"Description,omitempty"`
	Gatewaygrouparn interface{} `json:"GatewayGroupArn"`
	Name interface{} `json:"Name,omitempty"`
}

// ListDeviceEventsRequest represents the ListDeviceEventsRequest schema from the OpenAPI specification
type ListDeviceEventsRequest struct {
	Eventtype interface{} `json:"EventType,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Devicearn interface{} `json:"DeviceArn"`
}

// CreateProfileRequestMeetingRoomConfiguration represents the CreateProfileRequestMeetingRoomConfiguration schema from the OpenAPI specification
type CreateProfileRequestMeetingRoomConfiguration struct {
	Requirecheckin CreateMeetingRoomConfigurationRequireCheckIn `json:"RequireCheckIn,omitempty"`
	Roomutilizationmetricsenabled interface{} `json:"RoomUtilizationMetricsEnabled,omitempty"`
	Endofmeetingreminder CreateEndOfMeetingReminder `json:"EndOfMeetingReminder,omitempty"` // Creates settings for the end of meeting reminder feature that are applied to a room profile. The end of meeting reminder enables Alexa to remind users when a meeting is ending.
	Instantbooking CreateMeetingRoomConfigurationInstantBooking `json:"InstantBooking,omitempty"`
	Proactivejoin CreateProactiveJoin `json:"ProactiveJoin,omitempty"`
}

// ListTagsResponse represents the ListTagsResponse schema from the OpenAPI specification
type ListTagsResponse struct {
	Tags interface{} `json:"Tags,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// AssociateSkillGroupWithRoomResponse represents the AssociateSkillGroupWithRoomResponse schema from the OpenAPI specification
type AssociateSkillGroupWithRoomResponse struct {
}

// UpdateBusinessReportScheduleRequestRecurrence represents the UpdateBusinessReportScheduleRequestRecurrence schema from the OpenAPI specification
type UpdateBusinessReportScheduleRequestRecurrence struct {
	Startdate interface{} `json:"StartDate,omitempty"`
}

// Filter represents the Filter schema from the OpenAPI specification
type Filter struct {
	Key interface{} `json:"Key"`
	Values interface{} `json:"Values"`
}

// SearchRoomsResponse represents the SearchRoomsResponse schema from the OpenAPI specification
type SearchRoomsResponse struct {
	Totalcount interface{} `json:"TotalCount,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Rooms interface{} `json:"Rooms,omitempty"`
}

// GetInvitationConfigurationResponse represents the GetInvitationConfigurationResponse schema from the OpenAPI specification
type GetInvitationConfigurationResponse struct {
	Contactemail interface{} `json:"ContactEmail,omitempty"`
	Organizationname interface{} `json:"OrganizationName,omitempty"`
	Privateskillids interface{} `json:"PrivateSkillIds,omitempty"`
}

// SearchUsersRequest represents the SearchUsersRequest schema from the OpenAPI specification
type SearchUsersRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Sortcriteria interface{} `json:"SortCriteria,omitempty"`
	Filters interface{} `json:"Filters,omitempty"`
}

// DeleteGatewayGroupRequest represents the DeleteGatewayGroupRequest schema from the OpenAPI specification
type DeleteGatewayGroupRequest struct {
	Gatewaygrouparn interface{} `json:"GatewayGroupArn"`
}

// Content represents the Content schema from the OpenAPI specification
type Content struct {
	Ssmllist interface{} `json:"SsmlList,omitempty"`
	Textlist interface{} `json:"TextList,omitempty"`
	Audiolist interface{} `json:"AudioList,omitempty"`
}

// GetContactResponse represents the GetContactResponse schema from the OpenAPI specification
type GetContactResponse struct {
	Contact GetContactResponseContact `json:"Contact,omitempty"`
}

// SkillsStoreSkill represents the SkillsStoreSkill schema from the OpenAPI specification
type SkillsStoreSkill struct {
	Skilldetails SkillsStoreSkillSkillDetails `json:"SkillDetails,omitempty"`
	Skillid interface{} `json:"SkillId,omitempty"`
	Skillname interface{} `json:"SkillName,omitempty"`
	Supportslinking interface{} `json:"SupportsLinking,omitempty"`
	Iconurl interface{} `json:"IconUrl,omitempty"`
	Sampleutterances interface{} `json:"SampleUtterances,omitempty"`
	Shortdescription interface{} `json:"ShortDescription,omitempty"`
}

// CreateMeetingRoomConfigurationInstantBooking represents the CreateMeetingRoomConfigurationInstantBooking schema from the OpenAPI specification
type CreateMeetingRoomConfigurationInstantBooking struct {
	Durationinminutes interface{} `json:"DurationInMinutes"`
	Enabled interface{} `json:"Enabled"`
}

// GetConferenceProviderResponseConferenceProvider represents the GetConferenceProviderResponseConferenceProvider schema from the OpenAPI specification
type GetConferenceProviderResponseConferenceProvider struct {
	TypeField interface{} `json:"Type,omitempty"`
	Arn interface{} `json:"Arn,omitempty"`
	Ipdialin CreateConferenceProviderRequestIPDialIn `json:"IPDialIn,omitempty"`
	Meetingsetting CreateConferenceProviderRequestMeetingSetting `json:"MeetingSetting,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Pstndialin CreateConferenceProviderRequestPSTNDialIn `json:"PSTNDialIn,omitempty"`
}

// GetInvitationConfigurationRequest represents the GetInvitationConfigurationRequest schema from the OpenAPI specification
type GetInvitationConfigurationRequest struct {
}

// SkillsStoreSkillSkillDetails represents the SkillsStoreSkillSkillDetails schema from the OpenAPI specification
type SkillsStoreSkillSkillDetails struct {
	Skilltypes interface{} `json:"SkillTypes,omitempty"`
	Newinthisversionbulletpoints interface{} `json:"NewInThisVersionBulletPoints,omitempty"`
	Productdescription interface{} `json:"ProductDescription,omitempty"`
	Invocationphrase interface{} `json:"InvocationPhrase,omitempty"`
	Releasedate interface{} `json:"ReleaseDate,omitempty"`
	Generickeywords interface{} `json:"GenericKeywords,omitempty"`
	Bulletpoints interface{} `json:"BulletPoints,omitempty"`
	Developerinfo SkillDetailsDeveloperInfo `json:"DeveloperInfo,omitempty"`
	Enduserlicenseagreement interface{} `json:"EndUserLicenseAgreement,omitempty"`
	Reviews interface{} `json:"Reviews,omitempty"`
}

// DeleteRoomResponse represents the DeleteRoomResponse schema from the OpenAPI specification
type DeleteRoomResponse struct {
}

// StartSmartHomeApplianceDiscoveryResponse represents the StartSmartHomeApplianceDiscoveryResponse schema from the OpenAPI specification
type StartSmartHomeApplianceDiscoveryResponse struct {
}

// EndOfMeetingReminder represents the EndOfMeetingReminder schema from the OpenAPI specification
type EndOfMeetingReminder struct {
	Enabled interface{} `json:"Enabled,omitempty"`
	Reminderatminutes interface{} `json:"ReminderAtMinutes,omitempty"`
	Remindertype interface{} `json:"ReminderType,omitempty"`
}

// GetRoomSkillParameterRequest represents the GetRoomSkillParameterRequest schema from the OpenAPI specification
type GetRoomSkillParameterRequest struct {
	Skillid interface{} `json:"SkillId"`
	Parameterkey interface{} `json:"ParameterKey"`
	Roomarn interface{} `json:"RoomArn,omitempty"`
}

// GetNetworkProfileResponseNetworkProfile represents the GetNetworkProfileResponseNetworkProfile schema from the OpenAPI specification
type GetNetworkProfileResponseNetworkProfile struct {
	Networkprofilearn interface{} `json:"NetworkProfileArn,omitempty"`
	Networkprofilename interface{} `json:"NetworkProfileName,omitempty"`
	Ssid interface{} `json:"Ssid,omitempty"`
	Trustanchors interface{} `json:"TrustAnchors,omitempty"`
	Nextpassword interface{} `json:"NextPassword,omitempty"`
	Eapmethod interface{} `json:"EapMethod,omitempty"`
	Securitytype interface{} `json:"SecurityType,omitempty"`
	Certificateauthorityarn interface{} `json:"CertificateAuthorityArn,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Currentpassword interface{} `json:"CurrentPassword,omitempty"`
}

// DeleteSkillAuthorizationResponse represents the DeleteSkillAuthorizationResponse schema from the OpenAPI specification
type DeleteSkillAuthorizationResponse struct {
}

// ListTagsRequest represents the ListTagsRequest schema from the OpenAPI specification
type ListTagsRequest struct {
	Arn interface{} `json:"Arn"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// BusinessReportS3Location represents the BusinessReportS3Location schema from the OpenAPI specification
type BusinessReportS3Location struct {
	Bucketname interface{} `json:"BucketName,omitempty"`
	Path interface{} `json:"Path,omitempty"`
}

// MeetingRoomConfigurationRequireCheckIn represents the MeetingRoomConfigurationRequireCheckIn schema from the OpenAPI specification
type MeetingRoomConfigurationRequireCheckIn struct {
	Enabled interface{} `json:"Enabled,omitempty"`
	Releaseafterminutes interface{} `json:"ReleaseAfterMinutes,omitempty"`
}

// ListSkillsStoreCategoriesRequest represents the ListSkillsStoreCategoriesRequest schema from the OpenAPI specification
type ListSkillsStoreCategoriesRequest struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
}

// InstantBooking represents the InstantBooking schema from the OpenAPI specification
type InstantBooking struct {
	Durationinminutes interface{} `json:"DurationInMinutes,omitempty"`
	Enabled interface{} `json:"Enabled,omitempty"`
}

// CreateMeetingRoomConfigurationRequireCheckIn represents the CreateMeetingRoomConfigurationRequireCheckIn schema from the OpenAPI specification
type CreateMeetingRoomConfigurationRequireCheckIn struct {
	Enabled interface{} `json:"Enabled"`
	Releaseafterminutes interface{} `json:"ReleaseAfterMinutes"`
}

// GetDeviceResponse represents the GetDeviceResponse schema from the OpenAPI specification
type GetDeviceResponse struct {
	Device GetDeviceResponseDevice `json:"Device,omitempty"`
}

// GetGatewayResponseGateway represents the GetGatewayResponseGateway schema from the OpenAPI specification
type GetGatewayResponseGateway struct {
	Softwareversion interface{} `json:"SoftwareVersion,omitempty"`
	Arn interface{} `json:"Arn,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Gatewaygrouparn interface{} `json:"GatewayGroupArn,omitempty"`
	Name interface{} `json:"Name,omitempty"`
}

// SearchNetworkProfilesRequest represents the SearchNetworkProfilesRequest schema from the OpenAPI specification
type SearchNetworkProfilesRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Sortcriteria interface{} `json:"SortCriteria,omitempty"`
	Filters interface{} `json:"Filters,omitempty"`
}

// GetGatewayResponse represents the GetGatewayResponse schema from the OpenAPI specification
type GetGatewayResponse struct {
	Gateway GetGatewayResponseGateway `json:"Gateway,omitempty"`
}

// ListGatewaysResponse represents the ListGatewaysResponse schema from the OpenAPI specification
type ListGatewaysResponse struct {
	Gateways interface{} `json:"Gateways,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// DeleteAddressBookResponse represents the DeleteAddressBookResponse schema from the OpenAPI specification
type DeleteAddressBookResponse struct {
}

// ForgetSmartHomeAppliancesResponse represents the ForgetSmartHomeAppliancesResponse schema from the OpenAPI specification
type ForgetSmartHomeAppliancesResponse struct {
}

// CreateConferenceProviderRequestMeetingSetting represents the CreateConferenceProviderRequestMeetingSetting schema from the OpenAPI specification
type CreateConferenceProviderRequestMeetingSetting struct {
	Requirepin interface{} `json:"RequirePin"`
}

// CreateUserRequest represents the CreateUserRequest schema from the OpenAPI specification
type CreateUserRequest struct {
	Userid interface{} `json:"UserId"`
	Clientrequesttoken interface{} `json:"ClientRequestToken,omitempty"`
	Email interface{} `json:"Email,omitempty"`
	Firstname interface{} `json:"FirstName,omitempty"`
	Lastname interface{} `json:"LastName,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
}

// DisassociateContactFromAddressBookRequest represents the DisassociateContactFromAddressBookRequest schema from the OpenAPI specification
type DisassociateContactFromAddressBookRequest struct {
	Addressbookarn interface{} `json:"AddressBookArn"`
	Contactarn interface{} `json:"ContactArn"`
}

// DeleteContactRequest represents the DeleteContactRequest schema from the OpenAPI specification
type DeleteContactRequest struct {
	Contactarn interface{} `json:"ContactArn"`
}

// GetProfileRequest represents the GetProfileRequest schema from the OpenAPI specification
type GetProfileRequest struct {
	Profilearn interface{} `json:"ProfileArn,omitempty"`
}

// DeleteRoomSkillParameterRequest represents the DeleteRoomSkillParameterRequest schema from the OpenAPI specification
type DeleteRoomSkillParameterRequest struct {
	Roomarn interface{} `json:"RoomArn,omitempty"`
	Skillid interface{} `json:"SkillId"`
	Parameterkey interface{} `json:"ParameterKey"`
}

// SearchProfilesRequest represents the SearchProfilesRequest schema from the OpenAPI specification
type SearchProfilesRequest struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Sortcriteria interface{} `json:"SortCriteria,omitempty"`
	Filters interface{} `json:"Filters,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
}

// PutRoomSkillParameterRequest represents the PutRoomSkillParameterRequest schema from the OpenAPI specification
type PutRoomSkillParameterRequest struct {
	Skillid interface{} `json:"SkillId"`
	Roomarn interface{} `json:"RoomArn,omitempty"`
	Roomskillparameter PutRoomSkillParameterRequestRoomSkillParameter `json:"RoomSkillParameter"`
}

// GetSkillGroupResponse represents the GetSkillGroupResponse schema from the OpenAPI specification
type GetSkillGroupResponse struct {
	Skillgroup GetSkillGroupResponseSkillGroup `json:"SkillGroup,omitempty"`
}

// DisassociateSkillGroupFromRoomRequest represents the DisassociateSkillGroupFromRoomRequest schema from the OpenAPI specification
type DisassociateSkillGroupFromRoomRequest struct {
	Skillgrouparn interface{} `json:"SkillGroupArn,omitempty"`
	Roomarn interface{} `json:"RoomArn,omitempty"`
}

// GetGatewayRequest represents the GetGatewayRequest schema from the OpenAPI specification
type GetGatewayRequest struct {
	Gatewayarn interface{} `json:"GatewayArn"`
}

// SkillDetailsDeveloperInfo represents the SkillDetailsDeveloperInfo schema from the OpenAPI specification
type SkillDetailsDeveloperInfo struct {
	Developername interface{} `json:"DeveloperName,omitempty"`
	Email interface{} `json:"Email,omitempty"`
	Privacypolicy interface{} `json:"PrivacyPolicy,omitempty"`
	Url interface{} `json:"Url,omitempty"`
}

// GetRoomRequest represents the GetRoomRequest schema from the OpenAPI specification
type GetRoomRequest struct {
	Roomarn interface{} `json:"RoomArn,omitempty"`
}

// ContactData represents the ContactData schema from the OpenAPI specification
type ContactData struct {
	Sipaddresses interface{} `json:"SipAddresses,omitempty"`
	Contactarn interface{} `json:"ContactArn,omitempty"`
	Displayname interface{} `json:"DisplayName,omitempty"`
	Firstname interface{} `json:"FirstName,omitempty"`
	Lastname interface{} `json:"LastName,omitempty"`
	Phonenumber interface{} `json:"PhoneNumber,omitempty"`
	Phonenumbers interface{} `json:"PhoneNumbers,omitempty"`
}

// DeleteDeviceUsageDataResponse represents the DeleteDeviceUsageDataResponse schema from the OpenAPI specification
type DeleteDeviceUsageDataResponse struct {
}

// UpdateMeetingRoomConfigurationRequireCheckIn represents the UpdateMeetingRoomConfigurationRequireCheckIn schema from the OpenAPI specification
type UpdateMeetingRoomConfigurationRequireCheckIn struct {
	Enabled interface{} `json:"Enabled,omitempty"`
	Releaseafterminutes interface{} `json:"ReleaseAfterMinutes,omitempty"`
}

// GetRoomResponse represents the GetRoomResponse schema from the OpenAPI specification
type GetRoomResponse struct {
	Room GetRoomResponseRoom `json:"Room,omitempty"`
}

// CreateRoomResponse represents the CreateRoomResponse schema from the OpenAPI specification
type CreateRoomResponse struct {
	Roomarn interface{} `json:"RoomArn,omitempty"`
}

// UpdateRequireCheckIn represents the UpdateRequireCheckIn schema from the OpenAPI specification
type UpdateRequireCheckIn struct {
	Enabled interface{} `json:"Enabled,omitempty"`
	Releaseafterminutes interface{} `json:"ReleaseAfterMinutes,omitempty"`
}

// DeviceNetworkProfileInfo represents the DeviceNetworkProfileInfo schema from the OpenAPI specification
type DeviceNetworkProfileInfo struct {
	Certificatearn interface{} `json:"CertificateArn,omitempty"`
	Certificateexpirationtime interface{} `json:"CertificateExpirationTime,omitempty"`
	Networkprofilearn interface{} `json:"NetworkProfileArn,omitempty"`
}

// CreateGatewayGroupResponse represents the CreateGatewayGroupResponse schema from the OpenAPI specification
type CreateGatewayGroupResponse struct {
	Gatewaygrouparn interface{} `json:"GatewayGroupArn,omitempty"`
}

// SearchContactsResponse represents the SearchContactsResponse schema from the OpenAPI specification
type SearchContactsResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Totalcount interface{} `json:"TotalCount,omitempty"`
	Contacts interface{} `json:"Contacts,omitempty"`
}

// SearchAddressBooksResponse represents the SearchAddressBooksResponse schema from the OpenAPI specification
type SearchAddressBooksResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Totalcount interface{} `json:"TotalCount,omitempty"`
	Addressbooks interface{} `json:"AddressBooks,omitempty"`
}

// DeleteDeviceRequest represents the DeleteDeviceRequest schema from the OpenAPI specification
type DeleteDeviceRequest struct {
	Devicearn interface{} `json:"DeviceArn"`
}

// GetConferenceProviderResponse represents the GetConferenceProviderResponse schema from the OpenAPI specification
type GetConferenceProviderResponse struct {
	Conferenceprovider GetConferenceProviderResponseConferenceProvider `json:"ConferenceProvider,omitempty"`
}

// BusinessReportSchedule represents the BusinessReportSchedule schema from the OpenAPI specification
type BusinessReportSchedule struct {
	Schedulename interface{} `json:"ScheduleName,omitempty"`
	Contentrange CreateBusinessReportScheduleRequestContentRange `json:"ContentRange,omitempty"`
	Format interface{} `json:"Format,omitempty"`
	Lastbusinessreport BusinessReportScheduleLastBusinessReport `json:"LastBusinessReport,omitempty"`
	Recurrence UpdateBusinessReportScheduleRequestRecurrence `json:"Recurrence,omitempty"`
	S3bucketname interface{} `json:"S3BucketName,omitempty"`
	S3keyprefix interface{} `json:"S3KeyPrefix,omitempty"`
	Schedulearn interface{} `json:"ScheduleArn,omitempty"`
}

// UntagResourceResponse represents the UntagResourceResponse schema from the OpenAPI specification
type UntagResourceResponse struct {
}

// CreateConferenceProviderRequest represents the CreateConferenceProviderRequest schema from the OpenAPI specification
type CreateConferenceProviderRequest struct {
	Clientrequesttoken interface{} `json:"ClientRequestToken,omitempty"`
	Conferenceprovidername interface{} `json:"ConferenceProviderName"`
	Conferenceprovidertype interface{} `json:"ConferenceProviderType"`
	Ipdialin CreateConferenceProviderRequestIPDialIn `json:"IPDialIn,omitempty"`
	Meetingsetting CreateConferenceProviderRequestMeetingSetting `json:"MeetingSetting"`
	Pstndialin CreateConferenceProviderRequestPSTNDialIn `json:"PSTNDialIn,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
}

// SipAddress represents the SipAddress schema from the OpenAPI specification
type SipAddress struct {
	Uri interface{} `json:"Uri"`
	TypeField interface{} `json:"Type"`
}

// NetworkProfile represents the NetworkProfile schema from the OpenAPI specification
type NetworkProfile struct {
	Certificateauthorityarn interface{} `json:"CertificateAuthorityArn,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Currentpassword interface{} `json:"CurrentPassword,omitempty"`
	Networkprofilearn interface{} `json:"NetworkProfileArn,omitempty"`
	Networkprofilename interface{} `json:"NetworkProfileName,omitempty"`
	Ssid interface{} `json:"Ssid,omitempty"`
	Trustanchors interface{} `json:"TrustAnchors,omitempty"`
	Nextpassword interface{} `json:"NextPassword,omitempty"`
	Eapmethod interface{} `json:"EapMethod,omitempty"`
	Securitytype interface{} `json:"SecurityType,omitempty"`
}

// SendAnnouncementRequest represents the SendAnnouncementRequest schema from the OpenAPI specification
type SendAnnouncementRequest struct {
	Timetoliveinseconds interface{} `json:"TimeToLiveInSeconds,omitempty"`
	Clientrequesttoken interface{} `json:"ClientRequestToken"`
	Content SendAnnouncementRequestContent `json:"Content"`
	Roomfilters interface{} `json:"RoomFilters"`
}

// TagResourceRequest represents the TagResourceRequest schema from the OpenAPI specification
type TagResourceRequest struct {
	Arn interface{} `json:"Arn"`
	Tags interface{} `json:"Tags"`
}

// PutSkillAuthorizationResponse represents the PutSkillAuthorizationResponse schema from the OpenAPI specification
type PutSkillAuthorizationResponse struct {
}

// ApproveSkillResponse represents the ApproveSkillResponse schema from the OpenAPI specification
type ApproveSkillResponse struct {
}

// ListGatewayGroupsResponse represents the ListGatewayGroupsResponse schema from the OpenAPI specification
type ListGatewayGroupsResponse struct {
	Gatewaygroups interface{} `json:"GatewayGroups,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// UpdateAddressBookRequest represents the UpdateAddressBookRequest schema from the OpenAPI specification
type UpdateAddressBookRequest struct {
	Addressbookarn interface{} `json:"AddressBookArn"`
	Description interface{} `json:"Description,omitempty"`
	Name interface{} `json:"Name,omitempty"`
}

// DeleteBusinessReportScheduleRequest represents the DeleteBusinessReportScheduleRequest schema from the OpenAPI specification
type DeleteBusinessReportScheduleRequest struct {
	Schedulearn interface{} `json:"ScheduleArn"`
}

// Tag represents the Tag schema from the OpenAPI specification
type Tag struct {
	Key interface{} `json:"Key"`
	Value interface{} `json:"Value"`
}

// AssociateContactWithAddressBookRequest represents the AssociateContactWithAddressBookRequest schema from the OpenAPI specification
type AssociateContactWithAddressBookRequest struct {
	Addressbookarn interface{} `json:"AddressBookArn"`
	Contactarn interface{} `json:"ContactArn"`
}

// ListSkillsStoreCategoriesResponse represents the ListSkillsStoreCategoriesResponse schema from the OpenAPI specification
type ListSkillsStoreCategoriesResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Categorylist interface{} `json:"CategoryList,omitempty"`
}

// UpdateNetworkProfileRequest represents the UpdateNetworkProfileRequest schema from the OpenAPI specification
type UpdateNetworkProfileRequest struct {
	Certificateauthorityarn interface{} `json:"CertificateAuthorityArn,omitempty"`
	Currentpassword interface{} `json:"CurrentPassword,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Networkprofilearn interface{} `json:"NetworkProfileArn"`
	Networkprofilename interface{} `json:"NetworkProfileName,omitempty"`
	Nextpassword interface{} `json:"NextPassword,omitempty"`
	Trustanchors interface{} `json:"TrustAnchors,omitempty"`
}

// CreateGatewayGroupRequest represents the CreateGatewayGroupRequest schema from the OpenAPI specification
type CreateGatewayGroupRequest struct {
	Clientrequesttoken interface{} `json:"ClientRequestToken"`
	Description interface{} `json:"Description,omitempty"`
	Name interface{} `json:"Name"`
	Tags interface{} `json:"Tags,omitempty"`
}

// AssociateDeviceWithNetworkProfileRequest represents the AssociateDeviceWithNetworkProfileRequest schema from the OpenAPI specification
type AssociateDeviceWithNetworkProfileRequest struct {
	Devicearn interface{} `json:"DeviceArn"`
	Networkprofilearn interface{} `json:"NetworkProfileArn"`
}

// CreateConferenceProviderResponse represents the CreateConferenceProviderResponse schema from the OpenAPI specification
type CreateConferenceProviderResponse struct {
	Conferenceproviderarn interface{} `json:"ConferenceProviderArn,omitempty"`
}

// RejectSkillResponse represents the RejectSkillResponse schema from the OpenAPI specification
type RejectSkillResponse struct {
}

// AssociateDeviceWithRoomRequest represents the AssociateDeviceWithRoomRequest schema from the OpenAPI specification
type AssociateDeviceWithRoomRequest struct {
	Roomarn interface{} `json:"RoomArn,omitempty"`
	Devicearn interface{} `json:"DeviceArn,omitempty"`
}

// SkillDetails represents the SkillDetails schema from the OpenAPI specification
type SkillDetails struct {
	Generickeywords interface{} `json:"GenericKeywords,omitempty"`
	Bulletpoints interface{} `json:"BulletPoints,omitempty"`
	Developerinfo SkillDetailsDeveloperInfo `json:"DeveloperInfo,omitempty"`
	Enduserlicenseagreement interface{} `json:"EndUserLicenseAgreement,omitempty"`
	Reviews interface{} `json:"Reviews,omitempty"`
	Skilltypes interface{} `json:"SkillTypes,omitempty"`
	Newinthisversionbulletpoints interface{} `json:"NewInThisVersionBulletPoints,omitempty"`
	Productdescription interface{} `json:"ProductDescription,omitempty"`
	Invocationphrase interface{} `json:"InvocationPhrase,omitempty"`
	Releasedate interface{} `json:"ReleaseDate,omitempty"`
}

// GetRoomResponseRoom represents the GetRoomResponseRoom schema from the OpenAPI specification
type GetRoomResponseRoom struct {
	Roomname interface{} `json:"RoomName,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Profilearn interface{} `json:"ProfileArn,omitempty"`
	Providercalendarid interface{} `json:"ProviderCalendarId,omitempty"`
	Roomarn interface{} `json:"RoomArn,omitempty"`
}

// GetContactResponseContact represents the GetContactResponseContact schema from the OpenAPI specification
type GetContactResponseContact struct {
	Phonenumbers interface{} `json:"PhoneNumbers,omitempty"`
	Sipaddresses interface{} `json:"SipAddresses,omitempty"`
	Contactarn interface{} `json:"ContactArn,omitempty"`
	Displayname interface{} `json:"DisplayName,omitempty"`
	Firstname interface{} `json:"FirstName,omitempty"`
	Lastname interface{} `json:"LastName,omitempty"`
	Phonenumber interface{} `json:"PhoneNumber,omitempty"`
}

// BusinessReportContentRange represents the BusinessReportContentRange schema from the OpenAPI specification
type BusinessReportContentRange struct {
	Interval interface{} `json:"Interval"`
}

// GetSkillGroupResponseSkillGroup represents the GetSkillGroupResponseSkillGroup schema from the OpenAPI specification
type GetSkillGroupResponseSkillGroup struct {
	Description interface{} `json:"Description,omitempty"`
	Skillgrouparn interface{} `json:"SkillGroupArn,omitempty"`
	Skillgroupname interface{} `json:"SkillGroupName,omitempty"`
}

// CreateNetworkProfileResponse represents the CreateNetworkProfileResponse schema from the OpenAPI specification
type CreateNetworkProfileResponse struct {
	Networkprofilearn interface{} `json:"NetworkProfileArn,omitempty"`
}

// DisassociateSkillGroupFromRoomResponse represents the DisassociateSkillGroupFromRoomResponse schema from the OpenAPI specification
type DisassociateSkillGroupFromRoomResponse struct {
}

// RevokeInvitationResponse represents the RevokeInvitationResponse schema from the OpenAPI specification
type RevokeInvitationResponse struct {
}

// Sort represents the Sort schema from the OpenAPI specification
type Sort struct {
	Value interface{} `json:"Value"`
	Key interface{} `json:"Key"`
}

// MeetingSetting represents the MeetingSetting schema from the OpenAPI specification
type MeetingSetting struct {
	Requirepin interface{} `json:"RequirePin"`
}

// UpdateConferenceProviderResponse represents the UpdateConferenceProviderResponse schema from the OpenAPI specification
type UpdateConferenceProviderResponse struct {
}

// UpdateRoomRequest represents the UpdateRoomRequest schema from the OpenAPI specification
type UpdateRoomRequest struct {
	Description interface{} `json:"Description,omitempty"`
	Profilearn interface{} `json:"ProfileArn,omitempty"`
	Providercalendarid interface{} `json:"ProviderCalendarId,omitempty"`
	Roomarn interface{} `json:"RoomArn,omitempty"`
	Roomname interface{} `json:"RoomName,omitempty"`
}

// DisassociateDeviceFromRoomRequest represents the DisassociateDeviceFromRoomRequest schema from the OpenAPI specification
type DisassociateDeviceFromRoomRequest struct {
	Devicearn interface{} `json:"DeviceArn,omitempty"`
}

// ListDeviceEventsResponse represents the ListDeviceEventsResponse schema from the OpenAPI specification
type ListDeviceEventsResponse struct {
	Deviceevents interface{} `json:"DeviceEvents,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// ListConferenceProvidersRequest represents the ListConferenceProvidersRequest schema from the OpenAPI specification
type ListConferenceProvidersRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// UpdateBusinessReportScheduleRequest represents the UpdateBusinessReportScheduleRequest schema from the OpenAPI specification
type UpdateBusinessReportScheduleRequest struct {
	Schedulename interface{} `json:"ScheduleName,omitempty"`
	Format interface{} `json:"Format,omitempty"`
	Recurrence UpdateBusinessReportScheduleRequestRecurrence `json:"Recurrence,omitempty"`
	S3bucketname interface{} `json:"S3BucketName,omitempty"`
	S3keyprefix interface{} `json:"S3KeyPrefix,omitempty"`
	Schedulearn interface{} `json:"ScheduleArn"`
}

// CreateBusinessReportScheduleRequestRecurrence represents the CreateBusinessReportScheduleRequestRecurrence schema from the OpenAPI specification
type CreateBusinessReportScheduleRequestRecurrence struct {
	Startdate interface{} `json:"StartDate,omitempty"`
}

// ListBusinessReportSchedulesRequest represents the ListBusinessReportSchedulesRequest schema from the OpenAPI specification
type ListBusinessReportSchedulesRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// GetAddressBookResponse represents the GetAddressBookResponse schema from the OpenAPI specification
type GetAddressBookResponse struct {
	Addressbook GetAddressBookResponseAddressBook `json:"AddressBook,omitempty"`
}

// UpdateSkillGroupRequest represents the UpdateSkillGroupRequest schema from the OpenAPI specification
type UpdateSkillGroupRequest struct {
	Description interface{} `json:"Description,omitempty"`
	Skillgrouparn interface{} `json:"SkillGroupArn,omitempty"`
	Skillgroupname interface{} `json:"SkillGroupName,omitempty"`
}

// GetConferencePreferenceResponsePreference represents the GetConferencePreferenceResponsePreference schema from the OpenAPI specification
type GetConferencePreferenceResponsePreference struct {
	Defaultconferenceproviderarn interface{} `json:"DefaultConferenceProviderArn,omitempty"`
}

// GetConferenceProviderRequest represents the GetConferenceProviderRequest schema from the OpenAPI specification
type GetConferenceProviderRequest struct {
	Conferenceproviderarn interface{} `json:"ConferenceProviderArn"`
}

// ListSkillsResponse represents the ListSkillsResponse schema from the OpenAPI specification
type ListSkillsResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Skillsummaries interface{} `json:"SkillSummaries,omitempty"`
}

// DeleteRoomSkillParameterResponse represents the DeleteRoomSkillParameterResponse schema from the OpenAPI specification
type DeleteRoomSkillParameterResponse struct {
}

// UpdateContactRequest represents the UpdateContactRequest schema from the OpenAPI specification
type UpdateContactRequest struct {
	Contactarn interface{} `json:"ContactArn"`
	Displayname interface{} `json:"DisplayName,omitempty"`
	Firstname interface{} `json:"FirstName,omitempty"`
	Lastname interface{} `json:"LastName,omitempty"`
	Phonenumber interface{} `json:"PhoneNumber,omitempty"`
	Phonenumbers interface{} `json:"PhoneNumbers,omitempty"`
	Sipaddresses interface{} `json:"SipAddresses,omitempty"`
}

// GetProfileResponseProfile represents the GetProfileResponseProfile schema from the OpenAPI specification
type GetProfileResponseProfile struct {
	Address interface{} `json:"Address,omitempty"`
	Isdefault interface{} `json:"IsDefault,omitempty"`
	Distanceunit interface{} `json:"DistanceUnit,omitempty"`
	Pstnenabled interface{} `json:"PSTNEnabled,omitempty"`
	Timezone interface{} `json:"Timezone,omitempty"`
	Meetingroomconfiguration ProfileMeetingRoomConfiguration `json:"MeetingRoomConfiguration,omitempty"`
	Maxvolumelimit interface{} `json:"MaxVolumeLimit,omitempty"`
	Profilearn interface{} `json:"ProfileArn,omitempty"`
	Addressbookarn interface{} `json:"AddressBookArn,omitempty"`
	Locale interface{} `json:"Locale,omitempty"`
	Temperatureunit interface{} `json:"TemperatureUnit,omitempty"`
	Setupmodedisabled interface{} `json:"SetupModeDisabled,omitempty"`
	Wakeword interface{} `json:"WakeWord,omitempty"`
	Dataretentionoptin interface{} `json:"DataRetentionOptIn,omitempty"`
	Profilename interface{} `json:"ProfileName,omitempty"`
}

// DeleteRoomRequest represents the DeleteRoomRequest schema from the OpenAPI specification
type DeleteRoomRequest struct {
	Roomarn interface{} `json:"RoomArn,omitempty"`
}

// CreateProfileRequest represents the CreateProfileRequest schema from the OpenAPI specification
type CreateProfileRequest struct {
	Maxvolumelimit interface{} `json:"MaxVolumeLimit,omitempty"`
	Meetingroomconfiguration CreateProfileRequestMeetingRoomConfiguration `json:"MeetingRoomConfiguration,omitempty"`
	Setupmodedisabled interface{} `json:"SetupModeDisabled,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Wakeword interface{} `json:"WakeWord"`
	Dataretentionoptin interface{} `json:"DataRetentionOptIn,omitempty"`
	Distanceunit interface{} `json:"DistanceUnit"`
	Pstnenabled interface{} `json:"PSTNEnabled,omitempty"`
	Clientrequesttoken interface{} `json:"ClientRequestToken,omitempty"`
	Profilename interface{} `json:"ProfileName"`
	Timezone interface{} `json:"Timezone"`
	Address interface{} `json:"Address"`
	Locale interface{} `json:"Locale,omitempty"`
	Temperatureunit interface{} `json:"TemperatureUnit"`
}

// RoomData represents the RoomData schema from the OpenAPI specification
type RoomData struct {
	Roomname interface{} `json:"RoomName,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Profilearn interface{} `json:"ProfileArn,omitempty"`
	Profilename interface{} `json:"ProfileName,omitempty"`
	Providercalendarid interface{} `json:"ProviderCalendarId,omitempty"`
	Roomarn interface{} `json:"RoomArn,omitempty"`
}

// GatewayGroupSummary represents the GatewayGroupSummary schema from the OpenAPI specification
type GatewayGroupSummary struct {
	Arn interface{} `json:"Arn,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Name interface{} `json:"Name,omitempty"`
}

// SearchProfilesResponse represents the SearchProfilesResponse schema from the OpenAPI specification
type SearchProfilesResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Profiles interface{} `json:"Profiles,omitempty"`
	Totalcount interface{} `json:"TotalCount,omitempty"`
}

// PutConferencePreferenceRequestConferencePreference represents the PutConferencePreferenceRequestConferencePreference schema from the OpenAPI specification
type PutConferencePreferenceRequestConferencePreference struct {
	Defaultconferenceproviderarn interface{} `json:"DefaultConferenceProviderArn,omitempty"`
}

// GatewaySummary represents the GatewaySummary schema from the OpenAPI specification
type GatewaySummary struct {
	Name interface{} `json:"Name,omitempty"`
	Softwareversion interface{} `json:"SoftwareVersion,omitempty"`
	Arn interface{} `json:"Arn,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Gatewaygrouparn interface{} `json:"GatewayGroupArn,omitempty"`
}

// UpdateBusinessReportScheduleResponse represents the UpdateBusinessReportScheduleResponse schema from the OpenAPI specification
type UpdateBusinessReportScheduleResponse struct {
}

// GetGatewayGroupResponse represents the GetGatewayGroupResponse schema from the OpenAPI specification
type GetGatewayGroupResponse struct {
	Gatewaygroup GatewayGroup `json:"GatewayGroup,omitempty"` // The details of the gateway group.
}

// AssociateSkillWithSkillGroupResponse represents the AssociateSkillWithSkillGroupResponse schema from the OpenAPI specification
type AssociateSkillWithSkillGroupResponse struct {
}

// CreateSkillGroupResponse represents the CreateSkillGroupResponse schema from the OpenAPI specification
type CreateSkillGroupResponse struct {
	Skillgrouparn interface{} `json:"SkillGroupArn,omitempty"`
}

// GetDeviceResponseDevice represents the GetDeviceResponseDevice schema from the OpenAPI specification
type GetDeviceResponseDevice struct {
	Networkprofileinfo DeviceNetworkProfileInfo `json:"NetworkProfileInfo,omitempty"`
	Macaddress interface{} `json:"MacAddress,omitempty"`
	Softwareversion interface{} `json:"SoftwareVersion,omitempty"`
	Devicearn interface{} `json:"DeviceArn,omitempty"`
	Deviceserialnumber interface{} `json:"DeviceSerialNumber,omitempty"`
	Devicestatusinfo DeviceDeviceStatusInfo `json:"DeviceStatusInfo,omitempty"`
	Roomarn interface{} `json:"RoomArn,omitempty"`
	Devicename interface{} `json:"DeviceName,omitempty"`
	Devicestatus interface{} `json:"DeviceStatus,omitempty"`
	Devicetype interface{} `json:"DeviceType,omitempty"`
}

// TagResourceResponse represents the TagResourceResponse schema from the OpenAPI specification
type TagResourceResponse struct {
}

// GetSkillGroupRequest represents the GetSkillGroupRequest schema from the OpenAPI specification
type GetSkillGroupRequest struct {
	Skillgrouparn interface{} `json:"SkillGroupArn,omitempty"`
}

// DeleteSkillGroupRequest represents the DeleteSkillGroupRequest schema from the OpenAPI specification
type DeleteSkillGroupRequest struct {
	Skillgrouparn interface{} `json:"SkillGroupArn,omitempty"`
}

// BusinessReportRecurrence represents the BusinessReportRecurrence schema from the OpenAPI specification
type BusinessReportRecurrence struct {
	Startdate interface{} `json:"StartDate,omitempty"`
}

// UpdateMeetingRoomConfiguration represents the UpdateMeetingRoomConfiguration schema from the OpenAPI specification
type UpdateMeetingRoomConfiguration struct {
	Instantbooking UpdateMeetingRoomConfigurationInstantBooking `json:"InstantBooking,omitempty"`
	Proactivejoin UpdateProactiveJoin `json:"ProactiveJoin,omitempty"`
	Requirecheckin UpdateMeetingRoomConfigurationRequireCheckIn `json:"RequireCheckIn,omitempty"`
	Roomutilizationmetricsenabled interface{} `json:"RoomUtilizationMetricsEnabled,omitempty"`
	Endofmeetingreminder UpdateMeetingRoomConfigurationEndOfMeetingReminder `json:"EndOfMeetingReminder,omitempty"`
}

// ListGatewaysRequest represents the ListGatewaysRequest schema from the OpenAPI specification
type ListGatewaysRequest struct {
	Gatewaygrouparn interface{} `json:"GatewayGroupArn,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// CreateContactResponse represents the CreateContactResponse schema from the OpenAPI specification
type CreateContactResponse struct {
	Contactarn interface{} `json:"ContactArn,omitempty"`
}

// UpdateProfileRequestMeetingRoomConfiguration represents the UpdateProfileRequestMeetingRoomConfiguration schema from the OpenAPI specification
type UpdateProfileRequestMeetingRoomConfiguration struct {
	Proactivejoin UpdateProactiveJoin `json:"ProactiveJoin,omitempty"`
	Requirecheckin UpdateMeetingRoomConfigurationRequireCheckIn `json:"RequireCheckIn,omitempty"`
	Roomutilizationmetricsenabled interface{} `json:"RoomUtilizationMetricsEnabled,omitempty"`
	Endofmeetingreminder UpdateMeetingRoomConfigurationEndOfMeetingReminder `json:"EndOfMeetingReminder,omitempty"`
	Instantbooking UpdateMeetingRoomConfigurationInstantBooking `json:"InstantBooking,omitempty"`
}

// GetAddressBookRequest represents the GetAddressBookRequest schema from the OpenAPI specification
type GetAddressBookRequest struct {
	Addressbookarn interface{} `json:"AddressBookArn"`
}

// SearchNetworkProfilesResponse represents the SearchNetworkProfilesResponse schema from the OpenAPI specification
type SearchNetworkProfilesResponse struct {
	Networkprofiles interface{} `json:"NetworkProfiles,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Totalcount interface{} `json:"TotalCount,omitempty"`
}

// CreateBusinessReportScheduleRequest represents the CreateBusinessReportScheduleRequest schema from the OpenAPI specification
type CreateBusinessReportScheduleRequest struct {
	Clientrequesttoken interface{} `json:"ClientRequestToken,omitempty"`
	Contentrange CreateBusinessReportScheduleRequestContentRange `json:"ContentRange"`
	Format interface{} `json:"Format"`
	Recurrence CreateBusinessReportScheduleRequestRecurrence `json:"Recurrence,omitempty"`
	S3bucketname interface{} `json:"S3BucketName,omitempty"`
	S3keyprefix interface{} `json:"S3KeyPrefix,omitempty"`
	Schedulename interface{} `json:"ScheduleName,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
}

// SearchSkillGroupsRequest represents the SearchSkillGroupsRequest schema from the OpenAPI specification
type SearchSkillGroupsRequest struct {
	Sortcriteria interface{} `json:"SortCriteria,omitempty"`
	Filters interface{} `json:"Filters,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// CreateBusinessReportScheduleRequestContentRange represents the CreateBusinessReportScheduleRequestContentRange schema from the OpenAPI specification
type CreateBusinessReportScheduleRequestContentRange struct {
	Interval interface{} `json:"Interval"`
}

// DeveloperInfo represents the DeveloperInfo schema from the OpenAPI specification
type DeveloperInfo struct {
	Developername interface{} `json:"DeveloperName,omitempty"`
	Email interface{} `json:"Email,omitempty"`
	Privacypolicy interface{} `json:"PrivacyPolicy,omitempty"`
	Url interface{} `json:"Url,omitempty"`
}

// GetContactRequest represents the GetContactRequest schema from the OpenAPI specification
type GetContactRequest struct {
	Contactarn interface{} `json:"ContactArn"`
}

// UpdateGatewayGroupResponse represents the UpdateGatewayGroupResponse schema from the OpenAPI specification
type UpdateGatewayGroupResponse struct {
}

// ProfileMeetingRoomConfiguration represents the ProfileMeetingRoomConfiguration schema from the OpenAPI specification
type ProfileMeetingRoomConfiguration struct {
	Roomutilizationmetricsenabled interface{} `json:"RoomUtilizationMetricsEnabled,omitempty"`
	Endofmeetingreminder MeetingRoomConfigurationEndOfMeetingReminder `json:"EndOfMeetingReminder,omitempty"`
	Instantbooking MeetingRoomConfigurationInstantBooking `json:"InstantBooking,omitempty"`
	Proactivejoin ProactiveJoin `json:"ProactiveJoin,omitempty"`
	Requirecheckin MeetingRoomConfigurationRequireCheckIn `json:"RequireCheckIn,omitempty"`
}

// ConferencePreference represents the ConferencePreference schema from the OpenAPI specification
type ConferencePreference struct {
	Defaultconferenceproviderarn interface{} `json:"DefaultConferenceProviderArn,omitempty"`
}

// UpdateGatewayResponse represents the UpdateGatewayResponse schema from the OpenAPI specification
type UpdateGatewayResponse struct {
}

// CreateContactRequest represents the CreateContactRequest schema from the OpenAPI specification
type CreateContactRequest struct {
	Tags interface{} `json:"Tags,omitempty"`
	Clientrequesttoken interface{} `json:"ClientRequestToken,omitempty"`
	Displayname interface{} `json:"DisplayName,omitempty"`
	Firstname interface{} `json:"FirstName"`
	Lastname interface{} `json:"LastName,omitempty"`
	Phonenumber interface{} `json:"PhoneNumber,omitempty"`
	Phonenumbers interface{} `json:"PhoneNumbers,omitempty"`
	Sipaddresses interface{} `json:"SipAddresses,omitempty"`
}

// DisassociateSkillFromSkillGroupRequest represents the DisassociateSkillFromSkillGroupRequest schema from the OpenAPI specification
type DisassociateSkillFromSkillGroupRequest struct {
	Skillid interface{} `json:"SkillId"`
	Skillgrouparn interface{} `json:"SkillGroupArn,omitempty"`
}

// GetRoomSkillParameterResponseRoomSkillParameter represents the GetRoomSkillParameterResponseRoomSkillParameter schema from the OpenAPI specification
type GetRoomSkillParameterResponseRoomSkillParameter struct {
	Parametervalue interface{} `json:"ParameterValue"`
	Parameterkey interface{} `json:"ParameterKey"`
}

// GetConferencePreferenceRequest represents the GetConferencePreferenceRequest schema from the OpenAPI specification
type GetConferencePreferenceRequest struct {
}

// DisassociateContactFromAddressBookResponse represents the DisassociateContactFromAddressBookResponse schema from the OpenAPI specification
type DisassociateContactFromAddressBookResponse struct {
}

// UpdateDeviceResponse represents the UpdateDeviceResponse schema from the OpenAPI specification
type UpdateDeviceResponse struct {
}

// SendInvitationResponse represents the SendInvitationResponse schema from the OpenAPI specification
type SendInvitationResponse struct {
}

// SearchDevicesResponse represents the SearchDevicesResponse schema from the OpenAPI specification
type SearchDevicesResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Totalcount interface{} `json:"TotalCount,omitempty"`
	Devices interface{} `json:"Devices,omitempty"`
}

// DeleteDeviceResponse represents the DeleteDeviceResponse schema from the OpenAPI specification
type DeleteDeviceResponse struct {
}

// CreateNetworkProfileRequest represents the CreateNetworkProfileRequest schema from the OpenAPI specification
type CreateNetworkProfileRequest struct {
	Currentpassword interface{} `json:"CurrentPassword,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Networkprofilename interface{} `json:"NetworkProfileName"`
	Securitytype interface{} `json:"SecurityType"`
	Clientrequesttoken string `json:"ClientRequestToken"` // A unique, user-specified identifier for the request that ensures idempotency.
	Eapmethod interface{} `json:"EapMethod,omitempty"`
	Ssid interface{} `json:"Ssid"`
	Certificateauthorityarn interface{} `json:"CertificateAuthorityArn,omitempty"`
	Nextpassword interface{} `json:"NextPassword,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Trustanchors interface{} `json:"TrustAnchors,omitempty"`
}

// ResolveRoomResponse represents the ResolveRoomResponse schema from the OpenAPI specification
type ResolveRoomResponse struct {
	Roomarn interface{} `json:"RoomArn,omitempty"`
	Roomname interface{} `json:"RoomName,omitempty"`
	Roomskillparameters interface{} `json:"RoomSkillParameters,omitempty"`
}

// AssociateDeviceWithRoomResponse represents the AssociateDeviceWithRoomResponse schema from the OpenAPI specification
type AssociateDeviceWithRoomResponse struct {
}

// GetDeviceRequest represents the GetDeviceRequest schema from the OpenAPI specification
type GetDeviceRequest struct {
	Devicearn interface{} `json:"DeviceArn,omitempty"`
}

// UpdateDeviceRequest represents the UpdateDeviceRequest schema from the OpenAPI specification
type UpdateDeviceRequest struct {
	Devicearn interface{} `json:"DeviceArn,omitempty"`
	Devicename interface{} `json:"DeviceName,omitempty"`
}

// ListSkillsStoreSkillsByCategoryResponse represents the ListSkillsStoreSkillsByCategoryResponse schema from the OpenAPI specification
type ListSkillsStoreSkillsByCategoryResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Skillsstoreskills interface{} `json:"SkillsStoreSkills,omitempty"`
}

// DeleteConferenceProviderResponse represents the DeleteConferenceProviderResponse schema from the OpenAPI specification
type DeleteConferenceProviderResponse struct {
}

// AddressBook represents the AddressBook schema from the OpenAPI specification
type AddressBook struct {
	Name interface{} `json:"Name,omitempty"`
	Addressbookarn interface{} `json:"AddressBookArn,omitempty"`
	Description interface{} `json:"Description,omitempty"`
}

// DeviceStatusInfo represents the DeviceStatusInfo schema from the OpenAPI specification
type DeviceStatusInfo struct {
	Connectionstatus interface{} `json:"ConnectionStatus,omitempty"`
	Connectionstatusupdatedtime interface{} `json:"ConnectionStatusUpdatedTime,omitempty"`
	Devicestatusdetails interface{} `json:"DeviceStatusDetails,omitempty"`
}

// SearchAddressBooksRequest represents the SearchAddressBooksRequest schema from the OpenAPI specification
type SearchAddressBooksRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Sortcriteria interface{} `json:"SortCriteria,omitempty"`
	Filters interface{} `json:"Filters,omitempty"`
}

// UpdateRoomResponse represents the UpdateRoomResponse schema from the OpenAPI specification
type UpdateRoomResponse struct {
}

// AuthorizationResult represents the AuthorizationResult schema from the OpenAPI specification
type AuthorizationResult struct {
}

// RevokeInvitationRequest represents the RevokeInvitationRequest schema from the OpenAPI specification
type RevokeInvitationRequest struct {
	Enrollmentid interface{} `json:"EnrollmentId,omitempty"`
	Userarn interface{} `json:"UserArn,omitempty"`
}

// BusinessReportS3Location represents the BusinessReportS3Location schema from the OpenAPI specification
type BusinessReportS3Location struct {
	Bucketname interface{} `json:"BucketName,omitempty"`
	Path interface{} `json:"Path,omitempty"`
}

// Text represents the Text schema from the OpenAPI specification
type Text struct {
	Locale interface{} `json:"Locale"`
	Value interface{} `json:"Value"`
}

// UserData represents the UserData schema from the OpenAPI specification
type UserData struct {
	Email interface{} `json:"Email,omitempty"`
	Enrollmentid interface{} `json:"EnrollmentId,omitempty"`
	Enrollmentstatus interface{} `json:"EnrollmentStatus,omitempty"`
	Firstname interface{} `json:"FirstName,omitempty"`
	Lastname interface{} `json:"LastName,omitempty"`
	Userarn interface{} `json:"UserArn,omitempty"`
}

// RegisterAVSDeviceResponse represents the RegisterAVSDeviceResponse schema from the OpenAPI specification
type RegisterAVSDeviceResponse struct {
	Devicearn interface{} `json:"DeviceArn,omitempty"`
}

// PutRoomSkillParameterRequestRoomSkillParameter represents the PutRoomSkillParameterRequestRoomSkillParameter schema from the OpenAPI specification
type PutRoomSkillParameterRequestRoomSkillParameter struct {
	Parametervalue interface{} `json:"ParameterValue"`
	Parameterkey interface{} `json:"ParameterKey"`
}

// AssociateSkillGroupWithRoomRequest represents the AssociateSkillGroupWithRoomRequest schema from the OpenAPI specification
type AssociateSkillGroupWithRoomRequest struct {
	Skillgrouparn interface{} `json:"SkillGroupArn,omitempty"`
	Roomarn interface{} `json:"RoomArn,omitempty"`
}

// StartDeviceSyncRequest represents the StartDeviceSyncRequest schema from the OpenAPI specification
type StartDeviceSyncRequest struct {
	Devicearn interface{} `json:"DeviceArn,omitempty"`
	Features interface{} `json:"Features"`
	Roomarn interface{} `json:"RoomArn,omitempty"`
}

// ListGatewayGroupsRequest represents the ListGatewayGroupsRequest schema from the OpenAPI specification
type ListGatewayGroupsRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// Audio represents the Audio schema from the OpenAPI specification
type Audio struct {
	Locale interface{} `json:"Locale"`
	Location interface{} `json:"Location"`
}

// DeleteDeviceUsageDataRequest represents the DeleteDeviceUsageDataRequest schema from the OpenAPI specification
type DeleteDeviceUsageDataRequest struct {
	Devicearn interface{} `json:"DeviceArn"`
	Deviceusagetype interface{} `json:"DeviceUsageType"`
}

// GetNetworkProfileRequest represents the GetNetworkProfileRequest schema from the OpenAPI specification
type GetNetworkProfileRequest struct {
	Networkprofilearn interface{} `json:"NetworkProfileArn"`
}

// RequireCheckIn represents the RequireCheckIn schema from the OpenAPI specification
type RequireCheckIn struct {
	Enabled interface{} `json:"Enabled,omitempty"`
	Releaseafterminutes interface{} `json:"ReleaseAfterMinutes,omitempty"`
}

// DeleteProfileRequest represents the DeleteProfileRequest schema from the OpenAPI specification
type DeleteProfileRequest struct {
	Profilearn interface{} `json:"ProfileArn,omitempty"`
}

// GetRoomSkillParameterResponse represents the GetRoomSkillParameterResponse schema from the OpenAPI specification
type GetRoomSkillParameterResponse struct {
	Roomskillparameter GetRoomSkillParameterResponseRoomSkillParameter `json:"RoomSkillParameter,omitempty"`
}

// ApproveSkillRequest represents the ApproveSkillRequest schema from the OpenAPI specification
type ApproveSkillRequest struct {
	Skillid interface{} `json:"SkillId"`
}

// CreateUserResponse represents the CreateUserResponse schema from the OpenAPI specification
type CreateUserResponse struct {
	Userarn interface{} `json:"UserArn,omitempty"`
}

// Gateway represents the Gateway schema from the OpenAPI specification
type Gateway struct {
	Description interface{} `json:"Description,omitempty"`
	Gatewaygrouparn interface{} `json:"GatewayGroupArn,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Softwareversion interface{} `json:"SoftwareVersion,omitempty"`
	Arn interface{} `json:"Arn,omitempty"`
}

// ListSkillsStoreSkillsByCategoryRequest represents the ListSkillsStoreSkillsByCategoryRequest schema from the OpenAPI specification
type ListSkillsStoreSkillsByCategoryRequest struct {
	Categoryid interface{} `json:"CategoryId"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// UpdateInstantBooking represents the UpdateInstantBooking schema from the OpenAPI specification
type UpdateInstantBooking struct {
	Enabled interface{} `json:"Enabled,omitempty"`
	Durationinminutes interface{} `json:"DurationInMinutes,omitempty"`
}

// CreateRequireCheckIn represents the CreateRequireCheckIn schema from the OpenAPI specification
type CreateRequireCheckIn struct {
	Enabled interface{} `json:"Enabled"`
	Releaseafterminutes interface{} `json:"ReleaseAfterMinutes"`
}

// UpdateProfileRequest represents the UpdateProfileRequest schema from the OpenAPI specification
type UpdateProfileRequest struct {
	Profilearn interface{} `json:"ProfileArn,omitempty"`
	Profilename interface{} `json:"ProfileName,omitempty"`
	Temperatureunit interface{} `json:"TemperatureUnit,omitempty"`
	Distanceunit interface{} `json:"DistanceUnit,omitempty"`
	Timezone interface{} `json:"Timezone,omitempty"`
	Address interface{} `json:"Address,omitempty"`
	Setupmodedisabled interface{} `json:"SetupModeDisabled,omitempty"`
	Locale interface{} `json:"Locale,omitempty"`
	Maxvolumelimit interface{} `json:"MaxVolumeLimit,omitempty"`
	Meetingroomconfiguration UpdateProfileRequestMeetingRoomConfiguration `json:"MeetingRoomConfiguration,omitempty"`
	Pstnenabled interface{} `json:"PSTNEnabled,omitempty"`
	Wakeword interface{} `json:"WakeWord,omitempty"`
	Dataretentionoptin interface{} `json:"DataRetentionOptIn,omitempty"`
	Isdefault interface{} `json:"IsDefault,omitempty"`
}

// StartDeviceSyncResponse represents the StartDeviceSyncResponse schema from the OpenAPI specification
type StartDeviceSyncResponse struct {
}

// CreateProfileResponse represents the CreateProfileResponse schema from the OpenAPI specification
type CreateProfileResponse struct {
	Profilearn interface{} `json:"ProfileArn,omitempty"`
}

// AssociateContactWithAddressBookResponse represents the AssociateContactWithAddressBookResponse schema from the OpenAPI specification
type AssociateContactWithAddressBookResponse struct {
}

// ForgetSmartHomeAppliancesRequest represents the ForgetSmartHomeAppliancesRequest schema from the OpenAPI specification
type ForgetSmartHomeAppliancesRequest struct {
	Roomarn interface{} `json:"RoomArn"`
}

// UpdateEndOfMeetingReminder represents the UpdateEndOfMeetingReminder schema from the OpenAPI specification
type UpdateEndOfMeetingReminder struct {
	Enabled interface{} `json:"Enabled,omitempty"`
	Reminderatminutes interface{} `json:"ReminderAtMinutes,omitempty"`
	Remindertype interface{} `json:"ReminderType,omitempty"`
}

// ConferenceProvider represents the ConferenceProvider schema from the OpenAPI specification
type ConferenceProvider struct {
	TypeField interface{} `json:"Type,omitempty"`
	Arn interface{} `json:"Arn,omitempty"`
	Ipdialin CreateConferenceProviderRequestIPDialIn `json:"IPDialIn,omitempty"`
	Meetingsetting CreateConferenceProviderRequestMeetingSetting `json:"MeetingSetting,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Pstndialin CreateConferenceProviderRequestPSTNDialIn `json:"PSTNDialIn,omitempty"`
}

// ResolveRoomRequest represents the ResolveRoomRequest schema from the OpenAPI specification
type ResolveRoomRequest struct {
	Skillid interface{} `json:"SkillId"`
	Userid interface{} `json:"UserId"`
}

// DeviceStatusDetail represents the DeviceStatusDetail schema from the OpenAPI specification
type DeviceStatusDetail struct {
	Code interface{} `json:"Code,omitempty"`
	Feature interface{} `json:"Feature,omitempty"`
}

// SearchRoomsRequest represents the SearchRoomsRequest schema from the OpenAPI specification
type SearchRoomsRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Sortcriteria interface{} `json:"SortCriteria,omitempty"`
	Filters interface{} `json:"Filters,omitempty"`
}

// DeviceDeviceStatusInfo represents the DeviceDeviceStatusInfo schema from the OpenAPI specification
type DeviceDeviceStatusInfo struct {
	Connectionstatus interface{} `json:"ConnectionStatus,omitempty"`
	Connectionstatusupdatedtime interface{} `json:"ConnectionStatusUpdatedTime,omitempty"`
	Devicestatusdetails interface{} `json:"DeviceStatusDetails,omitempty"`
}

// UpdateMeetingRoomConfigurationEndOfMeetingReminder represents the UpdateMeetingRoomConfigurationEndOfMeetingReminder schema from the OpenAPI specification
type UpdateMeetingRoomConfigurationEndOfMeetingReminder struct {
	Enabled interface{} `json:"Enabled,omitempty"`
	Reminderatminutes interface{} `json:"ReminderAtMinutes,omitempty"`
	Remindertype interface{} `json:"ReminderType,omitempty"`
}

// AddressBookData represents the AddressBookData schema from the OpenAPI specification
type AddressBookData struct {
	Addressbookarn interface{} `json:"AddressBookArn,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Name interface{} `json:"Name,omitempty"`
}

// ListConferenceProvidersResponse represents the ListConferenceProvidersResponse schema from the OpenAPI specification
type ListConferenceProvidersResponse struct {
	Conferenceproviders interface{} `json:"ConferenceProviders,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// DisassociateSkillFromUsersRequest represents the DisassociateSkillFromUsersRequest schema from the OpenAPI specification
type DisassociateSkillFromUsersRequest struct {
	Skillid interface{} `json:"SkillId"`
}

// DeleteNetworkProfileRequest represents the DeleteNetworkProfileRequest schema from the OpenAPI specification
type DeleteNetworkProfileRequest struct {
	Networkprofilearn interface{} `json:"NetworkProfileArn"`
}

// SkillGroupData represents the SkillGroupData schema from the OpenAPI specification
type SkillGroupData struct {
	Description interface{} `json:"Description,omitempty"`
	Skillgrouparn interface{} `json:"SkillGroupArn,omitempty"`
	Skillgroupname interface{} `json:"SkillGroupName,omitempty"`
}

// ListSmartHomeAppliancesRequest represents the ListSmartHomeAppliancesRequest schema from the OpenAPI specification
type ListSmartHomeAppliancesRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Roomarn interface{} `json:"RoomArn"`
}

// UpdateNetworkProfileResponse represents the UpdateNetworkProfileResponse schema from the OpenAPI specification
type UpdateNetworkProfileResponse struct {
}

// PhoneNumber represents the PhoneNumber schema from the OpenAPI specification
type PhoneNumber struct {
	Number interface{} `json:"Number"`
	TypeField interface{} `json:"Type"`
}

// PutRoomSkillParameterResponse represents the PutRoomSkillParameterResponse schema from the OpenAPI specification
type PutRoomSkillParameterResponse struct {
}

// CreateEndOfMeetingReminder represents the CreateEndOfMeetingReminder schema from the OpenAPI specification
type CreateEndOfMeetingReminder struct {
	Remindertype interface{} `json:"ReminderType"`
	Enabled interface{} `json:"Enabled"`
	Reminderatminutes interface{} `json:"ReminderAtMinutes"`
}

// ListSkillsRequest represents the ListSkillsRequest schema from the OpenAPI specification
type ListSkillsRequest struct {
	Skillgrouparn interface{} `json:"SkillGroupArn,omitempty"`
	Skilltype interface{} `json:"SkillType,omitempty"`
	Enablementtype interface{} `json:"EnablementType,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// UpdateProactiveJoin represents the UpdateProactiveJoin schema from the OpenAPI specification
type UpdateProactiveJoin struct {
	Enabledbymotion bool `json:"EnabledByMotion"`
}

// GetNetworkProfileResponse represents the GetNetworkProfileResponse schema from the OpenAPI specification
type GetNetworkProfileResponse struct {
	Networkprofile GetNetworkProfileResponseNetworkProfile `json:"NetworkProfile,omitempty"`
}

// CreateMeetingRoomConfiguration represents the CreateMeetingRoomConfiguration schema from the OpenAPI specification
type CreateMeetingRoomConfiguration struct {
	Endofmeetingreminder CreateEndOfMeetingReminder `json:"EndOfMeetingReminder,omitempty"` // Creates settings for the end of meeting reminder feature that are applied to a room profile. The end of meeting reminder enables Alexa to remind users when a meeting is ending.
	Instantbooking CreateMeetingRoomConfigurationInstantBooking `json:"InstantBooking,omitempty"`
	Proactivejoin CreateProactiveJoin `json:"ProactiveJoin,omitempty"`
	Requirecheckin CreateMeetingRoomConfigurationRequireCheckIn `json:"RequireCheckIn,omitempty"`
	Roomutilizationmetricsenabled interface{} `json:"RoomUtilizationMetricsEnabled,omitempty"`
}

// AssociateDeviceWithNetworkProfileResponse represents the AssociateDeviceWithNetworkProfileResponse schema from the OpenAPI specification
type AssociateDeviceWithNetworkProfileResponse struct {
}

// MeetingRoomConfigurationInstantBooking represents the MeetingRoomConfigurationInstantBooking schema from the OpenAPI specification
type MeetingRoomConfigurationInstantBooking struct {
	Durationinminutes interface{} `json:"DurationInMinutes,omitempty"`
	Enabled interface{} `json:"Enabled,omitempty"`
}

// DeleteUserRequest represents the DeleteUserRequest schema from the OpenAPI specification
type DeleteUserRequest struct {
	Enrollmentid interface{} `json:"EnrollmentId"`
	Userarn interface{} `json:"UserArn,omitempty"`
}

// UpdateConferenceProviderRequest represents the UpdateConferenceProviderRequest schema from the OpenAPI specification
type UpdateConferenceProviderRequest struct {
	Ipdialin CreateConferenceProviderRequestIPDialIn `json:"IPDialIn,omitempty"`
	Meetingsetting CreateConferenceProviderRequestMeetingSetting `json:"MeetingSetting"`
	Pstndialin CreateConferenceProviderRequestPSTNDialIn `json:"PSTNDialIn,omitempty"`
	Conferenceproviderarn interface{} `json:"ConferenceProviderArn"`
	Conferenceprovidertype interface{} `json:"ConferenceProviderType"`
}

// GetConferencePreferenceResponse represents the GetConferencePreferenceResponse schema from the OpenAPI specification
type GetConferencePreferenceResponse struct {
	Preference GetConferencePreferenceResponsePreference `json:"Preference,omitempty"`
}

// Ssml represents the Ssml schema from the OpenAPI specification
type Ssml struct {
	Locale interface{} `json:"Locale"`
	Value interface{} `json:"Value"`
}

// BusinessReport represents the BusinessReport schema from the OpenAPI specification
type BusinessReport struct {
	Status interface{} `json:"Status,omitempty"`
	Deliverytime interface{} `json:"DeliveryTime,omitempty"`
	Downloadurl interface{} `json:"DownloadUrl,omitempty"`
	Failurecode interface{} `json:"FailureCode,omitempty"`
	S3location BusinessReportS3Location `json:"S3Location,omitempty"`
}

// GetAddressBookResponseAddressBook represents the GetAddressBookResponseAddressBook schema from the OpenAPI specification
type GetAddressBookResponseAddressBook struct {
	Name interface{} `json:"Name,omitempty"`
	Addressbookarn interface{} `json:"AddressBookArn,omitempty"`
	Description interface{} `json:"Description,omitempty"`
}

// MeetingRoomConfigurationEndOfMeetingReminder represents the MeetingRoomConfigurationEndOfMeetingReminder schema from the OpenAPI specification
type MeetingRoomConfigurationEndOfMeetingReminder struct {
	Enabled interface{} `json:"Enabled,omitempty"`
	Reminderatminutes interface{} `json:"ReminderAtMinutes,omitempty"`
	Remindertype interface{} `json:"ReminderType,omitempty"`
}
