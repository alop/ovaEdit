package ovf_struct

func init() {
}

type CDescription__ovf struct {
	XMLName xml.Name `xml:"Description,omitempty" json:"Description,omitempty"`
	string  string   `xml:",chardata" json:",omitempty"`
}

type CDisk__ovf struct {
	XMLName                             xml.Name `xml:"Disk,omitempty" json:"Disk,omitempty"`
	AttrOvfSpacecapacity                string   `xml:"http://schemas.dmtf.org/ovf/envelope/1 capacity,attr"  json:",omitempty"`
	AttrOvfSpacecapacityAllocationUnits string   `xml:"http://schemas.dmtf.org/ovf/envelope/1 capacityAllocationUnits,attr"  json:",omitempty"`
	AttrOvfSpacediskId                  string   `xml:"http://schemas.dmtf.org/ovf/envelope/1 diskId,attr"  json:",omitempty"`
	AttrOvfSpacefileRef                 string   `xml:"http://schemas.dmtf.org/ovf/envelope/1 fileRef,attr"  json:",omitempty"`
	AttrOvfSpaceformat                  string   `xml:"http://schemas.dmtf.org/ovf/envelope/1 format,attr"  json:",omitempty"`
	AttrOvfSpacepopulatedSize           string   `xml:"http://schemas.dmtf.org/ovf/envelope/1 populatedSize,attr"  json:",omitempty"`
}

type CDiskSection__ovf struct {
	XMLName    xml.Name    `xml:"DiskSection,omitempty" json:"DiskSection,omitempty"`
	CDisk__ovf *CDisk__ovf `xml:"http://schemas.dmtf.org/ovf/envelope/1 Disk,omitempty" json:"Disk,omitempty"`
	CInfo__ovf *CInfo__ovf `xml:"http://schemas.dmtf.org/ovf/envelope/1 Info,omitempty" json:"Info,omitempty"`
}

type CEnvelope__ovf struct {
	XMLName              xml.Name              `xml:"Envelope,omitempty" json:"Envelope,omitempty"`
	AttrXmlnscim         string                `xml:"xmlns cim,attr"  json:",omitempty"`
	AttrXmlnsovf         string                `xml:"xmlns ovf,attr"  json:",omitempty"`
	AttrXmlnsrasd        string                `xml:"xmlns rasd,attr"  json:",omitempty"`
	AttrXmlnsvmw         string                `xml:"xmlns vmw,attr"  json:",omitempty"`
	AttrXmlnsvssd        string                `xml:"xmlns vssd,attr"  json:",omitempty"`
	Attrxmlns            string                `xml:"xmlns,attr"  json:",omitempty"`
	AttrXmlnsxsi         string                `xml:"xmlns xsi,attr"  json:",omitempty"`
	CDiskSection__ovf    *CDiskSection__ovf    `xml:"http://schemas.dmtf.org/ovf/envelope/1 DiskSection,omitempty" json:"DiskSection,omitempty"`
	CNetworkSection__ovf *CNetworkSection__ovf `xml:"http://schemas.dmtf.org/ovf/envelope/1 NetworkSection,omitempty" json:"NetworkSection,omitempty"`
	CReferences__ovf     *CReferences__ovf     `xml:"http://schemas.dmtf.org/ovf/envelope/1 References,omitempty" json:"References,omitempty"`
	CVirtualSystem__ovf  *CVirtualSystem__ovf  `xml:"http://schemas.dmtf.org/ovf/envelope/1 VirtualSystem,omitempty" json:"VirtualSystem,omitempty"`
}

type CFile__ovf struct {
	XMLName          xml.Name `xml:"File,omitempty" json:"File,omitempty"`
	AttrOvfSpacehref string   `xml:"http://schemas.dmtf.org/ovf/envelope/1 href,attr"  json:",omitempty"`
	AttrOvfSpaceid   string   `xml:"http://schemas.dmtf.org/ovf/envelope/1 id,attr"  json:",omitempty"`
	AttrOvfSpacesize string   `xml:"http://schemas.dmtf.org/ovf/envelope/1 size,attr"  json:",omitempty"`
}

type CInfo__ovf struct {
	XMLName xml.Name `xml:"Info,omitempty" json:"Info,omitempty"`
	string  string   `xml:",chardata" json:",omitempty"`
}

type CItem__ovf struct {
	XMLName                    xml.Name                    `xml:"Item,omitempty" json:"Item,omitempty"`
	AttrOvfSpacerequired       string                      `xml:"http://schemas.dmtf.org/ovf/envelope/1 required,attr"  json:",omitempty"`
	CAddressOnParent__rasd     *CAddressOnParent__rasd     `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ResourceAllocationSettingData AddressOnParent,omitempty" json:"AddressOnParent,omitempty"`
	CAddress__rasd             *CAddress__rasd             `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ResourceAllocationSettingData Address,omitempty" json:"Address,omitempty"`
	CAllocationUnits__rasd     *CAllocationUnits__rasd     `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ResourceAllocationSettingData AllocationUnits,omitempty" json:"AllocationUnits,omitempty"`
	CAutomaticAllocation__rasd *CAutomaticAllocation__rasd `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ResourceAllocationSettingData AutomaticAllocation,omitempty" json:"AutomaticAllocation,omitempty"`
	CConfig__vmw               []*CConfig__vmw             `xml:"http://www.vmware.com/schema/ovf Config,omitempty" json:"Config,omitempty"`
	CConnection__rasd          *CConnection__rasd          `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ResourceAllocationSettingData Connection,omitempty" json:"Connection,omitempty"`
	CDescription__rasd         *CDescription__rasd         `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ResourceAllocationSettingData Description,omitempty" json:"Description,omitempty"`
	CElementName__rasd         *CElementName__rasd         `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ResourceAllocationSettingData ElementName,omitempty" json:"ElementName,omitempty"`
	CHostResource__rasd        *CHostResource__rasd        `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ResourceAllocationSettingData HostResource,omitempty" json:"HostResource,omitempty"`
	CInstanceID__rasd          *CInstanceID__rasd          `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ResourceAllocationSettingData InstanceID,omitempty" json:"InstanceID,omitempty"`
	CParent__rasd              *CParent__rasd              `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ResourceAllocationSettingData Parent,omitempty" json:"Parent,omitempty"`
	CResourceSubType__rasd     *CResourceSubType__rasd     `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ResourceAllocationSettingData ResourceSubType,omitempty" json:"ResourceSubType,omitempty"`
	CResourceType__rasd        *CResourceType__rasd        `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ResourceAllocationSettingData ResourceType,omitempty" json:"ResourceType,omitempty"`
	CVirtualQuantity__rasd     *CVirtualQuantity__rasd     `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ResourceAllocationSettingData VirtualQuantity,omitempty" json:"VirtualQuantity,omitempty"`
}

type CLabel__ovf struct {
	XMLName xml.Name `xml:"Label,omitempty" json:"Label,omitempty"`
	string  string   `xml:",chardata" json:",omitempty"`
}

type CName__ovf struct {
	XMLName xml.Name `xml:"Name,omitempty" json:"Name,omitempty"`
	string  string   `xml:",chardata" json:",omitempty"`
}

type CNetwork__ovf struct {
	XMLName           xml.Name           `xml:"Network,omitempty" json:"Network,omitempty"`
	AttrOvfSpacename  string             `xml:"http://schemas.dmtf.org/ovf/envelope/1 name,attr"  json:",omitempty"`
	CDescription__ovf *CDescription__ovf `xml:"http://schemas.dmtf.org/ovf/envelope/1 Description,omitempty" json:"Description,omitempty"`
}

type CNetworkSection__ovf struct {
	XMLName       xml.Name       `xml:"NetworkSection,omitempty" json:"NetworkSection,omitempty"`
	CInfo__ovf    *CInfo__ovf    `xml:"http://schemas.dmtf.org/ovf/envelope/1 Info,omitempty" json:"Info,omitempty"`
	CNetwork__ovf *CNetwork__ovf `xml:"http://schemas.dmtf.org/ovf/envelope/1 Network,omitempty" json:"Network,omitempty"`
}

type COperatingSystemSection__ovf struct {
	XMLName            xml.Name           `xml:"OperatingSystemSection,omitempty" json:"OperatingSystemSection,omitempty"`
	AttrOvfSpaceid     string             `xml:"http://schemas.dmtf.org/ovf/envelope/1 id,attr"  json:",omitempty"`
	AttrVmwSpaceosType string             `xml:"http://www.vmware.com/schema/ovf osType,attr"  json:",omitempty"`
	CDescription__ovf  *CDescription__ovf `xml:"http://schemas.dmtf.org/ovf/envelope/1 Description,omitempty" json:"Description,omitempty"`
	CInfo__ovf         *CInfo__ovf        `xml:"http://schemas.dmtf.org/ovf/envelope/1 Info,omitempty" json:"Info,omitempty"`
}

type CProduct__ovf struct {
	XMLName xml.Name `xml:"Product,omitempty" json:"Product,omitempty"`
	string  string   `xml:",chardata" json:",omitempty"`
}

type CProductSection__ovf struct {
	XMLName              xml.Name          `xml:"ProductSection,omitempty" json:"ProductSection,omitempty"`
	AttrOvfSpacerequired string            `xml:"http://schemas.dmtf.org/ovf/envelope/1 required,attr"  json:",omitempty"`
	CInfo__ovf           *CInfo__ovf       `xml:"http://schemas.dmtf.org/ovf/envelope/1 Info,omitempty" json:"Info,omitempty"`
	CProduct__ovf        *CProduct__ovf    `xml:"http://schemas.dmtf.org/ovf/envelope/1 Product,omitempty" json:"Product,omitempty"`
	CProperty__ovf       []*CProperty__ovf `xml:"http://schemas.dmtf.org/ovf/envelope/1 Property,omitempty" json:"Property,omitempty"`
}

type CProperty__ovf struct {
	XMLName                      xml.Name           `xml:"Property,omitempty" json:"Property,omitempty"`
	AttrOvfSpacekey              string             `xml:"http://schemas.dmtf.org/ovf/envelope/1 key,attr"  json:",omitempty"`
	AttrOvfSpacetype             string             `xml:"http://schemas.dmtf.org/ovf/envelope/1 type,attr"  json:",omitempty"`
	AttrOvfSpaceuserConfigurable string             `xml:"http://schemas.dmtf.org/ovf/envelope/1 userConfigurable,attr"  json:",omitempty"`
	AttrOvfSpacevalue            string             `xml:"http://schemas.dmtf.org/ovf/envelope/1 value,attr"  json:",omitempty"`
	CDescription__ovf            *CDescription__ovf `xml:"http://schemas.dmtf.org/ovf/envelope/1 Description,omitempty" json:"Description,omitempty"`
	CLabel__ovf                  *CLabel__ovf       `xml:"http://schemas.dmtf.org/ovf/envelope/1 Label,omitempty" json:"Label,omitempty"`
}

type CReferences__ovf struct {
	XMLName    xml.Name    `xml:"References,omitempty" json:"References,omitempty"`
	CFile__ovf *CFile__ovf `xml:"http://schemas.dmtf.org/ovf/envelope/1 File,omitempty" json:"File,omitempty"`
}

type CSystem__ovf struct {
	XMLName                        xml.Name                        `xml:"System,omitempty" json:"System,omitempty"`
	CElementName__vssd             *CElementName__vssd             `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_VirtualSystemSettingData ElementName,omitempty" json:"ElementName,omitempty"`
	CInstanceID__vssd              *CInstanceID__vssd              `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_VirtualSystemSettingData InstanceID,omitempty" json:"InstanceID,omitempty"`
	CVirtualSystemIdentifier__vssd *CVirtualSystemIdentifier__vssd `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_VirtualSystemSettingData VirtualSystemIdentifier,omitempty" json:"VirtualSystemIdentifier,omitempty"`
	CVirtualSystemType__vssd       *CVirtualSystemType__vssd       `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_VirtualSystemSettingData VirtualSystemType,omitempty" json:"VirtualSystemType,omitempty"`
}

type CVirtualHardwareSection__ovf struct {
	XMLName               xml.Name        `xml:"VirtualHardwareSection,omitempty" json:"VirtualHardwareSection,omitempty"`
	AttrOvfSpacetransport string          `xml:"http://schemas.dmtf.org/ovf/envelope/1 transport,attr"  json:",omitempty"`
	CConfig__vmw          []*CConfig__vmw `xml:"http://www.vmware.com/schema/ovf Config,omitempty" json:"Config,omitempty"`
	CInfo__ovf            *CInfo__ovf     `xml:"http://schemas.dmtf.org/ovf/envelope/1 Info,omitempty" json:"Info,omitempty"`
	CItem__ovf            []*CItem__ovf   `xml:"http://schemas.dmtf.org/ovf/envelope/1 Item,omitempty" json:"Item,omitempty"`
	CSystem__ovf          *CSystem__ovf   `xml:"http://schemas.dmtf.org/ovf/envelope/1 System,omitempty" json:"System,omitempty"`
}

type CVirtualSystem__ovf struct {
	XMLName                      xml.Name                      `xml:"VirtualSystem,omitempty" json:"VirtualSystem,omitempty"`
	AttrOvfSpaceid               string                        `xml:"http://schemas.dmtf.org/ovf/envelope/1 id,attr"  json:",omitempty"`
	CInfo__ovf                   *CInfo__ovf                   `xml:"http://schemas.dmtf.org/ovf/envelope/1 Info,omitempty" json:"Info,omitempty"`
	CName__ovf                   *CName__ovf                   `xml:"http://schemas.dmtf.org/ovf/envelope/1 Name,omitempty" json:"Name,omitempty"`
	COperatingSystemSection__ovf *COperatingSystemSection__ovf `xml:"http://schemas.dmtf.org/ovf/envelope/1 OperatingSystemSection,omitempty" json:"OperatingSystemSection,omitempty"`
	CProductSection__ovf         *CProductSection__ovf         `xml:"http://schemas.dmtf.org/ovf/envelope/1 ProductSection,omitempty" json:"ProductSection,omitempty"`
	CVirtualHardwareSection__ovf *CVirtualHardwareSection__ovf `xml:"http://schemas.dmtf.org/ovf/envelope/1 VirtualHardwareSection,omitempty" json:"VirtualHardwareSection,omitempty"`
}

type CAddress__rasd struct {
	XMLName xml.Name `xml:"Address,omitempty" json:"Address,omitempty"`
	string  string   `xml:",chardata" json:",omitempty"`
}

type CAddressOnParent__rasd struct {
	XMLName xml.Name `xml:"AddressOnParent,omitempty" json:"AddressOnParent,omitempty"`
	string  string   `xml:",chardata" json:",omitempty"`
}

type CAllocationUnits__rasd struct {
	XMLName xml.Name `xml:"AllocationUnits,omitempty" json:"AllocationUnits,omitempty"`
	string  string   `xml:",chardata" json:",omitempty"`
}

type CAutomaticAllocation__rasd struct {
	XMLName xml.Name `xml:"AutomaticAllocation,omitempty" json:"AutomaticAllocation,omitempty"`
	string  string   `xml:",chardata" json:",omitempty"`
}

type CConnection__rasd struct {
	XMLName xml.Name `xml:"Connection,omitempty" json:"Connection,omitempty"`
	string  string   `xml:",chardata" json:",omitempty"`
}

type CDescription__rasd struct {
	XMLName xml.Name `xml:"Description,omitempty" json:"Description,omitempty"`
	string  string   `xml:",chardata" json:",omitempty"`
}

type CElementName__rasd struct {
	XMLName xml.Name `xml:"ElementName,omitempty" json:"ElementName,omitempty"`
	string  string   `xml:",chardata" json:",omitempty"`
}

type CHostResource__rasd struct {
	XMLName xml.Name `xml:"HostResource,omitempty" json:"HostResource,omitempty"`
	string  string   `xml:",chardata" json:",omitempty"`
}

type CInstanceID__rasd struct {
	XMLName xml.Name `xml:"InstanceID,omitempty" json:"InstanceID,omitempty"`
	string  string   `xml:",chardata" json:",omitempty"`
}

type CParent__rasd struct {
	XMLName xml.Name `xml:"Parent,omitempty" json:"Parent,omitempty"`
	string  string   `xml:",chardata" json:",omitempty"`
}

type CResourceSubType__rasd struct {
	XMLName xml.Name `xml:"ResourceSubType,omitempty" json:"ResourceSubType,omitempty"`
	string  string   `xml:",chardata" json:",omitempty"`
}

type CResourceType__rasd struct {
	XMLName xml.Name `xml:"ResourceType,omitempty" json:"ResourceType,omitempty"`
	string  string   `xml:",chardata" json:",omitempty"`
}

type CVirtualQuantity__rasd struct {
	XMLName xml.Name `xml:"VirtualQuantity,omitempty" json:"VirtualQuantity,omitempty"`
	string  string   `xml:",chardata" json:",omitempty"`
}

type CElementName__vssd struct {
	XMLName xml.Name `xml:"ElementName,omitempty" json:"ElementName,omitempty"`
	string  string   `xml:",chardata" json:",omitempty"`
}

type CInstanceID__vssd struct {
	XMLName xml.Name `xml:"InstanceID,omitempty" json:"InstanceID,omitempty"`
	string  string   `xml:",chardata" json:",omitempty"`
}

type CVirtualSystemIdentifier__vssd struct {
	XMLName xml.Name `xml:"VirtualSystemIdentifier,omitempty" json:"VirtualSystemIdentifier,omitempty"`
	string  string   `xml:",chardata" json:",omitempty"`
}

type CVirtualSystemType__vssd struct {
	XMLName xml.Name `xml:"VirtualSystemType,omitempty" json:"VirtualSystemType,omitempty"`
	string  string   `xml:",chardata" json:",omitempty"`
}

type CConfig__vmw struct {
	XMLName              xml.Name `xml:"Config,omitempty" json:"Config,omitempty"`
	AttrVmwSpacekey      string   `xml:"http://www.vmware.com/schema/ovf key,attr"  json:",omitempty"`
	AttrOvfSpacerequired string   `xml:"http://schemas.dmtf.org/ovf/envelope/1 required,attr"  json:",omitempty"`
	AttrVmwSpacevalue    string   `xml:"http://www.vmware.com/schema/ovf value,attr"  json:",omitempty"`
}
